package binary

import (
	. "cbc-casper-go/casper/simulation"
	"encoding/json"
	"testing"
)

func TestSaveInitialEstimates(t *testing.T) {
	str := GenerateBinaryJsonString([]uint64{1, 2, 3, 4, 5}, "M-0-A", []int{1, 0, 1, 0, 1})
	var binaryJson JsonBase
	_ = json.Unmarshal([]byte(str), &binaryJson)
	p, err := NewBinaryProtocol(str, 0, 0, 1)
	if err != nil {
		t.Errorf("error")
	} else {
		p.Execute("")
		v0 := p.GlobalValidatorSet.GetValByName(0)
		if v0.Justification()[v0] != p.Messages["A"].Hash() {
			t.Errorf("error")
		}
	}

}
