package websocket

import "log"

func (ap *APIClientStruct) Disconnect() {
	if ap.conn != nil {
		err := ap.conn.Close()
		if err != nil {
			log.Fatal("AHAHAHHAHAHA")
		}
	}
}
