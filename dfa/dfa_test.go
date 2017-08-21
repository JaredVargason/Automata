package dfa

import "testing"

func TestEndsWithOneDfa(t *testing.T) {
	dfa, _ := ReadDfaFromFile("examples/endsWithOne.dfa")
	t.Run("000001", func(t *testing.T) {
		if !dfa.AcceptsString("000001") {
			t.Fail()
		}
	})
	t.Run("0100", func(t *testing.T) {
		if dfa.AcceptsString("00100") {
			t.Fail()
		}
	})
	t.Run("empty string", func(t *testing.T) {
		if dfa.AcceptsString("") {
			t.Fail()
		}
	})
}

func TestStartsWithOneOneDfa(t *testing.T) {
	dfa, _ := ReadDfaFromFile("examples/startsWithOneOne.dfa")
	t.Run("1100", func(t *testing.T) {
		if !dfa.AcceptsString("1100") {
			t.Fail()
		}
	})
	t.Run("0011", func(t *testing.T) {
		if dfa.AcceptsString("0011") {
			t.Fail()
		}
	})
}
