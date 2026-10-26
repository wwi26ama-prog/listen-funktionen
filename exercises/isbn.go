package exercises

// ISBN10CheckDigit erwartet die ersten `9` Ziffern einer ISBN-10
// und liefert die Prüfziffer, die als 10. Ziffer folgt.
//
// Für die Prüfziffer werden die Ziffern von links nach rechts mit
// `10`, `9`, ..., `2` gewichtet und addiert. Die Prüfziffer ergänzt
// diese gewichtete Summe zur nächsten durch `11` teilbaren Zahl.
// Sie kann deshalb auch `10` sein, was in einer ISBN als `X` steht.
//
// TODO
func ISBN10CheckDigit(digits []int) int {
	return 0
}

// IsValidISBN10 erwartet die Ziffern einer ISBN-10, mit `10` für
// die Prüfziffer `X`, und liefert `true`, falls es `10` Ziffern sind
// und die letzte die passende Prüfziffer ist.
//
// TODO
func IsValidISBN10(digits []int) bool {
	return false
}
