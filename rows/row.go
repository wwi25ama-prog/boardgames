package rows

import "strings"

type Row []string

// New erwartet eine Länge und einen String.
// Gibt eine neue `Row` zurück, die mit dem String gefüllt ist.
func New(length int, fill string) Row {
	s := strings.Repeat(fill, length)
	return strings.Split(s, "")
}

// String gibt die `Row` als String zurück.
// Die Elemente sind durch ` | ` getrennt und von `|` umgeben.
func (r Row) String() string {
	// TODO
	return ""
}

// ContainsOnly erwartet einen String.
// Gibt `true` zurück, wenn die `Row` nur aus dem String besteht.
func (r Row) ContainsOnly(s string) bool {
	// TODO
	return true
}
