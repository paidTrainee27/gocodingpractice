package allstrings

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
//  1. Open brackets must be closed by the same type of brackets.
//  2. Open brackets must be closed in the correct order.
//
// Example 1:
// Input: s = "()"
// Output: true
//
// Example 2:
// Input: s = "()[]{()}"
// Output: true
//
// Example 3:
// Input: s = "([)]"
// Output: false

//For multiple open then closed brackets
func validStrings1() {
	//str := "([])"
	str := "()[]{()}"
	//holds all closing brakets
	prevBrac := []string{}
	//used to break from loop if bracket match fails
	valid := true
	// lookup for opening brackets with corresponding closing
	lookUp := make(map[string]string) 
	lookUp["("] = ")"
	lookUp["["] = "]"
	lookUp["{"] = "}"

	for i := 0; i < len(str); i++ {
		//if str[i] found in lookup then add corresponding closing bracket from lookup in prevBrac and CONTINUE
		// prevBrac as slice bcoz we can have multiple open brackets
		if v, ok := lookUp[string(str[i])]; ok {
			prevBrac = append(prevBrac, v)
			continue
		}
		// this statement is executed means upper IF is false
		//means str[i] is closing bracket so check with prevBrac's last item
		prevLIndex := len(prevBrac) - 1

		//if last closing bracket in prevBrac is not same as str[i]
		//then turn valid flag as false and break
		if string(str[i]) != prevBrac[prevLIndex] {
			valid = false
			break
		//else means we have a closing brack for previous open brack i.e. {}
		//so remove the last closing brack from prevBrac as match is found
		} else {
			prevBrac = prevBrac[:prevLIndex]
		}
	}
	//if valid true the prevBrac should be empty or otherwise
	fmt.Println(valid, prevBrac)
}

//for only adjacent closing brackets

func validStrings() {
	str := "()[]"
	prevBrac := ""
	valid := true
	lookUp := make(map[string]string)
	lookUp["("] = ")"
	lookUp["["] = "]"
	for i := 0; i < len(str); i++ {
		if v, ok := lookUp[string(str[i])]; ok {
			prevBrac = v
			continue
		}
		if string(str[i]) != prevBrac {
			valid = false
			break
		}
	}
	fmt.Println(valid)
}
