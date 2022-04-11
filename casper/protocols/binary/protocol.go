package binary

import (
	"cbc-casper-go/casper"
	. "cbc-casper-go/casper/simulation"
	"encoding/json"
	"errors"
)

// Protocol 对一比特数据达成共识的协议
type Protocol struct {
	*casper.Protocol
}

func NewBinaryProtocol(jsonStr string, reportInterval uint64) (*Protocol, error) {
	parsedJson, err := parseJson(jsonStr)
	if parsedJson == nil || err != nil {
		return nil, err
	}
	protocol := casper.NewProtocol(parsedJson.Conf.Validators,
		casper.NewView(),
		nil,
		parsedJson.Exec.MsgPerRound*reportInterval)
	binaryProtocol := &Protocol{protocol}

	return binaryProtocol, nil
}

func parseJson(jsonStr string) (*JsonBase, error) {
	var parsedJson JsonBase
	err := json.Unmarshal([]byte(jsonStr), &parsedJson)
	if err != nil {
		return nil, err
	}
	if len(parsedJson.Conf.Estimates.([]interface{})) != len(parsedJson.Conf.Validators) {
		return nil, errors.New("len(validators) != len(estimates)")
	}
	for _, estimate := range parsedJson.Conf.Estimates.([]interface{}) {
		if !isValidEstimate(int(estimate.(float64))) {
			return nil, errors.New("estimate invalid")
		}
	}
	return &parsedJson, nil
}

func (p *Protocol) SetInitMsg(estimates []int) {
	for _, validator := range p.ValSet.Validators() {
		msg := &Bet{
			casper.NewMessage(estimates[validator.Name()], make(map[*casper.Validator]uint64), validator, 0, 0),
		}
		p.RegisterMessage(msg.Message, casper.GetRandomStr(10))
		validator.InitializeView(casper.NewView(), []casper.Messager{msg.Message})
	}
}
