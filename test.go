package main
import(
"fmt"
"bufio"
"os"
// "math"
// "bytes"
)
func main(){

	seen := make(map[string]bool) // a set of strings
	// input := bufio.NewScanner(os.Stdin)
	input := bufio.NewReader(os.Stdin)
	// for input.Scan() {
	// line := input.Text()
	for{
		line, n, err := input.ReadRune()
		fmt.Println(line,n, err, seen)
// 	if !seen[line]{
// 	seen[line] = true
// 	fmt.Println(line)
// 	}
// 	if line=="" {
// fmt.Fprintf(os.Stderr, "dedup: %v\n")
os.Exit(1)

}


//     // var u string = "my na$e is mogambo!"
//     // fmt.Printf("%h",u[5]) 
//     s := "abc"
//     fmt.Printf(s) 
// 	b := []byte(s)
// 	for _,d := range b{
// 		fmt.Printf("%f",d)
// 	}
// 	s2 := string(b)
// 	fmt.Printf(s2)
// var s = make([]int,10)
// var s =[]rune("hello bello")
// fmt.Println(len(s), cap(s), s)
// fmt.Println(s==nil)
// fmt.Printf("%T",s)


// var ages map[string]int = map[string]int{
// 	"bob":10,
// }
// var pages = make(map[string]int)

// tags :=map[string]int{
// 	"kob":25,
// }

// // append(ages,"bob")

// fmt.Println(ages==nil)
// fmt.Println(tags)
// pages["snip"]=50
// fmt.Println(pages)



//     // i  := 10
// // fmt.Println(i, i+1, i*i)
// }
// var x float32 = math.Pi
// var y float64 = math.Pi
// var z complex128 = math.Pi
	// type Currency int
	// const(
	// 	USD Currency=iota
	// 	EUR
	// 	GBP
	// 	RMP = 10
	// 	)
	
	// symbol :=[...]string{USD:"$", EUR:"@", GBP:"#", RMP:"&"}
	// fmt.Printf("%s",symbol)
// fmt.Println(x, y,z);
// func intsToString(values []int) string {
// var buf bytes.Buffer
// buf.WriteByte('[')
// for i, v:= range values {
// if i > 0 {
// buf.WriteString(", ")
// }
// fmt.Fprintf(&buf, "%d", v)
// }
// buf.WriteByte(']')
// return buf.String()
// }
// func main() {
// fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}