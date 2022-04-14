package blockchain

import (
	"cbc-casper-go/casper/simulation"
	"fmt"
)

func getProtocol(weights []uint64) *Protocol {
	if weights == nil {
		weights = []uint64{10, 9, 8, 7, 6}
	}
	estimates := make([]interface{}, len(weights))
	str := simulation.GenerateBlockchainJsonString(weights, "", estimates)
	p, err := NewBlockchainProtocol(str, 1)
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return nil
	}
	return p
}