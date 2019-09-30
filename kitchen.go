package rop

import "errors"

// Kitchen is used to prepare pancakes
type Kitchen struct{}

// ErrInvalidPancakeCount is returned if the number of pancakes is < 1 or > 100
var ErrInvalidPancakeCount = errors.New("at least one pancake must be in the stack")

// PreparePancakes prepares a stack of pancakes
func (k Kitchen) PreparePancakes(pancakeStr string) (Stack, error) {
	pancakeStack := NewStack()

	if pancakeStr == "" || len(pancakeStr) < 1 || len(pancakeStr) > 100 {
		return pancakeStack, ErrInvalidPancakeCount
	}

	for _, pancakeRune := range pancakeStr {
		pancake, err := NewPancake(pancakeRune)
		if err != nil {
			return nil, err
		}
		pancakeStack.AddPancake(pancake)
	}
	return pancakeStack, nil
}
