package subs

import (
	"encoding/json"
	"strings"

	"github.com/tucats/jaxon"
)

// Substitution is a function that replaces placeholders in a text string with
// their corresponding values. The function supports arbitrary values, and assumes
// the substitution expressions in the string match the json query language for
// Ego. The item(s) are retrieved from a JSON expression of the values passed in.
// Note that if multiple values are passed in, the query must assume that top
// level type is an array.
func Substitution(text string, values ...any) string {
	var (
		b   []byte
		err error
	)

	if len(values) == 0 {
		return text
	}

	if len(values) == 1 {
		if m, ok := values[0].(map[string]any); ok {
			return handleSubstitutionMap(text, m)
		}

		v := values[0]
		b, err = json.Marshal(v)
	} else {
		b, err = json.Marshal(values)
	}

	if err != nil {
		return "!" + err.Error() + "!"
	}

	// Right, it's an arbitrary item. Scan over the text looking for any substitution
	// operations and build a new map based on those strings from the parsed JSON data.
	m := map[string]any{}

	// Break the source text into segments based on the delimiters.
	const (
		Tag  = "<<-TAG->>"
		Head = "-HEAD-"
		Tail = "-TAIL-"
	)

	buffer := strings.ReplaceAll(text, "{{", Tag+Head)
	buffer = strings.ReplaceAll(buffer, "}}", Tail+Tag)
	parts := strings.Split(buffer, Tag)

	for _, part := range parts {
		if strings.HasPrefix(part, Head) && strings.HasSuffix(part, Tail) {
			expression := strings.TrimSpace(part[len(Head) : len(part)-len(Tail)])
			expressionElements := strings.SplitN(expression, "|", 2)
			key := expressionElements[0]

			value, err := jaxon.GetItem(string(b), key)
			if err != nil {
				return "!" + err.Error() + "!"
			}

			m[key] = value
		}
	}

	return SubstituteMap(text, m)
}
