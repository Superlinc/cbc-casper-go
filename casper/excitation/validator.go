package excitation

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Validator interface {
	GetCoin(award int) Validator
	Weight() int
	Expand() int
}

var (
	curGen   = 0
	interval = 100
)

func NewValidator(age bool) Validator {
	if age {
		return &AgeValidator{coins: make(map[int]int)}
	} else {
		return &NormalValidator{weight: 0}
	}
}

type NormalValidator struct {
	weight int
}

func (n *NormalValidator) GetCoin(award int) Validator {
	n.weight += award
	return n
}

func (n *NormalValidator) Weight() int {
	return n.weight
}

func (n *NormalValidator) Expand() int {
	num := rand.Intn(n.weight)
	n.weight -= num
	return num
}

func UpdateGeneration(generation int) {
	curGen = generation
}

type AgeValidator struct {
	coins map[int]int
}

func (a *AgeValidator) GetCoin(award int) Validator {
	a.coins[curGen] = award
	return a
}

func (a *AgeValidator) Weight() int {
	weight := 0
	for gen, coin := range a.coins {
		weight += ((curGen-gen)/interval + 1) * coin
	}
	return weight
}

func (a *AgeValidator) Expand() int {
	for gen, coin := range a.coins {
		delete(a.coins, gen)
		return coin
	}
	return 0
}
