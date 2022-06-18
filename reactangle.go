package golang_united_school_homework

// Rectangle must satisfy to Shape interface

// Прямоугольник должен соответствовать интерфейсу Shape
type Rectangle struct {
	Height, Weight float64
}

// CalcPerimeter возвращает результат расчёта периметра
func (a Rectangle) CalcPerimeter() float64 {
	return (a.Height + a.Weight) * 2
}

// CalcArea returns calculation result of area

// CalcArea возвращает результат расчёта площади
func (a Rectangle) CalcArea() float64 {
	return a.Height * a.Weight
}
