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

func main() {
	SaveCSV()
}

func simulation(lenVal, generation, interval int) []float64 {
	UpdateGeneration(0)
	validators := EqualWeights(lenVal, true)
	curve := make([]float64, 0, generation/interval)
	award := 1.0
	inflationRate := 1.0
	inflationInterval := 1000
	tradeInterval := 10000
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
			_, rate := RateOfWeightMoreThanThreshold(validators, 1.0/3)
			curve = append(curve, rate)
		}
		if i%inflationInterval == 0 {
			award *= inflationRate
		}
		if i%tradeInterval == 0 {
			x := rand.Intn(lenVal)
			y := rand.Intn(lenVal)
			num := validators[x].Remove()
			validators[y].GetCoin(num)
		}
	}
	return curve
}

func calAvg() {
	avg := 0.0
	times := 50
	for i := 0; i < times; i++ {
		curve := simulation(1000, 10000, 100)
		avg += curve[len(curve)-1]
		fmt.Println(i, curve[len(curve)-1])
	}
	avg /= float64(times)
	fmt.Println(avg)
}

func SaveCSV() {
	curve := simulation(1000, 10000, 100)
	fmt.Println(curve)
	file, _ := os.Create("优化有币龄无交易.csv")
	defer file.Close()
	interval := 100
	writer := csv.NewWriter(file)
	for i, v := range curve {
		writer.Write([]string{strconv.Itoa(i * interval), strconv.FormatFloat(v, 'E', -1, 64)})
	}
	writer.Flush()

}
