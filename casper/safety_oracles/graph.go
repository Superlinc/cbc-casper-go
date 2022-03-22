package safety_oracles

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"reflect"
	"sort"
)

type Graph struct {
	v       []interface{}
	m       map[interface{}]int
	e       [][]int
	cliques [][]int
}

func NewGraph(edges [][]interface{}) *Graph {
	v := make([]interface{}, 0, 4)
	m := make(map[interface{}]int)
	length := 0
	for _, edge := range edges {
		if _, ok := m[edge[0]]; !ok {
			v = append(v, edge[0])
			m[edge[0]] = length
			length++
		}
		if _, ok := m[edge[1]]; !ok {
			v = append(v, edge[1])
			m[edge[1]] = length
			length++
		}
	}
	e := make([][]int, 0, length)
	for i := 0; i < length; i++ {
		e = append(e, make([]int, length, length))
	}
	for _, edge := range edges {
		//fmt.Println(m[edge[0]], m[edge[1]])
		//fmt.Println(e[0][1])
		e[m[edge[0]]][m[edge[1]]] = 1
		e[m[edge[1]]][m[edge[0]]] = 1
	}
	for _, ints := range e {
		fmt.Println(ints)
	}
	return &Graph{
		v: v,
		m: m,
		e: e,
	}
}

func (g *Graph) FindMaximalClique() [][]interface{} {
	R := hashset.New()
	P := hashset.New()
	X := hashset.New()
	for i := 0; i < len(g.v); i++ {
		P.Add(i)
	}
	max := new(int)
	g.cliques = make([][]int, 0, 4)
	g.dfs(R, P, X, max)
	fmt.Println(*max)
	ans := make([][]interface{}, 0, len(g.cliques))
	for _, clique := range g.cliques {
		item := make([]interface{}, 0, len(clique))
		for _, index := range clique {
			item = append(item, g.v[index])
		}
		ans = append(ans, item)
	}
	return ans
}

func (g *Graph) dfs(R, P, X *hashset.Set, max *int) {
	if P.Empty() {
		if R.Size() == *max {
			clique := make([]int, 0, R.Size())
			for _, v := range R.Values() {
				clique = append(clique, v.(int))
			}
			sort.Ints(clique)
			for _, c := range g.cliques {
				if reflect.DeepEqual(c, clique) {
					return
				}
			}
			g.cliques = append(g.cliques, clique)
		} else if R.Size() > *max {
			*max = R.Size()
			g.cliques = make([][]int, 0, 4)
			clique := make([]int, 0, R.Size())
			for _, v := range R.Values() {
				clique = append(clique, v.(int))
			}
			sort.Ints(clique)
			for _, c := range g.cliques {
				if reflect.DeepEqual(c, clique) {
					return
				}
			}
			g.cliques = append(g.cliques, clique)
		}
		return
	}
	p := make([]int, 0, P.Size())
	for _, v := range P.Values() {
		p = append(p, v.(int))
	}
	sort.Ints(p)
	r := make([]int, 0, R.Size())
	for _, v := range R.Values() {
		r = append(r, v.(int))
	}
	sort.Ints(r)
	for _, index1 := range p {
		ok := true
		for _, index2 := range r {
			if index2 < index1 || g.e[index1][index2] != 1 {
				ok = false
				break
			}
		}
		P.Remove(index1)
		if ok {
			R.Add(index1)
			g.dfs(R, P, X, max)
			R.Remove(index1)
		} else {
			X.Add(index1)
			g.dfs(R, P, X, max)
			X.Remove(index1)
		}
		P.Add(index1)
	}
}
