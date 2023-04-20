package websocket

import (
	"apiclient.go/internal/config"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func (ap *APIClientStruct) Connection(cfg config.ServiceConfiguration) error {
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	connection, _, err := websocket.DefaultDialer.Dial(cfg.WebsocketRequest, header)
	if err != nil {
		return fmt.Errorf("unable to connect websocket: %w", err)
	}

	ap.conn, ap.cfg = connection, cfg
	return nil
}
