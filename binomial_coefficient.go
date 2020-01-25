package main

import (
	"fmt"
)

func binomialCoeff(n, k int) int {

	if k == 0 || k == n{
		return 1
	}
	return binomialCoeff(n - 1, k - 1) + binomialCoeff(n - 1, k)
	
}


func main(){
	n := 5
	k := 2

	fmt.Println("Value of Coefficient C(", n,k, ") is ", binomialCoeff(n, k))

} 