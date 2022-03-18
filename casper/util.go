package casper

import (
	"math/rand"
	"time"
)

func MaxUint(a, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func GetRandomStr(length uint64) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var res []byte
	for i := uint64(0); i < length; i++ {
		res = append(res, str[r.Intn(26)])
	}
	return string(res)
}

func IntSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func UInt64Sum(nums []uint64) uint64 {
	var sum uint64
	for _, num := range nums {
		sum += num
	}
	return sum
}
