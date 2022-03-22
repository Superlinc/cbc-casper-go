package integer

import "cbc-casper-go/casper"

type IntegerView struct {
	*casper.View
}

func (v *IntegerView) Estimate() int {
	return getEstimate(v.LatestMessages)
}

func (v IntegerView) UpdateSafeEstimate(valSet *casper.ValidatorSet) {

}
