package binary

import (
	. "cbc-casper-go/casper/simulation"
	"encoding/json"
	"testing"
)

func TestSaveInitialEstimates(t *testing.T) {
	str := GenerateBinaryJsonString([]float64{1, 2, 3, 4, 5}, []interface{}{1, 0, 1, 0, 1})
	var binaryJson JsonBase
	_ = json.Unmarshal([]byte(str), &binaryJson)
	p, err := NewProtocol(str, 1)
	if err != nil {
		t.Errorf("error")
	} else {
		p.Execute("M-0-A")
		v0 := p.ValSet.GetValByName(0)
		if v0.Justification()[v0] != p.Msgs["A"].Hash() {
			t.Errorf("error")
		}
	}

}
