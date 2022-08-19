package allstrings

import "fmt"

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

func checkPalindrome(s string) bool {
	isPalindrome := true
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			isPalindrome = false //break?
		}
	}
	return isPalindrome
}
