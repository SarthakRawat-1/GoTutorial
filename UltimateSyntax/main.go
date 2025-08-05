package main

import "fmt"

// func main() {
// 	var a int // We shouldn't specify bits like float usually and let the compiler decide them for us based on the architecture of computer.
// 	var b string
// 	var c float64 // float requires precision
// 	var d bool

// 	fmt.Println()

// 	fmt.Printf("var a int \t %T [%v]\n", a, a)
// 	fmt.Printf("var b string \t %T [%v]\n", b, b)
// 	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
// 	fmt.Printf("var d bool \t %T [%v]\n", d, d)


// 	// In Go, the "zero value" refers to the default value automatically assigned to a variable when it is declared without an explicit initial value. This ensures that all variables in Go always hold a valid, defined value, preventing potential errors that could arise from uninitialized variables

// 	// Declare variable and initialize.
// 	// Using the short variable declaration operator
// 	aa := 10

// 	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)

// 	// Specify type and perform a type conversion
// 	// Go emphasizes type conversion over type casting
// 	aaa := int32(10)

// 	fmt.Printf("aaa := int32(10) \t %T [%v]\n", aaa, aaa)
// }

// In Go a string is internally represented as a two-part data structure :- 
// type stringStruct struct {
//     str unsafe.Pointer // pointer to the underlying byte array
//     len int            // length of the string in bytes
// }

// // Struct in Go, which we can use to define user defined types
// type example struct {
// 	flag bool
// 	counter int16
// 	pi float32
// }

// func main() {
// 	var e1 example
// 	fmt.Printf("%+v\n", e1)

// 	// Struct Literal
// 	e2 := example {
// 		flag: true,
// 		counter: 10,
// 		pi: 3.1415,
// 	}

// 	fmt.Println("Flag ", e2.flag)
// 	fmt.Println("Counter ", e2.counter)
// 	fmt.Println("Pi ", e2.pi)

// 	// This is an empty literal construction, not a zero construction like var. Both may not result in same.
// 	// For code consistency, this may be bad.
// 	// e2 := example{} 
// }

type example struct {
	flag bool
	counter int16
	pi float32
}

type bill struct {
	flag bool
	counter int16
	pi float32
}

type erick struct {
	flag bool
	counter int16
	pi float32
}

func main() {
	// Declare a variable of an anonymous type and init using a struct literal
	e := struct {
		flag bool
		counter int16
		pi float32
	}{
		flag: true,
		counter: 10,
		pi: 3.1415,
	}

	// Create a value of type example
	var ex example

	// Assign value of unnamed struct type to the named struct type val
	ex = e // Perfectly allowed

	// Display the values
	fmt.Printf("%+v\n", ex) 

	var br bill
	var ei erick

	br = ei // Not allowed! But then why was ex = e allowed?

	// The diff is that in Type system in Go, once a variable is based on a named type, then Go will only allow assignment to happen if the variables are of the same named type.
	// But e is based on a literal type not named type. So it is doing an implicit conversion in case of ex = e
	// In case of br = ei, no implicit conversion is being done, so we have to do it ourselves.

	br = bill(ei) // Allowed!

	// Practical Example
	// Let's say we have. Both 8 bits.
	var signedInt int
	var unsignedInt uint

	// signedInt = unsignedInt // Not allowed
	signedInt = int(unsignedInt) // Allowed!
}

// Go provides value semantics (diff piece of code operates on it's own copy of data) and pointer semantics(diff piece of code shares the one copy of data).
// Value semantics -> Higher level of isolation
// Pointer semantics -> Lesser complexity

