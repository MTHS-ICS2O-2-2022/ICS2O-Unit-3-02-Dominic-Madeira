// Created by: Dominic M.
// Created on: March 2023
//
// This program finds the volume of a pyramid.
package main

import "fmt"

func main() {
	var length float64
	var width float64
	var height float64
	var volume float64

	// input
	fmt.Println("This program finds the volume of a pyramid.")
	fmt.Println()
	fmt.Print("Enter a length: ")
	fmt.Scanln(&length)
	fmt.Println()
	fmt.Print("Enter a width: ")
	fmt.Scanln(&width)
	fmt.Println()
	fmt.Print("Enter a height: ")
	fmt.Scanln(&height)

	// process
	volume = (length * width * height) / 3

	// output
	fmt.Println()
	fmt.Println("The area is:", volume, "mm³")

	fmt.Println("\nDone.")
}
