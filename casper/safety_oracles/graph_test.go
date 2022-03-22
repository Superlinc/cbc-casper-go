package safety_oracles

import (
	"testing"
)

func TestGraph_FindMaximalClique(t *testing.T) {
	edges := make([][]interface{}, 0, 2)
	edges = append(edges, []interface{}{1, 2})
	edges = append(edges, []interface{}{1, 4})
	edges = append(edges, []interface{}{1, 5})
	edges = append(edges, []interface{}{2, 3})
	edges = append(edges, []interface{}{2, 5})
	edges = append(edges, []interface{}{3, 5})
	edges = append(edges, []interface{}{4, 5})
	g := NewGraph(edges)
	cliques := g.FindMaximalClique()
	if len(cliques) != 3 {
		t.Errorf("error")
	}
}
