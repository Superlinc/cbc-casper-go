package order

import (
	"cbc-casper-go/casper"
	"container/list"
	"testing"
)

func TestIsValid(t *testing.T) {
	if isValidEstimate(1) {
		t.Errorf("error")
	}
	if !isValidEstimate(list.New()) {
		t.Errorf("error")
	}
}

func TestBet_ConflictWith(t *testing.T) {
	m1 := casper.NewMessage(list.New(), nil, nil, 0, 0)
	m2 := casper.NewMessage(list.New(), nil, nil, 0, 0)
	l1 := m1.Estimate.(*list.List)
	l2 := m2.Estimate.(*list.List)
	l1.PushBack(1)
	l2.PushBack(1)
	b1 := &Bet{m1}
	b2 := &Bet{m2}
	if ok, _ := b1.ConflictWith(b2.Message); ok {
		t.Errorf("error")
	}
	l1.PushBack(2)
	if ok, _ := b1.ConflictWith(b2.Message); !ok {
		t.Errorf("error")
	}
	l1.Init()
	l2.Init()
	l1.PushBack("hello")
	l2.PushBack("hello")
	if ok, _ := b1.ConflictWith(b2.Message); ok {
		t.Errorf("error")
	}
}
