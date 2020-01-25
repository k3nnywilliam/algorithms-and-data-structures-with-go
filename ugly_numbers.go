package main

import (
	"fmt"
)


func maxDivide(a, b int) int {

	for a%b == 0 {	
		a = a / b
	}

	return a
}


func isUgly(num int) int {

	num = maxDivide(num, 2)
	num = maxDivide(num, 3)
	num = maxDivide(num, 5)

	if num == 1 {
		return 1
	}else {
		return 0
	}
}


func getNthUgly(n int) int {

	var i int = 1
	var count int = 1

	for n > count {
		i++
		if(isUgly(i) == 1){
			count++
		}

	}

	return i
}

func main(){

	num := getNthUgly(150)
	fmt.Println("150th ugly no. is %d ",  num) //output should be 5832

} 