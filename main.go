package main

import "fmt"

func main() {

	inValConvert, cntConvert, outValConvert := getUserInput()
	calculateConvert(inValConvert, cntConvert, outValConvert)
}

func getUserInput() (string, float64, string) {
	var inValConvert string
	var outValConvert string
	var cntConvert float64
	fmt.Println("Какую валюту вы хотите обменять (rub/eur/usd)? ")
	fmt.Scan(&inValConvert)
	fmt.Println("Укажите сумму обмена: ")
	fmt.Scan(&cntConvert)
	switch {
	case inValConvert == "rub":
		fmt.Println("Укажите какую валюту вы хотите купить (eur/usd): ")
		fmt.Scan(&outValConvert)
	case inValConvert == "eur":
		fmt.Println("Укажите какую валюту вы хотите купить (rub/usd): ")
		fmt.Scan(&outValConvert)
	case inValConvert == "usd":
		fmt.Println("Укажите какую валюту вы хотите купить (eur/rub): ")
		fmt.Scan(&outValConvert)
	}
	return inValConvert, cntConvert, outValConvert
}

func calculateConvert(inValConvert string, cntConvert float64, outValConvert string) {
	const usdToEur = 0.94
	const usdToRub = 100.0
	const eurToRub = 110.0

	switch {
	case inValConvert == "usd" && outValConvert == "eur":
		resultConvert := cntConvert * usdToEur
		fmt.Printf("Результат конвертации USD в EUR: %.2f ", resultConvert)
	case inValConvert == "eur" && outValConvert == "usd":
		resultConvert := cntConvert / usdToEur
		fmt.Printf("Результат конвертации EUR в USD: %.2f ", resultConvert)
	case inValConvert == "usd" && outValConvert == "rub":
		resultConvert := cntConvert * usdToRub
		fmt.Printf("Результат конвертации USD в RUB: %.2f ", resultConvert)
	case inValConvert == "rub" && outValConvert == "usd":
		resultConvert := cntConvert / usdToRub
		fmt.Printf("Результат конвертации RUB в USD: %.2f ", resultConvert)
	case inValConvert == "eur" && outValConvert == "rub":
		resultConvert := cntConvert * eurToRub
		fmt.Printf("Результат конвертации EUR в RUB: %.2f ", resultConvert)
	case inValConvert == "rub" && outValConvert == "eur":
		resultConvert := cntConvert / eurToRub
		fmt.Printf("Результат конвертации RUB в EUR: %.2f ", resultConvert)
	}
}
