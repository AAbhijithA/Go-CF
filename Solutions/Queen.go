/*
Queen : 1143C
Rating : 1400
Concepts : Trees, dfs, sorting, traversals
*/

package main

import (
	"fmt"
	"sort"
)

func tree_deletions(node int,Tree *[][]int, noRespect *[]int, del *[]int, cSort *[]int) {
	nrc := 0
	for _, child := range (*Tree)[node] {
		nrc += (*noRespect)[child]
		tree_deletions(child,Tree,noRespect,del,cSort)
	}
	if nrc == len((*Tree)[node]) && (*noRespect)[node] == 1 {
		(*cSort)[node] = 1
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	Tree := make([][]int, n+1)
	noRespect := make([]int, n+1)
	cSort := make([]int, 100001)
	var del []int
	var root int
	for i := 1; i <= n; i++ {
		var p, c int
		fmt.Scan(&p, &c)
		if p == -1 {
			root = i
		} else {
		    Tree[p] = append(Tree[p],i)
		}
		noRespect[i] = c
	}
	tree_deletions(root, &Tree, &noRespect,&del,&cSort)
	sort.Ints(del)
	notNull := false
	for i := 1; i < 100001; i++  {
		if cSort[i] == 1 {
		    fmt.Printf("%d ",i)
		    notNull = true
		}
	}
	if !notNull {
		fmt.Printf("%d",-1)
	}
}
