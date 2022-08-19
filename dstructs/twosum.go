package dstructs

import "fmt"

func sumTarget()  {
	intArr := []int{5,7,13,9,2}
	targ := 9
	//since A+B=C,so A=C-B
	lookUp := make(map[int]int)
	for i := 0; i < len(intArr); i++ {
		op2 := targ - intArr[i]
		if _,ok :=lookUp[op2];ok{
			fmt.Println(i,lookUp[op2])
			return
		}
		lookUp[intArr[i]] = i
	}
}