package excitation

// GetRateSlice 生成weights切片的比例的前缀和的切片
func GetRateSlice(weights []float64) []float64 {
	sum := Sum(weights)
	rates := make([]float64, len(weights))
	for i, weight := range weights {
		if i == 0 {
			rates[0] = weight / sum
		} else {
			rates[i] = rates[i-1] + weight/sum
		}

	}
	return rates
}

func GetRateSliceWithLimit(weights []float64) []float64 {
	sum := Sum(weights)
	limitedWights := make([]float64, len(weights))
	limit := sum * 0.01
	for i, weight := range weights {
		if weight > limit {
			limitedWights[i] = limit
		} else {
			limitedWights[i] = weight
		}
	}
	limitedSum := Sum(limitedWights)
	rates := make([]float64, len(weights))
	for i, weight := range limitedWights {
		if i == 0 {
			rates[0] = weight / limitedSum
		} else {
			rates[i] = rates[i-1] + weight/limitedSum
		}
	}
	return rates
}
