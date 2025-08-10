package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func greetings(n string) {
	fmt.Println("Good Morning, ", n)
}
func severalGreatings(n []string, f func(string)) {
	for _, v := range n {

		f(v)
	}
}
func area(r float64) float64 {
	return math.Pi * r * r
}
func initials(n string) (string, string) {
	names := strings.ToUpper(n)
	name := strings.Split(names, " ")
	var initials []string
	for _, v := range name {
		initials = append(initials, v[:1])

	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

var file = "This is main file"

func updateName(n *string) {
	*n = "Joohn"
}
func main() {

	// VARIABLES
	first := "this is my first variable" //use this only in a function
	fmt.Println("The first variable using := ", first)
	var name string = "Joe"
	var age uint = 5
	var degrees int = -4
	var gpa float32 = 3.45
	var distance float64 = 456789876.5

	fmt.Println("My name is ", name, age, ", years of age")
	fmt.Print("Today the weather is chilly,", degrees, " degrees celsius \n")
	fmt.Printf("I studied %v miles away and scored %v \n", distance, gpa)
	fmt.Printf("I studied %v miles away and scored %f \n", distance, gpa)
	fmt.Printf("I studied %v miles away and scored %0.2f \n", distance, gpa)
	fmt.Printf("My name is %q, I studied %v miles away and scored %0.2f \n", name, distance, gpa)

	var statement = fmt.Sprintf("My name is %q, I studied %v miles away and scored %0.2f", name, distance, gpa)

	fmt.Print("The saved statement is: ", statement, "\n")
	fmt.Printf("The saved statement is: %q ", statement)

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")

	//ARRAYS => when you define the size
	subjects := [3]string{"english", "maths", "complex"}
	var names [3]string = [3]string{"Joe", "Jane"}
	ages := [4]int{12, 34, 56, 3}
	names[2] = "John"
	fmt.Println(names, ages)
	fmt.Println(names[0:2])
	fmt.Println(names[:2])
	fmt.Println(names[1:])

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")

	//Slices => when you do not specify array size
	marks := []int{45, 56, 76, 17}
	counties := []string{"nairobi", "nyeri", "angola", "burundi"}
	scores := append(marks, 89) //note append returns a new variable
	fmt.Println(marks, len(marks))
	fmt.Println(scores, len(scores))

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")

	//other packages
	//strings
	fmt.Println(strings.Contains(statement, "Joe"))           //strings is added to the import
	fmt.Println(strings.ReplaceAll(statement, "Joe", "John")) //does not alter original variable
	fmt.Println(statement)
	fmt.Println(strings.ToUpper(statement))
	fmt.Println(strings.Index(statement, "\"")) //prints position of ", \ is an escape character
	fmt.Println("The number of words in the statement are ", len(strings.Split(statement, " ")))

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")

	//sort
	sort.Ints(marks) //alters ariginal variable
	fmt.Println("The lowest mark was ", marks[0])

	exams := []int{65, 43, 23, 87, 54, 34, 66, 75, 66, 66, 66}
	below := sort.SearchInts(exams, 66)
	fmt.Println("Since 100 does not exist, its index is", len(exams), "To confirm ", sort.SearchInts(exams, 100)) //if you search an element with no index ie does not exist the return is one more index
	above := len(exams) - below
	fmt.Println("All students are ", len(exams))
	fmt.Println("Students bellow 66 were ", below)
	fmt.Println("Students above 66 were ", above)

	sort.Strings(counties)
	fmt.Println(counties)

	//you cannot sort an array directly first convert to slice

	people := names[:]
	fmt.Println("This is an unordered array", names) //This is truely unordered
	sort.Strings(people)
	fmt.Println("This is an ordered array", names)  //This will be ordered which is strange
	fmt.Println("This is an ordered slice", people) //note the strange act between the slice and array convertion and sorting
	//this is because No copying is done. Both people and names point to the same data in memory.

	//If you want to maintain independence
	units := append([]string(nil), subjects[:]...) // makes a copy
	fmt.Println("Original subjects array are ", subjects)
	fmt.Println("Unordered subjects slice are ", units)
	sort.Strings(units)
	fmt.Println("Ordered subjects slice are ", units)

	//Loop
	x := 0

	for x <= 5 {

		fmt.Println("Iteration number ", x)
		x++
	}

	for index, values := range ages {
		fmt.Printf("The county at index %v is %v \n", index, values)
	}

	for _, values := range ages { //to not show the index
		if values == 12 {
			fmt.Printf("The number is  %v \n", values)
		} else if values > 12 {
			continue
		} else {
			fmt.Println("Number below is ", values)
		}

	}

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")

	//Functions
	severalGreatings([]string{"Joe", "Jane"}, greetings)
	area := area(7)
	fmt.Println(area)

	initial1, initial2 := initials("Josphat Waweru")
	fmt.Println("The initials of your name are ", initial1, initial2)

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")

	//from other file provided same package name
	for _, v := range points {
		fmt.Println(v)
	}
	sayHello("Joe")

	fmt.Print("\n")

	//Maps => like dictionaries in python but can store only one data type
	//specify the type of key and value type also
	//maps are unordered collections — the iteration order is not guaranteed and can change between runs.
	//Go intentionally randomizes the iteration order for maps starting from Go 1.12, as a safety measure to make sure developers never rely on any specific order.

	menu := map[string]float32{
		"ugali": 45.89,
		"beans": 34.76,
		"rice":  78.43,
	}

	for i, v := range menu {
		fmt.Println(i, ":", v)
	}
	fmt.Println("Price for rice is ", menu["rice"])

	pupils := map[int]string{
		1: "Joe",
		2: "Elis",
		3: "Neema",
	}
	pupils[3] = "Wanjiku"
	fmt.Println(pupils)

	fmt.Println("\n")

	//Go is a pass by value language meaning the varibale passed to a function is just a copy of the variable with its own separate loaction in memory hence any modifications do not cahnge the original variable, unless the changes are returned by the function. This applies to int,string,struct,array,booleans and floats

	//functions, slices and maps are pass by reffrence meaning the pointer to them is the one passed meaning any change to them in the function applies

	//Pointers are data types in themselvs and when created are stored in a memory block of their own pointing to another memory loaction

	//To get the memory address of a variable use &

	address := &name
	fmt.Println("The memory address of may name is : ", address)
	fmt.Printf("The value stored at %v  is : %v \n", address, *address)

	fmt.Println("My name is : ", name)
	fmt.Println("Updating my name using pointer")
	updateName(address)
	fmt.Println("My name is : ", name)

	fmt.Println("\n")
	//Struct
	//creating a new struct instance
	c := employee{
		name: "John",
		deadline: map[string]float64{
			"skip":  32,
			"code":  24,
			"sleep": 8,
		},
		salary: 540,
	}
	c.addTask("go", 6)

	fmt.Println("salary befor update is ", c.salary)
	c.incSalary(567)
	fmt.Println("Updated salary is: ", c.salary)
	fmt.Println(c.task())

	fmt.Println(c)
	fmt.Println(makeEmployee("Joe"))

}
