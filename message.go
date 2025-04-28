package websocket

import "github.com/aivoicesystems/aivoice/common/writer"

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
	return writer.Value(m)
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
