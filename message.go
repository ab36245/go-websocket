package websocket

import (
	"fmt"
)

type Message struct {
	Kind MessageKind
	Data []byte
}

func (m Message) IsBinary() bool {
	return m.Kind == BinaryMessage
}

func (m Message) IsInvalid() bool {
	return m.Kind == InvalidMessage
}

func (m Message) IsText() bool {
	return m.Kind == TextMessage
}

func (m Message) Text() string {
	return string(m.Data)
}

func (m Message) String() string {
	return fmt.Sprintf("%s (%d bytes) %v", m.Kind, len(m.Data), m.Data)
}

type MessageKind int

const (
	InvalidMessage MessageKind = 0
	TextMessage    MessageKind = 1
	BinaryMessage  MessageKind = 2
	CloseMessage   MessageKind = 8
	PingMessage    MessageKind = 9
	PongMessage    MessageKind = 10
)

func (k MessageKind) String() string {
	switch k {
	case InvalidMessage:
		return "invalid"
	case TextMessage:
		return "text"
	case BinaryMessage:
		return "binary"
	case CloseMessage:
		return "close"
	case PingMessage:
		return "ping"
	case PongMessage:
		return "pong"
	default:
		return fmt.Sprintf("unknown (%d)", k)
	}
}
