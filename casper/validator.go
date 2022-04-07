package casper

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Validator struct {
	name   int
	weight uint64
	valSet ValidatorSetor
	view   Viewer
}

func (v *Validator) Weight() uint64 {
	return v.weight
}

func (v *Validator) Name() int {
	return v.name
}

func (v *Validator) View() Viewer {
	return v.view
}

func (v *Validator) InitializeView(messages []Messager) {
	v.view = NewView(messages)
}

func (v *Validator) ReceiveMessages(messages []Messager) {
	v.view.AddMessages(messages)
}

func (v *Validator) Estimate() interface{} {
	return v.view.Estimate()
}

func (v *Validator) MyLatestMessage() Messager {
	return v.view.LatestMsg()[v]
}

// MakeNewMessage 为该验证器生成一条最新消息
func (v *Validator) MakeNewMessage() Messager {
	newMsg := NewMessage(v.Estimate(), v.Justification(), v, v.NextSeqNum(), v.NextDisPlayHeight())
	v.view.AddMessages([]Messager{newMsg})
	return newMsg
}

// Justification 返回最新消息的哈希值
func (v *Validator) Justification() map[AbstractValidator]uint64 {
	latestMsgHashes := make(map[AbstractValidator]uint64)
	for validator := range v.view.LatestMsg() {
		latestMsgHashes[validator] = v.view.LatestMsg()[validator].Hash()
	}
	return latestMsgHashes
}

// NextSeqNum 返回该验证器下一个序列号
func (v *Validator) NextSeqNum() uint64 {
	msg := v.MyLatestMessage()
	if msg != nil {
		return msg.SeqNum() + 1
	} else {
		return 0
	}

}

// NextDisPlayHeight 返回下一个区块号
func (v *Validator) NextDisPlayHeight() uint64 {
	if len(v.view.LatestMsg()) == 0 {
		return 0
	}
	var max uint64
	for _, m := range v.view.LatestMsg() {
		max = MaxUint(m.DisHeight(), max)
	}
	return max + 1
}
