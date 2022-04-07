package casper

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	m := &Message{
		header: 1.0,
	}
	fmt.Println(m.Hash())
}
