package safety_oracles

import (
	. "cbc-casper-go/casper"
	"cbc-casper-go/casper/protocols/integer"
	"errors"
)

// CliqueOracle todo
type CliqueOracle struct {
	candidateEstimate *Message
	View              *View
	ValSet            *ValidatorSet
	candidates        []*Validator
}

func NewCliqueOracle(msg *Message, view *View, valSet *ValidatorSet) (*CliqueOracle, error) {
	if msg == nil {
		return nil, errors.New("cannot decide if safe without an estimate")
	}
	candidates := make([]*Validator, 0, 4)
	for _, v := range valSet.Validators() {
		if _, ok := view.LatestMessages[v]; ok {
			b := &integer.Bet{Message: msg}
			if ok, _ := b.ConflictWith(view.LatestMessages[v]); !ok {
				candidates = append(candidates, v)
			}
		}
	}
	c := &CliqueOracle{
		candidateEstimate: msg,
		View:              view,
		ValSet:            valSet,
		candidates:        candidates,
	}
	return c, nil
}

func (o *CliqueOracle) collectEdge() [][2]*Validator {
	edges := make([][2]*Validator, 0, 4)
	for i := 0; i < len(o.candidates); i++ {
		for j := i + 1; j < len(o.candidates); j++ {
			v1 := o.candidates[i]
			v2 := o.candidates[j]
			b := &integer.Bet{Message: o.candidateEstimate}
			msg1 := o.View.LatestMessages[v1]
			if _, ok := msg1.Justification[v2]; !ok {
				continue
			}
			hash := msg1.Justification[v2]
			msg2InV1 := o.View.JustifiedMessages[hash]

			if ok, _ := b.ConflictWith(msg2InV1); ok {
				continue
			}

			msg2 := o.View.LatestMessages[v2]
			if _, ok := msg2.Justification[v1]; !ok {
				continue
			}
			hash = msg2.Justification[v1]
			msg1InV2 := o.View.JustifiedMessages[hash]
			if ok, _ := b.ConflictWith(msg1InV2); ok {
				continue
			}

			if ExistFreeMsg(b, v2, msg2InV1.SeqNum, o.View) {
				continue
			}
			if ExistFreeMsg(b, v1, msg1InV2.SeqNum, o.View) {
				continue
			}
			edges = append(edges, [2]*Validator{v1, v2})
		}
	}
	return edges
}

// 找到最大的验证器集合,满足:
// 1. 他们的最新消息都满足候选消息
// 2. 他们互相看过最新消息
// 3. 他们当中没有不满足候选消息的最新消息
func (o *CliqueOracle) findBiggestClique() {
	if o.ValSet.Weight(o.candidates) < o.ValSet.Weight(nil)/2 {
		return
	}
	// edges := o.collectEdge()

}
