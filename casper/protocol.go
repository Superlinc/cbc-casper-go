package casper

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"regexp"
	"strconv"
	"strings"
)

type Protocol struct {
	GlobalValidatorSet  *ValidatorSet
	GlobalView          *View
	unexecuted          string
	executed            string
	messagePerRound     uint64
	messageThisRound    uint64
	Messages            map[string]*Message
	MessageFromHash     map[uint64]*Message
	MessageNameFromHash map[uint64]string
	handlers            map[string]func(*Protocol, *Validator, string)
}

func NewProtocol(weights []uint64, execution string, messagePerRound uint64, display, save interface{}) *Protocol {
	protocol := &Protocol{
		GlobalValidatorSet:  NewValidatorSet(weights),
		GlobalView:          NewView(nil),
		unexecuted:          execution,
		executed:            "",
		messagePerRound:     messagePerRound,
		messageThisRound:    0,
		Messages:            make(map[string]*Message),
		MessageFromHash:     make(map[uint64]*Message),
		MessageNameFromHash: make(map[uint64]string),
		handlers:            make(map[string]func(*Protocol, *Validator, string)),
	}
	protocol.RegisterHandler("M", (*Protocol).MakeMessage)
	protocol.RegisterHandler("S", (*Protocol).SendMessage)
	protocol.RegisterHandler("SJ", (*Protocol).SendAndJustify)
	return protocol
}

func (p *Protocol) RegisterHandler(token string, function func(*Protocol, *Validator, string)) {
	if _, ok := p.handlers[token]; ok {
		_ = fmt.Errorf("a function has been registered with this token")
	}
	p.handlers[token] = function
}

func (p *Protocol) RegisterMessage(message *Message, name string) {
	if _, ok := p.Messages[name]; ok {
		_ = fmt.Errorf("message with %s already exists", name)
	}
	if _, ok := p.MessageFromHash[message.Hash()]; ok {
		_ = fmt.Errorf("message with %d already exists", message.Hash())
	}
	p.Messages[name] = message
	p.MessageFromHash[message.Hash()] = message
	p.MessageNameFromHash[message.Hash()] = name
	p.GlobalView.AddMessages([]*Message{message})
}

// MakeMessage 使用该验证器生成消息
func (p *Protocol) MakeMessage(validator *Validator, messageName string) {
	newMessage := validator.MakeNewMessage()
	p.RegisterMessage(newMessage, messageName)
}

// SendMessage 给该验证器传递一个消息
func (p *Protocol) SendMessage(validator *Validator, messageName string) {
	message := p.Messages[messageName]
	validator.ReceiveMessages([]*Message{message})
}

// SendAndJustify 传递消息以及其依赖
func (p *Protocol) SendAndJustify(validator *Validator, messageName string) {
	message := p.Messages[messageName]
	messageToSend := p.MessagesNeededToJustify(message, validator)
	//for _, msg := range messageToSend {
	//	fmt.Println(p.MessageNameFromHash[msg.Hash()])
	//	for _, hash := range msg.Justification {
	//		fmt.Println(p.MessageNameFromHash[hash])
	//	}
	//}
	validator.ReceiveMessages(messageToSend)
}

// MessagesNeededToJustify 返回消息及未验证依赖
func (p *Protocol) MessagesNeededToJustify(message *Message, validator *Validator) []*Message {
	messageNeeded := hashset.New(message)
	messageHashes := hashset.New(message.Hash())
	for messageHashes.Size() != 0 {
		nextHashes := hashset.New()
		for _, m := range messageHashes.Values() {
			messageHash := m.(uint64)
			message = p.MessageFromHash[messageHash]
			messageNeeded.Add(message)
			for _, hash := range message.Justification {
				if _, ok := validator.View.justifiedMessages[hash]; !ok {
					nextHashes.Add(hash)
				}
			}
		}
		messageHashes = nextHashes
	}
	result := make([]*Message, 0, messageNeeded.Size())
	for _, v := range messageNeeded.Values() {
		result = append(result, v.(*Message))
	}
	return result
}

func (p *Protocol) Execute(additionalStr string) {
	if additionalStr != "" {
		p.unexecuted += additionalStr
	}
	for _, token := range strings.Split(p.unexecuted, " ") {
		comm, vali, name, _ := parseToken(token)
		validator := p.GlobalValidatorSet.GetValByName(vali)
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
