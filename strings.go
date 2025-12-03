package subs

import "strings"

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

func makeList(values any, format string) string {
	return strings.Join(makeArray(values, format), ", ")
}

func makeLines(values any, format string) string {
	return strings.Join(makeArray(values, format), "\n")
}
