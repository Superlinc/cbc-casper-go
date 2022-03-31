package order

import (
	"cbc-casper-go/casper"
	"container/list"
	"fmt"
	"testing"
)

func TestGetEstimate(t *testing.T) {
	weights := []uint64{5}
	latestEstimates := make(map[int]*list.List)
	latestEstimates[0] = list.New()
	latestEstimates[0].PushBack(1)
	latestEstimates[0].PushBack(5)
	latestEstimates[0].PushBack(6)
	estimate := list.New()
	estimate.PushBack(1)
	estimate.PushBack(5)
	estimate.PushBack(6)
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 6, 7}
	latestEstimates = make(map[int]*list.List)
	latestEstimates[0] = list.New()
	latestEstimates[0].PushBack(1)
	latestEstimates[0].PushBack(2)
	latestEstimates[1] = list.New()
	latestEstimates[1].PushBack(1)
	latestEstimates[1].PushBack(2)
	latestEstimates[2] = list.New()
	latestEstimates[2].PushBack(2)
	latestEstimates[2].PushBack(1)
	estimate.Init()
	estimate.PushBack(1)
	estimate.PushBack(2)
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 6, 7}
	latestEstimates = make(map[int]*list.List)
	latestEstimates[0] = list.New()
	latestEstimates[0].PushBack(1)
	latestEstimates[0].PushBack(2)
	latestEstimates[0].PushBack(3)
	latestEstimates[1] = list.New()
	latestEstimates[1].PushBack(2)
	latestEstimates[1].PushBack(3)
	latestEstimates[1].PushBack(1)
	latestEstimates[2] = list.New()
	latestEstimates[2].PushBack(2)
	latestEstimates[2].PushBack(1)
	latestEstimates[2].PushBack(3)
	estimate.Init()
	estimate.PushBack(2)
	estimate.PushBack(1)
	estimate.PushBack(3)
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 10, 14}
	latestEstimates = make(map[int]*list.List)
	latestEstimates[0] = list.New()
	latestEstimates[0].PushBack("fish")
	latestEstimates[0].PushBack("pig")
	latestEstimates[0].PushBack("horse")
	latestEstimates[0].PushBack("dog")
	latestEstimates[1] = list.New()
	latestEstimates[1].PushBack("dog")
	latestEstimates[1].PushBack("horse")
	latestEstimates[1].PushBack("pig")
	latestEstimates[1].PushBack("fish")
	latestEstimates[2] = list.New()
	latestEstimates[2].PushBack("pig")
	latestEstimates[2].PushBack("horse")
	latestEstimates[2].PushBack("fish")
	latestEstimates[2].PushBack("dog")
	estimate.Init()
	estimate.PushBack("pig")
	estimate.PushBack("horse")
	estimate.PushBack("dog")
	estimate.PushBack("fish")
	testGetEstimate(weights, latestEstimates, estimate, t)
	weights = []uint64{5, 6, 7, 8}
	latestEstimates = make(map[int]*list.List)
	latestEstimates[0] = list.New()
	latestEstimates[0].PushBack("fish")
	latestEstimates[0].PushBack("pig")
	latestEstimates[0].PushBack("horse")
	latestEstimates[1] = list.New()
	latestEstimates[1].PushBack("horse")
	latestEstimates[1].PushBack("pig")
	latestEstimates[1].PushBack("fish")
	latestEstimates[2] = list.New()
	latestEstimates[2].PushBack("pig")
	latestEstimates[2].PushBack("horse")
	latestEstimates[2].PushBack("fish")
	latestEstimates[3] = list.New()
	latestEstimates[3].PushBack("fish")
	latestEstimates[3].PushBack("horse")
	latestEstimates[3].PushBack("pig")
	estimate.Init()
	estimate.PushBack("horse")
	estimate.PushBack("fish")
	estimate.PushBack("pig")
	testGetEstimate(weights, latestEstimates, estimate, t)

}

func testGetEstimate(weights []uint64, latestEstimates map[int]*list.List, target *list.List, t *testing.T) {
	latestMsg := make(map[*casper.Validator]*casper.Message)
	valSet := casper.NewValidatorSet(weights)
	for name, esti := range latestEstimates {
		val := valSet.GetValByName(name)
		latestMsg[val] = &casper.Message{
			Estimate:      esti,
			Sender:        val,
			SeqNum:        1,
			DisplayHeight: 1,
		}
	}
	estimate := getEstimate(latestMsg)
	if !listEqual(target, estimate) {
		t.Errorf("error")
		iter := estimate.Front()
		for iter != nil {
			fmt.Print(iter.Value, " ")
			iter = iter.Next()
		}
		fmt.Println()
	}
}
