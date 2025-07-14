package websocket

import (
	"time"

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

func (s Socket) LocalAddr() string {
	return s.conn.LocalAddr().String()
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

func (s Socket) RemoteAddr() string {
	return s.conn.RemoteAddr().String()
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
