package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("welcome to fun calculator")
	fmt.Println("do addition,subtraction,multiplication or divide?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	vartext0 := scanner.Text()
	fmt.Println("please enter any number")
	scanner.Scan()
	vartext1 := scanner.Text()
	int1, _ := strconv.Atoi(vartext1)
	fmt.Println("your first number is", int1)
	scanner.Scan()
	vartext2 := scanner.Text()
	int2, _ := strconv.Atoi(vartext2)
	fmt.Println("your second number is", int2)
	if vartext0 == "addition" {
		operationaddition := int1 + int2
		// fmt.Println("your answer is", operationaddition)
		printanswer(operationaddition)
	} else if vartext0 == "subtraction" {
		operationsubtraction := int1 - int2
		// fmt.Println("your answer is", operationsubtraction) hi!
		// Im joining yeay!
		printanswer(operationsubtraction)
	} else if vartext0 == "multiplication" {
		operationmultiply := int1 * int2
		// fmt.Println("your answer is", operationmultiply)
		printanswer(operationmultiply)
	} else if vartext0 == "divide" {
		operationdivide := int1 / int2
		// fmt.Println("your answer is", operationdivide)
		printanswer(operationdivide)
	}

}
func printanswer(answer int) {
	fmt.Println("your answer is", answer)
}
