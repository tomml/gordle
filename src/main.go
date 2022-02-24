package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	WORD_MAX_LENGTH  = 5
	GUESS_MAX_LENGTH = 6
)

var (
	wordsList = []string{
		"CIGAR", "REBUT", "SISSY", "DWARF", "QUIET", "STAFF",
	}
)

func getuserInput() (string, error) {

	var input string = ""
	var inputerr error = nil

	fmt.Println("Enter your guess:")
	reader := bufio.NewReader(os.Stdin)
	input, erro := reader.ReadString('\n')

	if erro == nil && len(input) == 6 {
		// Input is correct
	} else {
		input = ""
		inputerr = errors.New(
			"Input is wrong, please provide a correct one.\nRemember: Your guess has to be a valid word of exactly six characters.\n")
	}

	return input, inputerr
}

func getRandomWordFromList() string {
	rand.Seed(time.Now().UnixNano())
	return wordsList[rand.Intn(len(wordsList))]
}

func printGreetingRules() {
	fmt.Println("### Hey, wanna play worlde? ###")
	fmt.Println("A perfect match is marked as 'g'")
	fmt.Println("A word in match is marked as 'y'")
	fmt.Println("       No match is marked as 'x'")
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
			if rune(guess[i]) == rune(wordtoguess[i]) {
				matched[i] = "g"
				matchedCnt++
				if matchedCnt == WORD_MAX_LENGTH {
					guessIsCorrect = true
					break
				}
				continue
			} else {
				matchedCnt = 0
				for j := 0; j < WORD_MAX_LENGTH; j++ {
					if rune(guess[i]) == rune(wordtoguess[j]) {
						matched[i] = "y"
						matchedCnt++
					}
				}
				if matchedCnt == 0 {
					matched[i] = "x"
				}
			}
			matchedCnt = 0
		}
		fmt.Printf("   Your guess: %s", guess)
		fmt.Printf("Guesses matched: %s\n", matched)
		attemptsCnt++
	}
	fmt.Printf("Word to guess: %s. GREAT!!!\n", wordtoguess)
	fmt.Printf("Your attempts: %d\n", attemptsCnt)
}
