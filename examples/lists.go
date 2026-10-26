package examples

// Diese Datei enthält Beispiele für Funktionen, die aus Listen
// neue Listen erzeugen. Die Übungsaufgaben bauen auf diesen Mustern auf.

// Double erwartet eine Liste von Zahlen und liefert eine neue Liste,
// in der jede Zahl verdoppelt ist.
func Double(numbers []int) []int {
	doubled := []int{}
	for _, number := range numbers {
		doubled = append(doubled, 2*number)
	}
	return doubled
	// Eine neue Liste aufbauen:
	// - Die Ergebnisliste startet leer.
	// - `append` liefert die um ein Element verlängerte Liste,
	//   die wieder in `doubled` gespeichert wird.
	// - Die ursprüngliche Liste bleibt unverändert.
}

// Add erwartet zwei gleich lange Listen von Zahlen und liefert eine
// neue Liste mit den Summen der Elemente an derselben Stelle.
func Add(a []int, b []int) []int {
	sums := []int{}
	for i := range a {
		sums = append(sums, a[i]+b[i])
	}
	return sums
	// Zwei Listen gleichzeitig durchlaufen:
	// - Die Schleife läuft über die Indizes der ersten Liste.
	// - Mit demselben Index wird auf beide Listen zugegriffen.
	// - Wäre `b` kürzer als `a`, bräche das Programm mit
	//   einem Fehler ab (`index out of range`).
}

// EveryOther erwartet eine Liste von Zahlen und liefert eine neue
// Liste mit dem ersten, dritten, fünften, ... Element.
func EveryOther(numbers []int) []int {
	selected := []int{}
	for i, number := range numbers {
		if i%2 == 0 {
			selected = append(selected, number)
		}
	}
	return selected
	// Ein Muster über den Index:
	// - `i%2` ist abwechselnd `0` und `1`.
	// - So lässt sich jedes zweite Element auswählen
	//   oder etwas abwechselnd tun.
}
