package rop

// Waiter receives and sorts pancake stacks
type Waiter struct{}

// ServePancakes sorts the pancakes and returns the number of times a flip was performed
// This waiter prefers to sort his pancakes from the top to bottom. Matching each pancake to
// what is beneath it. We could introduce a waiter interface and support multiple waiters
// who each sort their pancake stacks differently.
func (w Waiter) ServePancakes(pancakeStack Stack) (int, error) {
	flips := 0
	for i := 1; i < pancakeStack.Len(); i++ {

		facingSameWay, err := pancakeStack.AreFacingTheSameWay(i, i-1)
		if err != nil {
			return flips, err
		}

		if !facingSameWay {
			flips++
			pancakeStack.FlipN(i)
		}
	}

	if !pancakeStack.AreAllHappy() {
		flips++
		pancakeStack.FlipN(pancakeStack.Len())
	}

	return flips, nil
}
