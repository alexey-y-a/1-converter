package main

import (
	"fmt"
	"strings"
)

const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
)

func main() {
	fmt.Println("=== Калькулятор валют ===")

	fromCurrency := getCurrencyInput("Введите исходную валюту (USD, EUR, RUB): ")

	amount := getAmountInput("Введите сумму: ")

	toCurrency := getCurrencyInput("Введите целевую валюту (USD, EUR, RUB): ")

	result := calculate(amount, fromCurrency, toCurrency)
	fmt.Printf("\nРезультат: %.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)

}

func getCurrencyInput(prompt string) string {
	for {
		fmt.Print(prompt)
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		if currency == USD || currency == EUR || currency == RUB {
			return currency
		}
		fmt.Println("Ошибка: допустимые валюты - USD, EUR, RUB")
	}
}

func getAmountInput(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var amount float64
		_, err := fmt.Scan(&amount)
		if err == nil && amount > 0 {
			return amount
		}
		fmt.Println("Ошибка: введите положительное число")
		var discard string
		fmt.Scanln(&discard)
	}
}

func calculate(amount float64, from, to string) float64 {
	const (
		USD_TO_EUR = 0.85
		USD_TO_RUB = 75.50
		EUR_TO_RUB = 88.82
	)

	if from == to {
		return amount
	}

	switch from + "_TO_" + to {
	case "USD_TO_EUR":
		return amount * USD_TO_EUR
	case "USD_TO_RUB":
		return amount * USD_TO_RUB
	case "EUR_TO_RUB":
		return amount * EUR_TO_RUB
	case "EUR_TO_USD":
		return amount / USD_TO_EUR
	case "RUB_TO_USD":
		return amount / USD_TO_RUB
	case "RUB_TO_EUR":
		return amount / EUR_TO_RUB
	default:
		return 0
	}
}
