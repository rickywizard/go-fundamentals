package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
	// fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for Inflation): %.2f\n", futureValue, futureRealValue)
	// fmt.Printf(`Future Value: %.2f
	
	// Future Value (adjusted for Inflation): %.2f`, futureValue, futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	frv = fv / math.Pow(1 + inflationRate / 100, years)

	return fv, frv
	// return
}