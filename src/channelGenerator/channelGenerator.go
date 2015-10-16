package main

import (
	"fmt"
	"time"
)

func channelGenerator(msg string) <-chan string { //return a receive-only channel of strings.
	c := make(chan string)
	go func() { //we launch the goroutine from inside the generator function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg+" Chums...? ", i)
			time.Sleep(time.Duration(0.02*1e3) * time.Millisecond)
		}
	}()
	return c //return c to the caller
}
func main() {
	returnedChannel := channelGenerator("boring, isn't it, ") //function returning a channel
	joe := channelGenerator("is it Joe,")
	ann := channelGenerator("is it Ann,")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-returnedChannel)
		fmt.Printf("You say %q\n", <-joe)
		fmt.Printf("You say %q\n", <-ann)
	}
	fmt.Printf("You are too boring, I'm leaving!...\n")
}
