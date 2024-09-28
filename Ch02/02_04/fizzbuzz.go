/*
Print the numbers between 1 to 20, however
- If the number is divisible by 3, print "fizz" instead
- If the number is divisible by 5, print "buzz" instead
- If the number is divisible by 3 and 5, print "fizz buzz" instead
- Otherwise print the number
*/

package main

import (
	"fmt"
)

func main() {
	// fmt.Println(1 % 5)
	// fmt.Println(7 % 5)
	// fmt.Println(10 % 5)

	for i := 1; i < 21; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("fizz buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
