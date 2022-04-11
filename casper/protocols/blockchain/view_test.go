package blockchain

import (
	"fmt"
	"testing"
)

func TestNewView(t *testing.T) {
	view := NewView()
	fmt.Println(view.Estimate())
	// p := casper.NewProtocol([]uint64{1, 2}, view, "", 1)
}
