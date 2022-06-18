package golang_united_school_homework

import (
	"errors"
	"fmt"
	"reflect"
)

// box contains list of shapes and able to perform operations on them

// коробка содержит список фигур и возможность выполнять над ними операции
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
	// Максимальное количество фигур, которые могут быть внутри коробки
}

// NewBox creates new instance of box

// NewBox создает новый экземпляр коробки
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.

// AddShape добавляет фигуру в поле, возвращает ошибку,
// если она выходит за пределы диапазона shapeCapacity.

func (b *box) AddShape(shape Shape) error {
	err := errors.New("figure could not be added, box capacity exceeded")
	if len(b.shapes) == b.shapesCapacity {
		err = fmt.Errorf("%w", err)
		return err
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error

// GetByIndex позволяет получить фигуру по индексу,
// Если фигура по индексу не существует или индекс вышел за пределы диапазона,
// то он возвращает ошибку

func (b *box) GetByIndex(i int) (Shape, error) {
	err := errors.New("figure at index does not exist or index is out of capacity")
	if i >= len(b.shapes) {
		err = fmt.Errorf("%w", err)
		return nil, err
	}
	Shape := b.shapes[i]
	return Shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error

// ExtractByIndex позволяет получить фигуру по индексу и удаляет эту фигуру из списка.
// Если форма по индексу не существует или индекс вышел за пределы диапазона,
// то он возвращает ошибку
func (b *box) ExtractByIndex(i int) (Shape, error) {
	err := errors.New("figure at index does not exist or index is out of capacity")
	if i >= len(b.shapes) {
		err = fmt.Errorf("%w", err)
		return nil, err
	}
	Shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return Shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error

// ReplaceByIndex позволяет заменить форму по индексу и возвращает удалённую фигуру.
// Если форма по индексу не существует или индекс вышел за пределы диапазона,
// то он возвращает ошибку
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	err := errors.New("figure at index does not exist or index is out of capacity")
	if i >= len(b.shapes) {
		err = fmt.Errorf("%w", err)
		return nil, err
	}
	Shape := b.shapes[i]
	end := b.shapes[i+1:]
	b.shapes = append(b.shapes[:i], shape)
	b.shapes = append(b.shapes[:i+1], end...)
	return Shape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.

// SumPerimeter предоставляет суммарный периметр всех фигур в списке.
func (b *box) SumPerimeter() float64 {
	// Если фигур нет, то вернуть 0
	if len(b.shapes) == 0 {
		return 0
	}
	var sum float64
	for i, _ := range b.shapes {
		sum = sum + b.shapes[i].CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.

// SumArea предоставляет суммарную площадь всех фигур в списке.
func (b *box) SumArea() float64 {
	// Если фигур нет, то вернуть 0
	if len(b.shapes) == 0 {
		return 0
	}
	var sum float64
	for i, _ := range b.shapes {
		sum = sum + b.shapes[i].CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error

// RemoveAllCircles удаляет все круги в списке
// Если кружков в списке нет, то возвращает ошибку
func (b *box) RemoveAllCircles() error {
	// Если фигур нет, то вернуть 0
	if len(b.shapes) == 0 {
		return nil
	}

	found := false
	for i, shape := range b.shapes {
		// Сравниваем имя типа элемента с "Circle", в случае положительного результата
		// удаляем из массива
		if reflect.TypeOf(shape).Name() == "Circle" {
			found = true
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		}
	}

	if found {
		return nil
	}

	err := errors.New("circles not found")
	return err
}
