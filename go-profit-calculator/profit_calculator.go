package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	earningBeforeTax, earningAfterTax, ratio := calculateEarnings(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax: %.2f\n", earningBeforeTax)
	fmt.Printf("Earning after tax: %.2f\n", earningAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
	
	writeEarningsToFile(earningBeforeTax, earningAfterTax, ratio)
}

func getUserInput(prompt string) (float64, error) {
	var input float64

	fmt.Print(prompt)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return input, nil
}

func calculateEarnings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses

	earningAfterTax := earningBeforeTax * (1 - taxRate / 100)

	ratio := earningBeforeTax / earningAfterTax

	return earningBeforeTax, earningAfterTax, ratio
}

func writeEarningsToFile(earningBeforeTax, earningAfterTax, ratio float64) {
	earningsText := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", earningBeforeTax, earningAfterTax, ratio)
	os.WriteFile("earnings.txt", []byte(earningsText), 0644)
}