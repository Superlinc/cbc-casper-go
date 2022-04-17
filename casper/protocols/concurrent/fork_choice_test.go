package concurrent

import (
	"testing"
)

func TestGetSource(t *testing.T) {
	protocol := getProtocol(nil)
	getForkChoice(protocol.GlobalView.(*View).children, protocol.GlobalView.LatestMsg())
}
