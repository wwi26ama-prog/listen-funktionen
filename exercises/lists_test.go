package exercises

import "fmt"

func ExampleSum() {
	fmt.Println(Sum([]int{4, 8, 15, 16, 23, 42}))
	fmt.Println(Sum([]int{-3, 3, 5}))
	fmt.Println(Sum([]int{}))

	// Output:
	// 108
	// 5
	// 0
}

func ExampleProduct() {
	fmt.Println(Product([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(Product([]int{-2, 0}, []int{3, 7}))
	fmt.Println(Product([]int{}, []int{}))

	// Output:
	// [4 10 18]
	// [-6 0]
	// []
}

func ExampleRepeat() {
	fmt.Println(Repeat([]int{1, 3}, 5))
	fmt.Println(Repeat([]int{1, 2, 3}, 7))
	fmt.Println(Repeat([]int{1, 3}, 1))
	fmt.Println(Repeat([]int{1, 3}, 0))

	// Output:
	// [1 3 1 3 1]
	// [1 2 3 1 2 3 1]
	// [1]
	// []
}

func ExampleWeightedSum() {
	fmt.Println(WeightedSum([]int{4, 0, 6}, []int{1, 3, 1}))
	fmt.Println(WeightedSum([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(WeightedSum([]int{}, []int{}))

	// Output:
	// 10
	// 32
	// 0
}
