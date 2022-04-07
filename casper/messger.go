package casper

type Messager interface {
	Hash() uint64
	Sender() AbstractValidator
	Justification() map[AbstractValidator]uint64
	SeqNum() uint64
	DisHeight() uint64
}
