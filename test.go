package main
import(
    "strconv"
"fmt"
// "time"
"crypto/rand"
// "math/rand"
"math/big"
// "encoding/base64"
// "bytes"
)
func main() {
    fmt.Printf("Amount: %.2f", 50000.55876856)
    fmt.Println("break")
    fmt.Printf("Amount: %.2f", big.NewFloat(12559.98))
    // fmt.Println(RandNumber(5))

    // t := true
    // f:=0
    // d :=1

    // if t{
    //     fmt.Println("t is true")
    // }

    // if f{
    //     fmt.Println("f is true")
    // }

    // if d{
    //     fmt.Println(" is true")
    // }

	// n, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(9999))
	// if err != nil {
 //                panic(err)
 //        }

 //        fmt.Println(n)
	// // Uses default seed of 1, result will be 81:
 //        fmt.Println(rand.Intn(1000))

 //        // Seed the random number generator using the current time (nanoseconds since epoch):
 //        rand.Seed(time.Now().UnixNano())

 //        // Hard to predict...but is it possible? I know the day, and hour, probably minute...
 //        fmt.Println(rand.Intn(1000))
	// // rand.Seed(time.Now().UTC().UnixNano())
 //    // fmt.Println(randomString(5))

	// // go spinner(100 * time.Millisecond)
	// // const n = 45
	// // fibN := fib(n) // slow
	// // fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
// func randomString(l int) string {
//     bytes := make([]byte, l)
//     for i := 0; i < l; i++ {
//         bytes[i] = byte(randInt(65, 90))
//     }
//     return string(bytes)
// }

// func randInt(min int, max int) int {
//     return min + rand.Intn(max-min)
// }
// func spinner(delay time.Duration) {
// 	for{
// 		for _, r := range `-\|/`{
// 			fmt.Printf("\r%c", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }
// func fib(x int) int {
// 	if x < 2 {
// 		return x
// 	}
// 	return fib(x-1) + fib(x-2)
// }

// func randStr(len int) string {
//     buff := make([]byte, len)
//     rand.Read(buff)
//     str := base64.StdEncoding.EncodeToString(buff)
//     // Base 64 can be longer than len
//     return str[:len]
// }

func RandNumber(len int) *big.Int {
    numStr := "9"
    for i := 1; i < len; i++ {
        numStr += "9"
    }
    num, _ := strconv.Atoi(numStr)
    newNum, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
    return newNum
}