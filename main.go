package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// main is a starting point for a Go application.
func main() {
	intro()

	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	<-doneChan

	close(doneChan)

	fmt.Println("Goodbye!")
}

// intro is a function that displays welcome messages
func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

// readUserInput is a function that reads user entered input
func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

// checkNumbers is a function that filters user input
func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check if user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// convert user input to int
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

// prompt is a function that indicates that user must enter input
func prompt() {
	fmt.Print("-> ")
}

// isPrime is a function that checks if a number is prime.
func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// check if num is prime
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number, because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
