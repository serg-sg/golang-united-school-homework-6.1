package golang_united_school_homework

type Shape interface {
	// CalcPerimeter returns calculation result of perimeter

	// CalcPerimeter возвращает результат расчёта периметра
	CalcPerimeter() float64

	// CalcArea returns calculation result of area

	// CalcArea возвращает результат расчёта площади
	CalcArea() float64
}
