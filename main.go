/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-11-20
	* This file calculates the prices of nuts, bolts, and washers 
	*/
package main

import "fmt"

func main() {
    // Assign constants
    const boltPrice = 10
    const nutPrice = 5
    const washerPrice = 3

    // Variables to store user input
    var bolts int
		var nuts int
		var washers int

    // Input prompts
    fmt.Print("Enter the number of bolts: ")
    fmt.Scanln(&bolts)

    fmt.Print("Enter the number of nuts: ")
    fmt.Scanln(&nuts)

    fmt.Print("Enter the number of washers: ")
    fmt.Scanln(&washers)

    // Calculate total cost
    var totalCost = (bolts * boltPrice) + (nuts * nutPrice) + (washers * washerPrice)

    // Check order
    if nuts < bolts {
        fmt.Println("Check the Order, not enough nuts for the bolts you purchased.")
    }
    if washers < bolts {
        fmt.Println("Check the Order, not enough washers for the bolts you purchased.")
    }
    if nuts >= bolts && washers >= bolts {
        fmt.Println("Order is OK.")
    }

    // Output total cost
    fmt.Println("Your total cost of the order is", totalCost, "cents.")

		fmt.Println("\nDone.")
}