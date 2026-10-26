package exercises

// Scale erwartet eine Liste von Zahlen und einen Faktor und liefert
// eine neue Liste, in der jede Zahl mit dem Faktor multipliziert ist.
func Scale(numbers []int, factor int) []int {
	// Hinweis:
	// Gehen Sie vor wie in `Double` aus den Beispielen.

	scaled := []int{}
	for _, number := range numbers {
		scaled = append(scaled, factor*number)
	}
	return scaled
}

// AllDigits erwartet eine Liste von Zahlen und liefert `true`,
// falls jede Zahl eine Ziffer von `0` bis `9` ist.
func AllDigits(numbers []int) bool {
	// Hinweis:
	// - Sobald eine Zahl keine Ziffer ist, steht das Ergebnis fest.
	// - Welches Ergebnis ist für die leere Liste richtig?

	for _, number := range numbers {
		if number < 0 || number > 9 {
			return false
		}
	}
	return true
}

// MissingToMultiple erwartet zwei Zahlen `n >= 0` und `m > 0` und liefert,
// was zu `n` mindestens addiert werden muss, damit die Summe durch `m`
// teilbar ist. Für `n = 89` und `m = 10` ist das Ergebnis also `1`,
// für `n = 100` und `m = 10` ist es `0`.
func MissingToMultiple(n int, m int) int {
	// Hinweis:
	// - `n%m` gibt an, wie weit `n` über der letzten durch `m`
	//   teilbaren Zahl liegt.
	// - Ist dieser Rest `0`, muss auch das Ergebnis `0` sein, nicht `m`.

	return (m - n%m) % m
}
