package main

import (
	"fmt"
)

func catalan(n uint) uint64 {

	if n <= 1 {
		return 1
	}

	 var res uint64 = 0
	 var i uint

	 for i = 0; i < n; i++ {
			res = res + catalan(i) * catalan(n-i-1)
		}

		return res
}


func main(){

	var n uint
	for n = 0; n < 10; n++ {
		fmt.Println(catalan(n))
	}

	//should return 1 1 2 5 14 42 132 429 1430 4862

}