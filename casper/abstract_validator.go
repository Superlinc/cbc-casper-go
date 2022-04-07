package casper

type AbstractValidator interface {
	InitializeView(messages []Messager)
	Estimate() interface{}
	ReceiveMessages(messages []Messager)
	MakeNewMessage() Messager
	Justification() map[AbstractValidator]uint64
	NextSeqNum() uint64
	NextDisPlayHeight() uint64
	View() Viewer
	Weight() uint64
	Name() int
}
