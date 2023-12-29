package stocks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/spf13/viper"
	"github.com/stocks-selector/src/model"
)

func HandleSelectStocks(client *http.Client) ([]model.Stock, []model.Stock) {

	stockList, err := getStocks(client)

	if err != nil {
		log.Fatalln(err.Error())
	}
	var removedStocks []model.Stock
	stockList, removedStocks = filterStocks(stockList)
	rankStocks(stockList)

	return stockList, removedStocks
}

func rankStocks(stocks []model.Stock) {
	rankPl(stocks)
	rankRoa(stocks)
	rankEvEbit(stocks)
	rankGeral(stocks)
}

func filterStocks(stockList []model.Stock) ([]model.Stock, []model.Stock) {
	var removedStocks []model.Stock
	var r []model.Stock
	if viper.GetBool("FILTRAR_VOLUME_FINANC") {
		stockList, r = filtrarVolFin(stockList)
		removedStocks = append(removedStocks, r...)
	}
	if viper.GetBool("FILTRAR_MARGEM_EBIT") {
		stockList, r = filtrarMargemEbit(stockList)
		removedStocks = append(removedStocks, r...)
	}
	if viper.GetBool("FILTRAR_PL") {
		stockList, r = filtrarPl(stockList)
		removedStocks = append(removedStocks, r...)
	}
	if viper.GetBool("FILTRAR_ROA") {
		stockList, r = filtarRoa(stockList)
		removedStocks = append(removedStocks, r...)
	}

	sort.Slice(removedStocks, func(i, j int) bool {
		return removedStocks[i].Company_name < removedStocks[j].Company_name
	})
	return stockList, removedStocks
}

func rankGeral(stocks []model.Stock) {
	for i, s := range stocks {
		stocks[i].RankGeral = s.RankEvEbit + s.RankPl + s.RankRoa
	}

	sort.Slice(stocks, func(i, j int) bool {
		return stocks[i].RankGeral > stocks[j].RankGeral
	})
}

func rankEvEbit(stocks []model.Stock) {
	sort.Slice(stocks, func(i, j int) bool {
		return stocks[i].EV_Ebit < stocks[j].EV_Ebit
	})
	sz := len(stocks)
	pesoEvEbit := viper.GetFloat64("PESO_EV_EBIT")

	for i := range stocks {
		stocks[i].RankEvEbit = int(float64(sz-i) * pesoEvEbit)
	}
}

func rankRoa(stocks []model.Stock) {
	sort.Slice(stocks, func(i, j int) bool {
		return stocks[i].Roa < stocks[j].Roa
	})
	pesoRoa := viper.GetFloat64("PESO_ROA")

	for i := range stocks {
		stocks[i].RankRoa = int(float64(i) * pesoRoa)
	}
}

func rankPl(stocks []model.Stock) {
	sort.Slice(stocks, func(i, j int) bool {
		return stocks[i].P_L < stocks[j].P_L
	})
	sz := len(stocks)
	pesoPl := viper.GetFloat64("PESO_PL")

	for i := range stocks {
		stocks[i].RankPl = int(float64(sz-i) * pesoPl)
	}
}

func filtarRoa(stocks []model.Stock) ([]model.Stock, []model.Stock) {
	filteredStocks := []model.Stock{}
	removedStocks := []model.Stock{}
	roaMin := viper.GetFloat64("ROA_MINIMO")

	for _, s := range stocks {
		if s.Roa > roaMin {
			filteredStocks = append(filteredStocks, s)
		} else {
			s.ExcludedReason += "ROA"
			removedStocks = append(removedStocks, s)
		}
	}
	return filteredStocks, removedStocks
}

func filtrarPl(stocks []model.Stock) ([]model.Stock, []model.Stock) {
	filteredStocks := []model.Stock{}
	removedStocks := []model.Stock{}
	plMin := viper.GetFloat64("PL_MINIMO")

	for _, s := range stocks {
		if s.P_L > plMin {
			filteredStocks = append(filteredStocks, s)
		} else {
			s.ExcludedReason += "PL"
			removedStocks = append(removedStocks, s)
		}
	}
	return filteredStocks, removedStocks
}

func filtrarVolFin(stocks []model.Stock) ([]model.Stock, []model.Stock) {
	filteredStocks := []model.Stock{}
	removedStocks := []model.Stock{}
	liqMin := viper.GetFloat64("VOL_FIN_MIN")

	for _, s := range stocks {
		if s.LiquidezMediaDiaria > liqMin {
			filteredStocks = append(filteredStocks, s)
		} else {
			s.ExcludedReason += "Vol Fin"
			removedStocks = append(removedStocks, s)
		}

	}
	return filteredStocks, removedStocks
}

func filtrarMargemEbit(stocks []model.Stock) ([]model.Stock, []model.Stock) {
	filteredStocks := []model.Stock{}
	removedStocks := []model.Stock{}

	margMin := viper.GetFloat64("MARGEM_EBIT_MINIMA")

	for _, s := range stocks {
		if s.MargemEbit > margMin {
			filteredStocks = append(filteredStocks, s)
		} else {
			s.ExcludedReason += "Margem Ebit"
			removedStocks = append(removedStocks, s)
		}
	}
	return filteredStocks, removedStocks
}

func getStocks(client *http.Client) ([]model.Stock, error) {
	url := "https://statusinvest.com.br/category/advancedsearchresultpaginated?search={}&orderColumn=&isAsc=&page=0&take=10000&CategoryType=1"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Golang_Stoks_Selector/1.0")

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer res.Body.Close()
	// Create a map to hold the JSON response
	var result map[string]interface{}

	// Create a JSON decoder
	decoder := json.NewDecoder(res.Body)

	// Decode the JSON response into the map
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	// Extract the "list" attribute
	list, ok := result["list"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("could not convert list to []interface{}")
	}

	// Now you can decode the "list" into your stocksList
	var stocksList []model.Stock
	for _, item := range list {
		stock, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("could not convert item to map[string]interface{}")
		}
		// Convert map to JSON
		stockJson, err := json.Marshal(stock)
		if err != nil {
			return nil, err
		}
		// Convert JSON to Stock
		var modelStock model.Stock
		err = json.Unmarshal(stockJson, &modelStock)
		if err != nil {
			return nil, err
		}
		stocksList = append(stocksList, modelStock)
	}

	return stocksList, nil
}
