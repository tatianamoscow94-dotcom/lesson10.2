package main

import "fmt"

type PaymentProcessor interface {
	Process(amount float64) string
}

type CreditCard struct {
	Number string
	Name   string
}

func (c CreditCard) Process(amount float64) string {
	return fmt.Sprintf("Оплата %.2f по карте %s прошла успешно", amount, c.Number)
}

type CryptoWallet struct {
	Address  string
	Currency string
}

func (w CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf("Перевод %.2f %s на кошелек %s выполнен", amount, w.Currency, w.Address)
}

func main() {
	payments := []PaymentProcessor{
		CreditCard{Number: "22004567890", Name: "Кира Пластинина"},
		CryptoWallet{Address: "0xAB999", Currency: "BTC"},
	}

	for _, p := range payments {
		fmt.Println(p.Process(1000.50))
	}
}
