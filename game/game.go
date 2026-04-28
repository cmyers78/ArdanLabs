package main

import "fmt"

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i) // i: main.Item{X:0, Y:0}
	fmt.Printf("i: %v\n", i)  // i: {X:0, Y:0}
	// Use %#v for debugging/logging  
	a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n", a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)

	i = Item{
		x: 11,
		y: 22,
	}

	fmt.Printf("i: %#v\n", i)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(45, 2000))
}

// Factory
// func NewItem(x, y int) Item
// func NewItem(x, y int) *Item
// func NewItem(x, y int) (Item, error)
// func NewItem(x, y int) (*Item, error)

func NewItem(x, y int) (Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return Item{}, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	return Item{x: x, y: y}, nil 
}

const (
	maxX = 600
	maxY = 400
)

type Item struct {
	x int 
	y int

}