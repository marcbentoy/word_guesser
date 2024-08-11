package main

import (
	"fmt"

	"github.com/marcbentoy/wordguesser/utils"
)

func main() {
	fmt.Println("Word Guesser made with Genetic Algorithms")

	phrase := utils.RandomChars(10)

	fmt.Println(phrase)
}
