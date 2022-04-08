package blockchain

import (
	"cbc-casper-go/casper"
	"errors"
)

type Block struct {
	*casper.Message
}

func (b *Block) ConflictWith(message casper.Messager) (bool, error) {
	if !isValidEstimate(message) {
		return false, errors.New("error message")
	}
	return b.isInBlockChain(message), nil
}

func (b *Block) isInBlockChain(m casper.Messager) bool {
	block, ok := m.(*Block)
	if !ok {
		return false
	}
	if block == b {
		return true
	} else if block == nil {
		return false
	} else {
		return b.isInBlockChain(block.Estimate.(casper.Messager))
	}
}

func isValidEstimate(estimate interface{}) bool {
	_, ok := estimate.(*Block)
	return ok
}
