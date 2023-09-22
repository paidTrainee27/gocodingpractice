package allstrings

import (
	"fmt"
	"strings"
)


func checkIfAnagram(){
	str,str1 := "agftqw","tagwfq"
	isAnagram := true
	for i := 0; i < len(str); i++ {
		if !strings.Contains(str1,string(str[i])){
			isAnagram = false
			break
		}
	}
	fmt.Println(isAnagram)
}

func checkPossibleAnagrams(){
	str,str1 := "agftqw","tagwfq"

	anagrams := []string{}
	for i := 0; i < len(str); i++ {
		if strings.Contains(str1,string(str[i])){
			anagrams = append(anagrams, string(str[i]))
		}
	}
	fmt.Println(anagrams)
}