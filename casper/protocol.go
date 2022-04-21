package casper

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"regexp"
	"strconv"
	"strings"
)

type Protocol struct {
	ValSet           *ValidatorSet
	GlobalView       Viewer
	unexecuted       string
	executed         string
	messagePerRound  uint64
	messageThisRound uint64
	Msgs             map[string]Messager
	MsgsFromHash     map[uint64]Messager
	NamesFromHash    map[uint64]string
	handlers         map[string]func(*Protocol, *Validator, string)
}

func NewProtocol(weights []float64, view Viewer, views []Viewer, messagePerRound uint64) *Protocol {
	protocol := &Protocol{
		ValSet:           NewValidatorSet(weights, views),
		GlobalView:       view,
		unexecuted:       "",
		executed:         "",
		messagePerRound:  messagePerRound,
		messageThisRound: 0,
		Msgs:             make(map[string]Messager),
		MsgsFromHash:     make(map[uint64]Messager),
		NamesFromHash:    make(map[uint64]string),
		handlers:         make(map[string]func(*Protocol, *Validator, string)),
	}
	protocol.RegisterHandler("M", (*Protocol).makeMsg)
	protocol.RegisterHandler("S", (*Protocol).sendMsg)
	protocol.RegisterHandler("SJ", (*Protocol).sendAndJustify)
	return protocol
}

func (p *Protocol) RegisterHandler(token string, function func(*Protocol, *Validator, string)) {
	if _, ok := p.handlers[token]; ok {
		_ = fmt.Errorf("a function has been registered with this token")
	}
	p.handlers[token] = function
}

func (p *Protocol) RegisterMessage(message Messager, name string) {
	if _, ok := p.Msgs[name]; ok {
		_ = fmt.Errorf("message with %s already exists", name)
	}
	if _, ok := p.MsgsFromHash[message.Hash()]; ok {
		_ = fmt.Errorf("message with %d already exists", message.Hash())
	}
	p.Msgs[name] = message
	p.MsgsFromHash[message.Hash()] = message
	p.NamesFromHash[message.Hash()] = name
	p.GlobalView.AddMessages([]Messager{message})
}

// makeMsg 使用该验证器生成消息
func (p *Protocol) makeMsg(validator *Validator, messageName string) {
	msg := validator.MakeNewMessage()
	validator.ReceiveMessages([]Messager{msg})
	p.RegisterMessage(msg, messageName)
}

// sendMsg 给该验证器传递一个消息
func (p *Protocol) sendMsg(validator *Validator, messageName string) {
	message := p.Msgs[messageName]
	validator.ReceiveMessages([]Messager{message})
}

// sendAndJustify 传递消息以及其依赖
func (p *Protocol) sendAndJustify(validator *Validator, messageName string) {
	message := p.Msgs[messageName]
	messageToSend := p.msgNeededJustify(message, validator)
	//for _, msg := range messageToSend {
	//	fmt.Println(p.NamesFromHash[msg.Hash()])
	//	for _, hash := range msg.Justification {
	//		fmt.Println(p.NamesFromHash[hash])
	//	}
	//}
	validator.ReceiveMessages(messageToSend)
}

// MessagesNeededToJustify 返回消息及未验证依赖
func (p *Protocol) msgNeededJustify(message Messager, validator *Validator) []Messager {
	messageNeeded := hashset.New(message)
	messageHashes := hashset.New(message.Hash())
	for messageHashes.Size() != 0 {
		nextHashes := hashset.New()
		for _, m := range messageHashes.Values() {
			messageHash := m.(uint64)
			message = p.MsgsFromHash[messageHash]
			messageNeeded.Add(message)
			for _, hash := range message.Justification() {
				if _, ok := validator.View().JustifiedMsg()[hash]; !ok {
					nextHashes.Add(hash)
				}
			}
		}
		messageHashes = nextHashes
	}
	result := make([]Messager, 0, messageNeeded.Size())
	for _, v := range messageNeeded.Values() {
		result = append(result, v.(Messager))
	}
	return result
}

func (p *Protocol) Execute(additionalStr string) {
	if additionalStr != "" {
		p.unexecuted += additionalStr
	}
	for _, token := range strings.Split(p.unexecuted, " ") {
		comm, vali, name, _ := parseToken(token)
		validator := p.ValSet.GetValByName(vali)
		p.handlers[comm](p, validator, name)
		if comm == "M" {
			p.messageThisRound += 1
			if p.messageThisRound%p.messagePerRound == 0 {
				// todo
			}
		}
	}
	p.executed += p.unexecuted
	p.unexecuted = ""
}

func parseToken(token string) (comm string, vali int, name string, data string) {
	TokenPattern := "([A-Za-z]*)([-]*)([0-9]*)([-]*)([A-Za-z0-9]*)([-]*)([(){A-Za-z0-9,}]*)"
	matched, err := regexp.MatchString(TokenPattern, token)
	if matched || (err != nil) {
		_ = fmt.Errorf("pattern unmatch")
	}
	result := strings.Split(token, "-")
	comm = result[0]
	vali, _ = strconv.Atoi(result[1])
	name = result[2]
	if len(result) > 3 {
		data = result[3]
	}
	return
}
