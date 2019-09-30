package rop

import "errors"

// Pancake is an object that represents a pancake in a stack of pancakes.
type Pancake bool

// ErrInvalidRune is returned if the value used to create a pancake is invalid
var ErrInvalidRune = errors.New("a pancake must be either '-' or '+'")

// NewPancake creates a faceup pancake if the rune is '+' or a facedown pancake if the rune is '-'
// returns an error if another error is provided
func NewPancake(faceup rune) (Pancake, error) {
	if faceup == '+' {
		return Pancake(true), nil
	} else if faceup == '-' {
		return Pancake(false), nil
	}
	return Pancake(false), ErrInvalidRune
}

// Flip flips the pancake over
func (p *Pancake) Flip() {
	temp := !(*p)
	*p = temp
}

// IsHappy returns true if the pancake is face up
func (p Pancake) IsHappy() bool {
	return bool(p)
}
