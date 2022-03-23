package integer

import (
	"cbc-casper-go/casper"
	"sort"
)

// 在给定的消息中取得共识
func getEstimate(latestMessage map[*casper.Validator]*casper.Message) int {
	msgs := make([]*casper.Message, 0, len(latestMessage))
	for _, v := range latestMessage {
		msgs = append(msgs, v)
	}
	sort.SliceStable(msgs, func(i, j int) bool {
		return msgs[i].Estimate.(int) < msgs[j].Estimate.(int)
	})
	var half float64
	for v, _ := range latestMessage {
		half += float64(v.Weight)
	}
	half /= 2.0
	var prefixWeight uint64
	for _, bet := range msgs {
		prefixWeight += bet.Sender.Weight
		if float64(prefixWeight) >= half {
			return bet.Estimate.(int)
		}
	}
	return msgs[len(msgs)-1].Estimate.(int)
}
