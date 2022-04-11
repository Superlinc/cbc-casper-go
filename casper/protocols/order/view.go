package order

import (
	"cbc-casper-go/casper"
	"cbc-casper-go/casper/safety_oracles"
	"fmt"
)

type View struct {
	*casper.View
	lastFinalEstimate *casper.Message
}

func NewView() casper.Viewer {
	return &View{
		View:              casper.NewView().(*casper.View),
		lastFinalEstimate: nil,
	}
}

func (v *View) Estimate() interface{} {
	return getEstimate(v.LatestMsg())
}

func (v *View) UpdateSafeEstimates(valSet *casper.ValidatorSet) {
	for _, m := range v.LatestMsg() {
		message := m.(*casper.Message)
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
