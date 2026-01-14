package subs

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// handleFormat processes a substitution operator using the supplied list of
// replacement values in the map. IF the string passed does not have a format
// operator, the text is returned as-is. Otherwise, the substitution value
// is read from the map and any formatting operations applied to the value.
// The value is then returned as a string.
func handleFormat(text string, subs map[string]any) string {
	var (
		result string
		err    error
	)

	if !strings.HasPrefix(text, "{{") || !strings.HasSuffix(text, "}}") {
		return text
	}

	key := strings.TrimSuffix(strings.TrimPrefix(text, "{{"), "}}")

	key, format, found := strings.Cut(key, "|")
	if !found || format == "" {
		format = "%v"
	}

	value, ok := subs[key]
	if !ok {
		value = "!" + key + "!"
		format = "%s"
	}

	// Check for special cases in the format string
	formatParts := barUnescape(strings.Split(barEscape(format), "|"))

	label := ""
	format = "%v"

	for _, part := range formatParts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		switch {
		case strings.HasPrefix(part, "size "):
			sizeParm := strings.TrimSpace(part[len("size "):])
			if size, err := strconv.Atoi(sizeParm); err == nil && size > 4 {
				if format == "" {
					text = value.(string)
				} else {
					text = fmt.Sprintf(format, value)
				}

				if len(text) > size {
					value = text[:size-3] + "..."
				}
			} else {
				value = "!Invalid size: " + sizeParm + "!"
			}

			format = ""

		case part == "lines":
			value = makeLines(value, format)
			format = ""

		case strings.HasPrefix(part, "%"):
			format = part

		case strings.HasPrefix(part, "zero "):
			text := strings.TrimSpace(part[len("zero "):])
			if unquoted, err := strconv.Unquote(text); err == nil {
				text = unquoted
			}

			if getInt(value) == 0 {
				value = text
				label = ""
				format = "%s"
			}

		case strings.HasPrefix(part, "one "):
			text = strings.TrimSpace(part[len("one "):])
			if unquoted, err := strconv.Unquote(text); err == nil {
				text = unquoted
			}

			if getInt(value) == 1 {
				value = text
				label = ""
				format = "%s"
			}

		// Note, fully spelled out name "cardinal" must immediately be followed by
		// short form "card ".
		case strings.HasPrefix(part, "cardinal "):
			part = "card " + part[len("cardinal "):]

			fallthrough

		case strings.HasPrefix(part, "card "):
			parts := strings.Split(part[len("card "):], ",")
			if len(parts) == 2 {
				parts = []string{parts[1], parts[0], parts[1]}
			}

			if len(parts) < 1 || len(parts) > 3 {
				return "!Invalid card format: " + part + "!"
			}

			count := getInt(value)

			switch {
			case count == 0:
				text = strings.TrimSpace(parts[0])
				if unquotedText, err := strconv.Unquote(text); err == nil {
					value = unquotedText
				} else {
					value = text
				}

				label = ""
				format = "%s"

			case count == 1 && len(parts) > 1:
				text = strings.TrimSpace(parts[1])
				if unquotedText, err := strconv.Unquote(text); err == nil {
					value = unquotedText
				} else {
					value = text
				}

				label = ""
				format = "%s"

			case count > 1 && len(parts) > 2:
				text = strings.TrimSpace(parts[2])
				if unquotedText, err := strconv.Unquote(text); err == nil {
					value = unquotedText
				} else {
					value = text
				}

				label = ""
				format = "%s"

			default:
				value = count
				format = "%d"
			}

		case strings.HasPrefix(part, "many "):
			text = strings.TrimSpace(part[len("many "):])
			if unquoted, err := strconv.Unquote(text); err == nil {
				text = unquoted
			}

			if getInt(value) > 1 {
				value = text
				label = ""
				format = "%s"
			}

		case strings.HasPrefix(part, "label "):
			if !isZeroValue(value) {
				label = strings.TrimSpace(part[len("label "):])

				if unquoted, err := strconv.Unquote(label); err == nil {
					label = unquoted
				}
			} else {
				label = ""
				format = "%s"
				value = ""
			}

		case strings.HasPrefix(part, "pad "):
			pad := strings.TrimSpace(part[len("pad "):])
			if strings.HasPrefix(pad, "\"") {
				pad, _ = strconv.Unquote(pad)
			}

			var count int

			switch v := value.(type) {
			case int:
				count = v

			case float64:
				count = int(math.Round(v))

			case string:
				count, err = strconv.Atoi(v)
				if err != nil || count < 0 {
					return "!Invalid pad count: " + part + "!"
				}

			default:
				return "!Invalid pad type: " + part + "!"
			}

			if err != nil || count < 0 {
				return "!Invalid pad count: " + part + "!"
			}

			value = strings.Repeat(pad, count)

		case strings.HasPrefix(part, "left "):
			pad := strings.TrimSpace(part[len("left "):])

			count, err := strconv.Atoi(pad)
			if err != nil || count < 0 {
				return "!Invalid left count: " + part + "!"
			}

			var text string

			if format == "" {
				text = value.(string)
			} else {
				text = fmt.Sprintf(format, value)
			}

			for len(text) < count {
				text = text + " "
			}

			value = text

		case strings.HasPrefix(part, "right "):
			pad := strings.TrimSpace(part[len("left "):])

			count, err := strconv.Atoi(pad)
			if err != nil || count < 0 {
				return "!Invalid left count: " + part + "!"
			}

			var text string

			if format == "" {
				text = value.(string)
			} else {
				text = fmt.Sprintf(format, value)
			}

			for len(text) < count {
				text = " " + text
			}

			value = text

		case strings.HasPrefix(part, "center "):
			pad := strings.TrimSpace(part[len("center "):])

			count, err := strconv.Atoi(pad)
			if err != nil || count < 0 {
				return "!Invalid center count: " + part + "!"
			}

			var text string

			if format == "" {
				text = value.(string)
			} else {
				text = fmt.Sprintf(format, value)
			}

			isLeft := false

			for len(text) < count {
				if isLeft {
					text = " " + text
				} else {
					text = text + " "
				}

				isLeft = !isLeft
			}

			value = text
			format = ""

		case strings.HasPrefix(part, "format "):
			format = strings.TrimSpace(part[len("format "):])

		case strings.HasPrefix(part, "empty"):
			replacement := strings.TrimSpace(part[len("empty"):])

			if unquoted, err := strconv.Unquote(replacement); err == nil {
				replacement = unquoted
			}

			if isZeroValue(value) {
				value = replacement
				format = "%s"
			}

		case strings.HasPrefix(part, "list"):
			value = makeList(value, format)
			format = ""

		case strings.HasPrefix(part, "nonempty"):
			replacement := strings.TrimSpace(part[len("nonempty"):])

			if unquoted, err := strconv.Unquote(replacement); err == nil {
				replacement = unquoted
			}

			if !isZeroValue(value) {
				value = replacement
				format = "%s"
			}

		default:
			return "!Invalid format: " + part + "!"
		}
	}

	if format == "" {
		result = fmt.Sprintf("%s%s", label, value)
	} else {
		value = normalizeForFormat(format, value)
		result = fmt.Sprintf("%s"+format, label, value)
	}

	return result
}
