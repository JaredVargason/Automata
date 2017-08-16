package main

import (
	"fmt"
	"path/filepath"

	"github.com/jaredvargason/cse396/models"
)

func main() {
	absPath, _ := filepath.Abs("examples/dfa.dfa")
	fmt.Printf(absPath + "\n")
	dfa, _ := models.ReadDfaFromFile(filepath.FromSlash(absPath))
	fmt.Printf("%+v\n", dfa)

	fmt.Printf("%t", dfa.AcceptsString("aabdcdacbabbabc"))
}
