package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (s Rectangle) CalcPerimeter() float64 {
	return 2*s.Height + 2*s.Weight
}

func (s Rectangle) CalcArea() float64 {
	return s.Height * s.Weight
}
