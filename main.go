package main

import (
	"fmt"
	"os"
)

var input *os.File
var output *os.File

var A int
var B int
var C int

type Scheduler interface {
	Add(Car)
	Pop() Car
}

type Car struct {
	ID    int
	Rides []int

	Arrival int
	X       int
	Y       int
}

func solve() {
	fmt.Fprintf(output, "%d\n", C)
}

func main() {
	sample := os.Args[1]
	fileIn := sample + ".in"
	fileOut := sample + ".out"

	var err error
	input, err = os.Open(fileIn)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", fileIn, err))
	}
	output, err = os.Create(fileOut)
	if err != nil {
		panic(fmt.Sprintf("creating %s: %v", fileOut, err))
	}
	defer input.Close()
	defer output.Close()

	// Global
	A = readInt()
	B = readInt()
	C = readInt()

	solve()
}

func readInt() int {
	var i int
	fmt.Fscanf(input, "%d", &i)
	return i
}

func readString() string {
	var str string
	fmt.Fscanf(input, "%s", &str)
	return str
}

func readFloat() float64 {
	var x float64
	fmt.Fscanf(input, "%f", &x)
	return x
}
