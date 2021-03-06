package casper

import (
	"github.com/emirpasic/gods/sets/hashset"
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

func MaxFloat(a, b float64) float64 {
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

func UInt64SliceSum(nums []uint64) uint64 {
	var sum uint64
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Float64SliceSum(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

func UInt64SetSum(set *hashset.Set) uint64 {
	var sum uint64
	for _, v := range set.Values() {
		sum += v.(uint64)
	}
	return sum
}

func Float64SetSum(set *hashset.Set) float64 {
	var sum float64
	for _, v := range set.Values() {
		sum += v.(float64)
	}
	return sum
}

func Max(set *hashset.Set) float64 {
	var max float64
	for _, v := range set.Values() {
		max = MaxFloat(max, v.(float64))
	}
	return max
}

func Intersection(a, b *hashset.Set) *hashset.Set {
	result := hashset.New()
	for _, value := range a.Values() {
		if b.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

func StringListContain(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
