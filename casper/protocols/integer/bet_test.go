package integer

import (
	"cbc-casper-go/casper"
	"testing"
)

func TestIsValidEstimate(t *testing.T) {
	if !isValidEstimate(1) {
		t.Errorf("error")
	}
	if isValidEstimate(true) {
		t.Errorf("error")
	}
}

func TestBet_ConflictWith(t *testing.T) {
	message1 := casper.NewMessage(0, nil, nil, 0, 0)
	message2 := casper.NewMessage(0, nil, nil, 0, 0)
	bet1 := &Bet{message1}
	bet2 := &Bet{message2}
	if ok, _ := bet1.ConflictWith(bet2.Message); ok {
		t.Errorf("error")
	}
	bet2.Message.Estimate = 1
	if ok, _ := bet1.ConflictWith(bet2.Message); !ok {
		t.Errorf("error")
	}
	bet2.Message.Estimate = true
	if ok, _ := bet1.ConflictWith(bet2.Message); !ok {
		t.Errorf("error")
	}
}
