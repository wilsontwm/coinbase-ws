package main

import (
	"fmt"
	"os"

	"golang.org/x/net/websocket"
)

// WebsocketRequest :
type WebsocketRequest struct {
	Type       string   `json:"type"`
	ProductIDs []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

// WebsocketResponse :
type WebsocketResponse struct {
	Type      string `json:"type"`
	ProductID string `json:"product_id"`
	Price     string `json:"price"`
	Size      string `json:"size"`
}

const coinBaseURL = "ws-feed.exchange.coinbase.com"

// SubscribeToCoinBase :
func SubscribeToCoinBase(productIDs []string, c chan WebsocketResponse) {
	fmt.Println("Subscribing to coinbase websocket")
	ws, err := websocket.Dial(fmt.Sprintf("wss://%s", coinBaseURL), "", "http://localhost")
	if err != nil {
		fmt.Printf("Connect failed: %s\n", err.Error())
		os.Exit(1)
	}

	req := new(WebsocketRequest)
	req.Type = "subscribe"
	req.ProductIDs = productIDs
	req.Channels = []string{"matches"}

	err = websocket.JSON.Send(ws, req)
	if err != nil {
		fmt.Printf("Send failed: %s\n", err.Error())
		os.Exit(1)
	}

	go readClientMessages(ws, c)
}

func readClientMessages(ws *websocket.Conn, incomingMessages chan WebsocketResponse) {
	for {
		res := new(WebsocketResponse)
		err := websocket.JSON.Receive(ws, res)
		if err != nil {
			fmt.Printf("Receive failed: %s\n", err.Error())
			return
		}

		if res.Type == "match" {
			incomingMessages <- *res
		}
	}
}
