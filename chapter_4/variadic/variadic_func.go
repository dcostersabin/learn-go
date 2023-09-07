package main

import (
	"fmt"
	"os"
)

func main() {

	num := []int{123, 123, 111, 23, 123, 9, 1238, 0, 111111}

	num2 := []int{}

	test1, err := Intmin(num2...)

	if err != nil {
		fmt.Printf("%V", err)
		os.Exit(0)
	}

	test2, err := Intmin(num...)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(0)
	}

	fmt.Printf("%d", test1)
	fmt.Printf("%d", test2)

}

func Intmin(numbers ...int) (int, error) {
	if len(numbers) < 1 {
		return 0, fmt.Errorf("Slice should be greater than 0")

	}
	var min = numbers[0]
	for _, num := range numbers {
		if min > num {
			min = num
		}
	}

	return min, nil
}
