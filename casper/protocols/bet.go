package protocols

import "cbc-casper-go/casper"

type Bet interface {
	ConflictWith(message casper.Messager) (bool, error)
}
