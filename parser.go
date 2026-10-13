package cronstrue

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// cronParser parses and normalizes a cron expression, ported from
// cronstrue/src/cronParser.ts (upstream v3.x). parse returns the seven
// normalized expression parts (seconds, minutes, hours, day-of-month,
// month, day-of-week, year) or an error with the same message the
// JavaScript library produces.
type cronParser struct {
	expression              string
	dayOfWeekStartIndexZero bool
	monthStartIndexZero     bool
}

func newCronParser(expression string, dayOfWeekStartIndexZero, monthStartIndexZero bool) *cronParser {
	return &cronParser{
		expression:              expression,
		dayOfWeekStartIndexZero: dayOfWeekStartIndexZero,
		monthStartIndexZero:     monthStartIndexZero,
	}
}

var (
	spaceSplitRe = regexp.MustCompile(` +`)
	dowDigitsRe  = regexp.MustCompile(`(^\d)|([^#/\s]\d)`)
	// Note: Go's \s matches [\t\n\f\r ] like JavaScript's \s for ASCII input.
	monthDigitsRe = regexp.MustCompile(`(^\d{1,2})|([^#/\s\d]\d{1,2})`)
	yearAtEndRe   = regexp.MustCompile(`\d{4}$`)
	stepStartRe   = regexp.MustCompile(`^\*|-|,`)
	weekdayOffRe  = regexp.MustCompile(`^\d{1,2}W[+-]\d{1,2}$`)
)

var dowNameToNum = []struct {
	name string
	num  string
}{
	{"SUN", "0"}, {"MON", "1"}, {"TUE", "2"}, {"WED", "3"},
	{"THU", "4"}, {"FRI", "5"}, {"SAT", "6"},
}

var monthNameToNum = []struct {
	name string
	num  string
}{
	{"JAN", "1"}, {"FEB", "2"}, {"MAR", "3"}, {"APR", "4"},
	{"MAY", "5"}, {"JUN", "6"}, {"JUL", "7"}, {"AUG", "8"},
	{"SEP", "9"}, {"OCT", "10"}, {"NOV", "11"}, {"DEC", "12"},
}

// parse parses and normalizes a cron expression into an array of seven
// strings.
func (p *cronParser) parse() ([]string, error) {
	var parsed []string
	var err error

	expression := p.expression

	if expression == "@reboot" {
		return []string{"@reboot", "", "", "", "", "", ""}, nil
	}
	if strings.HasPrefix(expression, "@") {
		special, err := p.parseSpecial(p.expression)
		if err != nil {
			return nil, err
		}
		parsed, err = p.extractParts(special)
		if err != nil {
			return nil, err
		}
	} else {
		parsed, err = p.extractParts(p.expression)
		if err != nil {
			return nil, err
		}
	}

	if err := p.normalize(parsed); err != nil {
		return nil, err
	}
	if err := p.validate(parsed); err != nil {
		return nil, err
	}

	return parsed, nil
}

func (p *cronParser) parseSpecial(expression string) (string, error) {
	specialExpressions := map[string]string{
		"@yearly":   "0 0 1 JAN *",
		"@annually": "0 0 1 JAN *",
		"@monthly":  "0 0 1 * *",
		"@weekly":   "0 0 * * SUN",
		"@daily":    "0 0 * * *",
		"@midnight": "0 0 * * *",
		"@hourly":   "0 * * * *",
		"@reboot":   "@reboot",
	}

	special, ok := specialExpressions[expression]
	if !ok {
		return "", errors.New("Unknown special expression.")
	}
	return special, nil
}

func (p *cronParser) extractParts(expression string) ([]string, error) {
	if p.expression == "" {
		return nil, errors.New("cron expression is empty")
	}

	parsed := spaceSplitRe.Split(strings.TrimSpace(expression), -1)

	for i := range parsed {
		if strings.Contains(parsed[i], ",") {
			// Items are converted to numbers when possible (dropping
			// leading zeros, e.g. "05" -> "5") and sorted numerically;
			// non-numeric items keep their relative order, mirroring the
			// JavaScript comparator returning NaN (0) for mixed pairs.
			type item struct {
				num     float64
				str     string
				numeric bool
			}
			var items []item
			for _, raw := range strings.Split(parsed[i], ",") {
				raw = strings.TrimSpace(raw)
				if raw == "" {
					continue
				}
				if n, ok := jsNumber(raw); ok {
					items = append(items, item{num: n, str: jsNumberToString(n), numeric: true})
				} else {
					items = append(items, item{str: raw})
				}
			}
			if len(items) == 0 {
				items = append(items, item{str: "*"})
			}
			sort.SliceStable(items, func(a, b int) bool {
				if items[a].numeric && items[b].numeric {
					return items[a].num < items[b].num
				}
				return false
			})
			strs := make([]string, len(items))
			for j, it := range items {
				strs[j] = it.str
			}
			parsed[i] = strings.Join(strs, ",")
		}
	}

	switch {
	case len(parsed) < 5:
		plural := "s"
		if len(parsed) == 1 {
			plural = ""
		}
		return nil, fmt.Errorf("Expression has only %d part%s. At least 5 parts are required.", len(parsed), plural)
	case len(parsed) == 5:
		parsed = append([]string{""}, parsed...)
		parsed = append(parsed, "")
	case len(parsed) == 6:
		isYearWithNoSecondsPart := yearAtEndRe.MatchString(parsed[5]) || parsed[4] == "?" || parsed[2] == "?"
		if isYearWithNoSecondsPart {
			parsed = append([]string{""}, parsed...)
		} else {
			parsed = append(parsed, "")
		}
	case len(parsed) > 7:
		return nil, fmt.Errorf("Expression has %d parts; too many!", len(parsed))
	}

	return parsed, nil
}

func (p *cronParser) normalize(expressionParts []string) error {
	// Convert ? to * for DOM and DOW.
	expressionParts[3] = strings.Replace(expressionParts[3], "?", "*", 1)
	expressionParts[5] = strings.Replace(expressionParts[5], "?", "*", 1)

	// Convert ? to * for Hour. ? isn't valid for hour position but we can work around it.
	expressionParts[2] = strings.Replace(expressionParts[2], "?", "*", 1)

	// Convert 0/, 1/ to */
	if strings.HasPrefix(expressionParts[0], "0/") {
		expressionParts[0] = strings.Replace(expressionParts[0], "0/", "*/", 1)
	}
	if strings.HasPrefix(expressionParts[1], "0/") {
		expressionParts[1] = strings.Replace(expressionParts[1], "0/", "*/", 1)
	}
	if strings.HasPrefix(expressionParts[2], "0/") {
		expressionParts[2] = strings.Replace(expressionParts[2], "0/", "*/", 1)
	}
	if strings.HasPrefix(expressionParts[3], "1/") {
		expressionParts[3] = strings.Replace(expressionParts[3], "1/", "*/", 1)
	}
	if !p.monthStartIndexZero && strings.HasPrefix(expressionParts[4], "1/") {
		// Month (with monthStartIndexZero, 1/ starts in February so it is not the same as */)
		expressionParts[4] = strings.Replace(expressionParts[4], "1/", "*/", 1)
	}
	if strings.HasPrefix(expressionParts[6], "1/") {
		expressionParts[6] = strings.Replace(expressionParts[6], "1/", "*/", 1)
	}

	// Adjust DOW based on dayOfWeekStartIndexZero option.
	// Normalized DOW: 0=Sunday/6=Saturday
	expressionParts[5] = dowDigitsRe.ReplaceAllStringFunc(expressionParts[5], func(t string) string {
		// skip anything preceded by # or /
		dowDigits := extractDigits(t) // extract digit part (i.e. if "-2" or ",2", just take 2)
		dowDigitsAdjusted := dowDigits

		if p.dayOfWeekStartIndexZero {
			// "7" also means Sunday so we will convert to "0" to normalize it
			if dowDigits == "7" {
				dowDigitsAdjusted = "0"
			}
		} else {
			// If dayOfWeekStartIndexZero==false, Sunday is specified as 1 and Saturday is specified as 7.
			// To normalize, we will shift the DOW number down so that 1 becomes 0, 2 becomes 1, and so on.
			if n, ok := jsParseInt(dowDigits); ok {
				dowDigitsAdjusted = fmt.Sprintf("%d", n-1)
			}
		}

		return strings.Replace(t, dowDigits, dowDigitsAdjusted, 1)
	})

	// Convert DOW 'L' to '6' (Saturday)
	if expressionParts[5] == "L" {
		expressionParts[5] = "6"
	}

	// Convert DOM '?' to '*'
	if expressionParts[3] == "?" {
		expressionParts[3] = "*"
	}

	// i.e. 15W-2 or 15W+2 (offset from the nearest weekday)
	if strings.Contains(expressionParts[3], "W") &&
		(strings.Contains(expressionParts[3], ",") || strings.ContainsAny(expressionParts[3], "-+")) &&
		!weekdayOffRe.MatchString(expressionParts[3]) {
		return errors.New("The 'W' character can be specified only when the day-of-month is a single day, not a range or list of days.")
	}

	// Convert DOW SUN-SAT format to 0-6 format
	for _, d := range dowNameToNum {
		expressionParts[5] = regexp.MustCompile("(?i)"+d.name).ReplaceAllString(expressionParts[5], d.num)
	}

	// Adjust month based on monthStartIndexZero option.
	// Normalized Month: 1=JAN/12=DEC
	expressionParts[4] = monthDigitsRe.ReplaceAllStringFunc(expressionParts[4], func(t string) string {
		// skip anything preceded by # or /
		monthDigits := extractDigits(t)
		monthDigitsAdjusted := monthDigits

		if p.monthStartIndexZero {
			// if monthStartIndexZero==true, we will shift month number up so that JAN=1 and DEC=12
			if n, ok := jsParseInt(monthDigits); ok {
				monthDigitsAdjusted = fmt.Sprintf("%d", n+1)
			}
		}

		return strings.Replace(t, monthDigits, monthDigitsAdjusted, 1)
	})

	// Convert JAN-DEC format to 1-12 format
	for _, m := range monthNameToNum {
		expressionParts[4] = regexp.MustCompile("(?i)"+m.name).ReplaceAllString(expressionParts[4], m.num)
	}

	// Convert 0 second to (empty)
	if expressionParts[0] == "0" {
		expressionParts[0] = ""
	}

	// If time increment or * (every) is specified for seconds or minutes and hours part is single item, make it a "self-range" so
	// the expression can be interpreted as an increment / range. This will allow us to easily interpret an hour part as 'between' a second or minute duration.
	if !strings.ContainsAny(expressionParts[2], "*-,/") &&
		(strings.ContainsAny(expressionParts[1], "*/") || strings.ContainsAny(expressionParts[0], "*/")) {
		expressionParts[2] += "-" + expressionParts[2]
	}

	// Loop through all parts and apply global normalization
	for i := range expressionParts {
		// ignore empty characters around comma
		// if nothing left, set it to *
		if strings.Contains(expressionParts[i], ",") {
			var kept []string
			for _, s := range strings.Split(expressionParts[i], ",") {
				if s != "" {
					kept = append(kept, s)
				}
			}
			joined := strings.Join(kept, ",")
			if joined == "" {
				joined = "*"
			}
			expressionParts[i] = joined
		}

		// convert all '*/1' to '*'
		if expressionParts[i] == "*/1" {
			expressionParts[i] = "*"
		}

		// Convert Month,DOW,Year step values with a starting value (i.e. not '*') to range expressions.
		// This allows us to reuse the range expression handling for step values.
		//   For example:
		//   - month part '3/2' will be converted to '3-12/2' (every 2 months between March and December)
		//   - DOW part '3/2' will be converted to '3-6/2' (every 2 days between Tuesday and Saturday)
		if strings.Contains(expressionParts[i], "/") && !stepStartRe.MatchString(expressionParts[i]) {
			var stepRangeThrough string
			switch i {
			case 4:
				stepRangeThrough = "12"
			case 5:
				stepRangeThrough = "6"
			case 6:
				stepRangeThrough = "9999"
			}
			if stepRangeThrough != "" {
				parts := strings.Split(expressionParts[i], "/")
				expressionParts[i] = parts[0] + "-" + stepRangeThrough + "/" + parts[1]
			}
		}
	}

	return nil
}

var (
	domOffsetStripRe  = regexp.MustCompile(`^(\d{1,2}W)[+-]\d{1,2}$`)
	standardCharsRe   = regexp.MustCompile(`[^0-9,\-*/]+`)
	domAllowedCharsRe = regexp.MustCompile(`[^0-9,\-*/LW]+`)
	dowAllowedCharsRe = regexp.MustCompile(`[^0-9,\-*/L#]+`)
)

func (p *cronParser) validate(parsed []string) error {
	if err := p.validateOnlyExpectedCharactersFound(parsed[0], standardCharsRe); err != nil {
		return err
	}
	if err := p.validateOnlyExpectedCharactersFound(parsed[1], standardCharsRe); err != nil {
		return err
	}
	if err := p.validateOnlyExpectedCharactersFound(parsed[2], standardCharsRe); err != nil {
		return err
	}
	// DOM
	if err := p.validateOnlyExpectedCharactersFound(domOffsetStripRe.ReplaceAllString(parsed[3], "$1"), domAllowedCharsRe); err != nil {
		return err
	}
	if err := p.validateOnlyExpectedCharactersFound(parsed[4], standardCharsRe); err != nil {
		return err
	}
	// DOW
	if err := p.validateOnlyExpectedCharactersFound(parsed[5], dowAllowedCharsRe); err != nil {
		return err
	}
	if err := p.validateOnlyExpectedCharactersFound(parsed[6], standardCharsRe); err != nil {
		return err
	}

	return p.validateAnyRanges(parsed)
}

func (p *cronParser) validateAnyRanges(parsed []string) error {
	if err := secondRange(parsed[0]); err != nil {
		return err
	}
	if err := minuteRange(parsed[1]); err != nil {
		return err
	}
	if err := hourRange(parsed[2]); err != nil {
		return err
	}
	if err := dayOfMonthRange(parsed[3]); err != nil {
		return err
	}
	if err := monthRange(parsed[4], p.monthStartIndexZero); err != nil {
		return err
	}
	return dayOfWeekRange(parsed[5], p.dayOfWeekStartIndexZero)
}

func (p *cronParser) validateOnlyExpectedCharactersFound(cronPart string, invalidCharsRe *regexp.Regexp) error {
	invalidChars := invalidCharsRe.FindAllString(cronPart, -1)
	if len(invalidChars) > 0 {
		return fmt.Errorf("Expression contains invalid values: '%s'", strings.Join(invalidChars, ","))
	}
	return nil
}

// rangeValidator port (src/rangeValidator.ts). The functions return errors
// with the same messages as the JavaScript original.

func secondRange(parse string) error {
	for _, s := range strings.Split(parse, ",") {
		if second, ok := jsParseInt(s); ok {
			if second < 0 || second > 59 {
				return errors.New("seconds part must be >= 0 and <= 59")
			}
		}
	}
	return nil
}

func minuteRange(parse string) error {
	for _, s := range strings.Split(parse, ",") {
		if minute, ok := jsParseInt(s); ok {
			if minute < 0 || minute > 59 {
				return errors.New("minutes part must be >= 0 and <= 59")
			}
		}
	}
	return nil
}

func hourRange(parse string) error {
	for _, s := range strings.Split(parse, ",") {
		if hour, ok := jsParseInt(s); ok {
			if hour < 0 || hour > 23 {
				return errors.New("hours part must be >= 0 and <= 23")
			}
		}
	}
	return nil
}

func dayOfMonthRange(parse string) error {
	for _, s := range strings.Split(parse, ",") {
		if dayOfMonth, ok := jsParseInt(s); ok {
			if dayOfMonth < 1 || dayOfMonth > 31 {
				return errors.New("DOM part must be >= 1 and <= 31")
			}
		}
	}
	return nil
}

func monthRange(parse string, monthStartIndexZero bool) error {
	for _, s := range strings.Split(parse, ",") {
		if month, ok := jsParseInt(s); ok {
			if month < 1 || month > 12 {
				if monthStartIndexZero {
					return errors.New("month part must be >= 0 and <= 11")
				}
				return errors.New("month part must be >= 1 and <= 12")
			}
		}
	}
	return nil
}

func dayOfWeekRange(parse string, dayOfWeekStartIndexZero bool) error {
	for _, s := range strings.Split(parse, ",") {
		if dayOfWeek, ok := jsParseInt(s); ok {
			if dayOfWeek < 0 || dayOfWeek > 6 {
				if dayOfWeekStartIndexZero {
					return errors.New("DOW part must be >= 0 and <= 6")
				}
				return errors.New("DOW part must be >= 1 and <= 7")
			}
		}
	}
	return nil
}
