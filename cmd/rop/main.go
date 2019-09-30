package main

import (
	"fmt"
	"os"

	"github.com/bradleyshawkins/rop"
)

func main() {

	input := NewInput(os.Stdin)
	testCases, err := input.CreateTestCases()
	if err != nil {
		fmt.Println("Error received:", err)
		return
	}

	kitchen := rop.Kitchen{}
	waiter := rop.Waiter{}

	for i, testCase := range testCases {
		pancakeStack, err := kitchen.PreparePancakes(testCase)
		if err != nil {
			fmt.Printf("Error preparing pancakes. Error: %v", err)
			return
		}

		flips, err := waiter.ServePancakes(pancakeStack)
		if err != nil {
			fmt.Printf("Error serving pancakes. Error: %v", err)
			return
		}

		fmt.Printf("Test Case #%d: %d\n", i+1, flips)
	}
}
