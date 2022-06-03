package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(strings.Join(strings.Split(a, ""), "*"))

}
