package casper

// View 存储已经收到的消息
type View struct {
	justifiedMessages      map[uint64]*Message     // message hash => message
	pendingMessages        map[uint64]*Message     // message hash => message
	numMissingDependencies map[uint64]int          // message hash => number of message hashes
	dependenciesOfMessage  map[uint64][]uint64     // message hash => list(message hashes)
	latestMessages         map[*Validator]*Message // validator => message
}

func NewView(messages []*Message) *View {
	if messages == nil {
		messages = make([]*Message, 0, 4)
	}
	v := &View{
		justifiedMessages:      make(map[uint64]*Message),
		pendingMessages:        make(map[uint64]*Message),
		numMissingDependencies: make(map[uint64]int),
		dependenciesOfMessage:  make(map[uint64][]uint64),
		latestMessages:         make(map[*Validator]*Message),
	}
	v.AddMessages(messages)
	return v
}

func (v *View) Estimate() interface{} {
	// panic("need to be implemented")
	// todo
	return 0
}

func (v *View) updateSafeEstimate(set ValidatorSet) {
	panic("need to be implemented")
}

// AddMessages 添加新的message到pending或者justify
func (v *View) AddMessages(messages []*Message) {
	for _, message := range messages {
		if _, ok := v.pendingMessages[message.Hash()]; ok {
			continue
		}
		if _, ok := v.justifiedMessages[message.Hash()]; ok {
			continue
		}
		missingMessageHashes := v.MissingMessageInJustification(message)
		if len(missingMessageHashes) == 0 {
			v.ReceiveJustifiedMessage(message)
		} else {
			v.ReceivePendingMessage(message, missingMessageHashes)
		}
	}
}

// ReceiveJustifiedMessage 在收到已验证的消息后,处理等待队列并添加到View中
func (v *View) ReceiveJustifiedMessage(m *Message) {
	messages := v.GetNewlyJustifiedMessage(m)
	for _, message := range messages {
		v.AddToLatestMessage(message)
		v.AddJustifiedRemovePending(message)
		v.UpdateProtocolSpecificView(message)
	}
}

// ReceivePendingMessage 更新待验证消息
func (v *View) ReceivePendingMessage(m *Message, hashes []uint64) {
	h := m.Hash()
	v.pendingMessages[h] = m
	v.numMissingDependencies[h] = len(hashes)

	for _, hash := range hashes {
		if _, ok := v.dependenciesOfMessage[hash]; !ok {
			v.dependenciesOfMessage[hash] = make([]uint64, 0, 4)
		}
		v.dependenciesOfMessage[hash] = append(v.dependenciesOfMessage[hash], h)
	}

}

// GetNewlyJustifiedMessage 给定一个刚验证的信息, 得到所有因此得到验证的信息
func (v *View) GetNewlyJustifiedMessage(m *Message) []*Message {
	newlyJustifiedMessages := make([]*Message, 0, 4)
	newlyJustifiedMessages = append(newlyJustifiedMessages, m)
	for _, dependentHash := range v.dependenciesOfMessage[m.Hash()] {
		v.numMissingDependencies[dependentHash] -= 1
		if v.numMissingDependencies[dependentHash] == 0 {
			newMessage := v.pendingMessages[dependentHash]
			newlyJustifiedMessages = append(newlyJustifiedMessages, v.GetNewlyJustifiedMessage(newMessage)...)
		}
	}
	return newlyJustifiedMessages
}

func (v *View) UpdateProtocolSpecificView(m *Message) {
	return
}

// AddToLatestMessage 更新validator的最新消息
func (v *View) AddToLatestMessage(m *Message) {
	if _, ok := v.latestMessages[m.Sender]; !ok {
		v.latestMessages[m.Sender] = m
	} else if v.latestMessages[m.Sender].SeqNum < m.SeqNum {
		v.latestMessages[m.Sender] = m
	}
}

// AddJustifiedRemovePending 添加已验证的消息并删除相关数据
func (v *View) AddJustifiedRemovePending(m *Message) {
	h := m.Hash()
	v.justifiedMessages[h] = m
	if _, ok := v.numMissingDependencies[h]; ok {
		delete(v.numMissingDependencies, h)
	}
	if _, ok := v.dependenciesOfMessage[h]; ok {
		delete(v.dependenciesOfMessage, h)
	}
	if _, ok := v.pendingMessages[h]; ok {
		delete(v.pendingMessages, h)
	}
}

// MissingMessageInJustification 返回该消息中已验证但是本View中未验证消息的哈希值
func (v *View) MissingMessageInJustification(m *Message) []uint64 {
	values := make(map[uint64]bool)
	for _, value := range m.Justification {
		values[value] = true
	}
	result := make([]uint64, 0, 4)
	for value := range values {
		if _, ok := v.justifiedMessages[value]; !ok {
			result = append(result, value)
		}
	}
	return result
}
