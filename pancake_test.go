package rop_test

import (
	"testing"

	"github.com/bradleyshawkins/rop"
)

func TestNewPancakeSuccess(t *testing.T) {
	tests := []struct {
		Name    string
		Value   rune
		Pancake rop.Pancake
	}{
		{Name: "Happy", Value: '+', Pancake: rop.Pancake(true)},
		{Name: "Blank", Value: '-', Pancake: rop.Pancake(false)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			pancake, err := rop.NewPancake(test.Value)
			if err != nil {
				t.Errorf("Unexpected error creating pancake. Error: %v", err)
			}
			if pancake != test.Pancake {
				t.Errorf("Invalid pancake returned. Expected: %v, Got: %v", test.Pancake, pancake)
			}
		})
	}
}

func TestNewPancakeError(t *testing.T) {
	tests := []struct {
		Name  string
		Value rune
	}{
		{Name: "Space", Value: ' '},
		{Name: "Char", Value: 'a'},
		{Name: "Number", Value: '1'},
		{Name: "Symbol", Value: '#'},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := rop.NewPancake(test.Value)
			if err != rop.ErrInvalidRune {
				t.Errorf("Should have received ErrInvalidRune. Received: %v", err)
			}
		})
	}
}

func TestFlip(t *testing.T) {
	tests := []struct {
		Name            string
		InitPancake     rop.Pancake
		ExpectedPancake rop.Pancake
	}{
		{Name: "Happy Up", InitPancake: rop.Pancake(true), ExpectedPancake: rop.Pancake(false)},
		{Name: "Blank Up", InitPancake: rop.Pancake(false), ExpectedPancake: rop.Pancake(true)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.InitPancake.Flip()
			if test.InitPancake != test.ExpectedPancake {
				t.Errorf("Pancake did not flip Expected: %v, Got: %v", test.ExpectedPancake, test.InitPancake)
			}
		})
	}
}

func TestIsHappy(t *testing.T) {
	tests := []struct {
		Name        string
		InitPancake rop.Pancake
		Expected    bool
	}{
		{Name: "Happy Up", InitPancake: rop.Pancake(true), Expected: true},
		{Name: "Blank Up", InitPancake: rop.Pancake(false), Expected: false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			isHappy := test.InitPancake.IsHappy()
			if isHappy != test.Expected {
				t.Errorf("Unexpected IsHappy result. Expected: %v, Got: %v", test.Expected, isHappy)
			}
		})
	}
}
