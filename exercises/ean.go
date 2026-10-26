package exercises

// EANCheckDigit erwartet die ersten `12` Ziffern einer EAN-13
// (z.B. aus einem Strichcode) und liefert die Prüfziffer, die als
// 13. Ziffer folgt. ISBN-13 sind EAN-13, die mit `978` oder `979` beginnen.
//
// Für die Prüfziffer werden die Ziffern abwechselnd mit `1` und `3`
// gewichtet und addiert. Die Prüfziffer ergänzt diese gewichtete Summe
// zur nächsten durch `10` teilbaren Zahl: Für die Summe `89` ist sie `1`,
// für die Summe `100` ist sie `0`.
func EANCheckDigit(digits []int) int {
	// Hinweis:
	// Die Gewichte erhalten Sie mit `Repeat`, die gewichtete Summe
	// mit `WeightedSum` und die Prüfziffer mit `MissingToMultiple`.

	sum := WeightedSum(digits, Repeat([]int{1, 3}, len(digits)))
	return MissingToMultiple(sum, 10)
}

// IsValidEAN erwartet die Ziffern einer EAN-13 und liefert `true`,
// falls es `13` Ziffern sind und die letzte die passende Prüfziffer ist.
func IsValidEAN(digits []int) bool {
	// Hinweis:
	// - Prüfen Sie zuerst die Länge.
	// - Die gewichtete Summe aller `13` Ziffern, mit `1` für die Prüfziffer,
	//   ist bei einer gültigen EAN durch `10` teilbar.

	if len(digits) != 13 {
		return false
	}
	return WeightedSum(digits, Repeat([]int{1, 3}, 13))%10 == 0
	// Anmerkung:
	// Alternative:
	// `EANCheckDigit(digits[:12]) == digits[12]` vergleicht die
	// berechnete Prüfziffer direkt mit der letzten Ziffer.
}
