package Interface

import "fmt"

type Figure interface {
	pirimetric()
	ploshadic()
}

func perimetr(figure Figure) {
	fmt.Println("\nPerimetric")
	figure.pirimetric()
}
func ploshad(figure Figure) {
	fmt.Println("\nploshadic")
	figure.ploshadic()
}

func perimetrAndArea(figure Figure) {
	fmt.Println("\nperimetrAndArea")
	figure.pirimetric()
	figure.ploshadic()

}

type Square struct {
	side float64
}

func newSquare(side float64) *Square {
	return &Square{side}
}

// /
// Подсчет периметра у квадрата
// /
func (squ Square) pirimetric() {
	perimetr := squ.side * 4
	fmt.Printf("Пириметр квадрата составляет: %.2f\n", perimetr)
}

// /
// Подсчет площади у квадрата
// /
func (squ Square) ploshadic() {
	area := squ.side * squ.side
	fmt.Printf("Площадь квадрата составляет: %.2f\n", area)

}

type Tringle struct {
	sideA  float64
	sideB  float64
	sideC  float64
	height float64
}

func newTringle(
	sideA float64,
	sideB float64,
	sideC float64,
	height float64) *Tringle {
	return &Tringle{
		sideA,
		sideB,
		sideC,
		height}
}

// /
// Подсчет треугольника у квадрата
// /
func (trgl Tringle) pirimetric() {
	perimetr := trgl.sideA + trgl.sideB + trgl.sideC
	fmt.Printf("Пириметр треугольника составляет: %.2f\n", perimetr)
}

// /
// Подсчет площади у треугольника
// /
func (trgl Tringle) ploshadic() {
	area := (trgl.sideA * trgl.height) / 2
	fmt.Printf("Площадь треугольника составляет: %.2f\n", area)

}

func InterfaseMain() {

	oneSquare := newSquare(2.0)
	oneTriangle := newTringle(3.0, 3.0, 3.0, 1.0)

	ploshad(oneTriangle)
	ploshad(oneSquare)

	perimetr(oneTriangle)
	ploshad(oneSquare)

	perimetrAndArea(oneTriangle)
	perimetrAndArea(oneSquare)
}
