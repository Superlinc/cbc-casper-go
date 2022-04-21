package safety_oracles

import (
	. "cbc-casper-go/casper"
	"cbc-casper-go/casper/protocols"
)

func ExistFreeMsg(bet protocols.Bet, val *Validator, seqNum uint64, view *View) bool {
	curMsg := view.LatestMsg()[val]
	for curMsg.SeqNum() > seqNum {
		if ok, _ := bet.ConflictWith(curMsg); ok {
			return true
		}
		if curMsg.SeqNum() == 0 {
			break
		}
		nextHash := curMsg.Justification()[val]
		curMsg = view.JustifiedMsg()[nextHash]
	}
	return false
}

func GetWeight(validators ...interface{}) float64 {
	var weight float64
	for _, validator := range validators {
		weight += validator.(*Validator).Weight()
	}
	return weight
}

func InterfaceToValidator(values ...interface{}) []*Validator {
	validators := make([]*Validator, 0, len(values))
	for _, value := range values {
		validators = append(validators, value.(*Validator))
	}
	return validators
}
