package main

import (
	"fmt"
	"sort"
	"strings"
)

var nameTwo = []string{"Hello", "coders", "world"}

func Strings() {
	// Strings
	var nameOne string = "Hello coders world"

	fmt.Println(strings.Contains(nameOne, "coders")) // True or False
	fmt.Println(strings.Index(nameOne, "co"))        // Int if found or A number > string length
	fmt.Println(strings.Index(nameOne, "ko"))        // Int if found or A number > string length
	fmt.Println(strings.ReplaceAll(nameOne, "Hello", "Hi"))
	fmt.Println(strings.Split(nameOne, " "))
	fmt.Println(strings.Join(nameTwo, " "))
	fmt.Println(strings.ToUpper(nameOne))

}

func Sort() {
	// Strings
	var textsArr = []string{"10", "40", "30"}
	var ints = []int{10, 40, 30, 50, 47, 79, 89}

	fmt.Println(sort.SearchInts(ints, 25))             // True or False
	fmt.Println(sort.SearchStrings(textsArr, "Hello")) // Int if found or A number > string length
}

func Conditions(ages []int) {

	for index, value := range ages {
		if index == 1 {
			fmt.Printf("Continuing Index is %v , with %v \n", index, value)
			continue
		}
		if index == 3 {
			fmt.Printf("Breaking Index is %v , with %v \n", index, value)
			break
		}
		fmt.Printf("Index is Finally %v , with %v \n", index, value)
	}
}

func Loops(ages []int) {
	x := 0
	for x < len(ages) {
		fmt.Printf("Age1 is %v : at pos: %v \n", ages[x], x)
		x++
	}

	for i := 0; i < len(ages); i++ {
		fmt.Printf("Age2 is %v : at pos: %v \n", ages[i], i)
	}

	for index, value := range ages {
		fmt.Printf("Age3 is %v : at pos: %v \n", value, index)
	}
}

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
}
