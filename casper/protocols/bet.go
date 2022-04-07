package protocols

type Bet interface {
	ConflictWith(message interface{}) (bool, error)
}
