package excitation

// Sum 返回浮点切片的和
func Sum(arr []float64) float64 {
	var sum float64
	for _, value := range arr {
		sum += value
	}
	return sum
}

func TotalWeight(validators []Validator) float64 {
	weight := 0.0
	for _, validator := range validators {
		weight += float64(validator.Weight())
	}
	return weight
}
