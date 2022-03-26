package order

import (
	"cbc-casper-go/casper"
	"container/list"
	"errors"
)

type Bet struct {
	*casper.Message
}

func isValidEstimate(estimate interface{}) bool {
	_, ok := estimate.(*list.List)
	return ok
}

func listEqual(l1, l2 *list.List) bool {
	if l1.Len() != l2.Len() {
		return false
	}
	for x, y := l1.Front(), l2.Front(); x != nil; x, y = x.Next(), y.Next() {
		if x.Value != y.Value {
			return false
		}
	}
	return true
}

func (b *Bet) ConflictWith(message *casper.Message) (bool, error) {
	if !isValidEstimate(b.Estimate) {
		return true, errors.New("message should be list")
	}
	l1 := b.Estimate.(*list.List)
	l2 := message.Estimate.(*list.List)
	if listEqual(l1, l2) {
		return false, nil
	} else {
		return true, errors.New("list estimate conflict")
	}
}
