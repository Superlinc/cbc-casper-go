package order

import (
	"cbc-casper-go/casper/simulation"
	"container/list"
	"fmt"
	"testing"
)

func TestOrderProtocol_SetInitMsg(t *testing.T) {
	l1 := list.New()
	l1.PushBack(11)
	l2 := list.New()
	l2.PushFront(22)
	l3 := list.New()
	l3.PushFront(33)
	l4 := list.New()
	l4.PushBack(44)
	l5 := list.New()
	l5.PushBack(66)
	jsonString := simulation.GenerateOrderJsonString([]uint64{1, 2, 3, 4, 5}, "", []interface{}{l1, l2, l3, l4, l5})
	parsedJson, err := parseJson(jsonString)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if parsedJson.Conf.Estimates.([]*list.List)[0].Len() != 1 {
		fmt.Println(parsedJson.Conf.Estimates.([]*list.List)[0])
		t.Errorf("error: %s", "list length error")
	}
}
