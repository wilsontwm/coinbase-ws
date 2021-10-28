package main

import (
	"fmt"
	"strconv"
)

const (
	productBTCUSD = "BTC-USD"
	productETHUSD = "ETH-USD"
	productETHBTC = "ETH-BTC"
)

var productIDs = []string{productBTCUSD, productETHBTC, productETHUSD}

func main() {
	incomingWSResponse := make(chan WebsocketResponse)

	SubscribeToCoinBase(productIDs, incomingWSResponse)

	for {
		select {
		case message := <-incomingWSResponse:
			fmt.Printf("Match Received -> Product: %s Price: %s Size: %s\n", message.ProductID, message.Price, message.Size)
			price, err := strconv.ParseFloat(message.Price, 64)
			if err != nil {
				continue
			}

			size, err := strconv.ParseFloat(message.Size, 64)
			if err != nil {
				continue
			}

			addMatch(message.ProductID, &match{Size: size, Price: price})

			getAllVwap()
		}
	}
}

func getAllVwap() {
	for _, productID := range productIDs {
		if aggregator, ok := pairMatchAggregator[productID]; ok {
			fmt.Printf("%s -> %s\n", productID, aggregator.getString())
		}
	}
}
