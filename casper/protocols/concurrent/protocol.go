package concurrent

import (
	"cbc-casper-go/casper"
	. "cbc-casper-go/casper/simulation"
	"encoding/json"
	"errors"
	"github.com/emirpasic/gods/sets/hashset"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Protocol struct {
	*casper.Protocol
	selectRules map[string]selectFunc
	createRules map[string]createFunc
}

func NewProtocol(jsonStr string, reportInterval uint64) (*Protocol, error) {
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

	concurrentProtocol := &Protocol{
		Protocol:    protocol,
		selectRules: make(map[string]selectFunc),
		createRules: make(map[string]createFunc),
	}
	concurrentProtocol.selectRules["random"] = selectRandomOutputs
	concurrentProtocol.selectRules["all"] = selectAllOutputs
	concurrentProtocol.createRules["random"] = createRandomOutputs
	concurrentProtocol.createRules["all"] = createAllIncrementedOutputs
	concurrentProtocol.initRewriteRule(parsedJson.Conf.SelectOutputs, parsedJson.Conf.CreateOutputs)
	concurrentProtocol.SetInitMsg(parsedJson.Conf.Estimates, parsedJson.Conf.CreateOutputs)
	return concurrentProtocol, nil
}

func (p *Protocol) initRewriteRule(selectName, createName string) {
	selectOutputs := p.selectRules[selectName]
	createOutputs := p.createRules[createName]
	for _, validator := range p.ValSet.Validators() {
		validator.View().(*View).setRewriteRules(selectOutputs, createOutputs)
	}
	p.GlobalView.(*View).setRewriteRules(selectOutputs, createOutputs)
}

func (p *Protocol) SetInitMsg(genenis interface{}, createName string) {
	validator := p.ValSet.GetValByName(0)
	blocks := make([]*Block, 0)
	inputs := hashset.New()
	for _, v := range genenis.([]float64) {
		inputs.Add(int(v))
	}
	outputs := p.createRules[createName](inputs, inputs.Size())
	estimate := make(map[string]interface{})
	estimate["blocks"] = blocks
	estimate["inputs"] = inputs
	estimate["outputs"] = outputs
	initMsg := &Block{
		casper.NewMessage(estimate, make(map[*casper.Validator]uint64), validator, 0, 0),
	}
	p.RegisterMessage(initMsg, casper.GetRandomStr(10))
	for _, validator := range p.ValSet.Validators() {
		validator.InitializeView(NewView(), []casper.Messager{initMsg})
	}
}

func selectRandomOutputs(avaiOutputs *hashset.Set, sources map[interface{}]*Block) *hashset.Set {
	num := rand.Intn(avaiOutputs.Size()) + 1
	arr := avaiOutputs.Values()
	set := hashset.New()
	for num > 0 {
		index := rand.Intn(avaiOutputs.Size())
		if !set.Contains(arr[index]) {
			set.Add(arr[index])
			num--
		}
	}
	return set
}

func createRandomOutputs(oldOutputs *hashset.Set, num int) *hashset.Set {
	set := hashset.New()
	for i := 0; i < num; i++ {
		for {
			value := rand.Intn(1000000000)
			if !set.Contains(value) {
				set.Add(value)
				break
			}
		}
	}
	return set
}

func selectAllOutputs(avaiOutputs *hashset.Set, sources map[interface{}]*Block) *hashset.Set {
	set := hashset.New()
	set.Add(avaiOutputs.Values()...)
	return set
}

func createAllIncrementedOutputs(oldOutputs *hashset.Set, num int) *hashset.Set {
	set := hashset.New()
	for _, value := range oldOutputs.Values() {
		set.Add(value.(int) + 1)
	}
	return set
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
