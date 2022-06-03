package main

import (
	"fmt"
)

func main() {
	var a string
	var b int
	fmt.Scan(&a)
	b = int(a[0] - '0')
	for i := 0; i < len(a); i++ {
		if b < int(a[i]-'0') {
			b = int(a[i] - '0')
		}
	}
	fmt.Println(b)
}
