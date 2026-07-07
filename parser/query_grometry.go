package main

import "math"

func Distance(a, b PlayerState) float64 {

	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	dz := float64(a.Z - b.Z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func DistanceSquared(
	a PlayerState,
	b PlayerState,
) int {

	dx := int(a.X) - int(b.X)
	dy := int(a.Y) - int(b.Y)

	return dx*dx + dy*dy
}
