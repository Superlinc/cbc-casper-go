package casper

import (
	"testing"
)

func TestNewValidatorSet(t *testing.T) {
	weights := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vs := NewValidatorSet(weights, nil)
	valSet := vs
	if valSet.validators.Size() != len(weights) {
		t.Errorf("error")
	}
}

func TestValidatorSet_Weight(t *testing.T) {
	weights := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vs := NewValidatorSet(weights, nil)
	if vs.Weight(vs.Validators()) != UInt64Sum(weights) {
		t.Errorf("error")
	}
}

func TestValidatorSet_GetValByName(t *testing.T) {
	weights := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vs := NewValidatorSet(weights, nil)
	for _, validator := range vs.Validators() {
		if validator != vs.GetValByName(validator.Name()) {
			t.Errorf("error")
		}
	}

}
