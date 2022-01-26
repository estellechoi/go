package person

import (
	"fmt"
	"strings"
)

// can return values more than one
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func LenAndUpperNaked(name string) (length int, uppercasedName string) {
	defer fmt.Println("I'm done") // executes after the func is finished

	length = len(name)
	uppercasedName = strings.ToUpper(name)
	return // = return length, uppercasedName (naked return)
}

func RepeatMe(words ...string) {
	fmt.Println(words)
}

func CanIDrink(age int) bool {
	// can declare and use var inside the if scope
	if korAge := age + 2; korAge > 18 {
		return true
	}

	return false
}

func CanIDrink2(age int) bool {
	// can declare and use var inside the switch scope
	switch korAge := age + 2; korAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
