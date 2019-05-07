package main

import (
	"fmt"
	"github.com/jakereps/go-boy/cpu"
)

func main() {
	_ = cpu.LR35902{}
	for k, v := range cpu.InstructionSet {
		fmt.Printf("%08b (%x): %v\n", k, k, v)
	}
}
