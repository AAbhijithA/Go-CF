/*
Twins : 160A
Rating : 900
*/

package main
 
import (
	"fmt"
	"sort"
)
 

func main(){
    var n int
    var a []int
    fmt.Scan(&n)
    twinSum := 0
    Sum := 0
    for i := 0; i < n; i++ {
        var x int
        fmt.Scan(&x)
        a = append(a, x)
        twinSum += x
    }
    sort.Ints(a)
    for i := n - 1; i >= 0; i-- {
        Sum += a[i]
        twinSum -= a[i]
        if Sum > twinSum {
            fmt.Printf("%d", n - i)
            break
        }
    }
}
