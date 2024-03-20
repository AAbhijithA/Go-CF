package main
 
import (
	"fmt"
)
 

func main(){
    var C int
    fmt.Scan(&C)
	L := 0
	Lc := 0
	Ac := 1
	for {
	    Lc += Ac
	    if Lc > C {
	        break
	    }
	    C -= Lc
	    Ac += 1
	    L += 1
	}
	fmt.Printf("%d",L)
}
