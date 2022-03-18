package binary

import (
	"cbc-casper-go/casper"
	"encoding/json"
	"errors"
)

// BinaryProtocol 对一比特数据达成共识的协议
type BinaryProtocol struct {
	casper.Protocol
}

type BinaryJson struct {
	protocol string
	config   *struct {
		validators      []uint64
		initialEstimate []int
	}
	execution *struct {
		executionString string
		msgPerRound     uint64
	}
}

func NewBinaryProtocol(jsonStr string, display uint64, save interface{}, reportInterval uint64) (*BinaryProtocol, error) {
	parsedJson, err := parseJson(jsonStr)
	if parsedJson == nil || err != nil {
		return nil, err
	}
	protocol := casper.NewProtocol(parsedJson.config.validators,
		parsedJson.execution.executionString,
		parsedJson.execution.msgPerRound*reportInterval,
		display,
		save)
	binaryProtocol := &BinaryProtocol{*protocol}

	return binaryProtocol, nil
}

func parseJson(jsonStr string) (*BinaryJson, error) {
	var parsedJson BinaryJson
	err := json.Unmarshal([]byte(jsonStr), &parsedJson)
	if err != nil {
		return nil, err
	}
	if len(parsedJson.config.initialEstimate) != len(parsedJson.config.validators) {
		return nil, errors.New("len(validators) != len(estimates)")
	}
	for _, estimate := range parsedJson.config.initialEstimate {
		if !isValidEstimate(estimate) {
			return nil, errors.New("estimate invalid")
		}
	}
	return &parsedJson, nil
}

func (p *BinaryProtocol) SetInitMsg(estimates []int) {
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
