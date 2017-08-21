package nfa

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Nfa Representation of a Nondeterministic Finite Automaton (NFA)
type Nfa struct {
	states          int
	alphabet        []string //-1 denotes empty string
	transition      []map[string][]int
	startState      int
	acceptingStates []int
}

func (nfa *Nfa) isStringTestable(testString string) bool {
	for _, c := range testString {
		charInAlphabet := false

		for _, c2 := range nfa.alphabet {
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

// AcceptsString Returns whether or not the Nfa recognizes the string.
func (nfa *Nfa) AcceptsString(testString string) bool {
	if !nfa.isStringTestable(testString) {
		return false
	}
	initialExec, _ := nfa.getExecution(testString)

	return acceptsString(initialExec)
}

func acceptsString(exec Execution) bool {
	if len(exec.currentString) == 0 {
		for _, v := range exec.nfa.acceptingStates {
			if v == exec.currentState {
				return true
			}
		}

		return false
	}

	nfaTransition := exec.nfa.transition
	inputChar := string(exec.currentString[0])
	curString := exec.currentString[1:len(exec.currentString)]
	acceptsFlag := false
	for _, v := range nfaTransition[exec.currentState][inputChar] {
		//Make a new exec with changed values and do "acceptsString" on it
		newExec := Execution{currentState: v, currentString: curString, nfa: exec.nfa}
		if acceptsString(newExec) {
			acceptsFlag = true
		}
	}

	return acceptsFlag
}

// ReadNfaFromFile Reads an Nfa from a file. See examples/template.nfa for format.
func ReadNfaFromFile(filepath string) (*Nfa, error) {
	nfa := new(Nfa)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	numStatesLine, err := reader.ReadString('\n')
	numStates, err := strconv.Atoi(numStatesLine[0 : len(numStatesLine)-1])
	nfa.states = numStates

	alphabetLine, err := reader.ReadString('\n')
	alphabet := strings.Split(alphabetLine[0:len(alphabetLine)-1], " ")
	nfa.alphabet = alphabet

	numTransitionsLine, err := reader.ReadString('\n')
	numTransitions, err := strconv.Atoi(numTransitionsLine[0 : len(numTransitionsLine)-1])

	nfa.transition = make([]map[string][]int, numStates)
	for i := range nfa.transition {
		nfa.transition[i] = make(map[string][]int)
	}

	for i := 0; i < numTransitions; i++ {
		transitionLine, _ := reader.ReadString('\n')
		transitionLineSplit := strings.Split(transitionLine[0:len(transitionLine)-1], " ")
		fromState, _ := strconv.Atoi(transitionLineSplit[0])
		transitionString := transitionLineSplit[1]
		toState, _ := strconv.Atoi(transitionLineSplit[2])

		//	transition[fromState][transitionString] = append(transition[fromState][transitionString], toState)

		toStates := nfa.transition[fromState][transitionString]
		toStates = append(toStates, toState)
		nfa.transition[fromState][transitionString] = toStates
	}

	startingStateLine, err := reader.ReadString('\n')
	startingStateNum, err := strconv.Atoi(startingStateLine[0 : len(startingStateLine)-1])
	nfa.startState = startingStateNum

	acceptingStatesLine, err := reader.ReadString('\n')
	acceptingStatesSplit := strings.Split(acceptingStatesLine, " ")
	acceptingStatesIntArray := make([]int, len(acceptingStatesSplit))
	for i := 0; i < len(acceptingStatesSplit); i++ {
		acceptingStatesIntArray[i], err = strconv.Atoi(acceptingStatesSplit[i])
	}
	nfa.acceptingStates = acceptingStatesIntArray
	return nfa, nil
}

func (nfa *Nfa) getExecution(testString string) (Execution, error) {
	if !nfa.isStringTestable(testString) {
		return Execution{}, errors.New("test string does not belong in the NFA's alphabet")
	}
	exec := Execution{}
	exec.nfa = nfa
	exec.currentState = nfa.startState
	exec.currentString = testString

	return exec, nil
}

// Execution Represents an execution at a current state in the Nfa.
type Execution struct {
	currentString string
	currentState  int
	nfa           *Nfa
}
