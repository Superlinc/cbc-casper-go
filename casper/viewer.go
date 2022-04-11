package casper

type Viewer interface {
	Estimate() interface{}
	UpdateSafeEstimates(valSet *ValidatorSet)
	AddMessages(messages []Messager)
	LatestMsg() map[*Validator]Messager
	JustifiedMsg() map[uint64]Messager
}
