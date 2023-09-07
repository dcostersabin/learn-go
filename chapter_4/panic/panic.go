package main

import "fmt"

func panicTest(i int) {
	if i == 1 {
		panic(9)
	}
	fmt.Println("In Panic Test")
}

func test(str string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in test")
		}
	}()

	fmt.Println("calling panic test")
	panicTest(123)
	panicTest(1)
	fmt.Println("retruned normally from panic test")

}

func main() {
	test("hello")
}
