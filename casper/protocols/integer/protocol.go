package integer

import (
	"cbc-casper-go/casper"
	. "cbc-casper-go/casper/simulation"
	"encoding/json"
	"errors"
)

type Protocol struct {
	*casper.Protocol
}

func NewIntegerProtocol(jsonStr string, reportInterval uint64) (*Protocol, error) {
	parsedJson, err := parseJson(jsonStr)
	if parsedJson == nil || err != nil {
		return nil, err
	}
	views := make([]casper.Viewer, 0, len(parsedJson.Conf.Validators))
	for range parsedJson.Conf.Validators {
		views = append(views, NewView())
	}
	protocol := casper.NewProtocol(parsedJson.Conf.Validators,
		NewView(),
		views,
		parsedJson.Exec.MsgPerRound*reportInterval)
	integerProtocol := &Protocol{
		protocol,
	}
	return integerProtocol, nil
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
		validator.InitializeView(NewView(), []casper.Messager{msg.Message})
	}
}
