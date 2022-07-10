package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorFullCapacity        = errors.New("shapeCapacity is full")
	errorIndexOutOfRange     = errors.New("index out of range")
	errorCircleDoesNotExists = errors.New("circle does not exists")
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
	shapesCount := len(b.shapes)
	if shapesCount >= b.shapesCapacity {
		return fmt.Errorf("could not add shape: %w", errorFullCapacity)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	shapesCount := len(b.shapes)
	if i >= shapesCount {
		return nil, fmt.Errorf("could not get by index: %w", errorIndexOutOfRange)
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shapesCount := len(b.shapes)
	if i >= shapesCount {
		return nil, fmt.Errorf("could not extract by index: %w", errorIndexOutOfRange)
	}
	indexElement := b.shapes[i]
	b.shapes = removeByIndex(b.shapes, i)
	return indexElement, nil
}

// removeByIndex allows remove element by index from the slice and return new slice without this element
func removeByIndex(s []Shape, index int) []Shape {
	result := make([]Shape, 0)
	result = append(result, s[:index]...)
	return append(result, s[index+1:]...)
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	shapesCount := len(b.shapes)
	if i >= shapesCount {
		return nil, fmt.Errorf("could not replace by index: %w", errorIndexOutOfRange)
	}
	replacedShape := b.shapes[i]
	b.shapes[i] = shape
	return replacedShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	perimeterSum := .0
	for _, shape := range b.shapes {
		perimeterSum += shape.CalcPerimeter()
	}
	return perimeterSum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	areaSum := .0
	for _, shape := range b.shapes {
		areaSum += shape.CalcArea()
	}
	return areaSum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	circlesCount := 0
	for _, shape := range b.shapes {
		if isCircle(shape) {
			circlesCount += 1
		}
	}
	if circlesCount == 0 {
		return errorCircleDoesNotExists
	}
	shapesCount := len(b.shapes)
	for i := 0; i < shapesCount; i++ {
		if isCircle(b.shapes[i]) {
			b.shapes = removeByIndex(b.shapes, i)
			shapesCount -= 1
			i -= 1
		}
	}
	return nil
}

func isCircle(shape Shape) bool {
	_, ok := shape.(*Circle)
	if !ok {
		return false
	}
	return true
}
