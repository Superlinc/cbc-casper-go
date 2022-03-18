package casper

import "testing"

func TestValidator_InitializeView(t *testing.T) {
	v := &Validator{
		Name:         0,
		Weight:       1,
		ValidatorSet: nil,
	}
	v.InitializeView(nil)
	if v.View == nil {
		t.Errorf("error")
	}

}
