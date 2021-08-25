package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите массив целых чисел в строку через пробел")
	a, err := inputArray()
	if err != nil {
		log.Fatal(err)
	}
	a = insertionSort(a)
	fmt.Printf("%v\n", a)
}

func inputArray() ([]int, error) {
	var arr []int
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}
	arr, err := convertStringArrayToInt(strings.Split(text, " "))
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func convertStringArrayToInt(strArr []string) ([]int, error) {
	res := make([]int, len(strArr))
	for i, item := range strArr {
		n, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		res[i] = n
	}
	return res, nil
}

func insertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		current := a[i]
		var j int
		for j = i - 1; j >= 0 && a[j] > current; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = current
	}
	return a
}
