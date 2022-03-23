package safety_oracles

import (
	. "cbc-casper-go/casper"
	"cbc-casper-go/casper/protocols"
	"errors"
	"github.com/emirpasic/gods/sets/hashset"
)

// CliqueOracle todo
type CliqueOracle struct {
	candidateEstimate protocols.Bet
	View              *View
	ValSet            *ValidatorSet
	candidates        []*Validator
}

func NewCliqueOracle(bet protocols.Bet, view *View, valSet *ValidatorSet) (*CliqueOracle, error) {
	if bet == nil {
		return nil, errors.New("cannot decide if safe without an estimate")
	}
	candidates := make([]*Validator, 0, 4)
	for _, v := range valSet.Validators() {
		if _, ok := view.LatestMessages[v]; ok {
			if ok, _ := bet.ConflictWith(view.LatestMessages[v]); !ok {
				candidates = append(candidates, v)
			}
		}
	}
	c := &CliqueOracle{
		candidateEstimate: bet,
		View:              view,
		ValSet:            valSet,
		candidates:        candidates,
	}
	return c, nil
}

func (o *CliqueOracle) collectEdge() [][]interface{} {
	edges := make([][]interface{}, 0, 4)
	for i := 0; i < len(o.candidates); i++ {
		for j := i + 1; j < len(o.candidates); j++ {
			v1 := o.candidates[i]
			v2 := o.candidates[j]
			msg1 := o.View.LatestMessages[v1]
			if _, ok := msg1.Justification[v2]; !ok {
				continue
			}
			hash := msg1.Justification[v2]
			msg2InV1 := o.View.JustifiedMessages[hash]

			if ok, _ := o.candidateEstimate.ConflictWith(msg2InV1); ok {
				continue
			}

			msg2 := o.View.LatestMessages[v2]
			if _, ok := msg2.Justification[v1]; !ok {
				continue
			}
			hash = msg2.Justification[v1]
			msg1InV2 := o.View.JustifiedMessages[hash]
			if ok, _ := o.candidateEstimate.ConflictWith(msg1InV2); ok {
				continue
			}

			if ExistFreeMsg(o.candidateEstimate, v2, msg2InV1.SeqNum, o.View) {
				continue
			}
			if ExistFreeMsg(o.candidateEstimate, v1, msg1InV2.SeqNum, o.View) {
				continue
			}
			edges = append(edges, []interface{}{v1, v2})
		}
	}
	return edges
}

// 找到最大的验证器集合,满足:
// 1. 他们的最新消息都满足候选消息
// 2. 他们互相看过最新消息
// 3. 他们当中没有不满足候选消息的最新消息
func (o *CliqueOracle) findBiggestClique() ([]*Validator, uint64) {
	if o.ValSet.Weight(o.candidates) < o.ValSet.Weight(nil)/2 {
		return nil, 0
	}
	edges := o.collectEdge()
	g := NewGraph(edges...)
	cliques := g.FindMaximalClique()
	var maxWeight uint64
	var maxClique []*Validator
	for _, clique := range cliques {
		weight := GetWeight(clique...)
		if weight > maxWeight {
			maxWeight = weight
			maxClique = InterfaceToValidator(clique...)
		}
	}
	return maxClique, maxWeight
}

// CheckEstimateSafety todo
func (o *CliqueOracle) CheckEstimateSafety() (uint64, int) {
	biggestClique, cliqueWeight := o.findBiggestClique()
	faultTolerance := 2*cliqueWeight - o.ValSet.Weight(nil)

	if faultTolerance <= 0 {
		return 0, 0
	}

	equivocating := hashset.New()
	weights := hashset.New()
	diff := hashset.New()
	for _, validator := range biggestClique {
		weights.Add(validator.Weight)
		diff.Add(validator.Weight)
	}
	for Sum(equivocating) < faultTolerance {
		equivocating.Add(Max(diff))
		diff.Remove(Max(diff))
	}
	return faultTolerance, equivocating.Size() - 1
}
