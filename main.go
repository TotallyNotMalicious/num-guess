package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func genint() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(50)
}

func getthegoodies() (int, error) {
	var input string
	fmt.Scanln(&input)
	guess, error := strconv.Atoi(input)
	return guess, error
}

func main() {
	fmt.Println("\nTry To Guess The Number Between 1 And 50\n")
	number := genint()
	var guess int
	var err error
	for number != guess {
		fmt.Print("Your Guess: ")
		guess, err = getthegoodies()
		if err != nil {
			fmt.Println("\nYour Guess Needs To Be An Integer, Try Again\n")
		} else {
			checkguess(number, guess)
		}
	}
}

func guesslow(number int, guess int) bool {
	return number > guess
}

func guesshigh(number int, guess int) bool {
	return number < guess
}

func checkguess(number int, guess int) {
	if guesslow(number, guess) {
		fmt.Println("\nYou Guessed Too Low, Try And Guess Higher Next Time\n")
	} else if guesshigh(number, guess) {
		fmt.Println("\nYou Guessed Too High, Try And Guess Lower Next Time\n")
	} else {
		fmt.Println("\nYou Guessed Correctly, Congrats")
	}
}
