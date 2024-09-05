package main

import (
	"fmt"
	"log"
)

func main() {
	secretKey := "iwrupvqb"
	result, err := MineAdventCoin(secretKey, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The lowest number for secret key %s is: %d\n", secretKey, result)
}
