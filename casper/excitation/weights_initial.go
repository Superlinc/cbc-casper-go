package excitation

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// EqualWeights 返回长度为length, 权重都为1的切片
func EqualWeights(length int, age bool) []Validator {
	validators := make([]Validator, length)
	for i := 0; i < length; i++ {
		validators[i] = NewValidator(age).GetCoin(1)
	}
	return validators
}

// RandomWeights 返回长度为length, 权重随机分布于[0,10)的切片
func RandomWeights(length int, age bool) []Validator {
	validators := make([]Validator, length)
	for i := 0; i < length; i++ {
		validators[i] = NewValidator(age).GetCoin(rand.Intn(10) + 1)
	}
	return validators
}
