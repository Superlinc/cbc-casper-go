package integer

import (
	"cbc-casper-go/casper"
	"errors"
	"fmt"
)

type Bet struct {
	*casper.Message
}

func (b *Bet) ConflictWith(m interface{}) (bool, error) {
	message := m.(*casper.Message)
	if !isValidEstimate(message.Estimate) {
		_ = fmt.Errorf("estimate should be integer")
		return true, errors.New("message estimate error")
	}
	return b.Estimate != message.Estimate, nil
}

func isValidEstimate(estimate interface{}) bool {
	_, ok := estimate.(int)
	return ok
}
