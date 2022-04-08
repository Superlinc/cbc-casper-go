package casper

type Viewer interface {
	Estimate() interface{}
	UpdateSafeEstimates()
	AddMessages(messages []Messager)
	LatestMsg() map[AbstractValidator]Messager
	JustifiedMsg() map[uint64]Messager
}
