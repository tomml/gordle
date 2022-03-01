package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"gordle/words"
)

const guessMaxLength = 6

type guessMatchedType struct {
	m byte   // one char of a word
	s string // status of guess, either g, y or x
}

func getuserInput() (string, error) {

	var input string = ""
	var inputerr error = nil

	fmt.Print("Enter your guess: ")
	reader := bufio.NewReader(os.Stdin)
	input, erro := reader.ReadString('\n')

	if erro == nil && len(input) == 6 { // includes delimiter
		// Input is correct
	} else {
		input = ""
		inputerr = errors.New(
			`input is wrong, please provide a correct one.
       Remember: Your guess has to be a valid word of exactly five characters`)
	}

	return strings.ToUpper(input), inputerr
}

func printGreetingRules() {
	fmt.Println("### Hey, wanna play worlde? ###")
	fmt.Println("A perfect match is marked" + "\033[32m " + "green" + "\033[0m")
	fmt.Println("A word in match is marked" + "\033[33m " + "yellow" + "\033[0m")
	fmt.Println("       No match is marked" + "\033[37m " + "grey" + "\033[0m")
	fmt.Println()
}

func printMatch(matched []guessMatchedType) {
	for _, item := range matched {
		if item.s == "g" {
			fmt.Print("\033[42m\033[1;30m")
		} else if item.s == "y" {
			fmt.Print("\033[43m\033[1;30m")
		} else {
			fmt.Print("\033[40m\033[1;37m")
		}
		fmt.Printf(" %s ", string(item.m))
		fmt.Print("\033[m\033[m")
	}
	fmt.Println()
}

func main() {
	guessMatched := make([]guessMatchedType, 5)
	var guessIsCorrect bool
	var matchedCnt, attemptsCnt int
	var guess string
	var err error

	printGreetingRules()

	// Get a random word from the list
	wordtoguess := words.GetRandomWordFromList()

	for x := 0; x < guessMaxLength && !guessIsCorrect; x++ {
		for {
			if guess, err = getuserInput(); err != nil {
				fmt.Println(err)
			} else {
				break // User input is valid
			}
		}

		for i := range guessMatched {
			guessMatched[i] = guessMatchedType{m: 0, s: "x"}
		}

		for i := 0; i < words.WordMaxLength; i++ {
			if byte(guess[i]) == byte(wordtoguess[i]) { // check for perfect match
				guessMatched[i].s = "g"
				guessMatched[i].m = guess[i]
				matchedCnt++
				if matchedCnt == words.WordMaxLength {
					guessIsCorrect = true
					break
				}
				continue
			} else { // check for in-word match
				matchedCnt = 0
				for j := 0; j < words.WordMaxLength; j++ {
					if byte(guess[i]) == byte(wordtoguess[j]) && guessMatched[j].s != "g" {
						guessMatched[i].s = "y"
						guessMatched[i].m = guess[i]
						matchedCnt++
					}
				}
				if matchedCnt == 0 { // check for no match
					guessMatched[i].s = "x"
					guessMatched[i].m = guess[i]
				}
			}
			matchedCnt = 0
		}
		printMatch(guessMatched)
		attemptsCnt++
	}

	if guessIsCorrect {
		fmt.Printf("Word to guess: %s. GREAT!!!\n", wordtoguess)
		fmt.Printf("Your attempts: %d\n", attemptsCnt)
	} else {
		fmt.Printf("Try next time! The word you were looking for was: %s\n", wordtoguess)
	}
}
