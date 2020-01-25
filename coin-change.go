package main

import (

	"fmt"
	"unsafe"
)

func count(S [3]int, m uintptr, n int) int {

	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	if m <= 0 && n >= 1 {
		return 0
	}

	return count(S, m -1, n) + count(S, m, n-S[m - 1])
}

func main(){

	var arr = [3]int{1, 2, 3}
	var m uintptr = unsafe.Sizeof(arr) / unsafe.Sizeof(arr[0])

	fmt.Println(count(arr, m, 4))
} 