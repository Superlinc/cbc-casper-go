package casper

type Viewer interface {
	Estimate() interface{}
	UpdateSafeEstimates(valSet *ValidatorSet) Messager
	AddMessages(messages []Messager)
	LatestMsg() map[*Validator]Messager
	JustifiedMsg() map[uint64]Messager
}
