/*
Write a function that prints, in ascending order and on a single line: all unique combinations of
three different digits so that, the first digit is lower than the second, and the second is lower than the third.

These combinations are separated by a comma and a space.

Expected function
func PrintComb() {

}
Usage
Here is a possible program to test your function :

package main

import "piscine"

func main() {
	piscine.PrintComb()
}
This is the incomplete output :

$ go run . | cat -e
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789$
$
000 or 999 are not valid combinations because the digits are not different.

987 should not be shown because the first digit is not less than the second.

*/

package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i != '7' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
