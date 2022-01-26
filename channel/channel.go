package channel

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // channel
	people := [2]string{"Yujin", "Bomin"}

	for _, person := range people {
		go isGood(person, c) // send channel to routine
	}

	// go routine is valid only while main func works
	// time.Sleep(time.Second * 10)

	// main func will wait for channel to give feedback from routine
	// result1 := <-c
	// result2 := <-c
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	for i := 0; i < len(people); i += 1 {
		fmt.Println(<-c)
	}

	fmt.Println("==========================")
}

func isGood(person string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- person + " is good" // give feedback to channel
}
