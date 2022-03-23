package integer

import (
	"cbc-casper-go/casper"
	. "cbc-casper-go/casper/simulation"
	"encoding/json"
	"errors"
)

type IntegerProtocol struct {
	*casper.Protocol
}

func NewIntegerProtocol(jsonStr string, display uint64, save interface{}, reportInterval uint64) (*IntegerProtocol, error) {
	parsedJson, err := parseJson(jsonStr)
	if parsedJson == nil || err != nil {
		return nil, err
	}
	protocol := casper.NewProtocol(parsedJson.Conf.Validators,
		parsedJson.Exec.ExeStr,
		parsedJson.Exec.MsgPerRound*reportInterval,
		display,
		save)
	integerProtocol := &IntegerProtocol{
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

func (p IntegerProtocol) SetInitMsg(estimates []int) {
	for _, validator := range p.GlobalValidatorSet.Validators() {
		msg := &Bet{
			&casper.Message{
				Estimate:      estimates[validator.Name],
				Justification: make(map[*casper.Validator]uint64),
				Sender:        validator,
				SeqNum:        0,
				DisplayHeight: 0,
			},
		}
		p.RegisterMessage(msg.Message, casper.GetRandomStr(10))
		validator.InitializeView([]*casper.Message{msg.Message})
	}
}
