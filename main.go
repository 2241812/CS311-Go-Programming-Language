package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("\n Community Tools - Choose an option:")
		fmt.Println("1. Student Grade Calculator")
		fmt.Println("2. Barangay Resident Registry")
		fmt.Println("3. Water Consumption Estimator")
		fmt.Println("4. Garbage Collection Scheduler")
		fmt.Println("5. Simple Voting System")
		fmt.Println("0. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
		
		case 2:
			
		case 3:
		
		case 4:
		
		case 5:
	
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
