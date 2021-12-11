package main

/*
This solution is slower than [][]type on less than ~50k entries
but significantly faster on 10m+ data
*/

//Array2D - 2 dimensional array
type Array2D struct {
	data []int
	x    int
	y    int
}

//Create - create 2d array
func Create(x, y int) *Array2D {
	return &Array2D{data: make([]int, x*y), x: x, y: y}
}

//Put - add element to the array
func (a2d *Array2D) Put(x, y int, value int) {
	a2d.data[x*a2d.y+y] = value
}

//Get - get value from x y index
func (a2d *Array2D) Get(x, y int) int {
	return a2d.data[x*a2d.y+y]
}

//GetSize - returns the size of X and size of Y
func (a2d *Array2D) GetSize() (int, int) {
	return a2d.x, a2d.y
}

//GetPosition - returns the position of the element
func (a2d *Array2D) GetPosition(x, y int) int {
	return x*a2d.y + y
}

//Copy - returns a copy of the original array
func (a2d *Array2D) Copy() *Array2D {
	newArray := &Array2D{data: make([]int, a2d.x*a2d.y), x: a2d.x, y: a2d.y}
	for index, val := range a2d.data {
		newArray.data[index] = val
	}

	return newArray
}

//Equals - true if the arrays are identical
func (a2d *Array2D) Equals(other *Array2D) bool {
	if a2d.x != other.x || a2d.y != other.y {
		return false
	}

	for index, val := range a2d.data {
		if val != other.data[index] {
			return false
		}
	}

	return true
}

func (a2d *Array2D) increase() {
	for i := range a2d.data {
		a2d.data[i]++
	}
}

func (a2d *Array2D) zeroFlashes() {
	for i, val := range a2d.data {
		if val >= 10 {
			a2d.data[i] = 0
		}
	}
}

func (a2d *Array2D) getFlashes() int {
	flashCounter := 0
	xMax, yMax := a2d.GetSize()
	flashing := make([]pair, 0)
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			if a2d.Get(x, y) == 10 {
				flashCounter++
				flashing = append(flashing, pair{x: x, y: y})
			}
		}
	}

	for _, p := range flashing {
		flashCounter += a2d.flasher(p.x, p.y, 0)
	}

	return flashCounter
}

func (a2d *Array2D) flasher(x, y, flashCounter int) int {
	adjacents := a2d.getAdjacents(x, y)
	for _, p := range adjacents {
		newVal := a2d.Get(p.x, p.y) + 1
		a2d.Put(p.x, p.y, newVal)
		if newVal == 10 {
			flashCounter++
			flashCounter = a2d.flasher(p.x, p.y, flashCounter)
		}
	}

	return flashCounter
}

func (a2d *Array2D) getAdjacents(x, y int) []pair {
	xMax, yMax := x+1, y+1
	xMin, yMin := 0, 0

	if x > 0 {
		xMin = x - 1
	}

	if y > 0 {
		yMin = y - 1
	}

	limitX, limitY := a2d.GetSize()
	if xMax == limitX {
		xMax--
	}

	if yMax == limitY {
		yMax--
	}

	adjacents := make([]pair, 0)
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if i == x && j == y {
				continue
			}

			adjacents = append(adjacents, pair{x: i, y: j})
		}
	}

	return adjacents
}
