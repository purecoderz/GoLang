package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Band Name Generator.\nWhat's the name of the city you grew up in?")
	var city string
	fmt.Scan(&city)
	fmt.Println("What's your pet's name?")
	var pet string
	fmt.Scan(&pet)

	fmt.Printf("Your band name could be %s %s\n",city,pet)
}
