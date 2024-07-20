package main

import "fmt"

func main() {
	var inputs = []string{"1", "2"}
	for i, j := range inputs {
		fmt.Println(i, j)
	}
	s := 1
	var ss = []int{s}
	fmt.Print(ss)

	hello()
}
