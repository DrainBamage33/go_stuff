package pointers

import "fmt"

func Pointers_basic(){

	// pointers stores address to the data in memory.
	// "&" returns an address of a variable/function (is put in front of a variable)
	// "*" returns the data at an address, dereferencing. (is put in front of a pointer)
	
	var x int = 1
	var y int
	var ip *int // var ip is a pointer to int.

	ip = &x // var ip now points to the memory of x.
	y = *ip // value of y is now also 1.
	fmt.Printf("Value of x: %v, value of y: %v, value of ip: %v\n", x, y, ip)
}

func Pointers_new(){

	// new() function is also another way to declare a variable
	// new() declares the variable and returns a pointer to it.
	ptr := new(int)
	*ptr = 3
	fmt.Printf("Value of ptr: %v\n", *ptr)
}

func foo() *int {
	x := 1
	return &x
}

func Pointers_deallocation(){
	/* pointers make deallocation a bit tricky.
	   when somebody is returning pointer like function foo above,
	   one cannot simple deallocate the variables inside of it,
	   as they can still be used later. */
	var y *int
	y = foo()
	fmt.Printf("value of y which contains &x returned by foo(): %v\n", y) // will print new value each time.
}