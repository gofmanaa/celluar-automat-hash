package main

import (
	"fmt"
)

func main() {
	stringForEncoding := "Hello World"
	fmt.Println(stringForEncoding)
	buf := stringToBytes(stringForEncoding)
	fmt.Println("Binary presentation:")
	for _, n := range buf {
		fmt.Printf("%08b", n)
	}
	fmt.Println()

	run(buf)
}

// rule for cellular automaton.
func f(y1, y2, y3 uint) uint {
	return (y1 ^ (y2 | y3))
}

func run(buf []byte) {
	count := len(buf) * 8
	var y = make([]uint, count)
	var y1 = make([]uint, count)

	for i := 0; i < len(buf); i++ {
		for j := 0; j < 8; j++ {
			zeroOrOne := buf[i] >> (7 - j) & 1
			y[j+i*8] = uint(zeroOrOne)
		}
	}
	show(y)
	for gen := count; gen > 0; gen-- {
		for i := range y {
			// Use the % operator to wrap the indices around the count variable.
			y1[i] = f(y[(i-1+count)%count], y[i], y[(i+1)%count])
		}

		copy(y, y1)
		show(y)
	}

	by := parseBinToHex(y)
	for _, n := range by {
		fmt.Printf("%08b", n)
	}
	fmt.Println()
	fmt.Printf("Encode result: %X\n", by)
}

func show[T uint | byte](world []T) {
	for _, v := range world {
		if byte(v) == byte(1) {
			fmt.Printf("|")
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
}

func stringToBytes(str string) []byte {
	var buf []byte
	for _, ch := range str {
		buf = append(buf, byte(ch))
	}
	return buf
}

func parseBinToHex(arr []uint) []byte {
	var result []byte
	for i := 0; i < len(arr); i += 8 {
		// Get the current 8-character chunk as a byte.
		b := arr[i : i+8]

		// Convert the byte to an int.
		n := 0
		for j := 0; j < 8; j++ {
			n = n << 1
			if b[j] == 1 {
				n++
			}
		}
		result = append(result, byte(n))
	}

	return result
}
