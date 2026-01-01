package main

import "fmt"

func main() {
	user := "adeola"
	age := 25

	firstname := "oyetunji "
	lastname := "taofeek"

	fullname := firstname + lastname + " ololade"

	fmt.Println(fullname)
	fmt.Println(firstname + lastname + " ololade")

	fmt.Printf("%s is %d years old\n", user, age)

	age = 45

	fmt.Printf("%s is %d years old\n", user, age)

	// basic arit
	// +
	// -
	// *
	// /

	num1 := 45 //int

	num2 := 30

	problem := 20 + 10*2/3

	fmt.Println(problem)

	a := 30
	b := 20

	a, b = b, a

	fmt.Println(a, b)

	fmt.Println(num1 + num2 - 20 + num1 - num2)

	// boolean
	// true
	// false

}
