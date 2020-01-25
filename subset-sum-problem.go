package main

import (
	"fmt"
	"unsafe"
)

func isSubsetSum(set [6]int, n uintptr, sum int) bool{

	if sum == 0 {
		return true
	}
	if n == 0 && sum != 0{
		return false
	}

	if (set[n-1] > sum) {
		return isSubsetSum(set, n-1, sum)
	}

	return isSubsetSum(set, n-1, sum) ||  isSubsetSum(set, n-1, sum-set[n-1])
}


func main(){

	var set = [6]int{3, 34, 4, 12, 5, 2}
	var sum int = 9
	var n uintptr = unsafe.Sizeof(set) / unsafe.Sizeof(set[0])


	if isSubsetSum(set, n, sum) == true {
		fmt.Println("Found a subset with given sum")
	}else {
		fmt.Println("No subset with given sum")
	}

}