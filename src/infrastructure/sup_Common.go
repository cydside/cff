package infrastructure

//______________________________________________________________________________

// removeDupsInStringSlice Rimuovere gli elementi identici da un array di stringhe.
// ATTENZIONE: Non perdiamo la posizione degli elementi dell'array se sono
// ordinati e usiamo un altro array. Versione più chiara.
func removeDupsInStringSlice(s []string) []string {
	m := map[string]bool{}
	t := []string{}

	// Insert a map of all unique elements.
	for _, v := range s {
		if _, seen := m[v]; !seen {
			t = append(t, v)
			m[v] = true
		}
	}

	return t
}
