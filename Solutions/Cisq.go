/*
Can I Square? : 
Rating : 1000
Concepts : Binary Search, Geometry, Prefix Sums
*/

package main
 
import (
	  "fmt"
)
 
func solve() {
	  var n int
	  fmt.Scan(&n)
	  var allCubes int64 = 0
	  for i := 0; i < n; i++ {
		    var x int64
		    fmt.Scan(&x)
		    allCubes += x
	  }
	  var l int64 = 1
	  var r int64 = allCubes
	  possible := false
	  for l <= r {
		    var m int64 = (l + r) / 2
		    var q int64 = allCubes / m
		    if q == m && (allCubes % m) == 0 {
			      fmt.Println("Yes")
			      possible = true
			      break
		    } else if m > q {
			      r = m - 1
		    } else {
			      l = m + 1
		    }
  	}
  	if !possible {
  		  fmt.Println("No")
  	}
}
 
func main() {
	  var t int
	  fmt.Scan(&t)
	  for t > 0 {
		    solve()
		    t -= 1
	  }
}
 
