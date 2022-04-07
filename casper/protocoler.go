package casper

type Protocoler interface {
	MakeMessage(validator AbstractValidator, messageName string)
	SendMessage(validator AbstractValidator, messageName string)
	SendAndJustify(validator AbstractValidator, messageName string)
	MessagesNeededToJustify(message Messager, validator AbstractValidator) []Messager
	Execute(additionalStr string)
}
