package main

import (
	"apiclient.go/internal/config"
	"apiclient.go/internal/websocket"
	"fmt"
	"log"
)

const configName = "config.yaml"

func main() {
	conn := &websocket.APIClientStruct{}
	cfg, err := config.GetConfigFromFile(configName)
	fmt.Println(cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Connection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.SubscribeToChannel()
	if err != nil {
		log.Fatal(err)
	}
	conn.WriteMessagesToChannel()
	readChannel := make(chan websocket.BestOrderBook)
	conn.ReadMessagesFromChannel(readChannel)
	for {
		fmt.Println(<-readChannel)
	}
}
