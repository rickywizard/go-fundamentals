package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age:", *agePointer)
	
	// adultAge := getAdultYears(agePointer)
	// fmt.Println("Adult age:", adultAge)
	getAdultYears(agePointer)
	fmt.Println("Adult age:", age)
}

func getAdultYears(age *int) {
	// return *age - 18
	*age -= 18
}