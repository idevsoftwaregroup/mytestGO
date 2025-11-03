package main

import (
	"fmt"
)

// ...existing code...

// Min returns the smaller of a and b.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("New World of min in GO ! Hi , ...")
	// example usage (adjust if the package API differs)
	a, b := 3, 5
	fmt.Printf("min(%d, %d) = %d\n", a, b, Min(a, b))
}

// ...existing code...
