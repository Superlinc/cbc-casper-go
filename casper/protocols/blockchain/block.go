package blockchain

import (
	"cbc-casper-go/casper"
	"errors"
)

type Block struct {
	*casper.Message
}

func (b *Block) ConflictWith(message interface{}) (bool, error) {
	if !isValidEstimate(message) {
		return false, errors.New("error message")
	}
	return b.isInBlockChain(message.(*Block)), nil
}

func (b *Block) isInBlockChain(block *Block) bool {
	if block == b {
		return true
	} else if block == nil {
		return false
	} else {
		return b.isInBlockChain(block.Estimate.(*Block))
	}
}

func isValidEstimate(estimate interface{}) bool {
	_, ok := estimate.(*Block)
	return ok
}
