package integer

import (
	"cbc-casper-go/casper"
	"fmt"
	"testing"
)

func TestGetEstimate(t *testing.T) {
	weights := []uint64{5}
	latestEstimates := make(map[int]int)
	latestEstimates[0] = 5
	estimate := 5
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 6, 7}
	latestEstimates = make(map[int]int)
	latestEstimates[0] = 5
	latestEstimates[1] = 5
	latestEstimates[2] = 5
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 10, 14}
	latestEstimates = make(map[int]int)
	latestEstimates[0] = 0
	latestEstimates[1] = 5
	latestEstimates[2] = 10
	estimate = 5
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 11}
	latestEstimates = make(map[int]int)
	latestEstimates[0] = 0
	latestEstimates[1] = 6
	estimate = 6
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 10, 14}
	latestEstimates = make(map[int]int)
	latestEstimates[0] = 0
	latestEstimates[1] = 0
	latestEstimates[2] = 1
	estimate = 0
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 5}
	latestEstimates = make(map[int]int)
	latestEstimates[0] = 0
	latestEstimates[1] = 1
	estimate = 0
	testGetEstimate(weights, latestEstimates, estimate, t)

}

func testGetEstimate(weights []uint64, latestEstimates map[int]int, estimate int, t *testing.T) {
	latestMsg := make(map[*casper.Validator]casper.Messager)
	valSet := casper.NewValidatorSet(weights, nil)
	for name, esti := range latestEstimates {
		val := valSet.GetValByName(name)
		latestMsg[val] = casper.NewMessage(esti, nil, val, 1, 1)
	}
	if estimate != getEstimate(latestMsg) {
		t.Errorf("error")
		fmt.Println(getEstimate(latestMsg))
	}
}
