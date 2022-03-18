package casper

import (
	"hash/crc64"
	"strconv"
)

type Message struct {
	Sender        *Validator
	Estimate      interface{}
	Justification map[*Validator]uint64
	SeqNum        uint64
	DisplayHeight uint64
	Header        float64
}

func (m *Message) Hash() uint64 {
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	_, err := h.Write([]byte(strconv.FormatFloat(m.Header, 'f', 10, 64)))
	if err != nil {
		return 0
	}
	return h.Sum64()
}
