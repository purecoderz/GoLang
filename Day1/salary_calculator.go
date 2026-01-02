package main

import "fmt"

func main() {
	name := "Adeola"
	hourlyRate := 20
	hoursworked := 5
	salary := hourlyRate * hoursworked

	fmt.Printf("%s worked for %d hours, therefore his salary is $%d\n", name,hoursworked,salary)
	// fmt.Println(name,"worked for",hoursworked,"hours, therefore his salary is $",salary)

}

