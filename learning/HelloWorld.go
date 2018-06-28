package main
import "fmt"
import (
"os"
"strings"
)
func tind() {
	fmt.Println("Hello World")
}
func main() {
	// tind()
	var d ="trink"
	p := &d
	*p = "krink"
	// var s string

	fmt.Printf("%T\n%T",d,p)
	// pind()
}
func pind(){
	// var s, sep string
// for i := 1; i < len(os.Args); i++{
// for targ, arg := range os.Args[1:] {
// s += sep + arg
// sep = " "
fmt.Println(strings.Join(os.Args[1:]," "))

// fmt.Println(s)
}