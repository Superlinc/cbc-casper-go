package concurrent

import (
	"cbc-casper-go/casper"
	"errors"
)

type Block struct {
	*casper.Message
}

func (b *Block) ConflictWith(message casper.Messager) (bool, error) {
	if !isValidEstimate(message) {
		return true, errors.New("estimate error")
	}
	return !b.isInHistory(message.(*Block)), nil
}

func (b *Block) isInHistory(message *Block) bool {
	if b == message {
		return true
	}
	m := message.Estimate.(map[string]interface{})
	if len(m) == 1 {
		for _, block := range m["blocks"].([]*Block) {
			if block == nil {
				return false
			}
		}
	}

	for _, block := range m["blocks"].([]*Block) {
		if b.isInHistory(block) {
			return true
		}
	}
	return false
}

func isValidEstimate(estimate interface{}) bool {
	m, ok := estimate.(map[string]interface{})
	if !ok {
		return false
	}
	if len(m) != 3 {
		return false
	}
	for _, field := range []string{"blocks", "inputs", "outputs"} {
		if _, ok = m[field]; !ok {
			return false
		}
	}
	if blocks, ok := m["blocks"].([]*Block); !ok || len(blocks) < 1 {
		return false
	}
	return true
}
