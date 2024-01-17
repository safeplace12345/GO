package main

import (
	"fmt"
	"strings"
	"sort"
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

func multipleArgs(name string, age int){
	output := fmt.Sprintf("My name is %v and am %v years old", name , age)
	fmt.Println(output)
}

func multipleArgsWithReturnType(r int) int{
	var diameter int = 30 * r * r
	return diameter
}

func multipleReturnTypes(name string) (string, string){
	chars := strings.Split(name, " ")
	if len(chars) > 1 {
		return chars[0][:1], chars[1][:1]
	}
	return chars[0], chars[0][:1]
}