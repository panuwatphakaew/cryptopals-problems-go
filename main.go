package main

import (
	"crypopals/problems"
	"fmt"
	"log"
)

func main() {
	input01 := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	result01, err := problems.HexToBase64(input01)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result01)
}