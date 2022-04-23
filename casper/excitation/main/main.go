package main

import (
	. "cbc-casper-go/casper/excitation"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	weights := EqualWeights(1000)
	for range [1000000]struct{}{} {
		rate := rand.Float64()
		rates := GetRateSlice(weights)
		//fmt.Println(rates)
		index := sort.Search(len(rates), func(i int) bool {
			return rates[i] >= rate
		})
		//fmt.Println(rate, index)
		weights[index] += 1
	}
	sort.Float64s(weights)

	fmt.Println(weights)
	index, rate := RateOfWeightMoreThanThreshold(weights, 2.0/3)
	fmt.Println(index, rate)
}

func test() {

}
