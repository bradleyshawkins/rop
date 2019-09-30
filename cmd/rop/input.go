package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	newline = '\n'
	happy   = '+'
	blank   = '-'
)

// Input is used to receive user input on how many / what pancake test cases to run
type Input struct {
	reader *bufio.Reader
}

// NewInput creates a new Input struct reading strings from r
func NewInput(r io.Reader) Input {
	return Input{
		reader: bufio.NewReader(r),
	}
}

// CreateTestCases creates pancake test cases based on user input
func (r *Input) CreateTestCases() ([]string, error) {
	// Retrieve test count
	count, err := r.getTestCount()
	if err != nil {
		return nil, err
	}

	// Retrieve tests
	return r.getTestCases(count)
}

// getTestCount gets the desired number of test cases to run, verifying input along the way
func (r *Input) getTestCount() (int, error) {

	var valid bool
	var t int
	var err error
GetCount:
	for !valid {
		fmt.Printf("Enter Number of Test Cases: ")

		testCount, err := r.reader.ReadString(newline)
		if err != nil {
			fmt.Println("Invalid input. Error:", err)
			continue GetCount
		}

		c := strings.TrimSuffix(testCount, "\n")

		t, err = strconv.Atoi(c)
		if err != nil {
			fmt.Println("Invalid input. Must be an integer. Entered:", testCount)
			continue GetCount
			// 1 <= t <= 100
		} else if t < 1 || t > 100 {
			fmt.Println("Invalid input. Must be between 1 - 100. Entered:", t)
			continue GetCount
		}

		valid = true
	}

	return t, err
}

// getTestCases gets n test cases and verifies that all are valid input
func (r *Input) getTestCases(n int) ([]string, error) {

	testCases := make([]string, 0)
	num := 0
	// Read number of times
GetCases:
	for len(testCases) != n {
		fmt.Printf("Enter test case %v: ", len(testCases)+1)

		testCase, err := r.reader.ReadString(newline)
		if err != nil {
			return nil, fmt.Errorf("invalid input. Error: %v", err)
			// 2 - 101 because \n is included in the string
		} else if len(testCase) < 2 || len(testCase) > 101 {
			fmt.Println("Invalid test case. Must be between 1 - 100 characters")
			continue GetCases
		}

		testCase = strings.TrimSuffix(testCase, "\n")

		//verify length and content
		for _, c := range testCase {
			if c != happy && c != blank {
				fmt.Println("Invalid test case. Invalid character. Must be + / -. Provided:", c)
				continue GetCases
			}
		}

		testCases = append(testCases, testCase)
		num++
	}

	return testCases, nil
}
