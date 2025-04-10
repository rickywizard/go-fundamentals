package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"Movie", "Game", "Sports"}

	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])

	fmt.Println(hobbies[1:3])

	// 3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5
	courseGoals := []string{"Finish course", "Mastered basics"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "Can implement Go"

	courseGoals = append(courseGoals, "Learn all details")

	fmt.Println(courseGoals)

	// 7
	products := []Product{
		{"P001", "Body Soap", 3.5},
		{"P002", "Shampoo", 4.7},
	}

	fmt.Println(products)

	products = append(products, Product{"P003", "Conditioner", 4.3})

	fmt.Println(products)
}
