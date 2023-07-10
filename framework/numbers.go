package framework

import "math"

type numbers interface {
	int | float64 | Radian | Degrees
}

func Limited[T numbers](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func OverLimited[T numbers](x, min, max T) T {
	for x < min {
		x = max - (min - x - 1)
	}
	for x > max {
		x = min + (x - max - 1)
	}
	return x
}

func Stepped(from, to, stepSize float64) float64 {
	if math.Abs(from-to) <= stepSize {
		return to
	}
	if from > to {
		return from - stepSize
	}
	return from + stepSize
}
