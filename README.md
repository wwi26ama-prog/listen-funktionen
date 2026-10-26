# Funktionen auf Listen

In dieser Einheit entstehen kleine Funktionen auf Listen (Slices),
die aufeinander aufbauen und zum Schluss Prüfziffern von Strichcodes
und Büchern berechnen und prüfen.

Beispiele:

- [`examples/lists.go`](examples/lists.go): neue Listen aus einer
  oder zwei Listen erzeugen, jedes zweite Element auswählen.
- [`examples/check_digit.go`](examples/check_digit.go): eine einfache
  Prüfziffer ohne Gewichte und ihre Schwäche.

Übungsaufgaben:

- [`exercises/warmup.go`](exercises/warmup.go): zum Aufwärmen eine Liste
  skalieren, Ziffern erkennen und zum nächsten Vielfachen ergänzen.
- [`exercises/lists.go`](exercises/lists.go): Summe, elementweises Produkt,
  sich wiederholende Gewichte und gewichtete Summe.
- [`exercises/digits.go`](exercises/digits.go): einen Code wie
  `"978-3-16-148410-0"` in eine Liste von Ziffern umwandeln.
- [`exercises/ean.go`](exercises/ean.go): Prüfziffer einer EAN-13 (auch ISBN-13)
  berechnen und prüfen.
- [`exercises/isbn.go`](exercises/isbn.go): dasselbe für eine ISBN-10.

Die Aufgaben stehen in den Kommentaren der Funktionen, die erwarteten
Ergebnisse in den zugehörigen Tests. Bearbeiten Sie sie in dieser Reihenfolge,
spätere Funktionen verwenden die früheren.
Prüfen mit `go test ./...` in diesem Verzeichnis.
