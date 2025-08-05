package main

// func main() {
// 	var a int // We shouldn't specify bits like float usually because of portability and preciison.
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

// type example struct {
// 	flag bool
// 	counter int16
// 	pi float32
// }

// type bill struct {
// 	flag bool
// 	counter int16
// 	pi float32
// }

// type erick struct {
// 	flag bool
// 	counter int16
// 	pi float32
// }

// func main() {
// 	// Declare a variable of an anonymous type and init using a struct literal
// 	e := struct {
// 		flag bool
// 		counter int16
// 		pi float32
// 	}{
// 		flag: true,
// 		counter: 10,
// 		pi: 3.1415,
// 	}

// 	// Create a value of type example
// 	var ex example

// 	// Assign value of unnamed struct type to the named struct type val
// 	ex = e // Perfectly allowed

// 	// Display the values
// 	fmt.Printf("%+v\n", ex)

// 	var br bill
// 	var ei erick

// 	br = ei // Not allowed! But then why was ex = e allowed?

// 	// // The diff is that in Type system in Go, once a variable is based on a named type, then Go will only allow assignment to happen if the variables are of the same named type.
// 	// // But e is based on a literal type not named type. So it is doing an implicit conversion in case of ex = e
// 	// // In case of br = ei, no implicit conversion is being done, so we have to do it ourselves.

// 	br = bill(ei) // Allowed!

// 	// Practical Example
// 	// Let's say we have. Both 8 bits.
// 	var signedInt int
// 	var unsignedInt uint

// 	// signedInt = unsignedInt // Not allowed
// 	signedInt = int(unsignedInt) // Allowed!
// }

// Go provides value semantics (diff piece of code operates on it's own copy of data) and pointer semantics(diff piece of code shares the one copy of data).
// Value semantics -> Higher level of isolation
// Pointer semantics -> Lesser complexity

// Sample Program for Value and Pointer Semantics. To show the concept of pass by value and pass by reference.

// func main() {
// 	count := 10

// 	fmt.Println("count:\t value of[", count, "]\t addr of[", &count,"]")
// 	// increment(count)
// 	increment(&count) // For pointer semantics

// 	fmt.Println("count:\t value of[", count, "]\t addr of[", &count,"]")
// }

// // func increment(inc int) {
// // 	inc++
// // 	fmt.Println("inc:\t value of[", inc, "]\t addr of[", &inc,"]")
// // }

// // For pointer semantics
// func increment(inc *int) {
// 	(*inc)++
// 	fmt.Println("inc:\t value of[", *inc, "]\t addr of[", inc,"]")
// }

// func main() {
// 	// Literal type. Good if we don't want to name it because we are going to use it only here.
// 	// Variable of anonymous type set to it's zero value
// 	var e1 struct {
// 		flag bool
// 		counter int16
// 		pi float32
// 	}

// 	fmt.Printf("%+v\n", e1)

// 	// Variable of anonymous type set and init using a struct literal
// 	e2 := struct {
// 		flag bool
// 		count int16
// 		pi float32
// 	}{
// 		flag: true,
// 		count: 10,
// 		pi: 3.1415,
// 	}

// 	fmt.Printf("%+v\n", e2)
// }

// func main() {
// 	// Untyped constants
// 	// Constants of a kind can be implicitly converted by compiler
// 	const ui = 12345 // kind: integer
// 	const uf = 3.1415 // kind: floating-point

// 	// Typed constants still use constant type system but their precision is restricted
// 	const ti int = 12345 // type: int
// 	const tf float64 = 3.1415 // type: float64

// 	// Constant arithmetic supports different kinds
// 	// Kind promotion is used to determine kind in these scenarios

// 	// Variable answer will be of type float64
// 	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)

// 	// Constant third will be of kind floating-point
// 	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

// 	// Constant zero will be of kind integer
// 	const zero = 1 / 3 // KindInt(1) / KindInt(3)

//
// }

// func main() {
// 	// Writing a block so we don't have to write const over and over again, and iota works only in a block.
// 	// For every line in block, iota starts at 0 and then increments automatically by 1
// 	const(
// 		A1 = iota // 0 : Start at 0
// 		B1 = iota // 1 : Increment by 1
// 		C1 = iota // 2 : Increment by 1
// 	)

// 	fmt.Println("1: ", A1, B1, C1)

// 	// Don't have to list iota over and over again
// 	const(
// 		A2 = iota // 0 : Start at 0
// 		B2 // 1 : Increment by 1
// 		C2 // 2 : Increment by 1

// 	)
// 	// Iota always start at 0, if we want to start it at 1 :- A2 = iota + 1

// 	fmt.Println("2: ", A2, B2, C2)

// 	const(
// 		Ldate = 1 << iota // 1 : Shift 1 to left of 0
// 		Ltime // 2 : Shift 1 to left of 1
// 		Lmicroseconds // 4 : Shift 1 to left of 2
// 		Llongfile // 8 : Shift 1 to left of 3
// 		Lshortfile // 16 : Shift 1 to left of 4
// 		LUTC // 32 : Shift 1 to left of 5
// 	)

// 	fmt.Println("3: ", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
// }

// Go doesn't have built-in enum types like some other languages, but iota and custom types allow us to simulate them.