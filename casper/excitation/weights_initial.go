package excitation

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// EqualWeights 返回长度为length, 权重都为1的切片
func EqualWeights(length int) (weights []float64) {
	weights = make([]float64, length)
	for i := range weights {
		weights[i] = 1
	}
	return
}

// RandomWeights 返回长度为length, 权重随机分布于[0,10)的切片
func RandomWeights(length int) (weights []float64) {
	weights = make([]float64, length)
	for i := range weights {
		weights[i] = rand.Float64() * 10
	}
	return
}
