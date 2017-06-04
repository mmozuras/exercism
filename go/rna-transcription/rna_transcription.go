package strand

import "strings"

const testVersion = 3

var toRNA = strings.NewReplacer(
	"G", "C",
	"C", "G",
	"T", "A",
	"A", "U",
)

func ToRNA(s string) string {
	return toRNA.Replace(s)
}
