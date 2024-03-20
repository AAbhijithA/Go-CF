/*
Badge : 1020B
Rating : 1000
Concepts : Graphs, DFS, Maps
*/

package main
 
import (
	"fmt"
)
 
func dfs(visMap *map[int]int, idx int, a *[]int) int {
    if (*visMap)[idx] == 1 {
        return idx
    }
    (*visMap)[idx] = 1
    return dfs(visMap,(*a)[idx],a);
}

func main(){
    var n int
    var a []int
    fmt.Scan(&n)
    a = append(a, -1)
    for i := 1; i <= n; i++ {
        var x int
        fmt.Scan(&x)
        a = append(a, x)
    }
    for i := 1; i <= n; i++ {
        visMap := make(map[int]int)
        culprit := dfs(&visMap,i,&a);
        fmt.Printf("%d ",culprit);
    }
}
