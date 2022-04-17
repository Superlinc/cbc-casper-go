package concurrent

import (
	"cbc-casper-go/casper"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
)

type selectFunc func(*hashset.Set, map[interface{}]*Block) *hashset.Set
type createFunc func(*hashset.Set, int) *hashset.Set

type View struct {
	*casper.View
	children      map[*Block][]*Block
	selectOutputs selectFunc
	createOutputs createFunc
}

func NewView() casper.Viewer {
	return &View{
		View:     casper.NewView().(*casper.View),
		children: make(map[*Block][]*Block),
	}
}

func (v *View) Estimate() interface{} {
	avaiOutputs, sources := getForkChoice(v.children, v.LatestMsg())
	oldOutputs := v.selectOutputs(avaiOutputs, sources)
	newOutputs := v.createOutputs(oldOutputs, oldOutputs.Size())
	blocks := make([]*Block, 0, oldOutputs.Size())
	for _, output := range oldOutputs.Values() {
		blocks = append(blocks, sources[output])
	}
	estimate := make(map[string]interface{})
	estimate["blocks"] = blocks
	estimate["inputs"] = oldOutputs
	estimate["outputs"] = newOutputs
	return estimate
}

func (v *View) setRewriteRules(selectOutputs selectFunc, createOutputs createFunc) {
	v.selectOutputs = selectOutputs
	v.createOutputs = createOutputs
}

func (v *View) updateProtocolSpecificView(block *Block) {
	if _, ok := v.JustifiedMsg()[block.Hash()]; ok {
		_ = fmt.Errorf("block error")
	}
	for _, ancestor := range block.Estimate.(map[string]interface{})["blocks"].([]*Block) {
		if _, ok := v.children[ancestor]; !ok {
			v.children[ancestor] = make([]*Block, 0, 4)
		}
		v.children[ancestor] = append(v.children[ancestor], block)
	}
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
