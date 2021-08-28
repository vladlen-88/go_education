package main

import (
	"fmt"
)

func main() {

	fmt.Println("Чтобы найти число фибоначчи медленно, введите 1, быстро - введите 2")
	var n int
	fmt.Scanln(&n)

	var a int
	fmt.Print("Введите номер числа фибоначчи n = ")
	fmt.Scanln(&a)

	switch n {
	case 1:
		fmt.Println(LowFib(a))
	case 2:
		fmt.Println(HighFib(a))
	}

}

func LowFib(n int) int {
	if n < 3 {
		return 1
	}
	return LowFib(n-1) + LowFib(n-2)
}

func HighFib(n int) int {
	cache := map[int]int{1: 1, 2: 1}
	var f func(n int) int
	f = func(n int) int {
		if val, ok := cache[n]; ok {
			return val
		}
		cache[n] = f(n-1) + f(n-2)
		return cache[n]
	}
	return f(n)
}
