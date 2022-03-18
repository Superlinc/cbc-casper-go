package binary

import (
	"cbc-casper-go/casper"
	"errors"
	"fmt"
)

type Bet struct {
	*casper.Message
}

func (b *Bet) conflictWith(message *Bet) (bool, error) {
	if message.Estimate != 0 && message.Estimate != 1 {
		_ = fmt.Errorf("estimate should be binary")
		return true, errors.New("message estimate error")
	}
	return b.Estimate != message.Estimate, nil
}

func isValidEstimate(estimate int) bool {
	if estimate == 0 || estimate == 1 {
		return true
	} else {
		return false
	}
}
