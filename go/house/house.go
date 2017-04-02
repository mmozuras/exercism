package house

var testVersion = 1

var parts = []struct {
	name   string
	action string
}{
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

func Verse(num int) string {
	result := "This is "
	for i := num - 2; i >= 0; i-- {
		result += "the " + parts[i].name + "\n"
		result += "that " + parts[i].action + " "
	}
	return result + "the house that Jack built."
}

func Song() string {
	result := ""
	for i := 1; i <= 11; i++ {
		result += Verse(i) + "\n\n"
	}
	return result + Verse(12)
}
