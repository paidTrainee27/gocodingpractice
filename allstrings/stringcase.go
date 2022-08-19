package allstrings

import (
	"fmt"
	"strings"
)

func UpperCaseFirstLetter() {
	str := "goSAMples.dev is the best Go bLog in the world!"
	flds := strings.Fields(str)
	finalStr := ""
	for i := 0; i < len(flds); i++ {
		fLetter := strings.ToUpper(string(flds[i][0]))
		if i != 0 && i != len(flds)-1 {
			finalStr += " "
		}
		finalStr += fLetter + string(flds[i][1:])
	}
	fmt.Println(finalStr)
}
