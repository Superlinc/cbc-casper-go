package casper

// View 存储已经收到的消息
type View struct {
	JustifiedMessages      map[uint64]*Message     // message hash => message
	PendingMessages        map[uint64]*Message     // message hash => message
	NumMissingDependencies map[uint64]int          // message hash => number of message hashes
	DependenciesOfMessage  map[uint64][]uint64     // message hash => list(message hashes)
	LatestMessages         map[*Validator]*Message // validator => message
}

func NewView(messages []*Message) *View {
	if messages == nil {
		messages = make([]*Message, 0, 4)
	}
	v := &View{
		JustifiedMessages:      make(map[uint64]*Message),
		PendingMessages:        make(map[uint64]*Message),
		NumMissingDependencies: make(map[uint64]int),
		DependenciesOfMessage:  make(map[uint64][]uint64),
		LatestMessages:         make(map[*Validator]*Message),
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
		if _, ok := v.PendingMessages[message.Hash()]; ok {
			continue
		}
		if _, ok := v.JustifiedMessages[message.Hash()]; ok {
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
	v.PendingMessages[h] = m
	v.NumMissingDependencies[h] = len(hashes)

	for _, hash := range hashes {
		if _, ok := v.DependenciesOfMessage[hash]; !ok {
			v.DependenciesOfMessage[hash] = make([]uint64, 0, 4)
		}
		v.DependenciesOfMessage[hash] = append(v.DependenciesOfMessage[hash], h)
	}

}

// GetNewlyJustifiedMessage 给定一个刚验证的信息, 得到所有因此得到验证的信息
func (v *View) GetNewlyJustifiedMessage(m *Message) []*Message {
	newlyJustifiedMessages := make([]*Message, 0, 4)
	newlyJustifiedMessages = append(newlyJustifiedMessages, m)
	for _, dependentHash := range v.DependenciesOfMessage[m.Hash()] {
		v.NumMissingDependencies[dependentHash] -= 1
		if v.NumMissingDependencies[dependentHash] == 0 {
			newMessage := v.PendingMessages[dependentHash]
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
	if _, ok := v.LatestMessages[m.Sender]; !ok {
		v.LatestMessages[m.Sender] = m
	} else if v.LatestMessages[m.Sender].SeqNum < m.SeqNum {
		v.LatestMessages[m.Sender] = m
	}
}

// AddJustifiedRemovePending 添加已验证的消息并删除相关数据
func (v *View) AddJustifiedRemovePending(m *Message) {
	h := m.Hash()
	v.JustifiedMessages[h] = m
	if _, ok := v.NumMissingDependencies[h]; ok {
		delete(v.NumMissingDependencies, h)
	}
	if _, ok := v.DependenciesOfMessage[h]; ok {
		delete(v.DependenciesOfMessage, h)
	}
	if _, ok := v.PendingMessages[h]; ok {
		delete(v.PendingMessages, h)
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
		if _, ok := v.JustifiedMessages[value]; !ok {
			result = append(result, value)
		}
	}
	return result
}
