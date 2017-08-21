package nfa

import "testing"

func TestEndsWithZeroZero(t *testing.T) {
	nfa, _ := ReadNfaFromFile("examples/endsWithZeroZero.nfa")
	t.Run("00", func(t *testing.T) {
		if !nfa.AcceptsString("00") {
			t.Fail()
		}
	})
	t.Run("10100", func(t *testing.T) {
		if !nfa.AcceptsString("10100") {
			t.Fail()
		}
	})
	t.Run("1010", func(t *testing.T) {
		if !nfa.AcceptsString("1010") {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if nfa.AcceptsString("") {
			t.Fail()
		}
	})
}
