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

var exchangeRates = map[string]float64{
	USD: 1.0,
	EUR: 0.85,
	RUB: 75.50,
}

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
	if from == to {
		return amount
	}

	usdAmount := amount / exchangeRates[from]
	return usdAmount * exchangeRates[to]
}
