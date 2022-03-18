package casper

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	m := &Message{
		Header: 1.0,
	}
	fmt.Println(m.Hash())
}
