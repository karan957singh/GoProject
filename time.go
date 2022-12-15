package main

import (
	"fmt"
	"time"
	"math/rand"
	// "encoding/base64"
"encoding/hex"
"strings"
)

type a struct{
	abc int
	xyz string
}

type b struct{
	abc int
	xyz string
}
func main() {
	loc, _ := time.LoadLocation("UTC")
	times := time.Now().UTC()
	unixtime := times.Unix()

	fmt.Println(unixtime)

	fmt.Println(time.Unix(unixtime, 0).UTC())
	fmt.Println(time.Unix(unixtime, 0))
	fmt.Println(time.Unix(1592524045, 0))
	fmt.Println(time.Unix(1564170400, 0))
	fmt.Println(time.Unix(1599660000, 0).UTC())
	fmt.Println(time.Unix(1599660000, 0).In(loc))
	fmt.Println(time.Unix(1599696000, 0).UTC())
	fmt.Println(time.Unix(1599696000, 0))
	b := make([]byte, 4) //equals 8 charachters
rand.Read(b) 
s := hex.EncodeToString(b)
fmt.Println(s)
stt := RandStr()
fmt.Println(string(stt[2]), stt)
}

func RandStr() string {
	buff := make([]byte, 4)
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	// Base 64 can be longer than len
	return strings.ToUpper(str)
}
