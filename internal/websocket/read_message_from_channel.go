package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

//
//type Data struct {
//	Ts  int64     `json:"ts"`
//	Bid [2]string `json:"bid"`
//	Ask [2]string `json:"ask"`
//}

type BBOMessageResponse struct {
	//M      string `json:"m"`
	//Symbol string `json:"symbol"`
	//Data   Data   `json:"data"`
	M      string `json:"m"`
	Symbol string `json:"symbol"`
	Data   struct {
		Ts  int64     `json:"ts"`
		Bid [2]string `json:"bid"`
		Ask [2]string `json:"ask"`
	} `json:"data"`
}

func (ap *APIClientStruct) ReadMessagesFromChannel(ch chan<- BestOrderBook) {
	go func() {
		for {
			var response BBOMessageResponse

			messageType, message, err := ap.conn.ReadMessage()
			if err != nil {
				continue
			}
			ap.messageType = messageType

			err = json.Unmarshal(message, &response)
			if err != nil {
				continue
			}

			fmt.Println(response)
			askAmount, err := strconv.ParseFloat(response.Data.Ask[0], 64)
			fmt.Println(askAmount)
			if err != nil {
				log.Fatal(fmt.Errorf("askAmount failed parse: %w", err))
			}

			askPrice, err := strconv.ParseFloat(response.Data.Ask[1], 64)
			if err != nil {
				log.Fatal(fmt.Errorf("askPrice failed parse: %w", err))
			}

			bidAmount, err := strconv.ParseFloat(response.Data.Bid[0], 64)
			if err != nil {
				log.Fatal(fmt.Errorf("bidAmount failed parse: %w", err))
			}

			bidPrice, err := strconv.ParseFloat(response.Data.Bid[1], 64)
			if err != nil {
				log.Fatal(fmt.Errorf("bidPrice failed parse: %w", err))
			}

			//asks.Price > any bids.Price
			if askPrice <= bidPrice {
				continue
			}
			data := BestOrderBook{
				Ask: Order{
					Amount: askAmount,
					Price:  askPrice,
				},
				Bid: Order{
					Amount: bidAmount,
					Price:  bidPrice,
				},
			}

			ch <- data
		}
	}()
}
