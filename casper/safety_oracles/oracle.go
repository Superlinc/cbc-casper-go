package safety_oracles

type Oracle interface {
	CheckEstimateSafety() bool
}
