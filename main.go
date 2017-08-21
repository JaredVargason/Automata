package main

import (
	"fmt"
	"path/filepath"

	"github.com/jaredvargason/cse396/dfa"
	"github.com/jaredvargason/cse396/nfa"
)

func main() {
	absPath, _ := filepath.Abs("dfa/examples/dfa.dfa")
	fmt.Printf(absPath + "\n")
	dfa, _ := dfa.ReadDfaFromFile(filepath.FromSlash(absPath))
	fmt.Printf("%+v\n", dfa)

	fmt.Printf("%t", dfa.AcceptsString("aabdcd"))

	absPath, _ = filepath.Abs("nfa/examples/endsWithZeroZero.nfa")
	fmt.Printf(absPath + "\n")
	nfa, _ := nfa.ReadNfaFromFile(filepath.FromSlash(absPath))
	fmt.Printf("%+v\n", nfa)

	fmt.Printf("%t", nfa.AcceptsString("0000"))
}
