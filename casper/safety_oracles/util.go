package safety_oracles

import (
	. "cbc-casper-go/casper"
	"cbc-casper-go/casper/protocols"
)

func ExistFreeMsg(bet protocols.Bet, val *Validator, seqNum uint64, view *View) bool {
	curMsg := view.LatestMessages[val]
	for curMsg.SeqNum > seqNum {
		if ok, _ := bet.ConflictWith(curMsg); ok {
			return true
		}
		if curMsg.SeqNum == 0 {
			break
		}
		nextHash := curMsg.Justification[val]
		curMsg = view.JustifiedMessages[nextHash]
	}
	return false
}
