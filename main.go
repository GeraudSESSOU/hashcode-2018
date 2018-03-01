package main

import (
	"fmt"
	"os"
)

var input *os.File
var output *os.File

var R int
var C int
var F int
var N int
var B int
var T int

var Rides []*Ride

var Cars []*Car

var Sched Scheduler

type Ride struct {
	ID int

	a, b, x, y, s, f int

	used bool
}

type Scheduler interface {
	Add(*Car)
	Pop() *Car
}

type Car struct {
	ID    int
	Rides []int

	Arrival int
	X       int
	Y       int
}

func (c *Car) Update(r *Ride) {
	c.moveTo(r.a, r.b)
	c.moveTo(r.x, r.y)
	c.Rides = append(c.Rides, r.ID)
}

func (c *Car) moveTo(x, y int) {
	xdist := c.X - x
	if xdist < 0 {
		xdist = -xdist
	}
	ydist := c.Y - y
	if ydist < 0 {
		ydist = -ydist
	}

	c.Arrival += xdist + ydist
	c.X = x
	c.Y = y
}

func Choose(c *Car) *Ride { return nil }

func assign() bool {
	c := Sched.Pop()
	if c == nil {
		return false
	}
	r := Choose(c)
	if r == nil {
		return true
	}
	r.used = true
	c.Update(r)
	Sched.Add(c)

	return true
}

func solve() {
	Sched = &prioq{}

	// create cars
	for i := 0; i < F; i++ {
		c := &Car{
			ID:      i,
			Arrival: 0,
			X:       0,
			Y:       0,
		}
		Cars = append(Cars, c)
		Sched.Add(c)
	}

	for assign() {
	}

	for _, c := range Cars {
		fmt.Fprintf(output, "%d", len(c.Rides))
		for _, ri := range c.Rides {
			fmt.Fprintf(output, " %d", ri)
		}
		fmt.Fprintf(output, "\n")
	}
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
	R = readInt()
	C = readInt()
	F = readInt()
	N = readInt()
	B = readInt()
	T = readInt()

	for i := 0; i < N; i++ {
		Rides = append(Rides, &Ride{
			ID: i,
			a:  readInt(),
			b:  readInt(),
			x:  readInt(),
			y:  readInt(),
			s:  readInt(),
			f:  readInt(),
		})
	}

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
