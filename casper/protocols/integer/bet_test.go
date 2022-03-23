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
	message1 := &casper.Message{
		Sender:        nil,
		Estimate:      0,
		Justification: nil,
		SeqNum:        0,
		DisplayHeight: 0,
		Header:        0,
	}
	message2 := &casper.Message{
		Sender:        nil,
		Estimate:      0,
		Justification: nil,
		SeqNum:        0,
		DisplayHeight: 0,
		Header:        0,
	}
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
