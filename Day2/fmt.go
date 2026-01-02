package main

import "fmt"

func main() {
	// name := "Sam"
	// age := 25

	// // It puts a space between "Sam" and 25 automatically.
	// // It also moves to the next line after printing.
	// fmt.Println("User info:", name, age)

	// fmt.Println("This is on a new line.")

	// item := "Pizza"
	// price := 20

	// // Using %v for everything
	// fmt.Printf("Item: %v, Price: %v\n", item, price)

	// name := "sam"

	// msg := fmt.Sprintf("Welcome back, %s!\n", name)
	

	// fmt.Printf(msg)

	// fmt scan

	// var name string

	// fmt.Println("What is your first name:")
	// fmt.Scan(&name)

	// fmt.Println("Welcome",name)


	var name string
    var age int

    fmt.Println("Enter your name and age:")
    
    // The program pauses here until user types and hits Enter.
    // We use &name and &age so Scan can update them.
    fmt.Scan(&name, &age)

    fmt.Printf("Nice to meet you, %s. You are %d.\n", name, age)



}
