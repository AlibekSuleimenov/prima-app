package main

import (
	"bufio"
	"bytes"
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
		{
			name:     "prime",
			testNum:  7,
			expected: true,
			msg:      "7 is a prime number!",
		},
		{
			name:     "not a prime",
			testNum:  8,
			expected: false,
			msg:      "8 is not a prime number, because it is divisible by 2!",
		},
		{
			name:     "zero",
			testNum:  0,
			expected: false,
			msg:      "0 is not a prime number, by definition!",
		},
		{
			name:     "one",
			testNum:  1,
			expected: false,
			msg:      "1 is not a prime number, by definition!",
		},
		{
			name:     "negative number",
			testNum:  -11,
			expected: false,
			msg:      "Negative numbers are not prime, by definition!",
		},
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

	// create r/w pipe
	r, w, _ := os.Pipe()

	os.Stdout = w
	prompt()
	_ = w.Close()

	// reset os.Stdout
	os.Stdout = oldOut

	// read the output
	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> , but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create r/w pipe
	r, w, _ := os.Pipe()

	os.Stdout = w
	intro()
	_ = w.Close()

	// reset os.Stdout
	os.Stdout = oldOut

	// read the output
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "we'll tell you if it is") {
		t.Errorf("intro text not correct: got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty",
			input:    "",
			expected: "Please enter a whole number!",
		},
		{
			name:     "zero",
			input:    "0",
			expected: "0 is not a prime number, by definition!",
		},
		{
			name:     "one",
			input:    "1",
			expected: "1 is not a prime number, by definition!",
		},
		{
			name:     "two",
			input:    "2",
			expected: "2 is a prime number!",
		},
		{
			name:     "three",
			input:    "3",
			expected: "3 is a prime number!",
		},
		{
			name:     "negative",
			input:    "-1",
			expected: "Negative numbers are not prime, by definition!",
		},
		{
			name:     "typed",
			input:    "nine",
			expected: "Please enter a whole number!",
		},
		{
			name:     "decimal",
			input:    "1.1",
			expected: "Please enter a whole number!",
		},
		{
			name:     "quit",
			input:    "q",
			expected: "",
		},
		{
			name:     "QUIT",
			input:    "Q",
			expected: "",
		},
		{
			name:     "greek",
			input:    "αβγδ",
			expected: "Please enter a whole number!",
		},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)

		result, _ := checkNumbers(reader)

		if !strings.EqualFold(result, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, result)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)

}
