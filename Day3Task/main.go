package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the number guessing game!")

	// Generate a random number between 1 and 6
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(6) + 1
	//fmt.Println("Hint: The random number is:", randomNumber) // For debugging purposes, remove in production

	// Ask for user input
	fmt.Println("I have generated a random number between 1 and 6. Try to guess it!")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput) // Trim newline and extra spaces

	// Convert user input to an integer
	userGuess, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Invalid input! Please enter a valid number.")
		return
	}

	// Check if the input is within the valid range
	if userGuess < 1 || userGuess > 6 {
		fmt.Println("Out of range! Please guess a number between 1 and 6.")
		return
	}

	// Compare the user's guess with the random number
	if userGuess == randomNumber {
		fmt.Println("Guessed successfully!")
	} else {
		fmt.Println("Failed to guess. The correct number was:", randomNumber)
	}
}
