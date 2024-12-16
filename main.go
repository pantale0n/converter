package main

import "fmt"

func main() {
	const usdToEur = 0.94
	const usdToRub = 100.0
	euro := 1000.0
	rub := euro / usdToEur * usdToRub
	fmt.Printf("Сумма в рублях: %.2f", rub)
}
