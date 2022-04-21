package blockchain

import (
	"fmt"
	"testing"
)

func TestSingleValidator(t *testing.T) {
	weights := []float64{10, 9, 8, 7, 6}
	p := getProtocol(weights)
	p.Execute("M-0-A M-0-B M-0-C M-0-D M-0-E M-0-F M-0-G M-0-H")
	estimate := p.ValSet.GetValByName(0).Estimate()
	if estimate != p.Msgs["H"] {
		t.Errorf("error estimate")
	}
}

func TestTwoValidator(t *testing.T) {
	weights := []float64{10, 9, 8, 7, 6}
	p := getProtocol(weights)
	p.Execute("M-0-A SJ-1-A M-1-B SJ-0-B M-0-C SJ-1-C M-1-D SJ-0-D")
	estimate0 := p.ValSet.GetValByName(0).Estimate()
	estimate1 := p.ValSet.GetValByName(1).Estimate()
	if estimate0 != p.Msgs["D"] {
		t.Errorf("error estimate")
	}
	if estimate1 != p.Msgs["D"] {
		t.Errorf("error estimate")
	}
}

func TestZeroWeightValidator(t *testing.T) {
	weights := []float64{5, 0}
	p := getProtocol(weights)
	p.Execute("M-0-A SJ-1-A M-1-B SJ-0-B")
	for _, validator := range p.ValSet.GetValsByName([]int{0, 1}) {
		if validator.Estimate() != p.Msgs["A"] {
			t.Errorf("error")
		}
	}
}

func TestZeroWeightBlock(t *testing.T) {
	weights := []float64{10, 9, 8, 0.5}
	p := getProtocol(weights)
	p.Execute("M-0-A1 M-0-A2 M-1-B1 M-1-B2 SJ-3-B2 M-3-D1 SJ-3-A2 M-3-D2 SJ-2-B1 M-2-C1 SJ-1-D1 SJ-1-D2 SJ-1-C1")
	validator := p.ValSet.GetValByName(1)
	if validator.Estimate() != p.Msgs["B2"] {
		t.Errorf("error")
		fmt.Println(p.NamesFromHash[validator.Estimate().(*Block).Hash()])
	}
}

func TestReverse(t *testing.T) {
	weights := []float64{5, 6, 7, 8}
	p := getProtocol(weights)
	p.Execute("M-0-A SJ-1-A M-1-B SJ-0-B M-0-C SJ-1-C M-1-D SJ-0-D M-1-E SJ-0-E SJ-2-E SJ-3-A SJ-3-B SJ-3-C SJ-3-D SJ-3-E")
	validator := p.ValSet.GetValByName(3)
	if validator.Estimate() != p.Msgs["E"] {
		t.Errorf("error")
		fmt.Println(p.NamesFromHash[validator.Estimate().(*Block).Hash()])
	}
}
