package exercises

import "fmt"

func ExampleISBN10CheckDigit() {
	fmt.Println(ISBN10CheckDigit([]int{0, 3, 0, 6, 4, 0, 6, 1, 5}))
	fmt.Println(ISBN10CheckDigit([]int{3, 1, 6, 1, 4, 8, 4, 1, 0}))
	fmt.Println(ISBN10CheckDigit([]int{0, 0, 0, 0, 0, 0, 0, 0, 0}))

	// Output:
	// 2
	// 10
	// 0
}

func ExampleIsValidISBN10() {
	fmt.Println(IsValidISBN10([]int{0, 3, 0, 6, 4, 0, 6, 1, 5, 2}))
	fmt.Println(IsValidISBN10([]int{3, 1, 6, 1, 4, 8, 4, 1, 0, 10}))
	fmt.Println(IsValidISBN10([]int{0, 3, 0, 6, 4, 0, 6, 1, 5, 3}))
	fmt.Println(IsValidISBN10([]int{0, 3, 0, 6, 4, 0, 1, 6, 5, 2}))
	fmt.Println(IsValidISBN10([]int{0, 3, 0, 6, 4, 0, 6, 1, 5}))
	fmt.Println(IsValidISBN10([]int{}))

	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExampleIsValidISBN10_digits() {
	fmt.Println(IsValidISBN10(Digits("3-16-148410-X")))
	fmt.Println(IsValidISBN10(Digits("3-16-148410-9")))

	// Output:
	// true
	// false
}
