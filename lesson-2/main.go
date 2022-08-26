package main

import (
	"fmt"
)

func main() {
	//It's a function for calculation of Rectangle's square
	var a, b int
	fmt.Println("Введите длины сторон прямоугольника")
	fmt.Scanln(&a, &b)
	fmt.Println(squareOfRectangle(a, b))
}

func squareOfRectangle(a, b int) int {
	result := a * b
	return result
}
