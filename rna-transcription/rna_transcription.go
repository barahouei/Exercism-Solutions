// Package strand works with DNA and RNA.
package strand

// ToRNA takes a DNA strand and returns its RNA correspond.
func ToRNA(dna string) string {
	if dna == "" {
		return ""
	}

	nucleotides := map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}

	var rna string

	for _, v := range dna {
		nucleotide, ok := nucleotides[v]
		if ok {
			rna += nucleotide
		}
	}

	return rna
}
