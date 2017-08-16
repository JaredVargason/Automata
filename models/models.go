package models

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

// SetAlphabet Set an alphabet for the DFA.
func (dfa *Dfa) SetAlphabet(alphabet []string) {
	dfa.alphabet = alphabet
}

// SetNumStates Set the number of states the DFA has
func (dfa *Dfa) SetNumStates(num int) {
	dfa.states = num
}

// SetTransitionFunction Cartesian product of states and alphabet. QXE->Q
// M rows x N cols where M is number of states and N is alphabet size:
// With 3 states and alphabet {a,b}
// [  a b
// 0[ 0 2 ]
// 1[ 2 1 ]
// 2[ 0 1 ] ]
func (dfa *Dfa) SetTransitionFunction(transition []map[string]int) {
	dfa.transition = transition
}

// SetStartState Set the start state of the DFA
func (dfa *Dfa) SetStartState(startStart int) {
	dfa.startingState = startStart
}

// SetAcceptingStates Set the accepting states of the DFA
func (dfa *Dfa) SetAcceptingStates(acceptingStates []int) {
	dfa.acceptingStates = acceptingStates
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
	dfa.SetNumStates(numStates)

	alphabetLine, _ := reader.ReadString('\n')
	alphabet := strings.Split(alphabetLine[0:len(alphabetLine)-1], " ")
	dfa.SetAlphabet(alphabet)

	twoD := make([]map[string]int, numStates)
	for i := range twoD {
		twoD[i] = make(map[string]int)

		transitionLine, _ := reader.ReadString('\n')
		transitionLine = transitionLine[0 : len(transitionLine)-1]
		transitionStateString := strings.Split(transitionLine, " ")
		/*for j := 0; j < alphabetLength; j++ {
			transitionState, _ := strconv.Atoi(transitionStateString[j])
			twoD[i][j] = transitionState
		}*/
		for j, letter := range alphabet {
			transitionState, _ := strconv.Atoi(transitionStateString[j])
			twoD[i][letter] = transitionState
		}
	}
	dfa.SetTransitionFunction(twoD)

	startStateLine, _ := reader.ReadString('\n')
	startState, _ := strconv.Atoi(startStateLine[0 : len(startStateLine)-1])
	dfa.SetStartState(startState)

	acceptingStatesLine, _ := reader.ReadString('\n')
	acceptingStatesStringArray := strings.Split(acceptingStatesLine, " ")
	acceptingStatesIntArray := make([]int, len(acceptingStatesStringArray))
	for i, v := range acceptingStatesStringArray {
		acceptingState, _ := strconv.Atoi(v)
		acceptingStatesIntArray[i] = acceptingState
	}
	dfa.SetAcceptingStates(acceptingStatesIntArray)

	return dfa, nil
}

// AcceptsString Returns whether or not the string is accepted by the DFA.
func (dfa *Dfa) AcceptsString(testString string) bool {
	if !dfa.isStringTestable(testString) {
		return false
	}

	dfaExecution, _ := dfa.GetDFAExecution(testString)
	for dfaExecution.canStep() {
		dfaExecution.step()
	}

	for _, state := range dfa.acceptingStates {
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

// DfaExecution A tester for a string in the language
type DfaExecution struct {
	testString   string
	currentState int
	dfa          *Dfa
	modCountID   int
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
		exec.currentState = exec.dfa.transition[exec.currentState][string(inputChar)]
	}

}

//
func (exec *DfaExecution) canStep() bool {
	return len(exec.testString) > 0
}
