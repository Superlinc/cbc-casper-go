package casper

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Validator struct {
	Name         int
	Weight       uint64
	ValidatorSet *ValidatorSet
	View         *View
}

func (v *Validator) InitializeView(messages []*Message) {
	v.View = NewView(messages)
}

func (valSet *ValidatorSet) Contains(validator *Validator) bool {
	return valSet.validators.Contains(validator)
}

func (v *Validator) ReceiveMessages(messages []*Message) {
	v.View.AddMessages(messages)
}

func (v *Validator) Estimate() interface{} {
	return v.View.Estimate()
}

func (v *Validator) MyLatestMessage() *Message {
	return v.View.LatestMessages[v]
}

// MakeNewMessage 为该验证器生成一条最新消息
func (v *Validator) MakeNewMessage() *Message {
	newMessage := &Message{
		Estimate:      v.Estimate(),
		Justification: v.Justification(),
		Sender:        v,
		SeqNum:        v.NextSeqNum(),
		DisplayHeight: v.NextDisPlayHeight(),
		Header:        rand.Float64(),
	}
	v.View.AddMessages([]*Message{newMessage})
	return newMessage
}

// Justification 返回最新消息的哈希值
func (v *Validator) Justification() map[*Validator]uint64 {
	latestMessageHashes := make(map[*Validator]uint64)
	for validator := range v.View.LatestMessages {
		latestMessageHashes[validator] = v.View.LatestMessages[validator].Hash()
	}
	return latestMessageHashes
}

// NextSeqNum 返回该验证器下一个序列号
func (v *Validator) NextSeqNum() uint64 {
	latestMessage := v.MyLatestMessage()
	if latestMessage != nil {
		return latestMessage.SeqNum + 1
	} else {
		return 0
	}

}

// NextDisPlayHeight 返回下一个区块号
func (v *Validator) NextDisPlayHeight() uint64 {
	if len(v.View.LatestMessages) == 0 {
		return 0
	}
	var max uint64
	for _, m := range v.View.LatestMessages {
		max = MaxUint(m.DisplayHeight, max)
	}
	return max + 1
}
