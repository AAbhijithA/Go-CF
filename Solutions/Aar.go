/*
Alex and a Rhombus : 1180A
Rating : 800
Concepts : Geometry, DP, Logic
*/

package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    dp := make([]int, n)
    if n == 1 {
        fmt.Printf("1")
    } else {
        dp[0] = 1
        dp[1] = 5
        for i := 2; i < n; i++ {
            dp[i] = 4 + (2 * dp[i - 1]) - dp[i - 2]
        }
        fmt.Printf("%d",dp[n-1])
    }
}
