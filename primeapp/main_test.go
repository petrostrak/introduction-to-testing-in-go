// go test -coverprofile=coverage.out to create a coverage.out file
// go tool cover -html=coverage.out to render it to the browser as html
package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"zero", 0, false, "0 is not prime, by definition!\n"},
		{"one", 1, false, "1 is not prime, by definition!\n"},
		{"negative number", -1, false, "Negative numbers are not prime, by definition!"},
		{"prime", 7, true, "7 is a prime!\n"},
		{"not prime", 8, false, "8 is not prime because it is divisible by 2\n"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)

		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// st os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of the prompt() from read pipe
	out, _ := io.ReadAll(r)

	// perform test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_Intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to write pipe
	os.Stdout = w

	intro()

	// close writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of the intro() from read pipe
	out, _ := io.ReadAll(r)

	// perform test {
	if string(out) != "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> " {
		t.Errorf("incorrect prompt: expected Is it Prime?\n ------------\n Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> but got %s", string(out))
	}

}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "prime", input: "7", expected: "7 is a prime!\n"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime, by definition!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!\n"},
		{name: "non-prime", input: "9", expected: "9 is not prime because it is divisible by 3\n"},
		{name: "quit", input: "q", expected: ""},
	}

	for _, e := range tests {
		// emulate user input with strings.NewReader
		input := strings.NewReader(e.input)

		reader := bufio.NewScanner(input)

		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}

}
