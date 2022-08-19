package allstrings

import (
	"fmt"
	"strings"
)


func checkIfAnagram(){
	str,str1,str2 := "agftqw","tagwfq","qgarew"
	isAnagram := true
	for i := 0; i < len(str); i++ {
		if !strings.Contains(str1,string(str[i])){
			isAnagram = false
			break
		}
	}
	anagrams := []string{}
	for i := 0; i < len(str); i++ {
		if strings.Contains(str2,string(str[i])){
			anagrams = append(anagrams, string(str[i]))
		}
	}
	fmt.Println(isAnagram,anagrams)
}