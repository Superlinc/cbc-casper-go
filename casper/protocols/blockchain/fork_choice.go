package blockchain

import "cbc-casper-go/casper"

func getForkChoice(lastFinBlk *Block, children map[*Block][]*Block, latestMsg map[casper.AbstractValidator]casper.Messager) interface{} {
	scores := make(map[interface{}]uint64)
	for validator, curMsg := range latestMsg {
		curBlk := &Block{curMsg.(*casper.Message)}
		// todo 修改message抽象的结构
		for curBlk != nil && curBlk.Hash() != lastFinBlk.Hash() {
			scores[curBlk] += validator.Weight()
			curBlk = curBlk.Estimate.(*Block)
		}
	}
	bestBlk := lastFinBlk
	for {
		if _, ok := children[bestBlk]; !ok {
			break
		}
		curScores := make(map[*Block]uint64)
		var maxScore uint64
		for _, child := range children[bestBlk] {
			curScores[child] += scores[children]
			maxScore = casper.MaxUint(maxScore, curScores[child])
		}

		// 排除权重为0的孩子
		if maxScore == 0 {
			break
		}

		maxChildren := getMaxChild(curScores, maxScore)
		if len(maxChildren) != 1 {
			panic("length of max children != 1")
		}
		bestBlk = maxChildren[0]
	}
	return bestBlk
}

func getMaxChild(scores map[*Block]uint64, maxScore uint64) []*Block {
	res := make([]*Block, 0, 4)
	for blk, score := range scores {
		if score == maxScore {
			res = append(res, blk)
		}
	}
	return res
}
