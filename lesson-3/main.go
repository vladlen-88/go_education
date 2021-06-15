package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, r float64

	var op string

	fmt.Print("Введите число а: ")

	if _, err := fmt.Scanln(&a); err != nil {
		//Проверка на тип значения
		fmt.Print("Введено не число, работа программы завершена")
		return
	}

	fmt.Print("Введите число b: ")

	if _, err := fmt.Scanln(&b); err != nil {
		//Проверка на тип значения
		fmt.Print("Введено не число, работа программы завершена")
		return
	}

	//Арифметическая операция
	fmt.Print("Введите требуемую арифметическую операцию путем ввода одного символа из: +, -, *, /, % :")

	fmt.Scanln(&op)

	switch op {

	case "+":
		r = a + b

	case "-":
		r = a - b

	case "*":
		r = a * b

	case "/":

		if b == 0.0 {
			//Проверка на ноль
			fmt.Println("Делить на ноль нельзя! Работа программы завершена.")
			return
		}

		r = a / b

	case "%":

		r = math.Mod(a, b)

	default:
		fmt.Println("Операция выбрана неверно. Необходимо выбрать из символов: +, -, *, /, %")
		return
	}

	fmt.Println(a, op, b, "=", r)
}
