package main

import "math"

func Distance2D(x1, y1, x2, y2 float64) float64 {

	dx := x2 - x1
	dy := y2 - y1

	return math.Sqrt(dx*dx + dy*dy)
}

func Distance3D(
	x1, y1, z1,
	x2, y2, z2 float64,
) float64 {

	dx := x2 - x1
	dy := y2 - y1
	dz := z2 - z1

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func HeightDifferenceFloat(z1, z2 float64) float64 {

	if z1 > z2 {
		return z1 - z2
	}

	return z2 - z1
}
