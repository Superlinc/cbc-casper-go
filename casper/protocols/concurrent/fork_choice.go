package concurrent

import (
	"cbc-casper-go/casper"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"
)

func getAncestors(block *Block) *hashset.Set {
	ancestors := hashset.New()
	stack := arraystack.New()
	stack.Push(block)
	for stack.Size() > 0 {
		tmp, _ := stack.Pop()
		curr := tmp.(*Block)
		if curr == nil {
			continue
		}
		if !ancestors.Contains(curr) {
			ancestors.Add(curr)
			ancestors.Add(curr.Estimate.(map[string]interface{})["blocks"].([]interface{})...)
		}
	}
	return ancestors
}

func getScores(latestMsg map[*casper.Validator]casper.Messager) map[*Block]uint64 {
	scores := make(map[*Block]uint64)
	for validator, messager := range latestMsg {
		ancestors := getAncestors(messager.(*Block))
		for _, ancestor := range ancestors.Values() {
			scores[ancestor.(*Block)] += validator.Weight()
		}
	}
	return scores
}

func getOutputs(blocks []*Block) *hashset.Set {
	outputs := hashset.New()
	for _, block := range blocks {
		outputs.Add(block.Estimate.(map[string]interface{})["outputs"].(*hashset.Set).Values()...)
	}
	return outputs
}

func updateOutputs(outputs *hashset.Set, blocks []*Block) {
	for _, block := range blocks {
		m := block.Estimate.(map[string]interface{})
		outputs.Remove(m["inputs"].(*hashset.Set).Values()...)
		outputs.Add(m["outputs"].(*hashset.Set).Values()...)
	}
}

func trackOutputSources(sources map[interface{}]*Block, blocks []*Block) {
	for _, block := range blocks {
		for _, value := range block.Estimate.(map[string]interface{})["outputs"].(*hashset.Set).Values() {
			sources[value] = block
		}
	}
}

func isConsumable(block *Block, curBlks *hashset.Set, scores map[*Block]uint64, outputs *hashset.Set) bool {
	inputs := block.Estimate.(map[string]interface{})["inputs"].(*hashset.Set)
	for _, other := range curBlks.Values() {
		if casper.Intersection(inputs, other.(*Block).Estimate.(map[string]interface{})["inputs"].(*hashset.Set)).Size() > 0 {
			if scores[block] < scores[other.(*Block)] {
				return false
			}
		}
		for _, output := range inputs.Values() {
			if !outputs.Contains(output) {
				return false
			}
		}
	}
	return true
}

func getChildren(blocks []*Block, childrenMap map[*Block][]*Block) *hashset.Set {
	children := hashset.New()
	for _, block := range blocks {
		if _, ok := childrenMap[block]; ok {
			for _, child := range childrenMap[block] {
				children.Add(child)
			}
		}
	}
	return children
}

func getForkChoice(children map[*Block][]*Block, latestMsg map[*casper.Validator]casper.Messager) (*hashset.Set, map[interface{}]*Block) {
	sources := make(map[interface{}]*Block)
	avaiOutputs := hashset.New()
	for _, block := range children[nil] {
		avaiOutputs.Add(block.Estimate.(map[string]interface{})["inputs"].(*hashset.Set).Values()...)
	}
	scores := getScores(latestMsg)
	curBlks := children[nil]
	trackOutputSources(sources, curBlks)
	updateOutputs(avaiOutputs, curBlks)
	curChildren := getChildren(curBlks, children)

	for curChildren.Size() > 0 {
		nextBlks := make([]*Block, 0, 4)
		for _, value := range curChildren.Values() {
			if isConsumable(value.(*Block), curChildren, scores, avaiOutputs) {
				nextBlks = append(nextBlks, value.(*Block))
			}
		}
		curBlks = nextBlks
		trackOutputSources(sources, curBlks)
		updateOutputs(avaiOutputs, curBlks)
		curChildren = getChildren(curBlks, children)
	}
	return avaiOutputs, sources
}
