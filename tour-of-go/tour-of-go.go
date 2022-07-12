package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var c, python, java bool
var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 2i)
)

func main() {

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)

	fmt.Println(add(1, 2))
	fmt.Println(add2(5, 40))

	fmt.Println(swap("a", "b"))
	fmt.Println(swap2("a", "b", "c"))

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	k := 1
	x, y, z := 4, 5, 6
	fmt.Println(k, x, y, z)

	// Basic types
	/**
		bool
		string
		int int8 int16 int32 int64
		uint uint8 uint16 uint32 uint64 uintptr
		byte (alias for uint8)
		rune (alias for int32; represents a Unicode code point)
		float32 float64
		complex64 complex128
	**/

	// Zero values
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func swap2(x, y, z string) (string, string, string) {
	return y, z, x
}

// Using functions
func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}
