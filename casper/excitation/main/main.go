package main

import (
	. "cbc-casper-go/casper/excitation"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type GetPara func([]Validator) float64

func main() {
	CalAvg()
}

func simulation(lenVal, generation, interval int, f GetPara) []float64 {
	UpdateGeneration(0)
	validators := EqualWeights(lenVal, true)
	curve := make([]float64, 0, generation/interval)
	award := 1.0
	inflationRate := 1.0
	inflationInterval := 1000
	tradeInterval := 10
	for i := 0; i < generation; i++ {
		UpdateGeneration(i)
		rate := rand.Float64()
		rates := GetRateSliceWithLimit(validators)
		//fmt.Println(rates)
		index := sort.Search(len(rates), func(i int) bool {
			return rates[i] >= rate
		})
		//fmt.Println(rate, index)
		validators[index].GetCoin(int(award))
		if i%interval == 0 {
			rate := f(validators)
			curve = append(curve, rate)
		}
		if i%inflationInterval == 0 {
			award *= inflationRate
		}
		if i%tradeInterval == 0 {
			x := rand.Intn(lenVal)
			y := rand.Intn(lenVal)
			num := validators[x].Expand()
			validators[y].GetCoin(num)
		}
	}
	return curve
}

func CalAvg() {
	avg := 0.0
	times := 50
	for i := 0; i < times; i++ {
		curve := simulation(1000, 10000, 100, Gini)
		avg += curve[len(curve)-1]
		fmt.Println(i, curve[len(curve)-1])
	}
	avg /= float64(times)
	fmt.Println(avg)
}

func SaveCSV() {
	curve := simulation(1000, 10000, 100, Gini)
	fmt.Println(curve)
	file, _ := os.Create("csv/u-gini有币龄有交易.csv")
	defer file.Close()
	interval := 100
	writer := csv.NewWriter(file)
	for i, v := range curve {
		writer.Write([]string{strconv.Itoa(i * interval), strconv.FormatFloat(v, 'f', -1, 64)})
	}
	writer.Flush()

}

func LorenzCurve() {
	lenVal := 1000
	generation := 10000
	UpdateGeneration(0)
	validators := EqualWeights(lenVal, false)

	award := 1.0
	inflationRate := 1.0
	inflationInterval := 1000
	tradeInterval := 10000
	for i := 0; i < generation; i++ {
		UpdateGeneration(i)
		rate := rand.Float64()
		rates := GetRateSlice(validators)
		//fmt.Println(rates)
		index := sort.Search(len(rates), func(i int) bool {
			return rates[i] >= rate
		})
		//fmt.Println(rate, index)
		validators[index].GetCoin(int(award))
		if i%inflationInterval == 0 {
			award *= inflationRate
		}
		if i%tradeInterval == 0 {
			x := rand.Intn(lenVal)
			y := rand.Intn(lenVal)
			num := validators[x].Expand()
			validators[y].GetCoin(num)
		}
	}
	XValues := make([]float64, len(validators)+1)
	YValues := make([]float64, len(validators)+1)

	totalWeight := TotalWeight(validators)
	weights := make([]float64, len(validators))
	for i, validator := range validators {
		weights[i] = float64(validator.Weight())
	}
	sort.Float64s(weights)
	XValues[0] = 0.0
	YValues[0] = 0.0
	for i, weight := range weights {
		YValues[i+1] = YValues[i] + weight/totalWeight
		XValues[i+1] = float64(i+1) / float64(len(weights))
	}

	file, _ := os.Create("csv/Lorenz.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	for i := range XValues {
		writer.Write([]string{strconv.FormatFloat(XValues[i], 'f', -1, 64), strconv.FormatFloat(YValues[i], 'f', -1, 64)})
	}
	writer.Flush()
}
