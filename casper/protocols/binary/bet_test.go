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
		&casper.Message{
			Sender:        nil,
			Estimate:      0,
			Justification: nil,
			SeqNum:        0,
			DisplayHeight: 0,
			Header:        0,
		},
	}
	b1 := &Bet{
		&casper.Message{
			Sender:        nil,
			Estimate:      1,
			Justification: nil,
			SeqNum:        0,
			DisplayHeight: 0,
			Header:        0,
		},
	}
	ok, err := b0.conflictWith(b1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !ok {
		t.Errorf("error")
	}
	b1.Estimate = 0
	ok, err = b0.conflictWith(b1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if ok {
		t.Errorf("error")
	}
}
