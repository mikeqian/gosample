package main

import (
	"fmt"
	"os"
)

func main() {
	dir := os.Getenv("gopath")
	fmt.Println(dir)
}
