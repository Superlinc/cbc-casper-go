package simulation

import (
	"encoding/json"
)

type JsonBase struct {
	Protocol string    `json:"protocol"`
	Conf     Config    `json:"config"`
	Exec     Execution `json:"execution"`
}

type Config struct {
	Validators []uint64    `json:"validators"`
	Estimates  interface{} `json:"estimates,[]int"`
}

type Execution struct {
	MsgPerRound uint64 `json:"msg_per_round"`
	ExeStr      string `json:"exe_str"`
}

func makeBaseObj(protocol, exeStr string, weights []uint64, estimates []int) *JsonBase {
	return &JsonBase{
		Protocol: "binary",
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

func GenerateBinaryJsonString(weights []uint64, exeStr string, estimates []int) string {
	data := makeBaseObj("binary", exeStr, weights, estimates)
	var str string
	if bs, err := json.Marshal(data); err == nil {
		str = string(bs)
		// fmt.Println("generate successfully")
	}
	return str
}
