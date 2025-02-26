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

	result02, err := problems.FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result02)
}