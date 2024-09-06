package main

import (
	"fmt"
	"log"
)

func main() {
	secretKey := "iwrupvqb"

	// search for the sx leading zeroes
	result, err := MineAdventCoin(secretKey, 6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The lowest number for secret key %s with siz leading zeros is: %d\n", secretKey, result)
}
