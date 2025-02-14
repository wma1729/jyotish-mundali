package misc

import "math"

func AbsoluteDifference(q1, q2 float32) float32 {
	diff := q1 - q2
	return float32(math.Abs(float64(diff)))
}
