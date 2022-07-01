package ackhandler

import (
	"fmt"
	"time"

	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/utils"
)

type sentPacketHistory struct {
	rttStats              *utils.RTTStats
	outstandingPacketList *PacketList
	etcPacketList         *PacketList
	packetMap             map[protocol.PacketNumber]*PacketElement
	highestSent           protocol.PacketNumber
}

func newSentPacketHistory(rttStats *utils.RTTStats) *sentPacketHistory {
	return &sentPacketHistory{
		rttStats:              rttStats,
		outstandingPacketList: NewPacketList(),
		etcPacketList:         NewPacketList(),
		packetMap:             make(map[protocol.PacketNumber]*PacketElement),
		highestSent:           protocol.InvalidPacketNumber,
	}
}

func (h *sentPacketHistory) SentPacket(p *Packet, isAckEliciting bool) {
	if p.PacketNumber <= h.highestSent {
		panic("non-sequential packet number use")
	}
	// Skipped packet numbers.
	for pn := h.highestSent + 1; pn < p.PacketNumber; pn++ {
		el := h.etcPacketList.PushBack(Packet{
			PacketNumber:    pn,
			EncryptionLevel: p.EncryptionLevel,
			SendTime:        p.SendTime,
			skippedPacket:   true,
		})
		h.packetMap[pn] = el
	}
	h.highestSent = p.PacketNumber

	if isAckEliciting {
		var el *PacketElement
		if p.outstanding() {
			el = h.outstandingPacketList.PushBack(*p)
		} else {
			el = h.etcPacketList.PushBack(*p)
		}
		h.packetMap[p.PacketNumber] = el
	}
}

// Iterate iterates through all packets.
func (h *sentPacketHistory) Iterate(cb func(*Packet) (cont bool, err error)) error {
	cont := true
	nextOutstanding := h.outstandingPacketList.Front()
	nextEtc := h.etcPacketList.Front()
	var nextEl *PacketElement
	// whichever has the next packet number is returned first
	for cont {
		if nextOutstanding == nil || (nextEtc != nil && nextEtc.Value.PacketNumber < nextOutstanding.Value.PacketNumber) {
			nextEl = nextEtc
		} else {
			nextEl = nextOutstanding
		}
		if nextEl == nil {
			return nil
		} else if nextEl == nextOutstanding {
			nextOutstanding = nextOutstanding.Next()
		} else {
			nextEtc = nextEtc.Next()
		}
		var err error
		cont, err = cb(&nextEl.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

// FirstOutstanding returns the first outstanding packet.
func (h *sentPacketHistory) FirstOutstanding() *Packet {
	el := h.outstandingPacketList.Front()
	if el == nil {
		return nil
	} else {
		return &el.Value
	}
}

func (h *sentPacketHistory) Len() int {
	return len(h.packetMap)
}

func (h *sentPacketHistory) Remove(p protocol.PacketNumber) error {
	el, ok := h.packetMap[p]
	if !ok {
		return fmt.Errorf("packet %d not found in sent packet history", p)
	}
	h.outstandingPacketList.Remove(el)
	h.etcPacketList.Remove(el)
	delete(h.packetMap, p)
	return nil
}

func (h *sentPacketHistory) HasOutstandingPackets() bool {
	return h.outstandingPacketList.Len() > 0
}

func (h *sentPacketHistory) DeleteOldPackets(now time.Time) {
	maxAge := 3 * h.rttStats.PTO(false)
	var nextEl *PacketElement
	for el := h.etcPacketList.Front(); el != nil; el = nextEl {
		nextEl = el.Next()
		p := el.Value
		if p.SendTime.After(now.Add(-maxAge)) {
			break
		}
		delete(h.packetMap, p.PacketNumber)
		h.etcPacketList.Remove(el)
	}
}

func (h *sentPacketHistory) DeclareLost(p *Packet) {
	el, ok := h.packetMap[p.PacketNumber]
	if !ok {
		return
	}
	h.outstandingPacketList.Remove(el)
	h.etcPacketList.Remove(el)
	p.declaredLost = true
	// move it to the correct position in the etc list (based on the packet number)
	for el = h.etcPacketList.Back(); el != nil; el = el.Prev() {
		if el.Value.PacketNumber < p.PacketNumber {
			break
		}
	}
	if el == nil {
		h.packetMap[p.PacketNumber] = h.etcPacketList.PushFront(*p)
	} else {
		h.packetMap[p.PacketNumber] = h.etcPacketList.InsertAfter(*p, el)
	}
}
