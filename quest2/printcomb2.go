/*
Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.

These combinations are separated by a comma and a space.

Expected function
func PrintComb2() {

}
Usage
Here is a possible program to test your function :

package main

import "piscine"

	func main() {
		piscine.PrintComb2()
	}

This is the incomplete output :

$ go run . | cat -e
00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$
$
*/
package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for b := '0'; b <= '9'; b++ {
					if j < b {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(b)
						if !(i == '9' && j == '8' && k == '9' && b == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
}
