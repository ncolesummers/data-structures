package main

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct{}

// DrawShape struct has method draw Shape with float x and y coordinates

func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

type IContour interface {
	drawContour(x[5] float32, y[5] float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x[5] float32
	y[5] float32
	shape DrawShape
	factor int
}

func (contour DrawContour) drawContour(x[5] float32, y[5] float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

func main() {
	var (
		x = [5]float32{1,2,3,4,5}
		y = [5]float32{1,2,3,4,5}
		contour IContour = DrawContour{x,y, DrawShape{},2}
	)
	contour.drawContour(x,y)
	contour.resizeByFactor(2)
}