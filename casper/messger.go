package casper

type Messager interface {
	Hash() uint64
	Sender() *Validator
	Justification() map[*Validator]uint64
	SeqNum() uint64
	DisHeight() uint64
}
