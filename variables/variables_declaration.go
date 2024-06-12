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