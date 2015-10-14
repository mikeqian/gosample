package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var c = make(chan int, 3)

var body = `[{"Level":1,"AppNo":"Index","Ip":null,"Message":"TTTTTTT","CreateTime":"2015-10-14T16:22:38.0534627+08:00","Content":null}]`

func getremote(i int) {
	time.Sleep(time.Second)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://api.xxx.com/log/api/syslog/create", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := client.Do(req)

	fmt.Printf("%v\n", resp.StatusCode)
	defer resp.Body.Close()

	c <- i
}

func main() {
	var start = time.Now()

	maxNum := 5000

	for i := 0; i < maxNum; i++ {
		go getremote(i)
	}

	for i := 0; i < maxNum; i++ {
		fmt.Println(<-c)
	}

	var elapsed = time.Now().Sub(start).Seconds()
	fmt.Printf("cost: n% \n", int(elapsed))
}
