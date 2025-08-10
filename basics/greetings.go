// to share code between two files they need to share same package and also run both files go run main.go greetings.go
package main

import "fmt"

var points = []int{12, 3, 2, 4, 76}

// The fmt package itself is imported at the package level and can be used throughout the program after import.
//  However, specific fmt functions like fmt.Println cannot be used as standalone statements outside of a function body. The Go language requires that non-declaration statements, such as function calls, be contained within a function.
//  For example, placing fmt.Println("Hello") directly at the top level of a source file will result in a syntax error: "non-declaration statement outside function body".

// While fmt.Println cannot stand alone outside a function, it can be used within a variable declaration at the top level. For instance, var n, err = fmt.Println("I can do it") is valid because it is part of a variable declaration.
//
//	This allows the function call to execute during package initialization, which occurs before the main function runs.
//	The fmt package's functions are also commonly used within the main function, which is the entry point of the program.

func sayHello(n string) {
	fmt.Println("hello", n, ", I am from another file")
	fmt.Println("This is another county", file) //declared in another file
}
