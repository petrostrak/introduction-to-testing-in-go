package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// create a channel to indicate when a user wants to quit
	done := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(os.Stdin, done)

	// block until the done gets a value
	<-done

	// close the channel
	close(done)

	// say goodbye
	fmt.Println("Goodbye!")
}

func readUserInput(in io.Reader, done chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, ok := checkNumbers(scanner)
		if ok {
			done <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// chack to see if user wants to exit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert input to an int
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(num)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!\n", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not prime because it is divisible by %d\n", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime!\n", n)
}
