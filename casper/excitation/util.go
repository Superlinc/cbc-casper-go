package excitation

import "sort"

// Sum 返回浮点切片的和
func Sum(arr []float64) float64 {
	var sum float64
	for _, value := range arr {
		sum += value
	}
	return sum
}

// RateOfWeightMoreThanThreshold 返回切片中多大比例的总和超过阈值
func RateOfWeightMoreThanThreshold(validators []Validator, threshold float64) (length int, rate float64) {
	sum := TotalWeight(validators)
	sort.SliceStable(validators, func(i, j int) bool {
		return validators[i].Weight() > validators[j].Weight()
	})
	var prefix float64
	for i, validator := range validators {
		prefix += float64(validator.Weight())
		if prefix/sum > threshold {
			rate = float64(i+1) / float64(len(validators))
			return i + 1, rate
		}
	}
	return
}

func TotalWeight(validators []Validator) float64 {
	weight := 0.0
	for _, validator := range validators {
		weight += float64(validator.Weight())
	}
	return weight
}
