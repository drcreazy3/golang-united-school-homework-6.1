package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity == len(b.shapes) {
		return errors.New("Out of the shapesCapacity range")
	}
	b.shapesCapacity++
	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	element, err := b.getElementByIndex(i)
	if err != nil {
		return element, err
	}

	return element, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	element, err := b.getElementByIndex(i)
	if err != nil {
		return element, err
	}
	remove(b.shapes, i)

	return element, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	element, err := b.getElementByIndex(i)
	if err != nil {
		return element, err
	}
	b.shapes[i] = shape

	return shape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var newShapes []Shape
	circleExists := false
	for _, v := range b.shapes {
		if _, ok := v.(Circle); ok {
			circleExists = true
			continue
		}
		newShapes = append(newShapes, v)
	}

	if !circleExists {
		return errors.New("No circle exists")
	}

	b.shapes = newShapes

	return nil
}

func (b *box) getElementByIndex(index int) (Shape, error) {
	if index > len(b.shapes)-1 {
		return nil, errors.New("Out of the shapesCapacity range")
	}
	if !b.indexExists(index) {
		return nil, errors.New("Element index " + string(index) + "not exists")
	}

	return b.shapes[index], nil
}

func (b *box) indexExists(index int) bool {
	for k := range b.shapes {
		if k == index {
			return true
		}
	}
	return false
}

func remove(s []Shape, i int) []Shape {
	return append(s[:i], s[i+1:]...)
}
