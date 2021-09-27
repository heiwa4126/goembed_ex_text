package main

import (
	_ "embed"
)

//go:embed hello.txt
var m1 string

//go:embed goodbye.txt
var m2 string

func main() {
	print(m1)
	print(m2)
}
