package protein

const testVersion = 1

var codon = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(s string) string {
	return codon[s]
}

func FromRNA(rna string) []string {
	result := []string{}
	for i := 0; i < len(rna); i += 3 {
		protein := FromCodon(rna[i : i+3])
		if protein == "STOP" {
			return result
		}
		result = append(result, protein)
	}
	return result
}
