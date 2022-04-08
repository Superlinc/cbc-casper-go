package blockchain

import "cbc-casper-go/casper"

func getForkChoice(lastFinBlk *Block, children map[*Block][]*Block, latestMsg map[casper.AbstractValidator]casper.Messager) interface{} {
	scores := make(map[*Block]uint64)
	for validator, curBlk := range latestMsg {
		for curBlk != nil && curBlk != lastFinBlk {
			scores[curBlk.(*Block)] += validator.Weight()
			curBlk = curBlk.(*Block).Estimate.(*Block)
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
			curScores[child] += scores[child]
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
