package casper

import (
	"github.com/emirpasic/gods/sets/hashset"
	"sort"
)

type ValidatorSet struct {
	validators *hashset.Set
}

func NewValidatorSet(weights []uint64) ValidatorSetor {
	validatorSet := &ValidatorSet{validators: hashset.New()}
	for name, weight := range weights {
		validatorSet.validators.Add(&Validator{
			name:   name,
			weight: weight,
			valSet: validatorSet,
			view:   NewView(nil),
		})
	}
	return validatorSet
}

func (valSet *ValidatorSet) Weight(validators []AbstractValidator) uint64 {
	var sum uint64
	if validators == nil {
		validators = make([]AbstractValidator, 0, valSet.Size())
		for _, v := range valSet.validators.Values() {
			validators = append(validators, v.(AbstractValidator))
		}
	}
	for _, val := range validators {
		if valSet.validators.Contains(val) {
			sum += val.Weight()
		}
	}
	return sum
}

func (valSet *ValidatorSet) Contains(validator AbstractValidator) bool {
	return valSet.validators.Contains(validator)
}

func (valSet *ValidatorSet) SortedByName() []AbstractValidator {
	vals := make([]AbstractValidator, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		val := v.(AbstractValidator)
		vals = append(vals, val)
	}
	sort.SliceStable(vals, func(i, j int) bool {
		return vals[i].Name() < vals[j].Name()
	})
	return vals
}

func (valSet *ValidatorSet) SortedByWeight() []AbstractValidator {
	vals := make([]AbstractValidator, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		val := v.(AbstractValidator)
		vals = append(vals, val)
	}
	sort.SliceStable(vals, func(i, j int) bool {
		return vals[i].Weight() < vals[j].Weight()
	})
	return vals
}

func (valSet *ValidatorSet) GetValByName(name int) AbstractValidator {
	validators := valSet.GetValsByName([]int{name})
	if len(validators) == 0 {
		return nil
	} else {
		return validators[0]
	}
}

func (valSet *ValidatorSet) GetValsByName(names []int) []AbstractValidator {
	validators := make([]AbstractValidator, 0, 4)
	for _, name := range names {
		for _, v := range valSet.validators.Values() {
			validator := v.(AbstractValidator)
			if validator.Name() == name {
				validators = append(validators, validator)
				break
			}
		}
	}
	return validators
}

func (valSet ValidatorSet) Names() []int {
	names := make([]int, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		validator := v.(AbstractValidator)
		names = append(names, validator.Name())
	}
	return names
}

func (valSet ValidatorSet) Weights() []uint64 {
	weights := make([]uint64, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		validator := v.(AbstractValidator)
		weights = append(weights, validator.Weight())
	}
	return weights
}

func (valSet *ValidatorSet) Validators() []AbstractValidator {
	validators := make([]AbstractValidator, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		validator := v.(AbstractValidator)
		validators = append(validators, validator)
	}
	return validators
}

func (valSet *ValidatorSet) Size() uint64 {
	return valSet.Size()
}
