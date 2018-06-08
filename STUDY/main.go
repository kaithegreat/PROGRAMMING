package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("hello world!")
	x := 0
	// var unikl int
	// var unikl2 string
	// unikl = 0
	pencil := 10
	y := 5
	// x = 2
	z := x + y
	fmt.Println("z is ", z, "pencil is", pencil)
	x = 4
	// fmt.Println("for loop..")
	// if x == 1 {
	// 	fmt.Println("x is 1")
	// } else if x == 3 {
	// 	fmt.Println("xis 3")
	// } else {
	// 	fmt.Println("x is 2")
	// }
	for i := 0; i < 10; i++ {
		x += i
		// x=x+i
		fmt.Println("x is", x)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	vartext := scanner.Text()
	fmt.Println("my text is", vartext)

}
