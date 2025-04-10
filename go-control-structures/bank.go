package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	
	if err != nil {
		fmt.Println("Error:", err)
		// panic("Cannot continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 on", randomdata.PhoneNumber())

	for {
		presentOptions()
	
		var choice int
	
		fmt.Print(">> ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}
	
			accountBalance += depositAmount
	
			fmt.Println("Balance updated! New balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
	
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}
	
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Insufficient account balance to withdraw.")
				// return
				continue
			}
	
			accountBalance -= withdrawAmount
	
			fmt.Println("Balance updated! New balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our Bank!")
			return
			// break
		default:
			fmt.Println("Choose menu 1-4 only!")
		}
	}
}