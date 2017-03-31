package twelve

const testVersion = 1

var gifts = []struct {
	day, gift string
}{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	song := ""
	for i := 1; i <= len(gifts); i++ {
		song += Verse(i) + "\n"
	}
	return song
}

func Verse(verse int) string {
	result := "On the " + gifts[verse-1].day + " day of Christmas my true love gave to me, "
	for i := verse; i > 0; i-- {
		result += gifts[i-1].gift
		if i == 2 {
			result += ", and "
		} else if i > 2 {
			result += ", "
		}

	}
	return result + "."
}
