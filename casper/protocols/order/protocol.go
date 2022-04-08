package order

import (
	"cbc-casper-go/casper"
	. "cbc-casper-go/casper/simulation"
	"container/list"
	"encoding/json"
	"errors"
)

type Protocol struct {
	*casper.Protocol
}

func NewOrderProtocol(jsonStr string, reportInterval uint64) (*Protocol, error) {
	parsedJson, err := parseJson(jsonStr)
	if parsedJson == nil || err != nil {
		return nil, err
	}
	protocol := casper.NewProtocol(parsedJson.Conf.Validators,
		parsedJson.Exec.ExeStr,
		parsedJson.Exec.MsgPerRound*reportInterval)
	orderProtocol := &Protocol{
		protocol,
	}
	return orderProtocol, nil
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
	estimates := make([]*list.List, 0, len(parsedJson.Conf.Estimates.([]interface{})))
	for _, e := range parsedJson.Conf.Estimates.([]interface{}) {
		estimate := list.New()
		for _, v := range e.([]interface{}) {
			estimate.PushBack(v)
		}
		if !isValidEstimate(estimate) {
			return nil, errors.New("estimate invalid")
		}
		estimates = append(estimates, estimate)
	}
	parsedJson.Conf.Estimates = estimates
	return &parsedJson, nil
}

func (p *Protocol) SetInitMsg(estimates []*list.List) {
	for _, validator := range p.ValSet.Validators() {
		msg := &Bet{
			casper.NewMessage(estimates[validator.Name()], make(map[*casper.Validator]uint64), validator, 0, 0),
		}
		p.RegisterMessage(msg.Message, casper.GetRandomStr(10))
		validator.InitializeView([]casper.Messager{msg.Message})
	}
}
