package casper

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"testing"
)

func TestView_JustificationStoreHash(t *testing.T) {
	p := NewProtocol([]uint64{1, 2}, "", 10, 0, 0)
	p.Execute("M-0-A SJ-1-A M-1-B")
	v0 := p.GlobalValidatorSet.GetValByName(0)
	v1 := p.GlobalValidatorSet.GetValByName(1)
	j := v1.Justification()
	if p.Messages["A"].Hash() == p.Messages["B"].Hash() {
		t.Errorf("message error")
	}
	if j[v0] != p.Messages["A"].Hash() {
		t.Errorf("execute error")
	}
	if j[v1] != p.Messages["B"].Hash() {
		t.Errorf("execute error")
	}
}

func TestView_ReceiveJustifiedMessage(t *testing.T) {
	p := NewProtocol([]uint64{1, 2}, "", 10, 0, 0)
	p.Execute("M-0-A M-0-B S-1-B M-1-C")
	v0 := p.GlobalValidatorSet.GetValByName(0)
	v1 := p.GlobalValidatorSet.GetValByName(1)
	j := v1.Justification()
	set := hashset.New()
	for _, v := range j {
		//fmt.Println(k.Name, p.MessageNameFromHash[v])
		set.Add(v)
	}
	if set.Contains(p.Messages["A"]) {
		t.Errorf("error")
	}
	if set.Contains(p.Messages["B"]) {
		t.Errorf("error")
	}
	if j[v1] != p.Messages["C"].Hash() {
		t.Errorf("error")
	}
	p.Execute("SJ-1-B")
	//for hash := range v1.View.PendingMessages {
	//	fmt.Println(p.MessageNameFromHash[hash])
	//}
	j = v1.Justification()
	if j[v0] != p.Messages["B"].Hash() {
		t.Errorf("error")
	}
	if j[v1] != p.Messages["C"].Hash() {
		t.Errorf("error")
	}

}

func TestView_AddJustifiedMessage(t *testing.T) {
	p := NewProtocol([]uint64{1, 2}, "", 10, 0, 0)
	p.Execute("M-0-A M-0-B SJ-1-A")
	v0 := p.GlobalValidatorSet.GetValByName(0)
	v1 := p.GlobalValidatorSet.GetValByName(1)
	j0 := v0.View.JustifiedMessages
	j1 := v1.View.JustifiedMessages
	set0 := hashset.New()
	set1 := hashset.New()
	for _, v := range j0 {
		set0.Add(v)
		fmt.Print()
	}
	fmt.Println()
	for _, v := range j1 {
		set1.Add(v)
	}
	if !set0.Contains(p.Messages["A"]) {
		t.Errorf("error")
	}
	if !set1.Contains(p.Messages["A"]) {
		t.Errorf("error")
	}
	if !set0.Contains(p.Messages["B"]) {
		t.Errorf("error")
	}
	if set1.Contains(p.Messages["B"]) {
		t.Errorf("error")
	}
}

func TestView_MultipleMessage(t *testing.T) {
	p := NewProtocol([]uint64{1, 2}, "", 10, 0, 0)
	p.Execute("M-0-A SJ-1-A M-0-B M-0-C M-0-D M-0-E M-0-F S-1-F")
	v1 := p.GlobalValidatorSet.GetValByName(1)
	pendingSet := hashset.New()
	for _, v := range v1.View.PendingMessages {
		pendingSet.Add(v)
	}
	if !pendingSet.Contains(p.Messages["F"]) {
		t.Errorf("error")
	}
	for _, v := range p.GlobalView.JustifiedMessages {
		v1.ReceiveMessages([]*Message{v})
	}
	justifySet := hashset.New()
	for _, v := range v1.View.JustifiedMessages {
		justifySet.Add(v)
	}
	if !justifySet.Contains(p.Messages["F"]) {
		t.Errorf("error")
	}

}
