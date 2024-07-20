package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//python3 e.py
	cmd1 := exec.Command("python3", "e.py")
	out1, _ := cmd1.CombinedOutput()
	fmt.Printf("output: %s", out1)
}
