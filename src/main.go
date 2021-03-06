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

func getValidUserInput() (string, error) {
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

	return strings.ToUpper(strings.Trim(input, "\n")), inputerr
}

func printGreetingRules() {
	fmt.Println("### Hey, wanna play worlde? ###")
	fmt.Println("A perfect match is marked" + "\033[32m " + "green" + "\033[0m")
	fmt.Println("A word in match is marked" + "\033[33m " + "yellow" + "\033[0m")
	fmt.Println("       No match is marked" + "\033[37m " + "grey" + "\033[0m")
	fmt.Println()
}

func printResult(matched []guessMatchedType) {
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

func matchWords(guess, wordToGuess string) ([]guessMatchedType, bool) {
	guessMatched := make([]guessMatchedType, 5)
	var guessIsCorrect bool
	var matchCnt int

	// Init with guess string and mark all with x
	for i := range guessMatched {
		guessMatched[i] = guessMatchedType{m: guess[i], s: "x"}
	}

	for i := 0; i < words.WordMaxLength; i++ {
		// check for perfect match
		if guess[i] == wordToGuess[i] {
			guessMatched[i].s = "g"
			guessMatched[i].m = guess[i]
			// Remove character from word to guess to have no other match again
			wordToGuess = strings.Replace(wordToGuess, string(guess[i]), " ", 1)
			matchCnt++
			if matchCnt == words.WordMaxLength {
				guessIsCorrect = true
				break
			}
		}
	}

	// check for in-word match
	for i := 0; i < words.WordMaxLength; i++ {
		if strings.Contains(wordToGuess, string(guess[i])) && guessMatched[i].s != "g" {
			guessMatched[i].s = "y"
			guessMatched[i].m = guess[i]
			// Remove character from word to guess to have no other match again
			wordToGuess = strings.Replace(wordToGuess, string(guess[i]), " ", 1)
		}
	}

	return guessMatched, guessIsCorrect
}

func main() {
	var guess string
	var err error
	var attempts int = 1

	// Print the program greetings to the console
	printGreetingRules()
	// Get a random word from the list
	wordToGuess := words.GetRandomWordFromList()

	for {
		for {
			// Get a valid guess from the console
			if guess, err = getValidUserInput(); err != nil {
				fmt.Println(err)
			} else {
				if words.CheckGuessIsInList(guess) {
					break // user input is valid and guess is in the list
				} else {
					fmt.Println(`your guess is not in the wordlist. 
        Please provide a supported word.`)
				}
			}
		}
		// match the guess with the word to guess
		resultMatches, guessIsCorrect := matchWords(guess, wordToGuess)
		// ... and print the resulting match string
		printResult(resultMatches)

		if guessIsCorrect || attempts >= guessMaxLength {
			if guessIsCorrect {
				fmt.Printf("Word to guess: %s. GREAT!!!\n", wordToGuess)
				fmt.Printf("Your attempts: %d\n", attempts)
			} else {
				fmt.Printf("Try next time! The word you were looking for was: %s\n", wordToGuess)
			}
			break
		}
		attempts++
	}
}
