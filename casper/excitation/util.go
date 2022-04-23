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
func RateOfWeightMoreThanThreshold(weights []float64, threshold float64) (length int, rate float64) {
	sum := Sum(weights)
	sort.Sort(sort.Reverse(sort.Float64Slice(weights)))
	var prefix float64
	for i, weight := range weights {
		prefix += weight
		if prefix/sum > threshold {
			rate = float64(i+1) / float64(len(weights))
			return i + 1, rate
		}
	}
	return
}
