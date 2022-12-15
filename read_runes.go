package main
import (
"bufio"
"fmt"
"io"
"os"
"unicode"
"unicode/utf8"
)
func main() {
counts := make(map[rune]int) // counts of Unicode characters
var utflen [utf8.UTFMax + 1]int // count of lengths of UTF8 encodings
invalid := 0 // count of invalid UTF8 characters
in := bufio.NewReader(os.Stdin)
for {
r, n, err := in.ReadRune() // returns rune, nbytes, error
if err == io.EOF {
break
}
if err != nil {
fmt.Fprintf(os.Stderr, "charcount: %v", err)
os.Exit(1)
}
if r == unicode.ReplacementChar && n == 1 {
invalid++
continue
}
counts[r]++
utflen[n]++
}
fmt.Printf("rune\tcount\t")
for c, n := range counts {
fmt.Printf("%q\t%d\t", c, n)
}
fmt.Print("\nlen\tcount\t")
for i, n := range utflen {
if i > 0 {
fmt.Printf("%d\t%d\t", i, n)
}
}
if invalid > 0 {
fmt.Printf("%d invalid UTF8 characters", invalid)
}
}