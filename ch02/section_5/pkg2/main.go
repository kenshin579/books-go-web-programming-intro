package main

import (
	"fmt"
	"github.com/kenshin579/books-go-web-programming-intro/ch02/section_5/lib"
)

var v rune

func init() {
	fmt.Println("init called")
	v = '1'
}

func main() {
	fmt.Println(lib.IsDigit(v)) // lib 패키지의 IsDigit 함수 사용
}

func IsDigit(c int32) bool {
	return '0' <= c && c <= '9'
}
