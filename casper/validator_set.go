package casper

import (
	"github.com/emirpasic/gods/sets/hashset"
	"sort"
)

type ValidatorSet struct {
	validators *hashset.Set
}

func NewValidatorSet(weights []float64, views []Viewer) *ValidatorSet {
	validatorSet := &ValidatorSet{validators: hashset.New()}
	if views == nil {
		views = make([]Viewer, 0, len(weights))
		for range weights {
			views = append(views, NewView())
		}
	}
	for name, weight := range weights {
		validatorSet.validators.Add(&Validator{
			name:   name,
			weight: weight,
			valSet: validatorSet,
			view:   views[name],
		})
	}
	return validatorSet
}

func (valSet *ValidatorSet) Weight(validators []*Validator) float64 {
	var sum float64
	if validators == nil {
		validators = make([]*Validator, 0, valSet.Size())
		for _, v := range valSet.validators.Values() {
			validators = append(validators, v.(*Validator))
		}
	}
	for _, val := range validators {
		if valSet.validators.Contains(val) {
			sum += val.Weight()
		}
	}
	return sum
}

func (valSet *ValidatorSet) Contains(validator *Validator) bool {
	return valSet.validators.Contains(validator)
}

func (valSet *ValidatorSet) SortedByName() []*Validator {
	vals := make([]*Validator, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		val := v.(*Validator)
		vals = append(vals, val)
	}
	sort.SliceStable(vals, func(i, j int) bool {
		return vals[i].Name() < vals[j].Name()
	})
	return vals
}

func (valSet *ValidatorSet) SortedByWeight() []*Validator {
	vals := make([]*Validator, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		val := v.(*Validator)
		vals = append(vals, val)
	}
	sort.SliceStable(vals, func(i, j int) bool {
		return vals[i].Weight() < vals[j].Weight()
	})
	return vals
}

func (valSet *ValidatorSet) GetValByName(name int) *Validator {
	validators := valSet.GetValsByName([]int{name})
	if len(validators) == 0 {
		return nil
	} else {
		return validators[0]
	}
}

func (valSet *ValidatorSet) GetValsByName(names []int) []*Validator {
	validators := make([]*Validator, 0, 4)
	for _, name := range names {
		for _, v := range valSet.validators.Values() {
			validator := v.(*Validator)
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
		validator := v.(*Validator)
		names = append(names, validator.Name())
	}
	return names
}

func (valSet ValidatorSet) Weights() []float64 {
	weights := make([]float64, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		validator := v.(*Validator)
		weights = append(weights, validator.Weight())
	}
	return weights
}

func (valSet *ValidatorSet) Validators() []*Validator {
	validators := make([]*Validator, 0, valSet.validators.Size())
	for _, v := range valSet.validators.Values() {
		validator := v.(*Validator)
		validators = append(validators, validator)
	}
	return validators
}

func (valSet *ValidatorSet) Size() int {
	return valSet.validators.Size()
}
