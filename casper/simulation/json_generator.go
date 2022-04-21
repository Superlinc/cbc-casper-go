package simulation

import (
	"container/list"
	"encoding/json"
	"math/rand"
)

type JsonBase struct {
	Protocol string    `json:"protocol"`
	Conf     Config    `json:"config"`
	Exec     Execution `json:"execution"`
}

type Config struct {
	Validators      []float64   `json:"validators"`
	Estimates       interface{} `json:"estimates"`
	SelectOutputs   string      `json:"select_outputs"`
	CreateOutputs   string      `json:"create_outputs"`
	StartOutputs    []int       `json:"start_outputs"`
	GenesisEstimate []int       `json:"genesis_estimate"`
}

type Execution struct {
	MsgPerRound uint64 `json:"msg_per_round"`
	ExeStr      string `json:"exe_str"`
}

func makeBaseObj(protocol string, weights []float64, estimates []interface{}) *JsonBase {
	return &JsonBase{
		Protocol: protocol,
		Conf: Config{
			Validators: weights,
			Estimates:  estimates,
		},
		Exec: Execution{
			MsgPerRound: 1,
		},
	}
}

func GenerateBinaryJsonString(weights []float64, estimates []interface{}) string {
	data := makeBaseObj("binary", weights, estimates)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}

func GenerateIntegerJsonString(weights []float64, estimates []interface{}) string {
	data := makeBaseObj("integer", weights, estimates)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}

func GenerateOrderJsonString(weights []float64, estimates []interface{}) string {
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
	data := makeBaseObj("order", weights, wrap)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}

func GenerateBlockchainJsonString(weights []float64, estimates []interface{}) string {
	data := makeBaseObj("blockchain", weights, estimates)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}

func GenerateConcurrentJsonString(weights []float64, outputs []int, genEst []int, selectName, createName string) string {
	data := makeBaseObj("concurrent", weights, nil)
	if outputs == nil {
		outputs = make([]int, 0, 10)
		for range [10]struct{}{} {
			outputs = append(outputs, rand.Intn(100000))
		}
	}
	data.Conf.StartOutputs = outputs
	if genEst == nil {
		for i := 0; i < len(outputs); i += 2 {
			genEst = append(genEst, outputs[i])
		}
	}
	data.Conf.GenesisEstimate = genEst
	data.Conf.SelectOutputs = selectName
	data.Conf.CreateOutputs = createName
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}
