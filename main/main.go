package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.DigitLen(100, 2))
	fmt.Println(piscine.DigitLen(100, 7))
	fmt.Println(piscine.DigitLen(100, 1))
	fmt.Println(piscine.DigitLen(-100, 4))
}
