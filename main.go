package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		inValConvert, cntConvert, outValConvert := getUserInput()
		_, _, _, err := calculateConvert(inValConvert, cntConvert, outValConvert)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Завершение работы приложения")
			break
		}

		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}
}

func getUserInput() (string, float64, string) {
	var inValConvert string
	var outValConvert string
	var cntConvert float64

	fmt.Println("Какую валюту вы хотите продать: rub, eur, usd? ")
	fmt.Scan(&inValConvert)
	fmt.Println("Укажите сумму валюты для продажи: ")
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

func calculateConvert(inValConvert string, cntConvert float64, outValConvert string) (string, float64, string, error) {
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

	if inValConvert != "rub" && inValConvert != "eur" && inValConvert != "usd" {
		return "", 0, "", errors.New("Ошибка ввода валюты продажы!")
	} else if cntConvert <= 0 {
		return "", 0, "", errors.New("Ошибка ввода суммы валюты для продажи")
	} else if outValConvert != "rub" && outValConvert != "eur" && outValConvert != "usd" {
		return "", 0, "", errors.New("Ошибка ввода валюты покупки")
	} else {
		fmt.Println("Спасибо, что воспользовались услугами нашей комании!!!!")
	}
	return inValConvert, cntConvert, outValConvert, nil
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Println("Желаете продолжить (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" {
		return true
	}
	return false
}
