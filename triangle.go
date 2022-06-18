package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface

// Треугольник должен удовлетворять интерфейсу фигуры
type Triangle struct {
	Side float64
}

// CalcPerimeter возвращает результат расчёта периметра
func (a Triangle) CalcPerimeter() float64 {
	return a.Side * 3
}

// CalcArea returns calculation result of area

// CalcArea возвращает результат расчёта площади
func (a Triangle) CalcArea() float64 {
	return (math.Pow(3, 0.5) * math.Pow(a.Side, 2) / 4)
}
