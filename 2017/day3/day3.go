package main

import (
	"fmt"
	"math"
)

func main() {
	target := 289326

	// anchor is the smallest square which is larger than target.
	//   a = x^2 where x is a positive integer
	//   a > target
	// note:
	// 	- it will always reside on the bottom right of the square.
	var anchor int

	x1, y1 := 0, 0
	x2, y2 := 0, 0

	x := 1
	for {
		// update cartesian coordinates of potential anchor point
		x2++
		y2--

		x += 2
		if anchor = x * x; anchor > target {
			break
		}
	}

	// found the anchor point - we know our target is on this layer
	// now we need to figure out which "side" of the square it's on
	if target > anchor-(x-1) {
		// target is on bottom
		x2 -= (anchor - target) // y2 is the same as anchor

	} else if target > anchor-2*(x-1) {
		// target is on left
		x2 = -x2
		y2 += (anchor - 1*(x-1)) - target

	} else if target > anchor-3*(x-1) {
		// target is on top
		x2 = -x2 + (anchor - 2*(x-1)) - target
		y2 = -y2

	} else {
		// target is on right
		y2 = (anchor - 3*(x-1)) - target // x2 is the same as anchor
	}

	// now calculate manhattan distance between origin and target
	distance := math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2))

	fmt.Printf("Origin (%d, %d)\n", x1, y1)
	fmt.Printf("Target (%d, %d)\n", x2, y2)
	fmt.Printf("Distance: %f\n", distance)
}
