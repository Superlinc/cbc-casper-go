package integer

import (
	"cbc-casper-go/casper"
	"cbc-casper-go/casper/safety_oracles"
	"fmt"
)

type IntegerView struct {
	*casper.View
	lastFinalEstimate *casper.Message
}

func (v *IntegerView) Estimate() int {
	return getEstimate(v.LatestMessages)
}

func (v *IntegerView) UpdateSafeEstimate(valSet *casper.ValidatorSet) {
	for _, message := range v.LatestMessages {
		bet := &Bet{message}
		oracle, err := safety_oracles.NewCliqueOracle(bet, v.View, valSet)
		if err != nil {
			continue
		}
		faultTolerance, _ := oracle.CheckEstimateSafety()
		if faultTolerance > 0 {
			if v.lastFinalEstimate != nil {
				bet := &Bet{message}
				if ok, _ := bet.ConflictWith(message); ok {
					_ = fmt.Errorf("error")
				}
			}
			v.lastFinalEstimate = message
			break
		}
	}
}
