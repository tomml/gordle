package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	wordsList = []string{
		"CIGAR", "REBUT", "SISSY", "DWARF", "QUIET", "STAFF",
	}
)

func main() {

	var matchCntPerfect int8
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hey, wanna play worlde?")

	for wc, w := range wordsList {
		fmt.Println("Enter your guess:")
		guess, _ := reader.ReadString('\n')
		for i, c := range w {
			if rune(guess[i]) == c {
				matchCntPerfect++
			}
		}
		fmt.Printf("Word searched: %s\n", w)
		fmt.Printf("You got a total of %d matche(s) for your guess!\n", matchCntPerfect)
		matchCntPerfect = 0
		if wc == len(wordsList) {
			break
		}
	}

}
