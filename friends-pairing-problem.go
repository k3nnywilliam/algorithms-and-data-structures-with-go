package main

import (
	"fmt"
)

func countFriendsPairings(n int) int {

	dp := make([]int, n + 1)

	for i := 0; i <= n; i++ {
		if i <= 2 {
			dp[i] = i
		}else {
			dp[i] = dp[i - 1] + (i - 1) * dp[i - 2]; 
		}
	}
	return dp[n]
}

func main() {

	n := 4
	fmt.Println(countFriendsPairings(n))

} 