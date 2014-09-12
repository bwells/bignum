package main

import (
	"fmt"
)

func isEven(n uint8) bool {
	return n&1 == 0
}

func BitIsSet(n, p uint8) bool {
	return n&(1<<p) != 0
}

func SetBit(n, p uint8) uint8 {
	return n | 1<<p
}

func ClearBit(n, p uint8) uint8 {
	return n &^ (1 << p)
}

type Cell struct {
	value uint8
	next  *Cell
}

func PrintByte(n uint8) {
	for i := uint8(8); i > 0; i-- {
		if BitIsSet(n, i-1) {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
}

func bool2int(v bool) int {
	if v {
		return 1
	}
	return 0
}

func Add(a, b uint8) uint8 {
	result := uint8(0)
	remainder := 0
	a_val, b_val := 0, 0

	for i := uint8(0); i < 8; i++ {
		a_val = bool2int(BitIsSet(a, i))
		b_val = bool2int(BitIsSet(b, i))

		switch column_sum := a_val + b_val + remainder; column_sum {
		case 1:
			result = SetBit(result, i)
			remainder = 0
		case 2:
			remainder = 1
		case 3:
			result = SetBit(result, i)
		}
	}
	return result
}

func Add2(a, b uint8) uint8 {
	result := uint8(0)
	remainder := 0
	a_val, b_val := 0, 0

	for i := uint8(0); i < 8; i++ {
		a_val = bool2int(BitIsSet(a, i))
		b_val = bool2int(BitIsSet(b, i))

		switch column_sum := a_val + b_val + remainder; column_sum {
		case 1:
			result = SetBit(result, i)
			remainder = 0
		case 2:
			remainder = 1
		case 3:
			result = SetBit(result, i)
		}
	}
	return result
}

func main() {

	fmt.Printf("%b", uint8(2))
}
