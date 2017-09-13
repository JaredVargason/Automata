package pda

//Pda Represents a pushdown automata.
type Pda struct {
	numStates int
	inputAlphabet []string 
	stackAlphabet []string
	transition   [][][]struct{
		nextState int
		writeChar string
	}
//State X InputAlphabet X StackAlphabet -> (State, StackAlphabetChar to write) 
	startingState int
	acceptingStates []int
}

