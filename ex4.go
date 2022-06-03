package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	for _, elem := range a {
		fmt.Print(int(elem-'0') * int(elem-'0'))
	}
}
