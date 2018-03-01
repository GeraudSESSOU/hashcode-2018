package main

import (
	"fmt"
	"os"
	"sort"
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

func (r Ride) Length() int {
	xdist := r.a - r.x
	if xdist < 0 {
		xdist = -xdist
	}
	ydist := r.b - r.y
	if ydist < 0 {
		ydist = -ydist
	}

	return xdist + ydist
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (r *Ride) length() int {
	return abs(r.a-r.x) + abs(r.b-r.y)
}

type ByEndtime []*Ride

func (rs ByEndtime) Len() int      { return len(rs) }
func (rs ByEndtime) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }
func (rs ByEndtime) Less(i, j int) bool {
	return rs[i].f < rs[j].f || (rs[i].f == rs[j].f && rs[i].length() < rs[j].length())
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
	if c.Arrival < r.s {
		c.Arrival = r.s
	}
	c.moveTo(r.x, r.y)
	c.Rides = append(c.Rides, r.ID)
}

func (c *Car) EarliestFinish(r *Ride) int {
	copy := &Car{
		Arrival: c.Arrival,
		X:       c.X,
		Y:       c.Y,
	}
	c.moveTo(r.a, r.b)
	c.moveTo(r.x, r.y)
	return copy.Arrival
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

func (c *Car) distanceTo(x, y int) int {
	return abs(c.X-x) + abs(c.Y-y)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Choose(c *Car) *Ride {
	var bestRide *Ride
	bestLenOfRide := 0
	bestTotal := 0
	// fmt.Printf("car %d\n", c.ID)
	for _, r := range Rides {
		if r.used {
			continue
		}
		// if r.f < c.EarliestFinish(r) {
		// 	continue
		// }
		// fmt.Printf("%d %d -> %d %d\n", r.a, r.b, r.x, r.y)
		lenOfRide := r.length()
		total := max(c.distanceTo(r.a, r.b), r.s-c.Arrival) + lenOfRide
		// fmt.Printf("%d/%d\n", lenOfRide, total)
		if bestRide == nil || lenOfRide*bestTotal > total*bestLenOfRide {
			bestLenOfRide = lenOfRide
			bestTotal = total
			bestRide = r
		}
	}
	// if bestRide != nil {
	// 	fmt.Printf("Picking %d %d -> %d %d\n", bestRide.a, bestRide.b, bestRide.x, bestRide.y)
	// }
	return bestRide
}

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
	sort.Sort(ByEndtime(Rides))

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

// Prioq

// Invariant: both children are bigger

type prioq struct {
	bintree []*Car
}

func (pq *prioq) Add(car *Car) {
	pq.bintree = append(pq.bintree, car)

	// Rebalance tree to respect invariant
	var i = len(pq.bintree) - 1
	var p = (i - 1) / 2
	for p >= 0 && pq.bintree[p].Arrival > pq.bintree[i].Arrival {
		pq.bintree[p], pq.bintree[i] = pq.bintree[i], pq.bintree[p]
		i = p
		p = (i - 1) / 2
	}
}

func (pq *prioq) Pop() *Car {
	if len(pq.bintree) == 0 {
		return nil
	}

	if len(pq.bintree) == 1 {
		elem := pq.bintree[0]
		pq.bintree = pq.bintree[:0]
		return elem
	}

	elem := pq.bintree[0]
	// Put last element at root
	pq.bintree[0] = pq.bintree[len(pq.bintree)-1]
	// Remove last element
	pq.bintree = pq.bintree[:len(pq.bintree)-1]

	//        1                  9
	//    10     9	         10     12
	//  11 12   13 14  ->  11 12   13 14
	// 12

	// Rebalance tree to respect invariant
	len := len(pq.bintree)
	i, left, right := 0, 0, 0
	for {
		left = 2*i + 1
		right = 2*i + 2
		if left < len && right < len { // Two children
			if pq.bintree[left].Arrival <= pq.bintree[right].Arrival {
				if pq.bintree[i].Arrival <= pq.bintree[left].Arrival {
					break // Inferior to both children
				} else {
					pq.bintree[i], pq.bintree[left] = pq.bintree[left], pq.bintree[i]
					i = left
				}
			} else {
				if pq.bintree[i].Arrival <= pq.bintree[right].Arrival {
					break // Inferior to both children
				} else {
					pq.bintree[i], pq.bintree[right] = pq.bintree[right], pq.bintree[i]
					i = right
				}
			}
		} else if left < len { // One child (left)
			if pq.bintree[i].Arrival <= pq.bintree[left].Arrival {
				break // Inferior to only child
			}
			pq.bintree[i], pq.bintree[left] = pq.bintree[left], pq.bintree[i]
			i = left
		} else { // No child
			break
		}

	}

	return elem
}

func (pq *prioq) empty() bool {
	return len(pq.bintree) == 0
}
