package main

import (
	"fmt"
	
	"sort"
)

func main() {
	// fmt.Println("Lesson on slices in Go.")
	// cart := []string {"Apple", "Orange", "Banana"}

	// fmt.Println("len:", len(cart))
	// fmt.Println("cart[1]", cart[1])

	// for i := range cart {
	// 	fmt.Println(i)
	// }

	// for i, v := range cart {
	// 	fmt.Println("index:", i, "value:", v)
	// }

	// cart = append(cart, "milk")
	// fmt.Println(cart)

	// fruit := cart[:3]
	// fmt.Println("fruit:", fruit)

	// fruit = append(fruit, "lemon")
	// fmt.Println("fruit", fruit)
	// fmt.Println("Cart:", cart)
	// // a := make([]int, 6, 10) // creates a slice with no values, but a capacity of 4
	// // b := make([]string, 6, 15) // creates a slice of strings
	// // fmt.Println("a:", a)
	// // fmt.Println(len(a))
	// // fmt.Println(cap(a))

	// // fmt.Println("B:", b)
	// // fmt.Println(len(b))
	// // fmt.Println(cap(b))
	// // b[3] = "Chris"
	// // fmt.Printf("%q\n", b)

	// var slice []int 
	// for i := range 10_000 {
	// 	slice = appendInt(slice, i)

	// }
	// fmt.Println(slice[:10])
	// fmt.Println(slice[9999])

	// // Exercise: concat
	// out := concat([]string{"A", "B"}, []string{"C"})
	// fmt.Printf("%q/n", out) // [a, b, c]

	values := []float64{3, 7, 8, 5, 1}
	fmt.Println(median(values))

	values = []float64{3, 1, 2, 4, 6, 8}
	fmt.Println(median(values))
	fmt.Println(values)
}

/* TODO: find the median value
	- sort values
	- if odd number of values, return middle
	- if even number of values, return average of two middle values

	[1, 5, 9] -> 5
	[2,12, 9, 4] -> [2, 4, 9, 12] -> 4+9/2 -> 6.5
*/
func median(values []float64) float64 {
	// so we don't alter the original set of values
	vals := make([]float64, len(values))
	copy(vals, values)
	sort.Float64s(vals)
	
	mid := len(vals)/2
	fmt.Println("Mid: ", mid)
	if len(vals) % 2 == 1 {
		return vals[mid]
	}
	lowerBound := vals[mid - 1]
	upperBound := vals[mid]
	return (lowerBound + upperBound) / 2 
}

func concat(slice1, slice2 []string) []string {
	size := (len(slice1) + len(slice2)) // get the size necessary to hold all concatenated elements
	newSlice := make([]string, size) // create a new slice with the larger sized slice 
	copy(newSlice, slice1) // copy everything from slice 1 to the new slice
	copy(newSlice[len(slice1):], slice2) // this copies values from slice2 into the locations beginning at the particular location of newSlice (ns)
	return newSlice
}

// example of expanding the size of a slice, but also an efficient way to alter a copy of a slice without altering the underlying original slice. 
func appendInt(slice []int, value int) []int {
	i := len(slice)

	// if there is no more capactiy, then create a larger capacity and set the existing slice equal to the new slice, while copying everything over
	if len(slice) == cap(slice) {
		size := 2 * (len(slice) + 1)
		fmt.Println(cap(slice), "-->", size)
		newSlice := make([]int, size)
		copy(newSlice, slice)
		slice = newSlice[:len(slice)]
	}
	// otherwise, add an extra spot from the existing capacity
	slice = slice[:len(slice) + 1]
	// add the new value to the slice and then return to the slice
	slice[i] = value
	return slice 
}