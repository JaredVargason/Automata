package dfa

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Dfa The model for a Deterministic Finite Automaton.
type Dfa struct {
	alphabet        []string
	states          int
	transition      []map[string]int
	startingState   int
	acceptingStates []int
	modCount        int
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

// ReadDfaFromFile Reads a file and returns a Dfa from it
// In format:
// #num states
// 3
// #alphabet
// a b c d
// #transition function
// 0 1 2 1
// 1 1 2 0
// 2 0 1 1
// #start state
// 0
// #accepting states
// 1
func ReadDfaFromFile(filename string) (*Dfa, error) {
	dfa := new(Dfa)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	numStatesLine, _ := reader.ReadString('\n')
	numStates, _ := strconv.Atoi(numStatesLine[0 : len(numStatesLine)-1])
	dfa.states = numStates

	alphabetLine, _ := reader.ReadString('\n')
	alphabet := strings.Split(alphabetLine[0:len(alphabetLine)-1], " ")
	dfa.alphabet = alphabet

	twoD := make([]map[string]int, numStates)
	for i := range twoD {
		twoD[i] = make(map[string]int)

		transitionLine, _ := reader.ReadString('\n')
		transitionLine = transitionLine[0 : len(transitionLine)-1]
		transitionStateString := strings.Split(transitionLine, " ")

		for j, letter := range alphabet {
			transitionState, _ := strconv.Atoi(transitionStateString[j])
			twoD[i][letter] = transitionState
		}
	}
	dfa.transition = twoD

	startStateLine, _ := reader.ReadString('\n')
	startState, _ := strconv.Atoi(startStateLine[0 : len(startStateLine)-1])
	dfa.startingState = startState

	acceptingStatesLine, _ := reader.ReadString('\n')
	acceptingStatesStringArray := strings.Split(acceptingStatesLine, " ")
	acceptingStatesIntArray := make([]int, len(acceptingStatesStringArray))
	for i, v := range acceptingStatesStringArray {
		acceptingState, _ := strconv.Atoi(v)
		acceptingStatesIntArray[i] = acceptingState
	}

	dfa.acceptingStates = acceptingStatesIntArray

	return dfa, nil
}

// AcceptsString Returns whether or not the string is accepted by the DFA.
func (dfa *Dfa) AcceptsString(testString string) bool {
	if !dfa.isStringTestable(testString) {
		return false
	}

	dfaExecution, _ := dfa.GetExecution(testString)
	for dfaExecution.CanStep() {
		dfaExecution.Step()
	}

	for _, state := range dfa.acceptingStates {
		if state == dfaExecution.currentState {
			return true
		}
	}

	return false
}

// GetExecution Returns a new instance of a Dfa.Execution from a given string. If the string
// does not belong to the alphabet, then shit goes down.
func (dfa *Dfa) GetExecution(testString string) (*Execution, error) {
	if !dfa.isStringTestable(testString) {
		return nil, errors.New("test string does not belong in the DFA's alphabet")
	}

	dfaExecution := new(Execution)
	dfaExecution.testString = testString
	dfaExecution.currentState = dfa.startingState
	dfaExecution.dfa = dfa
	dfaExecution.modCountID = dfa.modCount
	return dfaExecution, nil
}

// Execution A tester for a string in the language
type Execution struct {
	testString   string
	currentState int
	dfa          *Dfa
	modCountID   int
}

// Machiney Something machiney is something cool
type Machiney interface {
	CanStep()
	Step()
	GetExecution(testString string)
	AcceptsString(testString string)
}

// Step One step through a Dfa
func (exec *Execution) Step() {
	if exec.CanStep() {
		//Grab current character off input string and decrease the length of the input string.
		//Get the next state based off the dfa transition function. Set current state to that state.
		inputChar := exec.testString[0]
		exec.testString = exec.testString[1:len(exec.testString)]
		exec.currentState = exec.dfa.transition[exec.currentState][string(inputChar)]
	}

}

// CanStep Returns whether or not stepping is even possible.
func (exec *Execution) CanStep() bool {
	return len(exec.testString) > 0
}
