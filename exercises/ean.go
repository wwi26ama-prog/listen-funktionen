package exercises

// EANCheckDigit erwartet die ersten `12` Ziffern einer EAN-13
// (z.B. aus einem Strichcode) und liefert die Prüfziffer, die als
// 13. Ziffer folgt. ISBN-13 sind EAN-13, die mit `978` oder `979` beginnen.
//
// Für die Prüfziffer werden die Ziffern abwechselnd mit `1` und `3`
// gewichtet und addiert. Die Prüfziffer ergänzt diese gewichtete Summe
// zur nächsten durch `10` teilbaren Zahl: Für die Summe `89` ist sie `1`,
// für die Summe `100` ist sie `0`.
//
// TODO
func EANCheckDigit(digits []int) int {
	return 0
}

// IsValidEAN erwartet die Ziffern einer EAN-13 und liefert `true`,
// falls es `13` Ziffern sind und die letzte die passende Prüfziffer ist.
//
// TODO
func IsValidEAN(digits []int) bool {
	return false
}
