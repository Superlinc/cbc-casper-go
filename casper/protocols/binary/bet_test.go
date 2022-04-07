package binary

import (
	"cbc-casper-go/casper"
	"testing"
)

func TestEstimate(t *testing.T) {
	if !isValidEstimate(1) || !isValidEstimate(0) {
		t.Errorf("error")
	}
	if isValidEstimate(2) || isValidEstimate(-1) {
		t.Errorf("error")
	}
}

func TestBet(t *testing.T) {
	b0 := &Bet{
		casper.NewMessage(0, nil, nil, 0, 0),
	}
	b1 := &Bet{
		casper.NewMessage(1, nil, nil, 0, 0),
	}
	ok, err := b0.ConflictWith(b1.Message)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !ok {
		t.Errorf("error")
	}
	b1.Estimate = 0
	ok, err = b0.ConflictWith(b1.Message)
	if err != nil {
		t.Errorf(err.Error())
	}
	if ok {
		t.Errorf("error")
	}
}
