package main

import (
	"fmt"

	"github.com/DrainBamage33/go_stuff/pointers"
	"github.com/DrainBamage33/go_stuff/variables"
)

func main() {
	fmt.Printf("=======go_stuff: Start=======\n")
	fmt.Println("\nRelated to variable declarations")
	fmt.Printf("============Start============\n")
	fmt.Printf("About: basics about variable declarations\n")
	variables.Variable_declaration_module1()
	fmt.Printf("=============End=============\n")
	fmt.Println("\nRelated to pointers")
	fmt.Printf("============Start============\n")
	fmt.Printf("About: Basics about pointers\n")
	pointers.Pointers_basic()
	fmt.Printf("=============End=============\n")
	fmt.Printf("============Start============\n")
	fmt.Printf("About: Basics about new() function\n")
	pointers.Pointers_new()
	fmt.Printf("=============End=============\n")
	fmt.Printf("============Start============\n")
	fmt.Printf("About: About dellocation and pointer functions\n")
	pointers.Pointers_deallocation()
	fmt.Printf("=============End=============\n")
	fmt.Printf("============Start============\n")
	fmt.Printf("About: About variable scan\n")
	variables.Variable_scan()
	fmt.Printf("=============End=============\n")
	fmt.Printf("============Start============\n")
	fmt.Printf("About: About dellocation and pointer functions\n")
	variables.ArraysAndSlices()
	fmt.Printf("=============End=============\n")

}