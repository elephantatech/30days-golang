package main

import (
	"bufio"
	"fmt"
	"os"
)

func dayonesimple() {
	var input string
	//fmt.Scanf("%q", &input)
	fmt.Sscanln("", &input)
	fmt.Println("Hello, World.")
	fmt.Println(input)
}

func dayonecomplex() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	fmt.Println("Hello, World.")
	fmt.Println(input)
}

func main() {
	dayonesimple()
}
