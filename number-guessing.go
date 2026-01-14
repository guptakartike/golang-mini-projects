package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("!! Welcome to Kartike's Number Guessing Game !!")

	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1

	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("You will be playing against Kartike. Good Luck!")
	fmt.Println("Kartike has selected a number between 1 and 100.")

	fmt.Print("Enter the maximum number of guesses which you would like to have: ")
	reader.Scan()
	maxGuess, err := strconv.Atoi(reader.Text())
	if err != nil || maxGuess <= 0 {
		fmt.Println("!! Not Valid !!")
		os.Exit(1)
	}

	attempts := 0

	for {
		fmt.Print("Enter your guess: ")
		reader.Scan()
		guess, err := strconv.Atoi(reader.Text())
		if err != nil {
			fmt.Println("!! Not a valid guess !!")
			continue
		}

		attempts++

		if guess < secretNumber {
			fmt.Println("You have guesses too low")
		} else if guess > secretNumber {
			fmt.Println("You have guessed too high")
		} else {
			fmt.Println("woohoo! you are rocking this game")
			fmt.Printf("You have used %d attempt(s)",attempts)
			break
		}

		if attempts >= maxGuess {
			fmt.Println("You have used all your attempts. The secret number was:", secretNumber)
			break
		}
	}
}
