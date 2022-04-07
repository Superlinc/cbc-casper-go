package order

import (
	"cbc-casper-go/casper"
	"container/list"
	"sort"
)

func getEstimate(latestMessage map[casper.AbstractValidator]casper.Messager) *list.List {
	weights := make(map[interface{}]uint64)

	for validator, m := range latestMessage {
		message := m.(*casper.Message)
		estimate := message.Estimate.(*list.List)
		for i, iter := 0, estimate.Front(); iter != nil; i, iter = i+1, iter.Next() {
			weights[iter.Value] += validator.Weight() * uint64(estimate.Len()-i)
		}
	}

	arr := make([]interface{}, 0, len(weights))
	for elem := range weights {
		arr = append(arr, elem)
	}

	sort.Slice(arr, func(i, j int) bool {
		return weights[arr[i]] < weights[arr[j]]
	})

	l := list.New()
	for _, elem := range arr {
		l.PushFront(elem)
	}

	return l
}
