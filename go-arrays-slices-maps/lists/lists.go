package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[:1])

	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99, 199.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 89.99, 20.95}
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}

// func main() {
// 	productNames := [4]string{"A book"}

// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Carpet"

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// note: sliced array only can be read to the end, so if the start slice from index 1, the 0 index cannot be read anymore
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
