package blockchain

import (
	"cbc-casper-go/casper"
	"cbc-casper-go/casper/safety_oracles"
	"fmt"
)

type View struct {
	*casper.View
	lastFinBlk *Block
	genesisBlk *Block
	children   map[*Block][]*Block
}

func NewView() casper.Viewer {
	return &View{
		View:       casper.NewView().(*casper.View),
		lastFinBlk: nil,
		genesisBlk: nil,
		children:   make(map[*Block][]*Block),
	}
}

func (v *View) Estimate() interface{} {
	return getForkChoice(v.lastFinBlk, v.children, v.LatestMsg())
}

func (v *View) UpdateSafeEstimates(valset *casper.ValidatorSet) casper.Messager {
	tip := v.Estimate().(*Block)
	for tip != nil && tip != v.lastFinBlk {
		oracle, err := safety_oracles.NewCliqueOracle(tip, v.View, valset)
		if err != nil {
			_ = fmt.Errorf("error")
			break
		}
		faultTolerance, _ := oracle.CheckEstimateSafety()
		if faultTolerance > 0 {
			v.lastFinBlk = tip
			break
		}
		tip = tip.Estimate.(*Block)
	}
	return v.lastFinBlk
}

func (v *View) updateProtocolSpecificView(block *Block) {
	if _, ok := v.JustifiedMsg()[block.Hash()]; ok {
		_ = fmt.Errorf("error")
	}
	if _, ok := v.children[block.Estimate.(*Block)]; !ok {
		v.children[block.Estimate.(*Block)] = make([]*Block, 0, 4)
	}
	v.children[block.Estimate.(*Block)] = append(v.children[block.Estimate.(*Block)], block)
}

// AddMessages 添加新的message到pending或者justify
func (v *View) AddMessages(msgs []casper.Messager) {
	for _, msg := range msgs {
		if v.Contain(msg) {
			continue
		}
		missMsgHashes := v.MissMsgInJustify(msg)
		if len(missMsgHashes) == 0 {
			v.ReceiveJustifiedMsg(msg)
		} else {
			v.ReceivePendingMsg(msg, missMsgHashes)
		}
	}
}

// ReceiveJustifiedMsg 在收到已验证的消息后,处理等待队列并添加到View中
func (v *View) ReceiveJustifiedMsg(m casper.Messager) {
	messages := v.GetJustifiedMsg(m)
	for _, message := range messages {
		v.AddToLatestMessage(message)
		v.AddJustifiedRemovePending(message)
		v.updateProtocolSpecificView(message.(*Block))
	}
}
