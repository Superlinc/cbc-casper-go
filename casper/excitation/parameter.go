package excitation

import "sort"

const threshold = 1.0 / 3

func Gini(validators []Validator) float64 {
	totalWeight := TotalWeight(validators)
	weights := make([]float64, len(validators))
	for i, validator := range validators {
		weights[i] = float64(validator.Weight())
	}
	sort.Float64s(weights)
	f := make([]float64, len(validators)+1)
	f[0] = 0
	for i, weight := range weights {
		f[i+1] = f[i] + weight/totalWeight
	}
	gini := 0.0
	length := len(f)
	for i := 0; i < length-1; i++ {
		gini += float64(i)/float64(length-1)*f[i+1] - float64(i+1)/float64(length-1)*f[i]
	}
	return gini
}

// RateOfWeightMoreThanThreshold 返回切片中多大比例的总和超过阈值
func RateOfWeightMoreThanThreshold(validators []Validator) (rate float64) {
	sum := TotalWeight(validators)
	sort.SliceStable(validators, func(i, j int) bool {
		return validators[i].Weight() > validators[j].Weight()
	})
	var prefix float64
	for i, validator := range validators {
		prefix += float64(validator.Weight())
		if prefix/sum > threshold {
			rate = float64(i+1) / float64(len(validators))
			return
		}
	}
	return
}
