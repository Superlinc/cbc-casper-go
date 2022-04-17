package concurrent

import "cbc-casper-go/casper/simulation"

func getProtocol(weights []uint64) *Protocol {
	if weights == nil {
		weights = []uint64{10, 9, 8, 7, 6}
	}
	str := simulation.GenerateConcurrentJsonString(weights, []int{0, 100, 200}, []int{0, 100}, "all", "all")
	protocol, _ := NewProtocol(str, 1)
	return protocol
}
