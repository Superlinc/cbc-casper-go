package casper

type ValidatorSetor interface {
	Weight(validators []AbstractValidator) uint64
	Contains(validator AbstractValidator) bool
	SortedByName() []AbstractValidator
	SortedByWeight() []AbstractValidator
	GetValByName(name int) AbstractValidator
	GetValsByName(name []int) []AbstractValidator
	Names() []int
	Weights() []uint64
	Validators() []AbstractValidator
	Size() uint64
}
