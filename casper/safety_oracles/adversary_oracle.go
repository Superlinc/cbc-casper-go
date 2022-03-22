package safety_oracles

import (
	"cbc-casper-go/casper"
	"errors"
)

type AdversaryOracle struct {
	candidateEstimate interface{}
	View              *casper.View
	ValSet            *casper.ValidatorSet
}

func NewAdversaryOracle(estimate interface{}, view *casper.View, set *casper.ValidatorSet) (*AdversaryOracle, error) {
	if estimate == nil {
		return nil, errors.New("candidate estimate can not be null")
	}
	return &AdversaryOracle{
		candidateEstimate: estimate,
		View:              view,
		ValSet:            set,
	}, nil
}

func (o *AdversaryOracle) getMsgAndView() {
	recentMsg := make(map[*casper.Validator]*casper.Message)
	for _, v := range o.ValSet.Validators() {
		if _, ok := o.View.LatestMessages[v]; !ok {
			recentMsg[v] = nil
		}
	}
}
