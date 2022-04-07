package casper

// View 存储已经收到的消息
type View struct {
	justifiedMsg  map[uint64]Messager            // message hash => message
	pendingMsg    map[uint64]Messager            // message hash => message
	numMissDepend map[uint64]int                 // message hash => number of message hashes
	msgDepend     map[uint64][]uint64            // message hash => list(message hashes)
	latestMsg     map[AbstractValidator]Messager // validator => message
}

func NewView(msg []Messager) Viewer {
	if msg == nil {
		msg = make([]Messager, 0, 4)
	}
	v := &View{
		justifiedMsg:  make(map[uint64]Messager),
		pendingMsg:    make(map[uint64]Messager),
		numMissDepend: make(map[uint64]int),
		msgDepend:     make(map[uint64][]uint64),
		latestMsg:     make(map[AbstractValidator]Messager),
	}
	v.AddMessages(msg)
	return v
}

func (v *View) JustifiedMsg() map[uint64]Messager {
	return v.justifiedMsg
}

func (v *View) LatestMsg() map[AbstractValidator]Messager {
	return v.latestMsg
}

func (v *View) Estimate() interface{} {
	return 0
}

func (v *View) UpdateSafeEstimates() {
	panic("implement me")
}

// AddMessages 添加新的message到pending或者justify
func (v *View) AddMessages(msgs []Messager) {
	for _, msg := range msgs {
		if _, ok := v.pendingMsg[msg.Hash()]; ok {
			continue
		}
		if _, ok := v.justifiedMsg[msg.Hash()]; ok {
			continue
		}
		missMsgHashes := v.missMsgInJustify(msg)
		if len(missMsgHashes) == 0 {
			v.ReceiveJustifiedMessage(msg)
		} else {
			v.ReceivePendingMessage(msg, missMsgHashes)
		}
	}
}

// ReceiveJustifiedMessage 在收到已验证的消息后,处理等待队列并添加到View中
func (v *View) ReceiveJustifiedMessage(m Messager) {
	messages := v.GetNewlyJustifiedMessage(m)
	for _, message := range messages {
		v.AddToLatestMessage(message)
		v.AddJustifiedRemovePending(message)
		v.UpdateProtocolSpecificView(message)
	}
}

// ReceivePendingMessage 更新待验证消息
func (v *View) ReceivePendingMessage(m Messager, hashes []uint64) {
	h := m.Hash()
	v.pendingMsg[h] = m
	v.numMissDepend[h] = len(hashes)

	for _, hash := range hashes {
		if _, ok := v.msgDepend[hash]; !ok {
			v.msgDepend[hash] = make([]uint64, 0, 4)
		}
		v.msgDepend[hash] = append(v.msgDepend[hash], h)
	}

}

// GetNewlyJustifiedMessage 给定一个刚验证的信息, 得到所有因此得到验证的信息
func (v *View) GetNewlyJustifiedMessage(m Messager) []Messager {
	newlyJustifiedMessages := make([]Messager, 0, 4)
	newlyJustifiedMessages = append(newlyJustifiedMessages, m)
	for _, dependentHash := range v.msgDepend[m.Hash()] {
		v.numMissDepend[dependentHash] -= 1
		if v.numMissDepend[dependentHash] == 0 {
			newMessage := v.pendingMsg[dependentHash]
			newlyJustifiedMessages = append(newlyJustifiedMessages, v.GetNewlyJustifiedMessage(newMessage)...)
		}
	}
	return newlyJustifiedMessages
}

func (v *View) UpdateProtocolSpecificView(m Messager) {
	// 未实现
}

// AddToLatestMessage 更新validator的最新消息
func (v *View) AddToLatestMessage(m Messager) {
	if _, ok := v.latestMsg[m.Sender()]; !ok {
		v.latestMsg[m.Sender()] = m
	} else if v.latestMsg[m.Sender()].SeqNum() < m.SeqNum() {
		v.latestMsg[m.Sender()] = m
	}
}

// AddJustifiedRemovePending 添加已验证的消息并删除相关数据
func (v *View) AddJustifiedRemovePending(m Messager) {
	h := m.Hash()
	v.justifiedMsg[h] = m
	if _, ok := v.numMissDepend[h]; ok {
		delete(v.numMissDepend, h)
	}
	if _, ok := v.msgDepend[h]; ok {
		delete(v.msgDepend, h)
	}
	if _, ok := v.pendingMsg[h]; ok {
		delete(v.pendingMsg, h)
	}
}

// missMsgInJustify 返回该消息中已验证但是本View中未验证消息的哈希值
func (v *View) missMsgInJustify(m Messager) []uint64 {
	values := make(map[uint64]bool)
	for _, value := range m.Justification() {
		values[value] = true
	}
	result := make([]uint64, 0, 4)
	for value := range values {
		if _, ok := v.justifiedMsg[value]; !ok {
			result = append(result, value)
		}
	}
	return result
}
