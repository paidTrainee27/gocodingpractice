package allstrings

import (
	"fmt"
	"strings"
)

//without using sort function
func sortAlphas() {
	intStr := "eWotQiDa"   //immutable
	str1 := []rune(intStr) //mutable to avaoid error on line 18
	//expectedd aDeioQtW

	for i := 0; i < len(str1); i++ {
		for j := i + 1; j < len(str1); j++ {
			if strings.ToLower(string(str1[i])) > strings.ToLower(string(str1[j])) {
				str1[i], str1[j] = str1[j], str1[i] //swpping
			}
		}

	}

	fmt.Println(string(str1))
}
