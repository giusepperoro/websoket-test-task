package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type quoteStream struct {
	Op string `json:"op"`
	Ch string `json:"ch"`
}

func (ap *APIClientStruct) SubscribeToChannel() error {
	qs := quoteStream{
		Op: ap.cfg.Op,
		Ch: ap.cfg.Ch,
	}

	toSub, err := json.Marshal(qs)
	if err != nil {
		return fmt.Errorf("error to masrshal in json: %w", err)
	}

	err = ap.conn.WriteMessage(websocket.TextMessage, toSub)
	if err != nil {
		return fmt.Errorf("error to write message: %w", err)
	}
	return nil
}
