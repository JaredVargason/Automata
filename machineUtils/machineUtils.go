package machineUtils

// IsStringTestable Returns whether or not all of the string's characters are in the alphabet
func IsStringTestable(alphabet []string, testString string) bool {
	for _, c := range testString {
		charInAlphabet := false

		for _, c2 := range alphabet {
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
