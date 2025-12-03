package subs

import "strings"

func SubstituteMap(text string, valueMap map[string]any) string {
	if len(valueMap) == 0 {
		return text
	}

	// Before we get cranking, fix any escaped newlines.
	text = strings.ReplaceAll(text, "\\n", "\n")

	return handleSubstitutionMap(text, valueMap)
}

func handleSubstitutionMap(text string, subs map[string]any) string {
	// Split the string into parts based on locating the placeholder tokens surrounded by "((" and "}}"
	if !strings.Contains(text, "{{") {
		return text
	}

	parts := splitOutFormats(text)

	for idx, part := range parts {
		if !strings.HasPrefix(part, "{{") || !strings.HasSuffix(part, "}}") {
			continue
		}

		parts[idx] = handleFormat(part, subs)
	}

	text = strings.Join(parts, "")

	return text
}
