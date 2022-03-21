package integer

import "cbc-casper-go/casper"

type IntegerView struct {
	*casper.View
}

func (v IntegerView) Estimate() int {

}
