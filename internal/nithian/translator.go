package nithian

import (
	"strings"
	"unicode"
)

// Ideogram represents a single Nithian ideogram mapping
type Ideogram struct {
	Phoneme  string `json:"phoneme"`
	Meaning  string `json:"meaning"`
	Ideogram string `json:"ideogram"`
}

// ideogramMap maps phonetic components to ideogram metadata
var ideogramMap = map[string]Ideogram{
	"ah": {Phoneme: "ah", Meaning: "father", Ideogram: "water"},
	"a":  {Phoneme: "a", Meaning: "mass", Ideogram: "hand"},
	"ay": {Phoneme: "ay", Meaning: "mace", Ideogram: "mace"},
	"eh": {Phoneme: "eh", Meaning: "mess", Ideogram: "health"},
	"ee": {Phoneme: "ee", Meaning: "machine", Ideogram: "reed"},
	"ih": {Phoneme: "ih", Meaning: "miss", Ideogram: "ear"},
	"oh": {Phoneme: "oh", Meaning: "most", Ideogram: "rope"},
	"u":  {Phoneme: "u", Meaning: "put", Ideogram: "foot"},
	"oo": {Phoneme: "oo", Meaning: "moose", Ideogram: "stool"},
	"uh": {Phoneme: "uh", Meaning: "must", Ideogram: "underworld"},
	"ai": {Phoneme: "ai", Meaning: "mice", Ideogram: "eye"},
	"au": {Phoneme: "au", Meaning: "mouse", Ideogram: "house"},
	"ow": {Phoneme: "ow", Meaning: "low", Ideogram: "house"},
	"oi": {Phoneme: "oi", Meaning: "moist", Ideogram: "hoist"},
	"b":  {Phoneme: "b", Meaning: "book", Ideogram: "bread"},
	"ch": {Phoneme: "ch", Meaning: "cheer", Ideogram: "chalice"},
	"d":  {Phoneme: "d", Meaning: "dear", Ideogram: "door bolt"},
	"f":  {Phoneme: "f", Meaning: "fear", Ideogram: "feather"},
	"g":  {Phoneme: "g", Meaning: "gear", Ideogram: "god"},
	"j":  {Phoneme: "j", Meaning: "jeer", Ideogram: "justice"},
	"h":  {Phoneme: "h", Meaning: "here", Ideogram: "head"},
	"y":  {Phoneme: "y", Meaning: "year", Ideogram: "yesterday"},
	"k":  {Phoneme: "k", Meaning: "king", Ideogram: "king"},
	"l":  {Phoneme: "l", Meaning: "leer", Ideogram: "life"},
	"m":  {Phoneme: "m", Meaning: "mere", Ideogram: "mouth"},
	"n":  {Phoneme: "n", Meaning: "near", Ideogram: "knot"},
	"ng": {Phoneme: "ng", Meaning: "seeing", Ideogram: "traveling"},
	"p":  {Phoneme: "p", Meaning: "peer", Ideogram: "protection"},
	"r":  {Phoneme: "r", Meaning: "rear", Ideogram: "rule"},
	"s":  {Phoneme: "s", Meaning: "sear", Ideogram: "sun"},
	"sh": {Phoneme: "sh", Meaning: "sheer", Ideogram: "shelter"},
	"t":  {Phoneme: "t", Meaning: "tier", Ideogram: "tail"},
	"th": {Phoneme: "th", Meaning: "theory", Ideogram: "Thoth"},
	"dh": {Phoneme: "dh", Meaning: "there", Ideogram: "square"},
	"v":  {Phoneme: "v", Meaning: "veer", Ideogram: "viper"},
	"w":  {Phoneme: "w", Meaning: "weird", Ideogram: "window"},
	"z":  {Phoneme: "z", Meaning: "zero", Ideogram: "zenith"},
	"zh": {Phoneme: "zh", Meaning: "pleasure", Ideogram: "pleasure"},
	".":  {Phoneme: "ta", Meaning: "full stop", Ideogram: "obelisk"},
}

var phonemeOrder = []string{
	"ng", "sh", "ch", "th", "dh", "zh", "ai", "au", "ow", "oi", "oo", "ee", "ay", "ah", "eh", "ih", "oh", "uh",
	"x", "qu",
	"b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "u", "v", "w", "y", "z", "a", "e", "i", "o",
}

// TranslateText converts a string to a sequence of ideograms using phoneme parsing
func TranslateText(input string) []Ideogram {
	input = strings.ToLower(input)
	var result []Ideogram
	words := strings.FieldsFunc(input, func(r rune) bool {
		return unicode.IsSpace(r) || r == ',' || r == ';' || r == ':'
	})

	substitutions := map[string]string{
		"qu": "kw",
		"x":  "ks",
	}

	for _, word := range words {
		// Preprocess substitutions
		w := word
		for old, new := range substitutions {
			w = strings.ReplaceAll(w, old, new)
		}

		for len(w) > 0 {
			found := false
			for _, ph := range phonemeOrder {
				if strings.HasPrefix(w, ph) {
					if ideogram, ok := ideogramMap[ph]; ok {
						result = append(result, ideogram)
						w = w[len(ph):]
						found = true
						break
					}
				}
			}
			if !found {
				w = w[1:] // skip unknown char
			}
		}
		if strings.HasSuffix(word, ".") {
			result = append(result, ideogramMap["."])
		}
	}
	return result
}
