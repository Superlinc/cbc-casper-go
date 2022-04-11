package casper

import "testing"

func TestValidator_InitializeView(t *testing.T) {
	v := &Validator{
		name:   0,
		weight: 1,
		valSet: nil,
	}
	v.InitializeView(NewView(), nil)
	if v.view == nil {
		t.Errorf("error")
	}

}
