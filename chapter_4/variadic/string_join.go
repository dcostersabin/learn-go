package main

import "fmt"

func main() {

	testString := []string{"hello", "-", "world"}

	fmt.Printf("%s", StringJoin(testString...))

}

func StringJoin(strs ...string) string {
	var st string
	for _, val := range strs {
		st += val
	}
	return st
}
