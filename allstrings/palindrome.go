package allstrings

import "fmt"

func checkPalindrome(s string) bool {
	// isPalindrome := true
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false //break?
		}
	}
	return true
}

func largestPalindrome() {
	str := "babjb"
	final := ""
	for i := len(str) - 1; i > 0; i-- {
		if checkPalindrome(str[:i+1]) {
			final = str[:i+1]
			break
		}
	}
	fmt.Println(final)
}

func findAllPalindrome() {
	str := "babad"
	final := []string{}
	i := 0
	for i < len(str)-1 {
	l1:
		for j := len(str) - 1; j > i; j = j - 1 {
			if checkPalindrome(str[i : j+1]) {
				final = append(final, str[i:j+1])
			}
			if j == i+1 {
				i++
				break l1
			}
		}
		// this gets executed before for loop hence moving up
		// if i == len(str)-2 {
		// 	fmt.Println("breaking final")
		// 	break
		// }
	}
	fmt.Println(final)
}

func findAllPalinSubstr() {
	str := "babad"
	res := make([]string, 0)
	for k := len(str) - 1; k > 0; k-- {
		for i, j := 0, k; j < len(str); i, j = i+1, j+1 {
			if str[i] == str[j] {
				sl := ""
				if j == len(str)-1 {
					sl = str[i:]
				} else {
					sl = str[i : j+1]
				}

				if checkPalindrome(sl) {
					res = append(res, sl)
				}
			}
		}
	}
	fmt.Println(res)
}

func findLargestPalinSubstr() {
	str := "babad"
	res := make([]string, 0)
	found := false
	for k := len(str) - 1; k > 0; k-- {
		if found {
			break
		}
		for i, j := 0, k; j < len(str); i, j = i+1, j+1 {
			if str[i] == str[j] {
				sl := ""
				if j == len(str)-1 {
					sl = str[i:]
				} else {
					sl = str[i : j+1]
				}

				if checkPalindrome(sl) {
					found = true
					res = append(res, sl)
				}
			}
		}
	}
	fmt.Println(res)
}
