package blockchain

import "testing"

func TestProtocol(t *testing.T) {
	p := getProtocol(nil)
	if len(p.GlobalView.JustifiedMsg()) != 1 {
		t.Errorf("error")
	}
	for _, validator := range p.ValSet.Validators() {
		if len(validator.Justification()) != 1 {
			t.Errorf("error")
		}
	}
}
