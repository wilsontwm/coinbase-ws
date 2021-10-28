# Coinbase VWAP Calculator
A simple real-time VWAP (Volume Weighted Average Price) calculator that receives real-time match stream from Coinbase websocket

## Installation and Usage
1. Clone this repo to your workspace
2. Navigate to the folder and run the following command and enjoys the VWAP on your terminal

```
go build coinbase-ws && go run coinbase-ws
```

## Program Explanation
The program is mainly made up of the following components:

**websocket.go**: Subscribe to Coinbase websocket 'match' channel and outputs the messages from the stream in an channel which is further processed

**vwap.go**: Instantiates pairMatchAggregator for every different product, add/remove the match received from the stream and calculates the VWAP for the cumulated matches, returns the VWAP result for each product

**main.go**: An entrypoint of the program that connects the components above

## Test
Execute the following command to run the tests
```
go test .
```