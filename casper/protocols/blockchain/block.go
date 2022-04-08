package blockchain

import (
	"cbc-casper-go/casper"
	"errors"
)

type Block struct {
	*casper.Message
	height uint64
}

func NewBlock(estimate *Block, justification map[*casper.Validator]uint64, sender *casper.Validator, seqNum, displayHeight uint64) *Block {
	block := &Block{
		Message: casper.NewMessage(estimate, justification, sender, seqNum, displayHeight),
		height:  1,
	}
	if estimate != nil {
		block.height = estimate.height + 1
	}
	return block
}

func (b *Block) ConflictWith(message casper.Messager) (bool, error) {
	if !isValidEstimate(message) {
		return false, errors.New("error message")
	}
	return b.isInBlockChain(message), nil
}

func (b *Block) isInBlockChain(m casper.Messager) bool {
	if m == nil {
		return true
	}
	block, ok := m.(*Block)
	if !ok {
		return true
	}
	if block == b {
		return false
	} else if block == nil || block.Estimate == nil {
		return true
	} else {
		return b.isInBlockChain(block.Estimate.(casper.Messager))
	}
}

func isValidEstimate(estimate interface{}) bool {
	if estimate == nil {
		return true
	}
	_, ok := estimate.(*Block)
	return ok
}
