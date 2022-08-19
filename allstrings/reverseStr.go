package allstrings

import "fmt"

func CallReverseStringInPlace() {
	const str = "Ethiopia"

	revStr := reverseString1(str)

	fmt.Println(revStr)
}

func reverseString1(s string) string {
	r := []rune(s)
	return swapString(r)
}

//prepending letters
func reverseString2(s string) string {
	var r string
	for _, v := range s {
		r = string(v) + r
	}

	// return the reversed string.
	return r
}

func swapString(r []rune) string {

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
