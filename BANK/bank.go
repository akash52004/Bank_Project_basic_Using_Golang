package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
func main() {
	var account_balance float64 = 10000
	fmt.Println("Welcome to Go Bank")
	fmt.Println("What u want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Withdraw Money")
	fmt.Println("3. Deposite Money")
	fmt.Println("4. Exit")
	var choice int
	fmt.Print("Enter ur choice : ")
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Println("Your balance is", account_balance)
	} else if choice == 2 {
		fmt.Print("Enter the amount to deposit : ")
		var deposite float64
		fmt.Scan(&deposite)
		if deposite <= 0 {
			fmt.Println("Enter the amount greater than zero")
			return
		}

		account_balance += deposite //this is equivalent to account_balance = account_balance+deposite
		fmt.Println("Balance Updated is : ", account_balance)
		writeBalanceToFile(account_balance)
	} else if choice == 3 {
		fmt.Print("Enter the Amout to withdraw : ")
		var withdraw float64
		fmt.Scan(&withdraw)
		if withdraw <= 0 {
			fmt.Println("Enter the amount greater than zero")
			return
		}
		if withdraw > account_balance {
			fmt.Println("Insufficient Balance")
			fmt.Println("Available balance is :", account_balance)
			return
		}
		account_balance -= withdraw
		fmt.Println("Balance Updated is : ", account_balance)
	} else {
		fmt.Println("Good Bye!! ")
	}
	//fmt.Print("Choice is : ", choice)
}
