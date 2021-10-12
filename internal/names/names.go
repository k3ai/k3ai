package names

import (
	"math/rand"
	"strings"
)

var names = []string{
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

var def = []string {
	"amazing",
	"super",
	"fast",
	"incredible",
	"awesome",
	"brilliant",
	"bravo",
	"cool",
}

func GeneratedName() string {
	randomIndexNames := rand.Intn(len(names))
	randomIndexDef := rand.Intn(len(def))
	pick := strings.ToLower(names[randomIndexNames] + "_" + def[randomIndexDef])
	return pick
}