package protocols

import "cbc-casper-go/casper"

type Bet interface {
	ConflictWith(message *casper.Message) (bool, error)
}
