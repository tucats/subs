package subs

import "strings"

// Search a string value for a "|" in quotes and if found convert it to "!BAR!".
func barEscape(text string) string {
	quote := false
	result := ""

	for _, char := range text {
		if char == '"' {
			quote = !quote
		}

		if char == '|' && quote {
			result += "!BAR!"
		} else {
			result += string(char)
		}
	}

	return result
}

// Search an array of strings and if any contain "!BAR!" convert it back to "|".
func barUnescape(parts []string) []string {
	result := make([]string, len(parts))

	for i, part := range parts {
		result[i] = strings.ReplaceAll(part, "!BAR!", "|")
	}

	return result
}
