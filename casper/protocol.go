package casper

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"regexp"
	"strconv"
	"strings"
)

type Protocol struct {
	ValSet           ValidatorSetor
	view             Viewer
	unexecuted       string
	executed         string
	messagePerRound  uint64
	messageThisRound uint64
	Msgs             map[string]Messager
	msgsFromHash     map[uint64]Messager
	namesFromHash    map[uint64]string
	handlers         map[string]func(*Protocol, AbstractValidator, string)
}

func NewProtocol(weights []uint64, execution string, messagePerRound uint64) *Protocol {
	protocol := &Protocol{
		ValSet:           NewValidatorSet(weights),
		view:             NewView(nil),
		unexecuted:       execution,
		executed:         "",
		messagePerRound:  messagePerRound,
		messageThisRound: 0,
		Msgs:             make(map[string]Messager),
		msgsFromHash:     make(map[uint64]Messager),
		namesFromHash:    make(map[uint64]string),
		handlers:         make(map[string]func(*Protocol, AbstractValidator, string)),
	}
	protocol.RegisterHandler("M", (*Protocol).MakeMessage)
	protocol.RegisterHandler("S", (*Protocol).SendMessage)
	protocol.RegisterHandler("SJ", (*Protocol).SendAndJustify)
	return protocol
}

func (p *Protocol) RegisterHandler(token string, function func(*Protocol, AbstractValidator, string)) {
	if _, ok := p.handlers[token]; ok {
		_ = fmt.Errorf("a function has been registered with this token")
	}
	p.handlers[token] = function
}

func (p *Protocol) RegisterMessage(message Messager, name string) {
	if _, ok := p.Msgs[name]; ok {
		_ = fmt.Errorf("message with %s already exists", name)
	}
	if _, ok := p.msgsFromHash[message.Hash()]; ok {
		_ = fmt.Errorf("message with %d already exists", message.Hash())
	}
	p.Msgs[name] = message
	p.msgsFromHash[message.Hash()] = message
	p.namesFromHash[message.Hash()] = name
	p.view.AddMessages([]Messager{message})
}

// MakeMessage 使用该验证器生成消息
func (p *Protocol) MakeMessage(validator AbstractValidator, messageName string) {
	newMessage := validator.MakeNewMessage()
	p.RegisterMessage(newMessage, messageName)
}

// SendMessage 给该验证器传递一个消息
func (p *Protocol) SendMessage(validator AbstractValidator, messageName string) {
	message := p.Msgs[messageName]
	validator.ReceiveMessages([]Messager{message})
}

// SendAndJustify 传递消息以及其依赖
func (p *Protocol) SendAndJustify(validator AbstractValidator, messageName string) {
	message := p.Msgs[messageName]
	messageToSend := p.MessagesNeededToJustify(message, validator)
	//for _, msg := range messageToSend {
	//	fmt.Println(p.namesFromHash[msg.Hash()])
	//	for _, hash := range msg.Justification {
	//		fmt.Println(p.namesFromHash[hash])
	//	}
	//}
	validator.ReceiveMessages(messageToSend)
}

// MessagesNeededToJustify 返回消息及未验证依赖
func (p *Protocol) MessagesNeededToJustify(message Messager, validator AbstractValidator) []Messager {
	messageNeeded := hashset.New(message)
	messageHashes := hashset.New(message.Hash())
	for messageHashes.Size() != 0 {
		nextHashes := hashset.New()
		for _, m := range messageHashes.Values() {
			messageHash := m.(uint64)
			message = p.msgsFromHash[messageHash]
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
