package exercises

// Sum erwartet eine Liste von Zahlen und liefert ihre Summe.
// Die Summe der leeren Liste ist `0`.
func Sum(numbers []int) int {
	// Hinweis:
	// Starten Sie mit einer Variablen für die Summe und addieren Sie
	// in einer Schleife jede Zahl der Liste dazu.

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Product erwartet zwei gleich lange Listen von Zahlen und liefert
// eine neue Liste mit den Produkten der Elemente an derselben Stelle.
// Für `[1, 2, 3]` und `[4, 5, 6]` ist das Ergebnis also `[4, 10, 18]`.
func Product(a []int, b []int) []int {
	// Hinweis:
	// Gehen Sie vor wie in `Add` aus den Beispielen.

	products := []int{}
	for i := range a {
		products = append(products, a[i]*b[i])
	}
	return products
}

// Repeat erwartet eine nicht leere Liste `pattern` und eine Länge und
// liefert eine Liste dieser Länge, in der sich `pattern` wiederholt.
// Für `[1, 3]` und die Länge `5` ist das Ergebnis also `[1, 3, 1, 3, 1]`.
func Repeat(pattern []int, length int) []int {
	// Hinweis:
	// - Das Element an Stelle `i` des Ergebnisses ist ein Element von
	//   `pattern`. Welcher Index gehört zu `i`?
	// - Der Rest bei der Division durch `len(pattern)` hilft dabei,
	//   so wie `i%2` in `EveryOther` aus den Beispielen.

	repeated := []int{}
	for i := range length {
		repeated = append(repeated, pattern[i%len(pattern)])
	}
	return repeated
}

// WeightedSum erwartet eine Liste von Zahlen und eine gleich lange
// Liste von Gewichten. Sie liefert die Summe der Zahlen, jeweils
// multipliziert mit dem Gewicht an derselben Stelle.
// Für `[4, 0, 6]` und `[1, 3, 1]` ist das Ergebnis also
// `4*1 + 0*3 + 6*1 = 10`.
func WeightedSum(numbers []int, weights []int) int {
	// Hinweis:
	// Sie brauchen keine eigene Schleife: Kombinieren Sie `Sum` und `Product`.

	return Sum(Product(numbers, weights))
}
