package blockchain

import (
	"cbc-casper-go/casper"
	"testing"
)

func TestView_Children(t *testing.T) {
	p0 := getProtocol([]float64{10, 9})
	p0.Execute("M-0-A S-1-A M-1-B S-0-B")
	children0 := make(map[string][]string)
	children0["A"] = []string{"B"}
	testChildren(t, children0, p0)
	p1 := getProtocol([]float64{10, 9, 8, 7, 6})
	p1.Execute("M-0-A S-1-A S-2-A S-3-A S-4-A M-1-B M-2-C M-3-D M-4-E S-0-B S-0-C S-0-D S-0-E")
	children1 := make(map[string][]string)
	children1["A"] = []string{"B", "C", "D", "E"}
	testChildren(t, children1, p1)
	p2 := getProtocol([]float64{10, 9, 8, 7, 6})
	p2.Execute("M-0-A S-1-A S-2-A M-1-B M-2-C S-0-B S-0-C M-0-D")
	children2 := make(map[string][]string)
	children2["A"] = []string{"B", "C"}
	children2["B"] = []string{"D"}
	testChildren(t, children2, p2)
}

func testChildren(t *testing.T, children map[string][]string, p *Protocol) {
	validator := p.ValSet.GetValByName(0)
	for name, names := range children {
		block := p.Msgs[name]
		for _, child := range validator.View().(*View).children[block.(*Block)] {
			if !casper.StringListContain(names, p.NamesFromHash[child.Hash()]) {
				t.Errorf("error")
			}
		}
	}
}

func TestView_UpdateSafeEstimates(t *testing.T) {
	p := getProtocol([]float64{2, 3})
	p.Execute("M-0-A1 M-1-B1 SJ-0-B1 M-0-A2 M-1-B2 SJ-0-B2 M-0-A3 M-0-A4 SJ-1-A2 M-1-B3")

	hash := p.GlobalView.UpdateSafeEstimates(p.ValSet).Hash()
	name := p.NamesFromHash[hash]
	if name != "B1" {
		t.Errorf("error safety oracle")
	}
}
