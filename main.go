/*
Connor Sanders
02/23/2019
Guess The Number v0.1
 */


package main


import (
	"fmt"
	"math/rand"
	"time"
)


// Function called in main to start playing the game
func PlayGuessingGame() {
	gg := GuessingGame {1, "", 0}
	gg.Greet()
	gg.ChooseNumber()
	gg.Play()
}


// Function that generate a random int32
func RandomNumber(min int32, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(max - min) + 1
}


// Define the GuessingGame struct and constants
type GuessingGame struct {
	Guesses      int
	Name         string
	Number       int32
}

const GuessName = "Guesser"
const GuessingTries = 6
const MaxNumber = 19


// GuessingGame method that determines the number that the user will try to guess
func (gg *GuessingGame) ChooseNumber() {
	if gg.Number == 0 {
		gg.Number = RandomNumber(0, MaxNumber)
	}
}


// GuessingGame method that handles the initial greeting
func (gg *GuessingGame) Greet() {
	fmt.Println("Hello! What is your Name?\n ")
	_, err := fmt.Scanf("%s", &gg.Name)
	if err != nil {
		fmt.Println(err)
	}
	if gg.Name == "" {
		gg.Name = GuessName
	}
	fmt.Printf("Well, %s, I am thinking of a number between 1 and 20.\n", gg.Name)
	fmt.Printf("You have %d tries to guess the number.\n", GuessingTries - 1)
}


// GuessingGame method that handles the users entered guess
func (gg *GuessingGame) makeGuess() (bool) {
	var guess int32
	guess = gg.askUser(guess)
	if guess < gg.Number {
		fmt.Printf("Your %d is too low.\n", guess)
	}
	if guess > gg.Number {
		fmt.Printf("Your %d is too high.\n", guess)
	}
	if guess == gg.Number {
		fmt.Printf("Good job, %s! You guessed my number in %d guesses!", gg.Name, gg.Guesses)
		return false
	}
	return true
}


// GuessingGame method that prompts the user to input a guess
func (gg *GuessingGame) askUser(guess int32) int32 {
	fmt.Println("\nTake a guess !!")
	_, err := fmt.Scanf("%d", &guess)
	if err != nil {
		fmt.Println(err)
	}
	return guess
}


// GuessingGame method that beging the Guess the Number game
func (gg *GuessingGame) Play() {
	for gg.Guesses < GuessingTries && gg.makeGuess() {
		gg.Guesses++
	}
	if gg.Guesses == GuessingTries {
		fmt.Printf("Sorry. The number I had in mind was %d\n", gg.Number)
	} else {
		fmt.Println("\nYou won!!")
	}
}


// Main Method
func main() {
	PlayGuessingGame()
}
