package main

import (
	"fmt"
	"math"
)

const PI float = 3.1472

func main() {
	fmt.Println(PI)
	const n = 500000

	const d = 3e20/n 
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}