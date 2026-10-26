package exercises

import "fmt"

func ExampleDigits() {
	fmt.Println(Digits("4006381333931"))
	fmt.Println(Digits("0306406152"))
	fmt.Println(Digits(""))

	// Output:
	// [4 0 0 6 3 8 1 3 3 3 9 3 1]
	// [0 3 0 6 4 0 6 1 5 2]
	// []
}

func ExampleDigits_separators() {
	fmt.Println(Digits("978-3-16-148410-0"))
	fmt.Println(Digits("3-16-148410-X"))
	fmt.Println(Digits("ISBN 0 306 40615 2"))

	// Output:
	// [9 7 8 3 1 6 1 4 8 4 1 0 0]
	// [3 1 6 1 4 8 4 1 0 10]
	// [0 3 0 6 4 0 6 1 5 2]
}
