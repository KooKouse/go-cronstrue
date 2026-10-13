package cronstrue

import (
	"regexp"
	"strconv"
	"strings"
)

// format replaces each "%s" placeholder in template with the next value from
// values, in order. It mirrors StringUtilities.format of the original
// JavaScript library, including its behavior of inserting the literal string
// "undefined" when a placeholder has no corresponding value.
func format(template string, values ...string) string {
	var b strings.Builder
	rest := template
	for {
		i := strings.Index(rest, "%s")
		if i < 0 {
			b.WriteString(rest)
			return b.String()
		}
		b.WriteString(rest[:i])
		if len(values) > 0 {
			b.WriteString(values[0])
			values = values[1:]
		} else {
			// Bug-compatible with the JavaScript original: a placeholder
			// without a corresponding value becomes "undefined".
			b.WriteString("undefined")
		}
		rest = rest[i+2:]
	}
}

// containsAny reports whether text contains any of searchStrings.
func containsAny(text string, searchStrings []string) bool {
	for _, s := range searchStrings {
		if strings.Contains(text, s) {
			return true
		}
	}
	return false
}

// jsNumber mimics JavaScript Number() conversion for the string shapes that
// can appear inside cron expressions: decimal integers/floats and scientific
// notation. It returns false when the string is not a valid number (NaN).
func jsNumber(s string) (float64, bool) {
	n, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0, false
	}
	return n, true
}

// jsNumberToString formats a number the way JavaScript's Number#toString
// does for finite decimal values (e.g. 5.0 -> "5", 5.5 -> "5.5").
func jsNumberToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

var jsIntPrefixRe = regexp.MustCompile(`^[ \t\r\n]*[+-]?[0-9]+`)

// jsParseInt mimics JavaScript parseInt(s, 10): it parses an optional sign
// followed by leading decimal digits and ignores anything after them. It
// returns false (NaN) when no digits are found at the start.
func jsParseInt(s string) (int, bool) {
	m := jsIntPrefixRe.FindString(s)
	if m == "" {
		return 0, false
	}
	n, err := strconv.Atoi(strings.TrimRight(m, " \t\r\n"))
	if err != nil {
		return 0, false
	}
	return n, true
}

// extractDigits returns only the ASCII digits of s (up to the first non-digit
// character being dropped is equivalent here because matched segments carry
// at most one digit group).
func extractDigits(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// padTwo left-pads a numeric string with "0" to a length of two, mirroring
// ("00" + s).substring(s.length) in the original library.
func padTwo(s string) string {
	for len(s) < 2 {
		s = "0" + s
	}
	return s
}

// upperFirst upper-cases the first character (Unicode aware) like
// charAt(0).toLocaleUpperCase() in the original library.
func upperFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	return strings.ToUpper(string(r[0])) + string(r[1:])
}
