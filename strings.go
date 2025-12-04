package subs

import "strings"

// splitOutFormats splits a string into segments of plain text and substitution expressions.
// The expressions are delimited by "{{" and "}}". The result is an array of each element of
// the original text. If the input text is an empty string, an empty array of strings is returned.
func splitOutFormats(text string) []string {
	parts := make([]string, 0)
	segments := strings.Split(text, "{{")

	for _, segment := range segments {
		if segment == "" {
			continue
		}

		if !strings.Contains(segment, "}}") {
			parts = append(parts, segment)

			continue
		}

		subparts := strings.SplitN(segment, "}}", 2)
		parts = append(parts, "{{"+subparts[0]+"}}")

		if len(subparts[1]) > 0 {
			parts = append(parts, subparts[1])
		}
	}

	return parts
}

// The value is formatted as a single comma-separated list in a string. If the
// item is a single value, it is just returned as a string. If it is an array,
// the array elements are a comma-separated list of array values expressed as
// strings. If the value is a map, the map elements are a comma-separated list
// of key: value pairs expressed as a string.
func makeList(values any, format string) string {
	return strings.Join(makeArray(values, format), ", ")
}

// The value is formatted as a string, separated by newlines if the element
// is an array or map with multiple values. If the item is a single value, it
// is just returned as a string. If it is an array, the array elements are
// separated by a newline character in a single string. If the value is a map,
// the map elements are a newline-separated list of key: value pairs expressed
// as a string.
func makeLines(values any, format string) string {
	return strings.Join(makeArray(values, format), "\n")
}
