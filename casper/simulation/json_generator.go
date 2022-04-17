package simulation

import (
	"container/list"
	"encoding/json"
)

type JsonBase struct {
	Protocol string    `json:"protocol"`
	Conf     Config    `json:"config"`
	Exec     Execution `json:"execution"`
}

type Config struct {
	Validators    []uint64    `json:"validators"`
	Estimates     interface{} `json:"estimates"`
	SelectOutputs string      `json:"select_outputs"`
	CreateOutputs string      `json:"create_outputs"`
}

type Execution struct {
	MsgPerRound uint64 `json:"msg_per_round"`
	ExeStr      string `json:"exe_str"`
}

func makeBaseObj(protocol, exeStr string, weights []uint64, estimates []interface{}) *JsonBase {
	return &JsonBase{
		Protocol: protocol,
		Conf: Config{
			Validators: weights,
			Estimates:  estimates,
		},
		Exec: Execution{
			MsgPerRound: 1,
			ExeStr:      exeStr,
		},
	}
}

func GenerateBinaryJsonString(weights []uint64, exeStr string, estimates []interface{}) string {
	data := makeBaseObj("binary", exeStr, weights, estimates)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}

func GenerateIntegerJsonString(weights []uint64, exeStr string, estimates []interface{}) string {
	data := makeBaseObj("integer", exeStr, weights, estimates)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}

func GenerateOrderJsonString(weights []uint64, exeStr string, estimates []interface{}) string {
	wrap := make([]interface{}, 0, len(estimates))
	for _, estimate := range estimates {
		l := estimate.(*list.List)
		tmp := make([]interface{}, 0, l.Len())
		iter := l.Front()
		for iter != nil {
			tmp = append(tmp, iter.Value)
			iter = iter.Next()
		}
		wrap = append(wrap, tmp)
	}
	data := makeBaseObj("order", exeStr, weights, wrap)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}

func GenerateBlockchainJsonString(weights []uint64, exeStr string, estimates []interface{}) string {
	data := makeBaseObj("blockchain", exeStr, weights, estimates)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}
