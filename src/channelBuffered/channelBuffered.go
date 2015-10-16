package main

import (
	"fmt"
	"time"
)

func main() {
	// Case-1: no buffer
	//chanMessage := make(chan string)
	// Case-2: with buffer ... the output changes
	chanMessage := make(chan string, 2)
	count := 4
	go func() {
		for i := 1; i <= count; i++ {
			fmt.Println("send message")
			// send to chanMessage
			chanMessage <- fmt.Sprintf("message %d", i)
		}
	}()
	// Pause the main to let the goroutine sends its messages
	time.Sleep(time.Second * 2)
	for i := 1; i <= count; i++ {
		// receive from chanMessage and print
		fmt.Println(<-chanMessage)
		time.Sleep(time.Second)
	}
}
