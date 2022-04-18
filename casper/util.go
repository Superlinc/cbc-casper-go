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

func Sum(set *hashset.Set) uint64 {
	var sum uint64
	for _, v := range set.Values() {
		sum += v.(uint64)
	}
	return sum
}

func Max(set *hashset.Set) uint64 {
	var max uint64
	for _, v := range set.Values() {
		max = MaxUint(max, v.(uint64))
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
