package main

import (
	"fmt"
	"net/http"
	"time"
)

var c = make(chan int, 3)

func getremote(i int) {
	time.Sleep(time.Second)

	res, err := http.Get("http://www.cnblogs.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", res.StatusCode)
		defer res.Body.Close()
	}

	c <- i
}

func main() {
	var start = time.Now()

	maxNum := 1000

	for i := 0; i < maxNum; i++ {
		go getremote(i)
	}

	for i := 0; i < maxNum; i++ {
		fmt.Println(<-c)
	}

	var elapsed = time.Now().Sub(start).Seconds()
	fmt.Printf("cost: n% \n", int(elapsed))
}
