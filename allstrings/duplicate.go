package allstrings

import (
	"fmt"
	"strings"
)

//Assumes no special character
//**Removing duplicate characters from entire string
func removeDuplicateFromString() {
	str := "rayaaz"
	var res string
	for i := 0; i < len(str); i++ {
		if !strings.Contains(res, string(str[i])) {
			res += fmt.Sprintf("%c", str[i])
		}
	}
	fmt.Println(res) //rayz
}

func removeDuplicateWithSpecialCharacter() {
	str := "rayaa♄z"
	rn := []rune(str) //here
	var res string
	for i := 0; i < len(rn); i++ {
		if !strings.Contains(res, string(rn[i])) {
			res += fmt.Sprintf("%c", rn[i])
		}

	}
	fmt.Println(res) //ray♄z
}

//Assumes no special character
//**Removing adjacent duplicate characters from string
func removeAdjacentDuplicateFromString() {
	str := "rayaazr"
	var res string
	for i := 0; i < len(str); i++ {
		if len(res) > 0 {
			if res[len(res)-1] != str[i] {
				res += fmt.Sprintf("%c", str[i])
			}
		} else {
			res += fmt.Sprintf("%c", str[i])
		}

	}
	fmt.Println(res) //rayaz
}
