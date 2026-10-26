package examples

// Diese Datei zeigt eine einfache Prüfziffer, die ohne Gewichte
// auskommt, und welche Fehler sie nicht erkennt.

// SimpleCheckDigit erwartet eine Liste von Ziffern und liefert
// die letzte Ziffer ihrer Summe als Prüfziffer.
func SimpleCheckDigit(digits []int) int {
	sum := 0
	for _, digit := range digits {
		sum += digit
	}
	return sum % 10
	// Eine Prüfziffer ohne Gewichte:
	// - Wird eine Ziffer falsch abgetippt, ändert sich die Summe
	//   und meist auch die Prüfziffer.
	// - Werden zwei Ziffern vertauscht, bleibt die Summe gleich:
	//   Für `[1, 2]` und `[2, 1]` ist die Prüfziffer jeweils `3`.
	// - Deshalb gewichten EAN und ISBN die Ziffern an verschiedenen
	//   Stellen unterschiedlich.
}
