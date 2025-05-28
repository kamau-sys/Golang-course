package main

import (
	"fmt"
)

func main() {

	 // PART 1: BASIC ARRAYS (1D)
	 
	fmt.Println("==complete go arrays tutorial==")
	fmt.Println("from basic arrays (1D) to 2D arrays")

	// PART 1.1: BASIC ARRAYS (1D)
	fmt.Println("1.1: declaring arrays")

	// Method 1: Declare with size, values are zero-initialized
	var numbers [5]int
	fmt.Println("Empty int array:", numbers)
	var names [3]string
	fmt.Println("Empty string array: ", names)

	//Method 2: Declare and initialize
	fruits := [3]string{"mango", "banana", "orange"}
	fmt.Printf("initialised array %v\n", fruits)
	// Method 3: let go count the elements
	colors := [...]string{"red", "yellow", "green"}
	fmt.Printf("Auto-sized array: %v\n", colors)

	//  Part 1.2 : ACCESSING AND MODIFYING ELEMENTS
	fmt.Println("Accessing and modifying elements")

	//Accessing elements (0-indexed)
	fmt.Printf("first fruit: %v\n", fruits[0])
	fmt.Println("last fruit: ", fruits[len(fruits)-1])

	//fill arrays with values
	for i := 0; i < len(numbers); i++ {
		numbers[i] = (i + 1) * 10
	}
	fmt.Println("filled number array", numbers)

	// 1.3 ARRAY PROPERTIES
	fmt.Println("**Array properties**")
	fmt.Println("length of fruits array: ", len(fruits))
	fmt.Println("length of numbers array: ", len(fruits))

	// Arrays have fixed size - this won't compile:
	// numbers[5] = 60 // This would cause "index out of range" error

	// Parts 1.4 looping thru arrays
	fmt.Println("looping thru arrays")
	// Method 1: Traditional for loop
	for i := 0; i < len(colors); i++ {
		fmt.Println("index", i, ":", colors[i])
	}

	// Method 2: Range (more idiomatic Go)
	for index,color := range colors{
		fmt.Println("index ",index,":", color)
	}

	// Method 3: Range with just values (ignore index)

	for _,fruit := range fruits{
		fmt.Println("fruit: ", fruit)
	}

	   // 1.5 ARRAY OPERATIONS
	   fmt.Println("common array operations")
	   
	   scores := [5] int {10,22,80,75,90}

	   //find highest score
	   max := scores[0]
	   for _,value := range scores{
		if value > max {
			max = value

		}
		

	   }
	   fmt.Println("max score: ", max)

	   //calculate average

	   sum := 0
	   for _,score := range scores {
		sum = sum + score
	   }

	   fmt.Println("Average is: ", float64(sum) / float64(len(scores)))

	   //search for a value
	   searchFor := 92
	   found := false
	   
	   for index,score := range scores{
		if score == searchFor {
			found = true
			fmt.Println("found ", score, "at index", index)
			break
		}
	   }
	   if !found{
		fmt.Println("could not find", searchFor)
	   }


	      // 1.6 COMPARING ARRAYS

		  arr1 := [3] int {3,6,8}
		  arr2 := [3] int {3,6,8}
		  arr3 := [3] int {3,7,8}
		  fmt.Println("arr1 == arr2 ?", arr1 == arr2)
		  fmt.Println("arr1 == arr3 ?", arr1 == arr3)


		     // PART 2: ARRAYS VS SLICES (Important Concept)
    //

	fmt.Println("\n\n Part2: Arrays vs Slices")
	//Arrays have a fixed size
	fixedArray := [3] int {1,2,3}
	fmt.Println("\nArray has a fixed size, length :", len(fixedArray))

	//slices are dynamic
	    dynamicSlice := []int{1, 2, 3}
    fmt.Printf("Slice (dynamic): %v, length: %d\n", dynamicSlice, len(dynamicSlice))




}
