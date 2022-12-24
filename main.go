package main

import (
	"fmt"
	"os/user"

	"github.com/stocks-selector/src/config"
	"github.com/stocks-selector/src/out"
	"github.com/stocks-selector/src/stocks"
)

func main() {

	showWelcomeMessage()
	config.ConfigureApp()
	client := config.HttpClient()

	stocks, removedStocks := stocks.HandleSelectStocks(client)
	out.PrintStocks(stocks, removedStocks)
}

func showWelcomeMessage() {
	config.ClearScreen()
	user, err := user.Current()
	if err != nil {
		user.Username = "fellow user"
	}
	fmt.Print(config.Yellow, "\nHello,", user.Username+"!\n", config.Reset)
	fmt.Print("This is a simple stocks selector. It woks by sorting and ranking stocks by ROA, EV/Ebit and P/l.\n\n")
	fmt.Print(config.Green, "\tDeveloped by: https://github.com/devlipe\n")
	fmt.Print("\tOn April 10th 2022\n")
	fmt.Print("\tVersion 1.0.0\n", config.Reset)
	fmt.Print(config.Red, "\nPress enter to continue...\n", config.Reset)
	var trash string
	fmt.Scanln(&trash)
}
