package excitation

// GetRateSlice 生成weights切片的比例的前缀和的切片
func GetRateSlice(validators []Validator) []float64 {
	sum := TotalWeight(validators)
	rates := make([]float64, len(validators))
	for i, validator := range validators {
		if i == 0 {
			rates[0] = float64(validator.Weight()) / sum
		} else {
			rates[i] = rates[i-1] + float64(validator.Weight())/sum
		}

	}
	return rates
}

func GetRateSliceWithLimit(validators []Validator) []float64 {
	sum := TotalWeight(validators)
	limitedWights := make([]float64, len(validators))
	limit := sum * 0.005
	for i, validator := range validators {
		if float64(validator.Weight()) > limit {
			limitedWights[i] = limit
		} else {
			limitedWights[i] = float64(validator.Weight())
		}
	}
	limitedSum := Sum(limitedWights)
	rates := make([]float64, len(validators))
	for i, weight := range limitedWights {
		if i == 0 {
			rates[0] = weight / limitedSum
		} else {
			rates[i] = rates[i-1] + weight/limitedSum
		}
	}
	return rates
}
