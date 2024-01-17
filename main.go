package main

import (
	"fmt"
)

func main() {
	// Strings
	var nameOne string = "nameOne"
	var nameTwo = "nameTwo"
	nameThree := "nameThree"

	// fmt Package
	fmt.Println(nameOne)
	fmt.Println(nameTwo)
	fmt.Println(nameThree)

	fmt.Print(nameOne)
	fmt.Print(nameTwo)

	fmt.Printf("The first name is %v \n", nameOne)
	fmt.Printf("The first and second names are %v and %v \n", nameOne, nameTwo)
	fmt.Printf("The  Third name is %q \n", nameThree)
	fmt.Printf("The Data Type of nameOne is %T \n", nameOne)

	formatted := fmt.Sprintf("The first and second names are %v and %v \n", nameOne, nameTwo)
	fmt.Println(formatted)

	//Floats
	var perc float32 = 28.6
	var percTwo float64 = 29463443464343.63463463463463
	// Integers
	var age int = 45
	var ageTwo int8 = 21
	var maxAge = 65
	minAge := 20
	fmt.Println(age, ageTwo, perc, percTwo)
	fmt.Printf("Sorry Max age is : %v \n", maxAge)
	fmt.Printf("Sorry Max age is : %v \n", minAge)

	// Arrays
	ages := [5]int{20, 30, 40, 50, 60}
	fmt.Println(ages[1])
	fmt.Println(ages[2:])
	fmt.Println(ages[:3])
	fmt.Println(ages[2:4])
	rangeExample := ages[2:4]
	fmt.Println("RangeExample is : ", rangeExample)

	//Length (arrays / slices)
	fmt.Println("Length is : ", len(ages))

	// Slices
	altitudes := []int{20, 30, 40, 50, 60}

	altitudes = append(altitudes, 78) //or
	latitudes := append(altitudes, 98)
	coords := append(altitudes, latitudes...)

	fmt.Println(altitudes, latitudes, coords)

	// Edit single Arr or slice element
	ages[2] = 78
	altitudes[1] = 58
	fmt.Println(altitudes, ages)

	//---------------------------------//
	Strings()
	Sort()
	Conditions(altitudes)
	Loops(altitudes)

	multipleReturnTypes("charlie")
	multipleReturnTypes("charlie harper")
	fmt.Println(multipleArgsWithReturnType(30,))
	multipleArgs("charlie", 45)
}
