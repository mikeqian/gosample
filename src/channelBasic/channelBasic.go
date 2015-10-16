package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    q := make(chan string)
    words := []string{"Enjoy", "Go", "Coding", "Patrick"}
    t := time.Now()
    for _, w := range words {
        // passing w to each goroutine to avoid repeating the same! Try not to.
        go func(w string) {
            time.Sleep(time.Duration(rand.Int63n(1e9)))
            q <- w
        }(w)
    }
    for i := 0; i < len(words); i++ {
        // reading from channel q
        fmt.Printf("%q\t\t", <-q)
        fmt.Println("Created in:", time.Now().Sub(t))
    }
}