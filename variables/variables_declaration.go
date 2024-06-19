package variables

import "fmt"

func Variable_declaration_module1(){
    // basic variable declaration.
    // uninitialized variables have zero value.
    var x int
    fmt.Printf("Printing variable with basic declaration: %v\n", x)

    // basic variable declaration, declaring multiple ones at the same time.
    var p, q int
    fmt.Printf("Printing variable with basic declaration, multiple declaration: %v %v\n ", p, q)

    // type declarations may improve clarity.
    type Celcius float64
    var temperature Celcius
    temperature = 12.66
    fmt.Printf("Printing temperature %v\n", temperature)

    // initialize by declaration.
    var a = 10
    fmt.Printf("Print variable a, by default int a=%v\n", a)

    // short variable declaration.
    // Can only be done when inside a function. Cannot do outside a function.
    b := 20
    fmt.Printf("Print variable b, by default int b=%v\n", b)
}

func Const_declaration_iota(){
    // iota is basically a way to give unique but random consts to the items.
    // starts at "1" and increments.

    type Grades int
    const (
        A Grades = iota
        B
        C
        D
        F
    )
}

func Variable_scan(){
    // Scan returns the number of tokens scanned and error if any.
    var nApples int

    fmt.Printf("Please enter a demo scan value: ")
    _, err := fmt.Scan(&nApples)
    if err != nil {
        fmt.Printf("Error reading into the user input: %v\n", err)
    }
    fmt.Printf("User input is: %v\n", nApples)
}

// when we accept "variable []type", it is accepting a slice
// when we want to accept an array, we use "variable [...]type"
func print_slice(array []int){
    // helper function to print array of ints.
    for index, element := range array {
        fmt.Printf("array[%v]=%v, ", index, element)
    }
    fmt.Println("")

}

func ArraysAndSlices(){
    // arrays are of fixed length in golang.
    // elements are accessed using subscript notation [].
    // elements are assigned to 0 by default.
    // note that []int{1,2,3} will create a slice wheres,
    // [...]int{1,2,3} will create an array.
    // use arr[:] to pass an array as a slice.

    // NOTE: slices are passed by reference.
    // Arrays are passed by values.

    // Slices are of variable length, arrays are of fixed length.
    var arr1 [5]int // this is an array.
    arr1[0] = 2 
    var arr2 = []int{1,2,3,4,5} // this is a slice. (slice literal)
    // "..." calculates the amount of array elements from the number of initializers
    var arr3 = [...]int{1,2,3,4,5,6} // this is an array.
    print_slice(arr1[:]) // passing the arr1 as a slice.
    print_slice(arr2)
    print_slice(arr3[:])
    // slice is [start:end+1] (0' indexed)
    print_slice(arr2[2:4]) // prints from arr[2] to arr[3] (0' indexed) 
    
    // slice has max capacity up-to the array from which it is derived from.
    // slice literal is by default up to the max capacity.
    slice1 := arr2 // arr2 is slice literal (from above). length=5, cap=5
    slice2 := arr3[1:3] // arr3 is an array. slice 2 is of length=2, cap=5
    fmt.Printf("slice1, length=%v, capacity=%v\n", len(slice1), cap(slice1))
    fmt.Printf("slice2, length=%v, capacity=%v\n", len(slice2), cap(slice2))
}
