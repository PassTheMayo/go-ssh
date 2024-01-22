package messages

const (
	Disconnect      MessageType = 1
	Ignore          MessageType = 2
	Unimplemented   MessageType = 3
	Debug           MessageType = 4
	ServiceRequest  MessageType = 5
	ServiceAccept   MessageType = 6
	KeyExchangeInit MessageType = 20
	NewKeys         MessageType = 21
)

type MessageType byte
