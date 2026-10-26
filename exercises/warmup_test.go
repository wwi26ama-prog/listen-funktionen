package exercises

import "fmt"

func ExampleScale() {
	numbers := []int{1, 2, 3}
	fmt.Println(Scale(numbers, 3))
	fmt.Println(numbers)
	fmt.Println(Scale([]int{4, -1}, 0))
	fmt.Println(Scale([]int{}, 3))

	// Output:
	// [3 6 9]
	// [1 2 3]
	// [0 0]
	// []
}

func ExampleAllDigits() {
	fmt.Println(AllDigits([]int{4, 0, 0, 6, 9}))
	fmt.Println(AllDigits([]int{4, 10, 6}))
	fmt.Println(AllDigits([]int{4, -1}))
	fmt.Println(AllDigits([]int{}))

	// Output:
	// true
	// false
	// false
	// true
}

func ExampleMissingToMultiple() {
	fmt.Println(MissingToMultiple(89, 10))
	fmt.Println(MissingToMultiple(100, 10))
	fmt.Println(MissingToMultiple(177, 11))
	fmt.Println(MissingToMultiple(0, 11))

	// Output:
	// 1
	// 0
	// 10
	// 0
}
