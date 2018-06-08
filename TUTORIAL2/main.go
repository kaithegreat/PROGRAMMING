package main

import (
	"fmt"
)

func main() {
	fmt.Println("channel program")
	// c := make(chan string)
	// go delay(c)
	delay2()
	fmt.Println("i am running")
	// fmt.Println(<-c)

}
func delay(c chan string) {
	tick := 0
	for i := 0; i < 1000000000000; i++ {
		tick += i

	}
	c <- "done"
}
func delay2(){
	tick := 0
	for i := 0; i < 1000000000000; i++ {
		tick += i

	}
}
