/*
Vanya and Lanterns : 492B
Rating : 1200
Concepts : Sorting, 2-Pointers and Geometry 
*/

package main
 
import (
	"fmt"
	"sort"
)
 

func main(){
    var n, l int
    var a []int
    fmt.Scan(&n,&l)
    for i := 0; i < n; i++ {
        var x int
        fmt.Scan(&x)
        a = append(a, x)
    }
    sort.Ints(a)
    left := a[0]
    maxSeparation := 0
    for i := 1; i < n; i++ {
        if maxSeparation < (a[i] - left) {
            maxSeparation = a[i] - left
        }
        left = a[i]
    }
    maxRadius := float64(maxSeparation)
    maxRadius /= 2
    if maxRadius < float64(a[0]) {
        maxRadius = float64(a[0])
    }
    if maxRadius < float64(l - left) {
        maxRadius = float64(l - left)
    }
    fmt.Printf("%f",maxRadius)
}
