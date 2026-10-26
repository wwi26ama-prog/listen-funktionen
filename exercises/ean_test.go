package exercises

import "fmt"

func ExampleEANCheckDigit() {
	fmt.Println(EANCheckDigit([]int{4, 0, 0, 6, 3, 8, 1, 3, 3, 3, 9, 3}))
	fmt.Println(EANCheckDigit([]int{5, 9, 0, 1, 2, 3, 4, 1, 2, 3, 4, 5}))
	fmt.Println(EANCheckDigit([]int{9, 7, 8, 3, 1, 6, 1, 4, 8, 4, 1, 0}))

	// Output:
	// 1
	// 7
	// 0
}

func ExampleIsValidEAN() {
	fmt.Println(IsValidEAN([]int{4, 0, 0, 6, 3, 8, 1, 3, 3, 3, 9, 3, 1}))
	fmt.Println(IsValidEAN([]int{9, 7, 8, 3, 1, 6, 1, 4, 8, 4, 1, 0, 0}))
	fmt.Println(IsValidEAN([]int{4, 0, 0, 6, 3, 8, 1, 3, 3, 3, 9, 3, 2}))
	fmt.Println(IsValidEAN([]int{4, 0, 0, 6, 3, 8, 1, 3, 3, 3, 3, 9, 1}))
	fmt.Println(IsValidEAN([]int{4, 0, 0, 6, 3, 8, 1, 3, 3, 3, 9, 3}))
	fmt.Println(IsValidEAN([]int{}))

	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExampleIsValidEAN_digits() {
	fmt.Println(IsValidEAN(Digits("978-3-16-148410-0")))
	fmt.Println(IsValidEAN(Digits("978-3-16-148410-1")))

	// Output:
	// true
	// false
}
