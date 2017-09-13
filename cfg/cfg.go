package cfg

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Cfg The object describing the formal definition of a Context-Free Grammar.
type Cfg struct {
	variables []string
	terminals []string
	rules     map[string][][]string
	startRule string
}

//CNFCfg A context-free grammar in Chomsky Normal Form.
//The main difference between Cfg and CNFCfg is
type CNFCfg struct {
	variables []string
	terminals []string
	rules     []CNFCfgRule
	startRule string
}

//CNFCfgRule Represents a rule of the form A -> BC
type CNFCfgRule struct {
	A string
	B string
	C string
}

func ReadCFGFromFile(filepath string) (*Cfg, error) {
	cfg := new(Cfg)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	variablesString, err := reader.ReadString('\n')
	variablesString = strings.TrimSpace(variablesString)
	variablesArray := strings.Split(variablesString, " ")
	cfg.variables = variablesArray

	terminalsString, err := reader.ReadString('\n')
	terminalsString = strings.TrimSpace(terminalsString)
	terminalsArray := strings.Split(terminalsString, " ")
	cfg.terminals = terminalsArray

	numRulesString, err := reader.ReadString('\n')
	numRulesString = strings.TrimSpace(numRulesString)
	numRules, err := strconv.Atoi(numRulesString)

	cfg.rules = make(map[string][][]string)

	for i := 0; i < numRules; i++ {
		ruleString, _ := reader.ReadString('\n')
		ruleString = strings.TrimSpace(ruleString)
		rule := strings.Split(ruleString, " ")
		val, ok := cfg.rules[rule[0]]
		if ok {
			cfg.rules[rule[0]] = append(val, rule[1:len(rule)])
		} else {
			cfg.rules[rule[0]] = append([][]string{}, rule[1:len(rule)])
		}
	}

	startingRuleString, err := reader.ReadString('\n')
	startingRuleString = strings.TrimSpace(startingRuleString)
	cfg.startRule = startingRuleString

	return cfg, nil
}

//AcceptsString Tests whether or not the grammar recognizes the string.
//This algorithm is known as the CYK algorithm.
//The grammar must be in Chomsky Normal Form in order to work.
func (cfg *CNFCfg) AcceptsString(testString string) bool {
	/*
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
	*/
	return false
}

func (cfg *Cfg) toCNF() *CNFCfg {
	//1. Add a new start variable.
	//2. Eliminate empty string rules
	//3. Remove unit rules.
	//4. Take care of rules with more than two terminals or variables.

	return nil
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

func (cfg *Cfg) generateStringList(depth int) []string {

	//for rule := range cfg.rules:

	return make([]string, 0)
}
