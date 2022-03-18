package casper

import (
	"testing"
)

func TestNewValidatorSet(t *testing.T) {
	weights := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vs := NewValidatorSet(weights)
	if vs.validators.Size() != len(weights) {
		t.Errorf("error")
	}
}

func TestValidatorSet_Weight(t *testing.T) {
	weights := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vs := NewValidatorSet(weights)
	if vs.Weight(vs.Validators()) != UInt64Sum(weights) {
		t.Errorf("error")
	}
}

func TestValidatorSet_GetValByName(t *testing.T) {
	weights := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vs := NewValidatorSet(weights)
	for _, validator := range vs.Validators() {
		if validator != vs.GetValByName(validator.Name) {
			t.Errorf("error")
		}
	}

}
