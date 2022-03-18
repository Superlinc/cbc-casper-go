package casper

import "testing"

func TestProtocol_Execute(t *testing.T) {
	p := NewProtocol([]uint64{1, 2}, "", 10, 0, 0)
	p.Execute("M-0-A")
	if p.executed != "M-0-A" || p.unexecuted != "" {
		t.Errorf("error")
	}
}
