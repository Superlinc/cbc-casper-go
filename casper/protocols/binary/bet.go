package binary

import (
	"cbc-casper-go/casper"
	"errors"
	"fmt"
)

type Bet struct {
	*casper.Message
}

func (b *Bet) ConflictWith(message *casper.Message) (bool, error) {
	if message.Estimate != 0 && message.Estimate != 1 {
		_ = fmt.Errorf("estimate should be binary")
		return true, errors.New("message estimate error")
	}
	return b.Estimate != message.Estimate, nil
}

func isValidEstimate(estimate interface{}) bool {
	value, ok := estimate.(int)
	if !ok {
		return false
	} else if value != 0 && value != 1 {
		return false
	} else {
		return true
	}
}
