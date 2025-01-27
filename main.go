package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Israel"
	fmt.Printf("Hello, %v\n", name)

	totalCars := 50
	fmt.Println("Our fleet consist of ", totalCars, " cars")

	//Booleans
	insuranceIncluded := true
	fmt.Println("Insurance included: ", insuranceIncluded)

	str1 := "Go Programming"
	str2 := "go programming"
	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToUpper(str2))
	fmt.Println(strings.Contains(str2, "go"))

}