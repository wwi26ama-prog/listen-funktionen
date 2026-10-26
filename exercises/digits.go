package exercises

// Digits erwartet einen Code wie `"978-3-16-148410-0"` und liefert
// die Liste seiner Ziffern als Zahlen. Ein `X`, die Prüfziffer `10`
// einer ISBN-10, wird zu `10`. Alle anderen Zeichen, z.B. Bindestriche
// und Leerzeichen, werden übersprungen.
func Digits(code string) []int {
	// Hinweis:
	// - Eine Schleife mit `range` über einen String liefert die
	//   einzelnen Zeichen (`rune`), die Sie z.B. mit `'0'` vergleichen können.
	// - Zeichen lassen sich wie Zahlen vergleichen und subtrahieren:
	//   Die Ziffern `'0'` bis `'9'` liegen direkt hintereinander.
	//   Was ergibt also `'7' - '0'`?
	// - Mit `int(...)` wird aus dem Ergebnis eine Zahl vom Typ `int`.

	digits := []int{}
	for _, char := range code {
		if char >= '0' && char <= '9' {
			digits = append(digits, int(char-'0'))
		} else if char == 'X' {
			digits = append(digits, 10)
		}
	}
	return digits
}
