package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (s Triangle) CalcPerimeter() float64 {
	return 3 * s.Side
}

func (s Triangle) CalcArea() float64 {
	return math.Sqrt(3) / 4 * (s.Side * s.Side)
}
