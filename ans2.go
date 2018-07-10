package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//CalcIn to store the responses
type CalcIn struct {
	Operation string
	Split     []string
}

func main() {
	fmt.Println("welcome to fun calculator")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()

	ci := CalcIn{}
	ci.calcStr(t)

	int1, _ := strconv.Atoi(ci.Split[0])
	ex := ci.Operation
	int2, _ := strconv.Atoi(ci.Split[1])

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
	r1, _ := regexp.Compile(`\D`)

	s := r1.FindString(t)

	(*ci).Operation = s
	(*ci).Split = strings.Split(t, s)
}
