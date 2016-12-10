package main
import (
	"fmt"
	"runtime"
)
func main(){
	s:="world"
	fmt.Printf(s+"\n")
	fmt.Printf("%T\n",s)
	fmt.Printf("%d",runtime.NumCPU())
}
