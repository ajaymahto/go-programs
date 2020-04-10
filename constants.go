// Author: Ajay Kumar Mahto <aj.mahto01@gmail.com>
// Description: Printing constants

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	// s = "vidhi"
	// # command-line-arguments
	// ./constants.go:16:4: cannot assign to s

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
