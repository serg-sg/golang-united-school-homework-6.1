package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface

// Круг должен соответствовать интерфейсу фигуры
type Circle struct {
	Radius float64
}

// CalcPerimeter возвращает результат расчёта периметра
func (a Circle) CalcPerimeter() float64 {
	return math.Pi * 2 * a.Radius
}

// CalcArea returns calculation result of area

// CalcArea возвращает результат расчёта площади
func (a Circle) CalcArea() float64 {
	return math.Pi * math.Pow(a.Radius, 2)
}
