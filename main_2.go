package main

import (
	"fmt"
	"math"
)

func main() {

	/*2. Напишите программу, вычисляющую диаметр и
	//длину окружности по заданной площади круга.
	//Площадь круга должна вводиться пользователем с клавиатуры.*/

	var s float64
	fmt.Println("Программа рассчитает диаметр и длину окружности по заданной площади круга. ")
	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&s)

	//Длина окружности
	l := math.Sqrt(s * (4 * math.Pi))

	//Диаметр окружности
	d := l / math.Pi

	fmt.Println("Длина окружности:", l)
	fmt.Println("Диаметр окружности:", d)
}
