package exercises

// ISBN10CheckDigit erwartet die ersten `9` Ziffern einer ISBN-10
// und liefert die Prüfziffer, die als 10. Ziffer folgt.
//
// Für die Prüfziffer werden die Ziffern von links nach rechts mit
// `10`, `9`, ..., `2` gewichtet und addiert. Die Prüfziffer ergänzt
// diese gewichtete Summe zur nächsten durch `11` teilbaren Zahl.
// Sie kann deshalb auch `10` sein, was in einer ISBN als `X` steht.
func ISBN10CheckDigit(digits []int) int {
	// Hinweis:
	// Gehen Sie vor wie in `EANCheckDigit`, aber mit anderen Gewichten
	// und mit `11` statt `10`. Die Gewichte können Sie direkt als
	// Liste angeben.

	sum := WeightedSum(digits, []int{10, 9, 8, 7, 6, 5, 4, 3, 2})
	return MissingToMultiple(sum, 11)
}

// IsValidISBN10 erwartet die Ziffern einer ISBN-10, mit `10` für
// die Prüfziffer `X`, und liefert `true`, falls es `10` Ziffern sind
// und die letzte die passende Prüfziffer ist.
func IsValidISBN10(digits []int) bool {
	// Hinweis:
	// Gehen Sie vor wie in `IsValidEAN`. Die Prüfziffer erhält das
	// Gewicht `1`, die Summe muss durch `11` teilbar sein.

	if len(digits) != 10 {
		return false
	}
	return WeightedSum(digits, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})%11 == 0
}
