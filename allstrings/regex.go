package allstrings

import "regexp"

//check for duplicate
func RemoveSplChar() {
	strArr := []string{"123 55478", "(123) 55478", "(123) 55-478", "(123)-55478"}

	reg, _ := regexp.Compile("[^a-zA-z0-9]+")

	res := []string{}

	for i := 0; i < len(strArr); i++ {
		temp := reg.ReplaceAllString(strArr[i], "")
		res = append(res, temp)
	}
}
