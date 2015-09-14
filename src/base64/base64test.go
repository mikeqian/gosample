package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	input := "hello world2"
	enbyte := base64Encode([]byte(input))

	debyte, err := base64Decode(enbyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if input != string(debyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
	fmt.Println(string(debyte))
}
