package models

import (
	"errors"
)

// Dfa The model for a Deterministic Finite Automaton.
type Dfa struct {
	alphabet        []string
	states          int
	transition      [][]int
	startingState   int
	acceptingStates []int
	modCount        int
}

// DfaExecution A tester for a string in the language
type DfaExecution struct {
	testString   string
	currentState int
	dfa          *Dfa
	modCountID   int
}

func (dfa *Dfa) isStringTestable(testString string) bool {
	for _, c := range testString {
		charInAlphabet := false

		for _, c2 := range dfa.alphabet {
			if string(c) == c2 {
				charInAlphabet = true
			}
		}

		if !charInAlphabet {
			return false
		}
	}

	return true
}

func (dfa *Dfa) acceptsString(testString string) bool {
	if !dfa.isStringTestable(testString) {
		return false
	}

	dfaExecution, _ := dfa.GetDFAExecution(testString)
	for dfaExecution.canStep() {
		dfaExecution.step()
	}

	for state := range dfa.acceptingStates {
		if state == dfaExecution.currentState {
			return true
		}
	}

	return false
}

// GetDFAExecution Returns a new instance of a DFAExecution from a given string. If the string
// does not belong to the alphabet, then shit goes down.
func (dfa *Dfa) GetDFAExecution(testString string) (*DfaExecution, error) {
	if !dfa.isStringTestable(testString) {
		return nil, errors.New("test string does not belong in the DFA's alphabet")
	}

	dfaExecution := new(DfaExecution)
	dfaExecution.testString = testString
	dfaExecution.currentState = dfa.startingState
	dfaExecution.dfa = dfa
	dfaExecution.modCountID = dfa.modCount
	return dfaExecution, nil
}

// Steppy Something steppy is something that can be stepped through.
type Steppy interface {
	canStep()
	step()
}

func (exec *DfaExecution) step() {
	if exec.canStep() {
		//Grab current character off input string and decrease the length of the input string.
		//Get the next state based off the dfa transition function. Set current state to that state.
		inputChar := exec.testString[0]
		exec.testString = exec.testString[1:len(exec.testString)]
		exec.currentState = exec.dfa.transition[exec.currentState][inputChar]
	}

}

//
func (exec *DfaExecution) canStep() bool {
	return len(exec.testString) > 0
}
