package cfg

//Cfg The object describing the formal definition of a Context-Free Grammar.
type Cfg struct {
	variables  []string
	terminals  []string
	rules      map[string][]string
	startState string
}

//CNFCfg A context-free grammar in Chomsky Normal Form.
//The main difference between Cfg and CNFCfg is
type CNFCfg struct {
	variables  []string
	terminals  []string
	rules      []CNFCfgRule
	startState string
}

//CNFCfgRule Represents a rule of the form A -> BC
type CNFCfgRule struct {
	A string
	B string
	C string
}

//AcceptsString Tests whether or not the grammar recognizes the string.
//This algorithm is known as the CYK algorithm.
//The grammar must be in Chomsky Normal Form in order to work.
func (cfg *CNFCfg) AcceptsString(testString string) bool {

	stringLength := len(testString)
	dynamicMap := make([][][]bool, stringLength)

	//Initialize my dynamic programming map
	for i := 0; i < stringLength; i++ {
		dynamicMap[i] = make([][]bool, stringLength)
		for j := 0; j < stringLength; j++ {
			dynamicMap[i][j] = make([]bool, len(cfg.variables))
		}
	}

	//Initializing the first unit productions
	for s := 0; s < stringLength; s++ {
		v := 0
		for _, outputArray := range cfg.rules {
			for _, symbol := range outputArray {
				v++
				if symbol == string(testString[s]) {
					dynamicMap[1][s][v] = true
				}
			}
		}
	}

	for l := 2; l < stringLength; l++ {
		for s := 1; s < stringLength-l+1; s++ {
			for p := 1; p < l-1; p++ {

			}
		}
	}

	return false
}

func (cfg *Cfg) isStringTestable(testString string) bool {
	for _, c := range testString {
		flag := false
		for _, t := range cfg.terminals {
			if string(c) == t {
				flag = true
			}
		}

		if !flag {
			return false
		}
	}

	return true
}
