package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var words = [8]string{"apple", "grape", "melon", "peach", "berry", "orange", "lemon", "kiwi"}

// Function that returns a random word from the words array
func randomWord() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	return words[randomIndex]
}

// Function to check which letters of the guessed word are correct and return their positions
func checkLetters(word string, guess string) string {
	var positions []string
	for i := 0; i < len(word) && i < len(guess); i++ {
		if word[i] == guess[i] {

			positions = append(positions, strconv.Itoa(i+1))
		}
	}

	return fmt.Sprint(positions)
}

func main() {

	var guess string
	word := randomWord()
	attempts := 0

	fmt.Println("Welcome to the word guessing game!")
	fmt.Println("Try to guess a 5-letter word.")
	for {
		attempts++
		fmt.Println("Guess the word:")
		fmt.Scanln(&guess)
		if guess == word {
			fmt.Println("Congratulations! You guessed the word in", attempts, "attempts.")
			break
		} else {
			correctPositions := checkLetters(word, guess)
			fmt.Println("Correct letters in positions:", correctPositions)
			fmt.Println("Try again!")
		}
	}
}
