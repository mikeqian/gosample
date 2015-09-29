///todo: replace the flag to config.json
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const VERSION = "0.2.0"

var (
	help    = flag.Bool("h", false, "Helps")
	client  = flag.Int("c", 10, "Clients")
	seconds = flag.Int64("t", 60, "Seconds")
	url     = flag.String("url", "", "URL")
	file    = flag.String("f", "", "URL list file")
	urls    []string
)

func fetch(url string, c chan bool) {
	status := false
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if http.StatusOK == resp.StatusCode {
		status = true
	}

	c <- status
}

func readLines(filename string) (lines []string, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return
}

func main() {
	flag.Parse()

	fmt.Println(*client, "clients, run", seconds, "seconds")

	success, fail := 0, 0
	c := make(chan bool)

	endTime := time.Now().Add(time.Duration(*seconds) * time.Second)
	for endTime.Before(time.Now()) {
		for i := 0; i < *client; i++ {
			go fetch(*url, c)
		}
		status := <-c
		if status {
			success += 1
		} else {
			fail += 1
		}
	}

	total := success + fail
	fmt.Println("Total:", total)
	fmt.Println("Success:", success)
	fmt.Println("Fail:", fail)
}
