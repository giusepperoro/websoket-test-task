package websocket

import (
	"fmt"
	"time"
)

const ping = 10

func (ap *APIClientStruct) WriteMessagesToChannel() {
	go func() {
		select {
		case <-time.NewTicker(ping * time.Second).C:
			err := ap.conn.WriteMessage(ap.messageType, nil)
			if err != nil {
				fmt.Errorf("writemessage error:%w", err)
				return
			}
		}
	}()
}
