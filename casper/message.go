package casper

import (
	"hash/crc64"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Message struct {
	sender        *Validator
	Estimate      interface{}
	justification map[*Validator]uint64
	seqNum        uint64
	disHeight     uint64
	header        float64
}

func NewMessage(estimate interface{}, justification map[*Validator]uint64, sender *Validator, seqNum uint64, disHeight uint64) *Message {
	return &Message{
		Estimate:      estimate,
		justification: justification,
		sender:        sender,
		seqNum:        seqNum,
		disHeight:     disHeight,
		header:        rand.Float64(),
	}
}

func (m *Message) Sender() *Validator {
	return m.sender
}

func (m *Message) Justification() map[*Validator]uint64 {
	return m.justification
}

func (m *Message) SeqNum() uint64 {
	return m.seqNum
}

func (m *Message) DisHeight() uint64 {
	return m.disHeight
}

func (m *Message) Hash() uint64 {
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	_, err := h.Write([]byte(strconv.FormatFloat(m.header, 'f', 10, 64)))
	if err != nil {
		return 0
	}
	return h.Sum64()
}
