package main

import (
	"fmt"
)


func bellNumber(n, length int) int {
	
	bell := make([][]int, length + 1)

	for i := 0; i < length + 1; i++ {
	bell[i] = make([]int, length+ 1)
	}

	bell[0][0] = 1

	for i := 1; i <= n; i++ {
		bell[i][0] = bell[i-1][i-1]

		for j:=1; j <=i; j++ {
			bell[i][j] = bell[i-1][j-1] + bell[i][j-1]
		}
	}

	return bell[n][0]

}



func main(){
	length :=5
	for n:= 0; n <=length; n++ {
		fmt.Println("Bell number", n, "is ", bellNumber(n, length))
	}

} 