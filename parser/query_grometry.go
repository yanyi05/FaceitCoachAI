package main

import "math"

func DistanceSquared(
	a PlayerState,
	b PlayerState,
) int {

	dx := int(a.X) - int(b.X)
	dy := int(a.Y) - int(b.Y)

	return dx*dx + dy*dy
}

func Distance(
	a PlayerState,
	b PlayerState,
) int {

	return int(math.Sqrt(float64(DistanceSquared(a, b))))
}

func HeightDifference(
	a PlayerState,
	b PlayerState,
) int {

	dz := int(a.Z) - int(b.Z)

	if dz < 0 {
		dz = -dz
	}

	return dz
}
