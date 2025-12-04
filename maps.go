package subs

import "strings"

// SubstituteMap is a function that replaces placeholders in a text string with
// their corresponding values. The second parameter is a map[string]any that contains
// the key-value pairs representing the placeholders and their corresponding values.
//
// The result is a string containing the input text with all replacement and formatting
// operations performed. If errors occur during the substitution, an error message is
// returned as the replacement value, delimited by "!" characters.
func SubstituteMap(text string, valueMap map[string]any) string {
	if len(valueMap) == 0 {
		return text
	}

	// Before we get cranking, fix any escaped newlines.
	text = strings.ReplaceAll(text, "\\n", "\n")

	return handleSubstitutionMap(text, valueMap)
}

// This does the work of substituting placeholders from the map. The input string
// is broken into segments of plain text and operators, and for each operator the
// formatting function is called to process the associated value from the map.
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
