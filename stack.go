package rop

import (
	"errors"
)

// Stack is a stack of pancakes
type Stack []Pancake

// ErrOutOfBounds is returned if the parameter provided isn't valid
var ErrOutOfBounds = errors.New("n must be over 0 and less than stack height")

// NewStack creates a new stack that can be used to stack pancakes
func NewStack() Stack {
	return make(Stack, 0)
}

// Len returns the number of pancakes in the pancake stack
func (s Stack) Len() int {
	return len(s)
}

// PeekTop returns the top pancake without removing it from the pancake stack
func (s Stack) PeekTop() Pancake {
	return s[s.Len()-1]
}

// AddPancake adds a pancake to the top of the pancake stack
func (s *Stack) AddPancake(pancake Pancake) {
	*s = append(*s, pancake)
}

// AreFacingTheSameWay compares the two pancakes at indices i and j in the stack
// returns true if the pancakes are facing the same way
func (s Stack) AreFacingTheSameWay(i, j int) (bool, error) {
	if !s.isInBounds(i) || !s.isInBounds(j) {
		return false, ErrOutOfBounds
	}
	return s[i] == s[j], nil
}

// FlipN flips over the top n pancakes of the pancake stack
func (s Stack) FlipN(n int) error {
	// if !s.isInBounds(n) {
	// 	return ErrOutOfBounds
	// }

	// Pull top n pancakes off into a temp stack in reverse order
	tempStack := make(Stack, 0)
	for i := n - 1; i >= 0; i-- {
		tempStack = append(tempStack, s[i])
	}

	// Insert pancakes back into the stack
	for i, pancake := range tempStack {
		s[i] = pancake
		s[i].Flip()
	}
	return nil
}

// AreAllHappy checks to see if all pancakes in the stack are happy side up
func (s Stack) AreAllHappy() bool {
	isHappy := true

	for _, pancake := range s {
		isHappy = pancake.IsHappy() && isHappy
	}
	return isHappy
}

// isInBounds checks to ensure the index exists in the pancake stack
func (s Stack) isInBounds(i int) bool {
	return i < s.Len() && i >= 0
}
