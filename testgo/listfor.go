package main

import "fmt"

func main() {
	aa := []string{"aaaaaa", "bbbbb"}
	fmt.Println(aa[1])
	for a, b := range aa {
		fmt.Println(a, b)
	}
}
