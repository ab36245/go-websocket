package websocket

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aivoicesystems/aivoice/common/stream"

	"github.com/gorilla/websocket"
)

type Socket struct {
	conn *websocket.Conn
}

func (s Socket) Close() error {
	if err := s.conn.Close(); err != nil {
		return CloseError.Wrap(err)
	}
	return nil
}

func (s Socket) Read() (Message, error) {
	type_, data, err := s.conn.ReadMessage()
	if err != nil {
		if e, ok := err.(*websocket.CloseError); ok {
			err = ClosedError.Mesg("code %d", e.Code)
		} else {
			err = ReadError.Wrap(err)
		}
		return Message{}, err
	}
	var kind MessageKind
	switch type_ {
	case websocket.TextMessage:
		kind = TextMessage
	case websocket.BinaryMessage:
		kind = BinaryMessage
	case websocket.CloseMessage:
		kind = CloseMessage
	case websocket.PingMessage:
		kind = PingMessage
	case websocket.PongMessage:
		kind = PongMessage
	default:
		kind = InvalidMessage
	}
	return Message{
		Kind: kind,
		Data: data,
	}, nil
}

func (s Socket) Sink(upstream *stream.Output[Message]) *stream.Output[int] {
	return stream.Add(
		upstream,
		func(input <-chan Message, output func(int)) error {
			for n := 0; ; n++ {
				m, ok := <-input
				if !m.IsInvalid() {
					err := s.Write(m)
					if err != nil {
						return err
					}
				}
				if !ok {
					output(n)
					return nil
				}
			}
		},
	)
}

func (s Socket) SinkBinary(upstream *stream.Output[[]byte]) *stream.Output[int] {
	return s.Sink(
		stream.Add(
			upstream,
			func(input <-chan []byte, output func(Message)) error {
				for {
					data, ok := <-input
					if data != nil {
						output(Message{Kind: BinaryMessage, Data: data})
					}
					if !ok {
						return nil
					}
				}
			},
		),
	)
}

func (s Socket) Stream() *stream.Output[Message] {
	return stream.New(
		func(output func(m Message)) error {
			for {
				m, err := s.Read()
				if err != nil {
					return err
				}
				output(m)
			}
		},
		func() {
			s.Close()
		},
	)
}

func (s Socket) StreamBinary() *stream.Output[[]byte] {
	return stream.Add(
		s.Stream(),
		func(input <-chan Message, output func(bytes []byte)) error {
			for {
				m, ok := <-input
				if m.IsBinary() {
					output(m.Data)
				}
				if !ok {
					return nil
				}
			}
		},
	)
}

func (s Socket) Write(message Message) error {
	switch message.Kind {
	case BinaryMessage:
		return s.write(websocket.BinaryMessage, message.Data)
	case TextMessage:
		return s.write(websocket.TextMessage, message.Data)
	default:
		return WriteError.Mesg("can't write message of kind %d", message.Kind)
	}
}

func (s Socket) WriteBinary(data []byte) error {
	return s.write(websocket.BinaryMessage, data)
}

func (s Socket) WriteClose() error {
	err := s.conn.WriteControl(websocket.CloseMessage, nil, time.Time{})
	if err != nil {
		return CloseError.Wrap(err)
	}
	return nil
}

func (s Socket) WriteText(text string) error {
	return s.write(websocket.TextMessage, []byte(text))
}

func checkOrigin(r *http.Request) bool {
	// TODO: this is a hack to allow web connections until I work out Origin
	// issues
	fmt.Printf("Horrible hack in CheckOrigin in websocket upgrade must be removed!!!\n")
	return true
}

func (s Socket) write(type_ int, data []byte) error {
	err := s.conn.WriteMessage(type_, data)
	if err != nil {
		if e, ok := err.(*websocket.CloseError); ok {
			err = ClosedError.Mesg("code %d", e.Code)
		} else {
			err = WriteError.Wrap(err)
		}
	}
	return err
}
