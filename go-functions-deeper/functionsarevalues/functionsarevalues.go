package functionsarevalues

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	transformFunc1 := getTransformerFunction(&numbers)
	transformFunc2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFunc1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFunc2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
