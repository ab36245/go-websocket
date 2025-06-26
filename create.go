package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func Connect(uri string) (Socket, error) {
	conn, _, err := websocket.DefaultDialer.Dial(uri, nil)
	if err != nil {
		return Socket{}, ConnectError.Wrap(err)
	}
	return Socket{conn}, nil
}

func Upgrade(w http.ResponseWriter, r *http.Request) (Socket, error) {
	upgrader := websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return Socket{}, UpgradeError.Wrap(err)
	}
	return Socket{conn}, nil
}

func checkOrigin(r *http.Request) bool {
	// TODO: this is a hack to allow web connections until I work out Origin
	// issues
	fmt.Printf("Horrible hack in CheckOrigin in websocket upgrade must be removed!!!\n")
	return true
}
