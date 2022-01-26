package main // to compile

import (
	"fmt"

	"github.com/estellechoi/go/calc"
	"github.com/estellechoi/go/person"
	"github.com/estellechoi/go/something"
)

func main() {
	// methods
	fmt.Println("Hello World") // Upper case function means it's exported
	something.SayHello()       // only uppercased methods are exported

	// constant
	const name string = "Estelle"

	// variable
	nick := "Yujin"  // var nick string = "Yujin"
	isFemale := true // var isFemale bool = true

	fmt.Println(name)
	fmt.Println(nick)
	fmt.Println(isFemale)
	fmt.Println(calc.Multiply(2, 2))

	// funcs can return multiple values
	nameLen1, upperName1 := person.LenAndUpper("Yujin")
	fmt.Println(nameLen1, upperName1)

	nameLen2, _ := person.LenAndUpper("Bomin") // _ is not a variable, it ignores value
	fmt.Println(nameLen2)

	nameLen3, upperName3 := person.LenAndUpperNaked("Nico")
	fmt.Println(nameLen3, upperName3)

	// funcs can get multiple params through like ...string
	person.RepeatMe("Yujin", "Bomin") // [Yujin Bomin]

	total := calc.Add(1, 2, 23, 4, 4, 5, 5)
	fmt.Println(total)

	canDrink := person.CanIDrink(18)
	fmt.Println(canDrink)

	// pointers
	a := 2
	b := &a // b got the memory address of a
	a = 5
	*b = 20 // can change the value of a through b pointing a
	fmt.Println(a, b)
	fmt.Println(*b) // access the value of the address b is pointing (the value of a)

	// array
	names := [5]string{"Yujin", "Bomin"} // string array with length 5
	fmt.Println(names)

	// slice
	numbers := []int{1, 2, 3}    // slice is array without length
	numbers = append(numbers, 4) // append does not change the original slice
	fmt.Println(numbers)

	// map
	data := map[string]string{"name": "Yujin", "age": "29"}
	fmt.Println(data)

	for key, value := range data {
		fmt.Println(key, value)
	}

	// struct does not have constructor built-in
	type person struct {
		name  string
		age   int
		items []string
	}

	// yujin := person{"Yujin", 29, []string{"ring"}}
	yujin := person{name: "Yujin", age: 29, items: []string{"ring"}}
	fmt.Println(yujin)
}
