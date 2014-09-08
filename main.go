package main

import (
	"fmt"
)

func IsEven(n uint8) bool {
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
	data uint8
	next *Cell
}

func PrintByte(n uint8) {
	for i := uint8(8); i > 0; i-- {
		if BitIsSet(n, i-1) {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
	fmt.Print("\n")
}

func Add(a, b uint8) uint8 {
	// var c, d Cell

	result := uint8(0)

	remainder := uint8(0)
	a_set, b_set := uint8(0), uint8(0)
	for i := uint8(0); i < 8; i++ {
		a_set = BitIsSet(a, i)
		b_set = BitIsSet(b, i)

		if !a_set && !b_set {
			if remainder > 0 {
				remainder = 0
				result = SetBit(result, i)
			} else {
				result = ClearBit(result, i)
			}
		} else if a_set && b_set {
			if remainder > 0 {
				result = ClearBit(result, i)
			} else {
				result = ClearBit(result, i)
				remainder = 1
			}
		} else {
			if remainder > 0 {
				result = SetBit(result, i)
				remainder = 0
			} else {
				result = SetBit(result, i)
			}
		}

	}

	return result
}

func main() {

	// var a uint8 = 0

	// var MaxInt uint8 = 1<<8 - 1

	PrintByte(Add(uint8(12), uint8(22)))
	PrintByte(uint8(34))

}
