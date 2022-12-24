package out

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/stocks-selector/src/config"
	"github.com/stocks-selector/src/model"
	"github.com/xuri/excelize/v2"
)

func PrintStocks(stocks []model.Stock, removedStocks []model.Stock) {
	fileName := time.Now().Format("02-Jan-2006")

	switch viper.GetString("OUTPUT") {
	case "xlsx":
		writeToxlsx(stocks, removedStocks, fileName)
	case "txt":
		writeToTxt(stocks, removedStocks, fileName)
	case "cli":
		writeToCli(stocks, removedStocks, fileName)
	default:
		fmt.Println("Output not supported")
	}
}

func writeToCli(stocks []model.Stock, removedStocks []model.Stock, date string) {
	config.ClearScreen()
	printSelectedToCli(date, stocks)
	var awnser string
	for awnser == "" {
		fmt.Print("\nDo you want to see the removed stocks? [yes/no]: ")
		fmt.Scanf("%s", &awnser)
	}
	awnser = strings.ToLower(awnser)
	if awnser == "yes" || awnser == "y" {

		printRemovedToCli(removedStocks)
	}
}

func writeToxlsx(stocks []model.Stock, removedStocks []model.Stock, fileName string) {
	file := excelize.NewFile()
	defer file.Close()

	printSelectedToXlsx(stocks, fileName, file)
	printRemovedToXlsx(removedStocks, fileName, file)

	err := os.MkdirAll("results/xlsx/", 0777)
	if err != nil {
		log.Fatalln("Unable to create directory", err.Error())
	}

	fileName += ".xlsx"
	err = file.SaveAs("results/xlsx/" + fileName)
	if err != nil {
		log.Fatalln("Unable to create xlsx file", err.Error())
	}

}

func printSelectedToXlsx(stocks []model.Stock, fileName string, file *excelize.File) {

	file.SetSheetName("Sheet1", fileName)

	file.SetCellValue(fileName, "A1", "Rank")
	file.SetCellValue(fileName, "B1", "Ticker")
	file.SetCellValue(fileName, "C1", "Empresa")
	file.SetCellValue(fileName, "D1", "Preço")
	file.SetCellValue(fileName, "E1", "PL")
	file.SetCellValue(fileName, "F1", "Ev/Ebit")
	file.SetCellValue(fileName, "G1", "Roa")
	file.SetCellValue(fileName, "H1", "Liquidez")
	file.SetCellValue(fileName, "I1", "Geral")
	file.SetCellValue(fileName, "J1", "Rank EvEbit")
	file.SetCellValue(fileName, "K1", "Rank Roa")
	file.SetCellValue(fileName, "L1", "Rank Pl")

	for i, s := range stocks {

		file.SetCellValue(fileName, fmt.Sprintf("A%d", i+2), i+1)
		file.SetCellValue(fileName, fmt.Sprintf("B%d", i+2), s.Ticker)
		file.SetCellValue(fileName, fmt.Sprintf("C%d", i+2), s.Company_name)
		file.SetCellValue(fileName, fmt.Sprintf("D%d", i+2), s.Price)
		file.SetCellValue(fileName, fmt.Sprintf("E%d", i+2), s.P_L)
		file.SetCellValue(fileName, fmt.Sprintf("F%d", i+2), s.EV_Ebit)
		file.SetCellValue(fileName, fmt.Sprintf("G%d", i+2), s.Roa)
		file.SetCellValue(fileName, fmt.Sprintf("H%d", i+2), s.LiquidezMediaDiaria)
		file.SetCellValue(fileName, fmt.Sprintf("I%d", i+2), s.RankGeral)
		file.SetCellValue(fileName, fmt.Sprintf("J%d", i+2), s.RankEvEbit)
		file.SetCellValue(fileName, fmt.Sprintf("K%d", i+2), s.RankRoa)
		file.SetCellValue(fileName, fmt.Sprintf("L%d", i+2), s.RankPl)
	}
}

func printRemovedToXlsx(removedStocks []model.Stock, fileName string, file *excelize.File) {
	deletedFileName := fileName + "_deleted"

	file.NewSheet(deletedFileName)
	file.SetCellValue(deletedFileName, "A1", "Rank")
	file.SetCellValue(deletedFileName, "B1", "Ticker")
	file.SetCellValue(deletedFileName, "C1", "Empresa")
	file.SetCellValue(deletedFileName, "D1", "Preço")
	file.SetCellValue(deletedFileName, "E1", "PL")
	file.SetCellValue(deletedFileName, "F1", "Ev/Ebit")
	file.SetCellValue(deletedFileName, "G1", "Roa")
	file.SetCellValue(deletedFileName, "H1", "Liquidez")
	file.SetCellValue(deletedFileName, "H1", "Motivo")
	for i, s := range removedStocks {

		file.SetCellValue(deletedFileName, fmt.Sprintf("A%d", i+2), i+1)
		file.SetCellValue(deletedFileName, fmt.Sprintf("B%d", i+2), s.Ticker)
		file.SetCellValue(deletedFileName, fmt.Sprintf("C%d", i+2), s.Company_name)
		file.SetCellValue(deletedFileName, fmt.Sprintf("D%d", i+2), s.Price)
		file.SetCellValue(deletedFileName, fmt.Sprintf("E%d", i+2), s.P_L)
		file.SetCellValue(deletedFileName, fmt.Sprintf("F%d", i+2), s.EV_Ebit)
		file.SetCellValue(deletedFileName, fmt.Sprintf("G%d", i+2), s.Roa)
		file.SetCellValue(deletedFileName, fmt.Sprintf("H%d", i+2), s.LiquidezMediaDiaria)
		file.SetCellValue(deletedFileName, fmt.Sprintf("H%d", i+2), s.ExcludedReason)
	}

}

func writeToTxt(stocks []model.Stock, removedStocks []model.Stock, fileName string) {
	err := os.MkdirAll("results/txt/"+fileName, 0777)
	if err != nil {
		log.Fatalln("Unable to create directory", err.Error())
	}
	printSelectToTxt(stocks, fileName)
	printRemovedToTxt(removedStocks, fileName)

}

func printRemovedToTxt(remoevedStocks []model.Stock, fileName string) {
	removed, err := os.Create("results/txt/" + fileName + "/Removed-" + fileName + ".txt")
	if err != nil {
		log.Fatalln("Unable to create removed txt file", err.Error())
	}

	defer removed.Close()

	removed.WriteString(fmt.Sprintf("%-6s %-15s %-9s | %-8s %-8s %-8s %-8s | %10s  | %-8s\n",
		"Ticker", "Empresa", "Preço",
		"PL", "EV/Ebit", "ROA", "M Ebit",
		"Liquidez",
		"Mot. Excl."))

	for _, s := range remoevedStocks {
		removed.WriteString(fmt.Sprintf("%-6s %-15.14s R$%-7.2f | %-8.2f %-8.2f %-8.2f %-8.2f | %8.f K  | %-8s\n",
			s.Ticker, s.Company_name, s.Price,
			s.P_L, s.EV_Ebit, s.Roa, s.MargemEbit,
			s.LiquidezMediaDiaria/1000,
			s.ExcludedReason))
	}
}

func printSelectToTxt(stocks []model.Stock, fileName string) {
	selected, err := os.Create("results/txt/" + fileName + "/Selected-" + fileName + ".txt")
	if err != nil {
		log.Fatalln("Unable to create selected txt file", err.Error())
	}
	defer selected.Close()

	selected.WriteString(fmt.Sprintf("%-4s | %-6s %-15s %-15s %-8s | %-8s %-8s %-8s %-8s | %-10s | %-5s\n",
		"Rank", "Ticker", "Empresa", "Setor", "Preço", "PL", "EV/Ebit", "ROA", "M Ebit", "Liquidez", "Geral"))

	for i, s := range stocks {
		selected.WriteString(fmt.Sprintf("%-4d | %-6s %-15.14s %-15s R$%-6.2f | %-8.2f %-8.2f %-8.2f %-8.2f | %8.f K | %-5d \n",
			i+1,
			s.Ticker, s.Company_name, "___", s.Price,
			s.P_L, s.EV_Ebit, s.Roa, s.MargemEbit,
			s.LiquidezMediaDiaria/1000,
			s.RankGeral))
	}
}

func printSelectedToCli(date string, stocks []model.Stock) {
	fmt.Print(config.Yellow, "Selecting stocks -- ", date, " -- Hope you make good use of them 🔥🔥\n\n")
	fmt.Printf("%s%s %-4s | %s %-6s %-15s %-8s | %s %-8s %-8s %-8s %-8s | %s %10s | %s %-5s %s\n",
		config.BackBlack, config.White, "Rank",
		config.Cyan, "Ticker", "Empresa", "Preço",
		config.Red, "PL", "EV/Ebit", "ROA", "M Ebit",
		config.Yellow, "Liquidez",
		config.White, "Geral", config.Reset)

	for i, s := range stocks {
		fmt.Printf("%s %-4d | %s %-6s %-15.14s R$%-6.2f | %s %-8.2f %-8.2f %-8.2f %-8.2f | %s %8.f K | %s %-5d \n",
			config.White, i+1,
			config.Cyan, s.Ticker, s.Company_name, s.Price,
			config.Red, s.P_L, s.EV_Ebit, s.Roa, s.MargemEbit,
			config.Yellow, s.LiquidezMediaDiaria/1000,
			config.Reset, s.RankGeral)
	}
}

func printRemovedToCli(removedStocks []model.Stock) {
	fmt.Println("\n", config.BackBlack, config.White, "❌❌ Listing Stocks that have been removed ❌❌", config.Reset)
	fmt.Printf("\n%s%s %-6s %-15.13s %-9s | %s %-8s %-8s %-8s %-8s | %s %10s  | %s %-8s%s\n",
		config.BackBlack, config.Cyan, "Ticker", "Empresa", "Preço",
		config.Red, "PL", "EV/Ebit", "ROA", "M Ebit",
		config.Yellow, "Liquidez",
		config.White, "Mot. Excl.", config.Reset)

	for _, s := range removedStocks {
		fmt.Printf("%s %-6s %-15.14s R$%-7.2f | %s %-8.2f %-8.2f %-8.2f %-8.2f | %s %8.f K  | %s %-8s\n",
			config.Cyan, s.Ticker, s.Company_name, s.Price,
			config.Red, s.P_L, s.EV_Ebit, s.Roa, s.MargemEbit,
			config.Yellow, s.LiquidezMediaDiaria/1000,
			config.Reset, s.ExcludedReason)
	}
}
