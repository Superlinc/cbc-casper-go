package concurrent

import "testing"

func TestGenesis(t *testing.T) {
	protocol := getProtocol(nil)
	if len(protocol.GlobalView.JustifiedMsg()) != 1 {
		t.Errorf("length error")
	}
}
