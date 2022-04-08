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

func (v *View) Estimate() interface{} {
	return getForkChoice(v.lastFinBlk, v.children, v.LatestMsg())
}

func (v *View) UpdateSafeEstimates(valset *casper.ValidatorSet) {
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
		}
		tip = tip.Estimate.(*Block)
	}
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
