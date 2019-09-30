package rop_test

import (
	"reflect"
	"testing"

	"github.com/bradleyshawkins/rop"
)

func TestAddPancakeNewPancakesOnTop(t *testing.T) {
	stack := rop.NewStack()

	stack.AddPancake(rop.Pancake(true))
	if stack.Len() != 1 {
		t.Errorf("Pancake was not successfully added to stack. Expected length to be %v but was %v", 1, stack.Len())
	}
	pancake1 := stack.PeekTop()
	if !pancake1.IsHappy() {
		t.Errorf("Just added a happy pancake to the top of the stack, but found an unhappy one")
	}

	stack.AddPancake(rop.Pancake(false))
	pancake2 := stack.PeekTop()
	if pancake2.IsHappy() {
		t.Errorf("Just added an unhappy pancake to the top of the stack, but found a happy one")
	}
}

func TestFlipN(t *testing.T) {
	tests := []struct {
		Name     string
		Pancakes []bool
		FlipN    int
		Expected []bool
	}{
		{Name: "-", Pancakes: []bool{false}, FlipN: 1, Expected: []bool{true}},
		{Name: "-+", Pancakes: []bool{false, true}, FlipN: 1, Expected: []bool{true, true}},
		{Name: "-+-", Pancakes: []bool{false, true, false}, FlipN: 2, Expected: []bool{false, true, false}},
		{Name: "-+-", Pancakes: []bool{false, true, false}, FlipN: 3, Expected: []bool{true, false, true}},
		{Name: "-+--", Pancakes: []bool{false, true, false, false}, FlipN: 3, Expected: []bool{true, false, true, false}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			testStack := buildStack(test.Pancakes...)
			expected := buildStack(test.Expected...)
			err := testStack.FlipN(test.FlipN)
			if err != nil {
				t.Errorf("Unexpected error. Error: %v", err)
			}

			if !reflect.DeepEqual(testStack, expected) {
				t.Errorf("unexpected flip result. Expected: %v, Got: %v", expected, testStack)
			}
		})
	}
}

func TestAreFacingTheSameWay(t *testing.T) {
	tests := []struct {
		Name     string
		Pancakes []bool
		i        int
		j        int
		Expected bool
	}{
		{Name: "-+", Pancakes: []bool{false, true}, i: 0, j: 1, Expected: false},
		{Name: "-+-", Pancakes: []bool{false, true, false}, i: 0, j: 1, Expected: false},
		{Name: "---", Pancakes: []bool{false, false, false}, i: 0, j: 1, Expected: true},
		{Name: "-+--", Pancakes: []bool{false, true, false, false}, i: 0, j: 2, Expected: true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			testStack := buildStack(test.Pancakes...)

			facingSameWay, err := testStack.AreFacingTheSameWay(test.i, test.j)
			if err != nil {
				t.Errorf("Unexpected error. Error: %v", err)
			}

			if facingSameWay != test.Expected {
				t.Errorf("unexpected result. Expected: %v, Got: %v", test.Expected, facingSameWay)
			}
		})
	}
}

func TestAreFacingTheSameWayError(t *testing.T) {
	tests := []struct {
		Name     string
		Pancakes []bool
		i        int
		j        int
	}{
		{Name: "Out of bounds - Above", Pancakes: []bool{false, true}, i: 1, j: 3},
		{Name: "Out of bounds - Below", Pancakes: []bool{false, true, false}, i: -1, j: 1},
		{Name: "Out of bounds - Both", Pancakes: []bool{false, false, false}, i: -1, j: 3},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			testStack := buildStack(test.Pancakes...)

			_, err := testStack.AreFacingTheSameWay(test.i, test.j)
			if err == nil {
				t.Error("Expected an error but didn't get one.")
			}
		})
	}
}

func TestAreAllHappy(t *testing.T) {
	tests := []struct {
		Name     string
		Pancakes []bool
		Expected bool
	}{
		{Name: "All Happy", Pancakes: []bool{true, true, true}, Expected: true},
		{Name: "One Unhappy", Pancakes: []bool{true, true, false}, Expected: false},
		{Name: "All Unhappy", Pancakes: []bool{false, false, false}, Expected: false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			stack := buildStack(test.Pancakes...)
			areHappy := stack.AreAllHappy()
			if areHappy != test.Expected {
				t.Errorf("Expected %v, Got: %v", test.Expected, areHappy)
			}
		})
	}
}

func buildStack(pancakeBools ...bool) rop.Stack {
	stack := rop.Stack{}
	for _, pancake := range pancakeBools {
		stack.AddPancake(rop.Pancake(pancake))
	}
	return stack
}
