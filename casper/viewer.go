package casper

type Viewer interface {
	Estimate() interface{}
	UpdateSafeEstimates()
	AddMessages(messages []Messager)
	ReceiveJustifiedMessage(message Messager)
	ReceivePendingMessage(message Messager, hashes []uint64)
	LatestMsg() map[AbstractValidator]Messager
	JustifiedMsg() map[uint64]Messager
}
