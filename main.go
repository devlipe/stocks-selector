package main

import (
	"github.com/stocks-selector/src/config"
	"github.com/stocks-selector/src/out"
	"github.com/stocks-selector/src/stocks"
)

func main() {
	config.GenerateFlags()

	client := config.HttpClient()

	stocks, removedStocks := stocks.HandleSelectStocks(client)
	out.PrintStocks(stocks, removedStocks)
}
