package nfa

import (
	"sync"
)

type state uint
type TransitionFunction func(st state, act rune) []state

// look for true results in the channel
func TrueInResult(resultChannel chan bool) bool{
	for r := range resultChannel {
		if r {
			return true
		}
	}
	return false
}

func ReachableHelper(transitions TransitionFunction, start, final state, input []rune, wg *sync.WaitGroup, resultChannel chan bool) {
	if len(input) == 0 {
		// put true in the channel if you found a path (if not, default is false)
		if start == final {
			resultChannel <- true
		}
	} else {
		// everything below is the same except wait group has to be updated as to not lock
		nextStates := transitions(start, input[0])
		for _, state := range nextStates {
			wg.Add(1)
			go ReachableHelper(transitions, state, final, input[1:len(input)], wg, resultChannel)
		}
	}
	wg.Done()
}

// this one is like the main function: call first thread and wait for return 
func Reachable(transitions TransitionFunction, start, final state, input []rune, ) bool {
	// make a waitgroup to keep track of where the return values should be generated 
	var wg sync.WaitGroup

	// make a channel of size 10,000 
	resultChannel := make(chan bool, 10000)

	wg.Add(1)

	ReachableHelper(transitions, start, final, input, &wg, resultChannel)

	wg.Wait()
	close(resultChannel)

	// if there are any trues in result channel
	return (TrueInResult(resultChannel))
}