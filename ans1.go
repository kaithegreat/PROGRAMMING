package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//CalcIn to store the responses
type CalcIn struct {
	String string
	Split  []string
}

func main() {
	fmt.Println("welcome to fun calculator")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()

	ci := CalcIn{}
	ci.calcStr(t)

	int1, _ := strconv.Atoi(ci.Split[0])
	ex := ci.Split[1]
	int2, _ := strconv.Atoi(ci.Split[2])

	switch ex {
	case "+":
		fmt.Println(int1 + int2)
	case "-":
		fmt.Println(int1 - int2)
	case "*":
		fmt.Println(int1 * int2)
	case "/":
		fmt.Println(int1 / int2)
	}
}

func (ci *CalcIn) calcStr(t string) {
	(*ci).String = t
	(*ci).Split = strings.Split(t, " ")
}
