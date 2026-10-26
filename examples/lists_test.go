package examples

import "fmt"

func ExampleDouble() {
	numbers := []int{1, 2, 3}
	fmt.Println(Double(numbers))
	fmt.Println(numbers)
	fmt.Println(Double([]int{}))

	// Output:
	// [2 4 6]
	// [1 2 3]
	// []
}

func ExampleAdd() {
	fmt.Println(Add([]int{1, 2, 3}, []int{10, 20, 30}))
	fmt.Println(Add([]int{}, []int{}))

	// Output:
	// [11 22 33]
	// []
}

func ExampleEveryOther() {
	fmt.Println(EveryOther([]int{4, 0, 0, 6, 3}))
	fmt.Println(EveryOther([]int{7}))
	fmt.Println(EveryOther([]int{}))

	// Output:
	// [4 0 3]
	// [7]
	// []
}
