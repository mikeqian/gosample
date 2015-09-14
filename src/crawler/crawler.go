package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getremote() {
	resp, err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%s", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%s", body)

}

func main() {
	getremote()
}
