package examples

import "fmt"

func ExampleSimpleCheckDigit() {
	fmt.Println(SimpleCheckDigit([]int{4, 0, 0, 6}))
	fmt.Println(SimpleCheckDigit([]int{4, 0, 0, 7}))
	fmt.Println(SimpleCheckDigit([]int{1, 2}))
	fmt.Println(SimpleCheckDigit([]int{2, 1}))

	// Output:
	// 0
	// 1
	// 3
	// 3
}
