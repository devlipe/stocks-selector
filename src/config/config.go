package config

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	_FILTRAR_VOLUME_FINANC bool = true
	_FILTRAR_MARGEM_EBIT   bool = true
	_FILTRAR_PL            bool = true
	_FILTRAR_ROA           bool = true

	_VOL_FIN_MIN        int     = 200000
	_MARGEM_EBIT_MINIMA float64 = 0.0
	_PL_MINIMO          float64 = 1.5
	_ROA_MINIMO         float64 = 5.0

	_PESO_PL      float64 = 1.5
	_PESO_ROA     float64 = 1.0
	_PESO_EV_EBIT float64 = 2.0

	_OUTPUT string = "cli"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())

	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ConfigureApp() {
	var configuring bool = true
	var errorMessage string
	for configuring {
		printConfig(errorMessage)
		fmt.Print(": ")

		var choice string
		var newValue string
		fmt.Scanln(&choice, &newValue)
		if choice == "" {
			break
		}
		newValue = strings.TrimSpace(newValue)
		newValue = strings.ToLower(newValue)
		newValue = strings.ReplaceAll(newValue, ",", ".")
		err := changeConfig(choice, newValue)
		if err != nil {
			errorMessage = err.Error()
		} else {
			errorMessage = ""
		}
	}
	setEnvVariablesWithViper()
}

func changeConfig(choice, value string) error {
	switch choice {
	case "?":
		return showHelpMessage(value)
	case "1":
		_FILTRAR_VOLUME_FINANC = !_FILTRAR_VOLUME_FINANC
	case "2":
		_FILTRAR_MARGEM_EBIT = !_FILTRAR_MARGEM_EBIT
	case "3":
		_FILTRAR_ROA = !_FILTRAR_ROA
	case "4":
		_FILTRAR_PL = !_FILTRAR_PL
	case "5":
		temp, err := strconv.Atoi(value)
		if err == nil && temp > 0 {
			_VOL_FIN_MIN = temp
		} else {
			return errors.New("the value of Vol Min must be a integer and greater than 0")
		}
	case "6":
		temp, err := strconv.ParseFloat(value, 64)
		if err == nil {
			_MARGEM_EBIT_MINIMA = temp
		} else {
			return errors.New("the value of minimun Ebit margin must be a number")
		}
	case "7":
		temp, err := strconv.ParseFloat(value, 64)
		if err == nil {
			_ROA_MINIMO = temp
		} else {
			return errors.New("the value of minimum Roa must be a number")
		}
	case "8":
		temp, err := strconv.ParseFloat(value, 64)
		if err == nil {
			_PL_MINIMO = temp
		} else {
			return errors.New("the value of minimum P/L must be a number")
		}
	case "9":
		temp, err := strconv.ParseFloat(value, 64)
		if err == nil && temp > 0 {
			_PESO_EV_EBIT = temp
		} else {
			return errors.New("the value of Peso Ev/Ebit must be a number and greater than 0")
		}
	case "10":
		temp, err := strconv.ParseFloat(value, 64)
		if err == nil && temp > 0 {
			_PESO_ROA = temp
		} else {
			return errors.New("the value of Peso Roa must be a number and greater than 0")
		}
	case "11":
		temp, err := strconv.ParseFloat(value, 64)
		if err == nil && temp > 0 {
			_PESO_PL = temp
		} else {
			return errors.New("the value of Peso P/L must be a number and greater than 0")
		}
	case "12":
		value = strings.ToLower(value)
		if value == "cli" || value == "txt" || value == "xlsx" {
			_OUTPUT = value
		} else {
			return errors.New("the value of output must be one of the following (cli, txt, xlsx)")
		}
	default:
		return errors.New("something went wrong")
	}
	return nil
}

func showHelpMessage(value string) error {
	if val, ok := helpMessage[value]; ok {
		ClearScreen()
		fmt.Print(val)
		fmt.Print(Blue, "\nPress enter to go back\n", Reset)

		fmt.Scanln(&value)
	} else {
		return errors.New("There are no help message for config " + value)
	}
	return nil
}

func printConfig(errorMessage string) {
	ClearScreen()
	fmt.Println("We need to set the variables that the program uses to generate the report.")
	fmt.Println("Digit the config number follow by the new value ie `8 2.3` to change it, than press enter again to run!!")
	fmt.Println("You can also digit '?' followed by the config number for a help message ie `? 5`")
	fmt.Print("\nThe current configuration is:\n\n")
	fmt.Println(Ternary(_FILTRAR_VOLUME_FINANC, Green, Red), "1. Filter Fin. Volume :\t", _FILTRAR_VOLUME_FINANC)
	fmt.Println(Ternary(_FILTRAR_MARGEM_EBIT, Green, Red), "2. Filter Marg. Ebit :\t\t", _FILTRAR_MARGEM_EBIT)
	fmt.Println(Ternary(_FILTRAR_ROA, Green, Red), "3. Filter Roa :\t\t", _FILTRAR_ROA)
	fmt.Println(Ternary(_FILTRAR_PL, Green, Red), "4. Filter P/L :\t\t", _FILTRAR_PL)
	fmt.Println(Cyan)
	fmt.Println(" 5. Min Fin Volume :\t", _VOL_FIN_MIN)
	fmt.Println(" 6. Min Ebit Margin :\t", _MARGEM_EBIT_MINIMA)
	fmt.Println(" 7. Min Roa :\t\t", _ROA_MINIMO)
	fmt.Println(" 8. Min P/L :\t\t", _PL_MINIMO)
	fmt.Println(Yellow)
	fmt.Println(" 9. Peso EV/Ebit :\t", _PESO_EV_EBIT)
	fmt.Println(" 10. Peso Roa :\t\t", _PESO_ROA)
	fmt.Println(" 11. Peso P/L :\t\t", _PESO_PL)
	fmt.Println(Blue)
	fmt.Println(" 12. Output mode : ", _OUTPUT)
	if errorMessage != "" {
		fmt.Println(Red)
		fmt.Println(errorMessage)
	}
	fmt.Println(Reset)
}

func setEnvVariablesWithViper() {
	viper.Set("FILTRAR_VOLUME_FINANC", _FILTRAR_VOLUME_FINANC)
	viper.Set("FILTRAR_MARGEM_EBIT", _FILTRAR_MARGEM_EBIT)
	viper.Set("FILTRAR_PL", _FILTRAR_PL)
	viper.Set("FILTRAR_ROA", _FILTRAR_ROA)
	viper.Set("VOL_FIN_MIN", _VOL_FIN_MIN)
	viper.Set("MARGEM_EBIT_MINIMA", _MARGEM_EBIT_MINIMA)
	viper.Set("PL_MINIMO", _PL_MINIMO)
	viper.Set("ROA_MINIMO", _ROA_MINIMO)
	viper.Set("PESO_PL", _PESO_PL)
	viper.Set("PESO_ROA", _PESO_ROA)
	viper.Set("PESO_EV_EBIT", _PESO_EV_EBIT)
	viper.Set("OUTPUT", _OUTPUT)
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("Your platform is unsupported! I can't clear terminal screen, sorry for the bad experience :(")
	}
}

func HttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}

	return client
}
