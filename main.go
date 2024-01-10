// go run main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello! Welcome to this demo go program.")

	var name string
	var age uint // uint = unsign int = positive number

	var v1 = "variable_name_assignment1"
	v2 := "variable_name_assignment2"

	println("variable name assignment 1 ", v1)
	println("variable name assignment 2", v2)

	// User input ----------------

	fmt.Print("Please enter your name: ")
	fmt.Scan(&name) // & is the pointer specifying the memory address of that var

	fmt.Print("Please enter your age: ")
	fmt.Scan(&age)

	fmt.Println()
	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %v\n", age)

	// Array ----------------
	// Array have fixed size -> syntax = var variable_name [size]variable_type -> cant mix type
	var list = [50]string{}
	var list2 = [25]string{"Hi"}

	fmt.Println("list", list)
	fmt.Println("list2", list2)

	// Array assignment
	list[10] = name

	fmt.Printf("The whole array: %v\n", list)
	fmt.Printf("The 10th value: %v\n", list[10])
	fmt.Printf("Array type: %T\n", list)
	fmt.Printf("Array length: %v\n", len(list))

	// Slices ----------------
	// Arraylist -> array with dynamic size
	var slice = []string{}
	slice = append(slice, name)
	fmt.Printf("The whole slice: %v\n", slice)

	// Loops ----------------
	// only for loop no while/do while/for each
	ls := []string{"first", "second", "thrid", "forth"}
	for index, element := range ls{
		fmt.Printf("element %v = %v = %v\n", index, element, ls[index])
	}

	// for {
	// 	var l1 string

	// 	fmt.Print("Looping with user input: ")
	// 	fmt.Scan(&l1)
	// 	fmt.Printf("You just entered %v\n\n", l1)
	// }

	// Field ---------------- ????????
	f := "string with spaces between"
	sf := strings.Fields(f)
	fmt.Println(sf)
}