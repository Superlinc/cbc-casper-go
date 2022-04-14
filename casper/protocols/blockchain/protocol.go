package blockchain

import (
	"cbc-casper-go/casper"
	. "cbc-casper-go/casper/simulation"
	"encoding/json"
	"errors"
)

type Protocol struct {
	*casper.Protocol
}

func NewBlockchainProtocol(jsonStr string, reportInterval uint64) (*Protocol, error) {
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
	blockchainProtocol := &Protocol{
		protocol,
	}
	blockchainProtocol.RegisterHandler("M", makeBlk)
	blockchainProtocol.SetInitMsg()
	return blockchainProtocol, nil
}

func makeBlk(p *casper.Protocol, validator *casper.Validator, messageName string) {
	newMsg := validator.MakeNewMessage().(*casper.Message)
	newBlk := &Block{
		Message: newMsg,
		height:  1,
	}
	if newMsg.Estimate.(*Block) != nil {
		newBlk.height = newMsg.Estimate.(*Block).height + 1
	}
	validator.ReceiveMessages([]casper.Messager{newBlk})
	p.RegisterMessage(newBlk, messageName)
}

func parseJson(jsonStr string) (*JsonBase, error) {
	// todo 改造为解析block
	var parsedJson JsonBase
	err := json.Unmarshal([]byte(jsonStr), &parsedJson)
	if err != nil {
		return nil, err
	}
	if len(parsedJson.Conf.Estimates.([]interface{})) != len(parsedJson.Conf.Validators) {
		return nil, errors.New("len(validators) != len(estimates)")
	}
	blocks := make([]*Block, len(parsedJson.Conf.Estimates.([]interface{})))
	parsedJson.Conf.Estimates = blocks
	return &parsedJson, nil
}

func (p *Protocol) SetInitMsg() {
	genesis := NewBlock(nil, make(map[*casper.Validator]uint64), p.ValSet.GetValByName(0), 0, 0)
	p.RegisterMessage(genesis, casper.GetRandomStr(10))
	for _, validator := range p.ValSet.Validators() {
		validator.InitializeView(NewView(), []casper.Messager{genesis})
	}
}
