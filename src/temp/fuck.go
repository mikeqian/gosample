package main

import (
	"fmt"
)

func main() {
	fmt.Println([]byte("aaa")[0])
	fmt.Println([]byte("ccc")[0])

	if []byte("aaa")[0] == []byte("ccc")[0] {
		fmt.Println([]byte("aaa")[0])
		fmt.Println([]byte("ccc")[0])
		fmt.Println("ok")
	} else {
		fmt.Println("在Go 1.5.1中，此荒唐BUG已修复")
	}
}
