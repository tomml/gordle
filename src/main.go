package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const GUESS_MAX_LENGTH = 6

func getuserInput() (string, error) {

	var input string = ""
	var inputerr error = nil

	fmt.Println("Enter your guess:")
	reader := bufio.NewReader(os.Stdin)
	input, erro := reader.ReadString('\n')

	if erro == nil && len(input) == 6 { // includes delimiter
		// Input is correct
	} else {
		input = ""
		inputerr = errors.New(
			`Input is wrong, please provide a correct one.
       Remember: Your guess has to be a valid word of exactly five characters.`)
	}

	return input, inputerr
}

func printGreetingRules() {
	fmt.Println("### Hey, wanna play worlde? ###")
	fmt.Println("A perfect match is marked as 'g'")
	fmt.Println("A word in match is marked as 'y'")
	fmt.Println("       No match is marked as 'x'")
	fmt.Println()
}

func main() {

	matched := make([]string, 5)
	var guessIsCorrect bool
	var matchedCnt, attemptsCnt int
	var guess string
	var err error

	printGreetingRules()
	// Get a random word from the list
	wordtoguess := getRandomWordFromList()

	for x := 0; x < GUESS_MAX_LENGTH && guessIsCorrect == false; x++ {
		for {
			if guess, err = getuserInput(); err != nil {
				fmt.Println(err)
			} else {
				break // User input is valid
			}
		}

		for i := 0; i < WORD_MAX_LENGTH; i++ {
			if byte(guess[i]) == byte(wordtoguess[i]) { // check for perfect match
				matched[i] = "g"
				matchedCnt++
				if matchedCnt == WORD_MAX_LENGTH {
					guessIsCorrect = true
					break
				}
				continue
			} else { // check for in-word match
				matchedCnt = 0
				for j := 0; j < WORD_MAX_LENGTH; j++ {
					if byte(guess[i]) == byte(wordtoguess[j]) {
						matched[i] = "y"
						matchedCnt++
					}
				}
				if matchedCnt == 0 { // check for no match
					matched[i] = "x"
				}
			}
			matchedCnt = 0
		}
		fmt.Printf("   Your guess: %s", guess)
		fmt.Printf("Guesses matched: %s\n\n", matched)
		attemptsCnt++
	}

	if guessIsCorrect == true {
		fmt.Printf("Word to guess: %s. GREAT!!!\n", wordtoguess)
		fmt.Printf("Your attempts: %d\n", attemptsCnt)
	} else {
		fmt.Printf("Try next time! The word you were looking for was: %s\n", wordtoguess)
	}
}
