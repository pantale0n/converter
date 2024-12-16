package main

import "fmt"

func main() {
	const usdToEur = 0.94
	const usdToRub = 100.0
	euro := getUserInput()
	rub := euro / usdToEur * usdToRub
	fmt.Printf("Сумма в рублях: %.2f", rub)
}

func getUserInput() float64 {
	var euro float64
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&euro)
	return euro
}

func convert(usNum float64, val1 string, val2 string) float64 {

}
