package main

import "fmt"

func main() {
	const (
		USD_TO_EUR = 0.85
		USD_TO_RUB = 75.50
	)

	const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR
}

func getUserInput() int {
	var num int
	fmt.Scan(&num)
	return num
}

func calculate(num1 int, num2, num3 float64) {}
