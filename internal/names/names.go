package names

import (
	"math/rand"
	"strconv"
	"strings"
)

var left = [...]string{
"Michael",
"Mark",
"Matt",
"Caeleb",
"Ryan",
"Gary",
"Ian",
"Aaron",
"Nathan",
"Tom",
"Don",
"Johnny",
"Alexander",
"Roland",
"Jason",
"Jenny",
"Katie",
"Kristin",
"Amy",
"Emma",
"Krisztina",
"Dana",
"Missy",
"Dara",
"Dawn",
"Kornelia",
"Allison",
"Inge",
"Cate",
"Libby",
}

var right = [...]string {
	"amazing",
	"super",
	"fast",
	"incredible",
	"awesome",
	"brilliant",
	"bravo",
	"cool",
}

func GeneratedName(retry int) string {
	begin:
		name := left[rand.Intn(len(left))] + right[rand.Intn(len(right))] //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
		if name == "boring_wozniak" /* Steve Wozniak is not boring */ {
			goto begin
		}
	
		if retry > 0 {
			name += strconv.Itoa(rand.Intn(10)) //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
		}
		name=strings.ToLower(name)
		return name
	}