package cronstrue

import (
	"regexp"
	"strconv"
	"strings"
)

// expressionDescriptor converts a parsed cron expression into a human
// readable description, ported from cronstrue/src/expressionDescriptor.ts.
type expressionDescriptor struct {
	expression      string
	expressionParts []string
	options         *Options
	i18n            locale
}

var specialCharacters = []string{"/", "-", ",", "*"}

func newExpressionDescriptor(expression string, options *Options) *expressionDescriptor {
	localeName := options.LocaleName
	if localeName == "" {
		localeName = defaultLocaleName
	}

	i18n, ok := locales[localeName]
	if !ok {
		// fall back to the default locale (English), mirroring upstream
		// which falls back to the first registered locale.
		localeName = defaultLocaleName
		i18n = locales[defaultLocaleName]
	}
	options.LocaleName = localeName

	if !options.Use24HourTimeFormatSet {
		// if use24HourTimeFormat not specified, set based on locale default
		options.Use24HourTimeFormat = i18n.use24HourTimeFormatByDefault()
	}

	return &expressionDescriptor{
		expression: expression,
		options:    options,
		i18n:       i18n,
	}
}

// getFullDescription returns the description of the expression or an error
// describing why the expression could not be parsed.
func (d *expressionDescriptor) getFullDescription() (string, error) {
	description := ""

	parser := newCronParser(d.expression, d.options.DayOfWeekStartIndexZero, d.options.MonthStartIndexZero)
	parts, err := parser.parse()
	if err != nil {
		return "", err
	}
	d.expressionParts = parts

	// Special handling for @reboot
	if d.expressionParts[0] == "@reboot" {
		if s, ok := d.i18n.atReboot(); ok {
			return s, nil
		}
		return "Run once, at startup", nil
	}

	timeSegment := d.getTimeOfDayDescription()
	dayOfMonthDesc := d.getDayOfMonthDescription()
	monthDesc := d.getMonthDescription()
	dayOfWeekDesc := d.getDayOfWeekDescription()
	yearDesc := d.getYearDescription()

	description += timeSegment + dayOfMonthDesc + dayOfWeekDesc + monthDesc + yearDesc
	description = d.transformVerbosity(description, d.options.Verbose)

	// uppercase first character
	description = upperFirst(description)

	return description, nil
}

func (d *expressionDescriptor) getTimeOfDayDescription() string {
	secondsExpression := d.expressionParts[0]
	minuteExpression := d.expressionParts[1]
	hourExpression := d.expressionParts[2]

	description := ""

	switch {
	case !containsAny(minuteExpression, specialCharacters) &&
		!containsAny(hourExpression, specialCharacters) &&
		!containsAny(secondsExpression, specialCharacters):
		// specific time of day (i.e. 10 14)
		description += d.i18n.atSpace() + d.formatTime(hourExpression, minuteExpression, secondsExpression)
	case secondsExpression == "" &&
		strings.Contains(minuteExpression, "-") &&
		!strings.Contains(minuteExpression, ",") &&
		!strings.Contains(minuteExpression, "/") &&
		!containsAny(hourExpression, specialCharacters):
		// minute range in single hour (i.e. 0-10 11)
		minuteParts := strings.Split(minuteExpression, "-")
		description += format(
			d.i18n.everyMinuteBetweenX0AndX1(),
			d.formatTime(hourExpression, minuteParts[0], ""),
			d.formatTime(hourExpression, minuteParts[1], ""),
		)
	case secondsExpression == "" &&
		strings.Contains(hourExpression, ",") &&
		!strings.Contains(hourExpression, "-") &&
		!strings.Contains(hourExpression, "/") &&
		!containsAny(minuteExpression, specialCharacters):
		// hours list with single minute (i.e. 30 6,14,16)
		hourParts := strings.Split(hourExpression, ",")
		description += d.i18n.at()

		for i := 0; i < len(hourParts); i++ {
			if description != "" {
				description += " "
			}
			description += d.formatTime(hourParts[i], minuteExpression, "")

			if i < len(hourParts)-2 {
				description += ","
			}

			if i == len(hourParts)-2 {
				description += d.i18n.spaceAnd()
			}
		}
	default:
		// default time description
		secondsDescription := d.getSecondsDescription()
		minutesDescription := d.getMinutesDescription()
		hoursDescription := d.getHoursDescription()

		description += secondsDescription

		if description != "" && minutesDescription != "" {
			description += ", "
		}

		description += minutesDescription

		if minutesDescription == hoursDescription {
			return description
		}

		if description != "" && hoursDescription != "" {
			description += ", "
		}

		description += hoursDescription
	}

	return description
}

func (d *expressionDescriptor) getSecondsDescription() string {
	return d.getSegmentDescription(
		d.expressionParts[0],
		d.i18n.everySecond(),
		func(s string, form int) string {
			return s
		},
		func(s string) string {
			return format(d.i18n.everyX0Seconds(s), s)
		},
		func(s string) string {
			return d.i18n.secondsX0ThroughX1PastTheMinute()
		},
		func(s string) string {
			if s == "0" {
				return ""
			}
			n, _ := jsParseInt(s)
			if n < 20 {
				return d.i18n.atX0SecondsPastTheMinute(s)
			}
			if t, ok := d.i18n.atX0SecondsPastTheMinuteGt20(); ok {
				return t
			}
			return d.i18n.atX0SecondsPastTheMinute(s)
		},
		&segmentRange{min: 0, max: 59},
	)
}

func (d *expressionDescriptor) getMinutesDescription() string {
	secondsExpression := d.expressionParts[0]
	hourExpression := d.expressionParts[2]

	return d.getSegmentDescription(
		d.expressionParts[1],
		d.i18n.everyMinute(),
		func(s string, form int) string {
			return s
		},
		func(s string) string {
			return format(d.i18n.everyX0Minutes(s), s)
		},
		func(s string) string {
			return d.i18n.minutesX0ThroughX1PastTheHour()
		},
		func(s string) string {
			if s == "0" && !strings.Contains(hourExpression, "/") && secondsExpression == "" {
				return d.i18n.everyHour()
			}
			if s == "0" {
				if t, ok := d.i18n.onTheHour(); ok {
					return t
				}
				return d.i18n.atX0MinutesPastTheHour(s)
			}
			n, _ := jsParseInt(s)
			if n < 20 {
				return d.i18n.atX0MinutesPastTheHour(s)
			}
			if t, ok := d.i18n.atX0MinutesPastTheHourGt20(); ok {
				return t
			}
			return d.i18n.atX0MinutesPastTheHour(s)
		},
		&segmentRange{min: 0, max: 59},
	)
}

func (d *expressionDescriptor) getHoursDescription() string {
	expression := d.expressionParts[2]

	// Collect the end values of hour ranges (e.g., "7-19" -> "19") to
	// determine if ":59" should be used instead of ":00" for the last minute
	// of the range.
	hourIndex := 0
	type rangeEndValue struct {
		value string
		index int
	}
	var rangeEndValues []rangeEndValue
	for _, r := range strings.Split(strings.Split(expression, "/")[0], ",") {
		rangeParts := strings.Split(r, "-")
		if len(rangeParts) == 2 {
			rangeEndValues = append(rangeEndValues, rangeEndValue{value: rangeParts[1], index: hourIndex + 1})
		}
		hourIndex += len(rangeParts)
	}

	evaluationIndex := 0
	return d.getSegmentDescription(
		expression,
		d.i18n.everyHour(),
		func(s string, form int) string {
			isRangeEnd := false
			for _, r := range rangeEndValues {
				if r.value == s && r.index == evaluationIndex {
					isRangeEnd = true
					break
				}
			}
			isRangeEndWithNonZeroMinute := isRangeEnd && d.expressionParts[1] != "0"

			evaluationIndex++

			if isRangeEndWithNonZeroMinute {
				return d.formatTime(s, "59", "")
			}
			return d.formatTime(s, "0", "")
		},
		func(s string) string {
			return format(d.i18n.everyX0Hours(s), s)
		},
		func(s string) string {
			return d.i18n.betweenX0AndX1()
		},
		func(s string) string {
			return d.i18n.atX0()
		},
		&segmentRange{min: 0, max: 23},
	)
}

func (d *expressionDescriptor) getDayOfWeekDescription() string {
	daysOfWeekNames := d.i18n.daysOfTheWeek()

	var description string
	if d.expressionParts[5] == "*" {
		// DOW is specified as * so we will not generate a description and defer to DOM part.
		// Otherwise, we could get a contradiction like "on day 1 of the month, every day"
		// or a dupe description like "every day, every day".
		description = ""
	} else {
		description = d.getSegmentDescription(
			d.expressionParts[5],
			d.i18n.commaEveryDay(),
			func(s string, form int) string {
				exp := s
				if i := strings.Index(s, "#"); i > -1 {
					exp = s[:i]
				} else if strings.Contains(s, "L") {
					exp = strings.Replace(exp, "L", "", 1)
				}

				parsedExp, _ := jsParseInt(exp)

				dayDescription := dayOfWeekName(daysOfWeekNames, parsedExp)
				if names, ok := d.i18n.daysOfTheWeekInCase(form); ok {
					dayDescription = dayOfWeekName(names, parsedExp)
				}

				if i := strings.Index(s, "#"); i > -1 {
					dayOfWeekOfMonthNumber := s[i+1:]
					dayOfWeekNumber := s[:i]

					// Note: when the nth-week number is out of the 1-5
					// range, the JavaScript original concatenates null
					// (the string "null"); we keep that behavior.
					dayOfWeekOfMonthDescription := "null"
					switch dayOfWeekOfMonthNumber {
					case "1":
						dayOfWeekOfMonthDescription = d.i18n.first(dayOfWeekNumber)
					case "2":
						dayOfWeekOfMonthDescription = d.i18n.second(dayOfWeekNumber)
					case "3":
						dayOfWeekOfMonthDescription = d.i18n.third(dayOfWeekNumber)
					case "4":
						dayOfWeekOfMonthDescription = d.i18n.fourth(dayOfWeekNumber)
					case "5":
						dayOfWeekOfMonthDescription = d.i18n.fifth(dayOfWeekNumber)
					}

					dayDescription = dayOfWeekOfMonthDescription + " " + dayDescription
				}

				return dayDescription
			},
			func(s string) string {
				if n, ok := jsParseInt(s); ok && n == 1 {
					// rather than "every 1 days" just return empty string
					return ""
				}
				return format(d.i18n.commaEveryX0DaysOfTheWeek(s), s)
			},
			func(s string) string {
				// If both DOM and DOW are specified, the cron will execute at either time.
				beginFrom := s[:strings.Index(s, "-")]

				domSpecified := d.expressionParts[3] != "*"
				if domSpecified {
					return d.i18n.commaAndX0ThroughX1(beginFrom)
				}
				return d.i18n.commaX0ThroughX1(beginFrom)
			}, func(s string) string {
				if strings.Contains(s, "#") {
					dayOfWeekOfMonthNumber := s[strings.Index(s, "#")+1:]
					dayOfWeek := s[:strings.Index(s, "#")]
					return strings.TrimSpace(d.i18n.commaOnThe(dayOfWeekOfMonthNumber, dayOfWeek)) + d.i18n.spaceX0OfTheMonth()
				}
				if strings.Contains(s, "L") {
					return d.i18n.commaOnTheLastX0OfTheMonth(strings.Replace(s, "L", "", 1))
				}
				// If both DOM and DOW are specified, the cron executes depending on the cron format.
				// For the standard crontab format, the cron will execute at either time. For some
				// custom formats, the cron will execute when both times are satisfied.
				domSpecified := d.expressionParts[3] != "*"
				if !domSpecified || d.options.LogicalAndDayFields {
					return d.i18n.commaOnlyOnX0(s)
				}
				return d.i18n.commaAndOnX0()
			},
			nil,
		)
	}

	return description
}

// dayOfWeekName returns the weekday name at index i, guarding against out of
// range values the way JavaScript array indexing yields undefined.
func dayOfWeekName(names []string, i int) string {
	if i < 0 || i >= len(names) {
		return "undefined"
	}
	return names[i]
}

func (d *expressionDescriptor) getMonthDescription() string {
	monthNames := d.i18n.monthsOfTheYear()

	return d.getSegmentDescription(
		d.expressionParts[4],
		"",
		func(s string, form int) string {
			n, _ := jsParseInt(s)
			if form != 0 {
				if names, ok := d.i18n.monthsOfTheYearInCase(form); ok {
					return monthName(names, n-1)
				}
			}
			return monthName(monthNames, n-1)
		},
		func(s string) string {
			if n, ok := jsParseInt(s); ok && n == 1 {
				// rather than "every 1 months" just return empty string
				return ""
			}
			return format(d.i18n.commaEveryX0Months(s), s)
		},
		func(s string) string {
			if t, ok := d.i18n.commaMonthX0ThroughMonthX1(); ok {
				return t
			}
			// Note: the upstream descriptor calls commaX0ThroughX1() with NO
			// argument here (unlike the day-of-week range format).
			return d.i18n.commaX0ThroughX1("")
		},
		func(s string) string {
			if t, ok := d.i18n.commaOnlyInMonthX0(); ok {
				return t
			}
			return d.i18n.commaOnlyInX0()
		},
		nil,
	)
}

// monthName returns the month name at index i (0-based), guarding against
// out of range values.
func monthName(names []string, i int) string {
	if i < 0 || i >= len(names) {
		return "undefined"
	}
	return names[i]
}

var (
	weekDayNumberRe = regexp.MustCompile(`^(?:(\d{1,2})W|W(\d{1,2}))(?:([+-])(\d{1,2}))?$`)
	lastDayOffSetRe = regexp.MustCompile(`L-(\d{1,2})`)
	digitsOnlyRe    = regexp.MustCompile(`^\d+$`)
)

func (d *expressionDescriptor) getDayOfMonthDescription() string {
	expression := d.expressionParts[3]

	switch expression {
	case "L":
		return d.i18n.commaOnTheLastDayOfTheMonth()
	case "WL", "LW":
		return d.i18n.commaOnTheLastWeekdayOfTheMonth()
	}

	// i.e. 3W, W2, 15W-2 or 15W+2
	if m := weekDayNumberRe.FindStringSubmatch(expression); m != nil {
		dayNumberStr := m[1]
		if dayNumberStr == "" {
			dayNumberStr = m[2]
		}
		dayNumber, _ := strconv.Atoi(dayNumberStr)

		dayString := format(d.i18n.weekdayNearestDayX0(), strconv.Itoa(dayNumber))
		if dayNumber == 1 {
			dayString = d.i18n.firstWeekday()
		}

		offsetDays := m[4]
		switch {
		case offsetDays == "":
			return format(d.i18n.commaOnTheX0OfTheMonth(), dayString)
		case m[3] == "-":
			template := ", %s days before the %s of the month"
			if t, ok := d.i18n.commaDaysBeforeTheX0OfTheMonth(offsetDays); ok {
				template = t
			}
			return format(template, offsetDays, dayString)
		default:
			template := ", %s days after the %s of the month"
			if t, ok := d.i18n.commaDaysAfterTheX0OfTheMonth(offsetDays); ok {
				template = t
			}
			return format(template, offsetDays, dayString)
		}
	}

	// Handle "last day offset" (i.e. L-5: "5 days before the last day of the month")
	if m := lastDayOffSetRe.FindStringSubmatch(expression); m != nil {
		offSetDays := m[1]
		return format(d.i18n.commaDaysBeforeTheLastDayOfTheMonth(offSetDays), offSetDays)
	}

	if expression == "*" && d.expressionParts[5] != "*" {
		// * dayOfMonth and dayOfWeek specified so use dayOfWeek verbiage instead
		return ""
	}

	return d.getSegmentDescription(
		expression,
		d.i18n.commaEveryDay(),
		func(s string, form int) string {
			if s == "L" {
				return d.i18n.lastDay()
			}
			if t, ok := d.i18n.dayX0(); ok {
				return format(t, s)
			}
			return s
		},
		func(s string) string {
			if s == "1" {
				return d.i18n.commaEveryDay()
			}
			return d.i18n.commaEveryX0Days(s)
		},
		func(s string) string {
			return d.i18n.commaBetweenDayX0AndX1OfTheMonth(s)
		},
		func(s string) string {
			return d.i18n.commaOnDayX0OfTheMonth(s)
		},
		nil,
	)
}

func (d *expressionDescriptor) getYearDescription() string {
	return d.getSegmentDescription(
		d.expressionParts[6],
		"",
		func(s string, form int) string {
			if digitsOnlyRe.MatchString(s) {
				n, _ := strconv.Atoi(s)
				// The JavaScript original builds new Date(year, 1), and
				// JavaScript maps years 0-99 to 1900+year.
				if n < 100 {
					n += 1900
				}
				return strconv.Itoa(n)
			}
			return s
		},
		func(s string) string {
			return format(d.i18n.commaEveryX0Years(s), s)
		},
		func(s string) string {
			if t, ok := d.i18n.commaYearX0ThroughYearX1(); ok {
				return t
			}
			// Note: called with NO argument upstream, like the month variant.
			return d.i18n.commaX0ThroughX1("")
		},
		func(s string) string {
			if t, ok := d.i18n.commaOnlyInYearX0(); ok {
				return t
			}
			return d.i18n.commaOnlyInX0()
		},
		nil,
	)
}

type segmentRange struct {
	min int
	max int
}

// singleItemDescriptionFunc describes a single value; form is 0 for single
// items and 1 ("from") / 2 ("through") inside ranges.
type singleItemDescriptionFunc func(s string, form int) string

func (d *expressionDescriptor) getSegmentDescription(
	expression string,
	allDescription string,
	getSingleItemDescription singleItemDescriptionFunc,
	getIncrementDescriptionFormat func(s string) string,
	getRangeDescriptionFormat func(s string) string,
	getDescriptionFormat func(s string) string,
	rng *segmentRange,
) string {
	var description string
	doesExpressionContainIncrement := strings.Contains(expression, "/")
	doesExpressionContainRange := strings.Contains(expression, "-")
	doesExpressionContainMultipleValues := strings.Contains(expression, ",")

	switch {
	case expression == "":
		// Empty
		description = ""
	case expression == "*":
		// * (All)
		description = allDescription
	case !doesExpressionContainIncrement && !doesExpressionContainRange && !doesExpressionContainMultipleValues:
		// Simple
		description = format(getDescriptionFormat(expression), getSingleItemDescription(expression, 0))
	case doesExpressionContainMultipleValues:
		// Multiple Values
		segments := strings.Split(expression, ",")
		descriptionContent := ""
		conjunction := d.i18n.spaceAnd()
		for i := 0; i < len(segments); i++ {
			if i > 0 && len(segments) > 2 {
				descriptionContent += ","

				if i < len(segments)-1 {
					descriptionContent += " "
				}
			}

			if i > 0 && len(segments) > 1 && (i == len(segments)-1 || len(segments) == 2) {
				if conjunction == "," && len(segments) > 2 {
					descriptionContent += " "
				} else {
					descriptionContent += conjunction + " "
				}
			}

			if strings.Contains(segments[i], "/") || strings.Contains(segments[i], "-") {
				// Multiple Values with Increment or Range
				isSegmentRangeWithoutIncrement := strings.Contains(segments[i], "-") && !strings.Contains(segments[i], "/")

				rangeFormatFunc := getRangeDescriptionFormat
				if isSegmentRangeWithoutIncrement {
					rangeFormatFunc = func(s string) string { return d.i18n.commaX0ThroughX1(s) }
				}

				currentDescriptionContent := d.getSegmentDescription(
					segments[i],
					allDescription,
					getSingleItemDescription,
					getIncrementDescriptionFormat,
					rangeFormatFunc,
					getDescriptionFormat,
					nil,
				)

				if isSegmentRangeWithoutIncrement {
					// strip leading comma since the caller manages separators
					currentDescriptionContent = strings.Replace(currentDescriptionContent, ", ", "", 1)
				}

				descriptionContent += currentDescriptionContent
			} else if !doesExpressionContainIncrement {
				descriptionContent += getSingleItemDescription(segments[i], 0)
			} else {
				segmentDescription := d.getSegmentDescription(
					segments[i],
					allDescription,
					getSingleItemDescription,
					getIncrementDescriptionFormat,
					getRangeDescriptionFormat,
					getDescriptionFormat,
					nil,
				)
				// Strip leading comma since we manage separators in this loop
				segmentDescription = strings.TrimPrefix(segmentDescription, ", ")
				descriptionContent += segmentDescription
			}
		}

		if !doesExpressionContainIncrement {
			description = format(getDescriptionFormat(expression), descriptionContent)
		} else {
			description = descriptionContent
		}
	case doesExpressionContainIncrement:
		// Increment
		segments := strings.Split(expression, "/")
		description = format(getIncrementDescriptionFormat(segments[1]), segments[1])

		if strings.Contains(segments[0], "-") {
			// Range with Increment (ex: 2-59/3)
			rangeSegmentDescription := d.generateRangeSegmentDescription(
				segments[0],
				getRangeDescriptionFormat,
				getSingleItemDescription,
			)

			if !strings.HasPrefix(rangeSegmentDescription, ", ") {
				description += ", "
			}

			description += rangeSegmentDescription
		} else if !strings.Contains(segments[0], "*") {
			rangeItemDescription := format(
				getDescriptionFormat(segments[0]),
				getSingleItemDescription(segments[0], 0),
			)

			// remove any leading comma
			rangeItemDescription = strings.Replace(rangeItemDescription, ", ", "", 1)

			description += format(d.i18n.commaStartingX0(), rangeItemDescription)
		}

		description += d.getWrappingIncrementDescription(
			expression,
			rng,
			allDescription,
			getSingleItemDescription,
			getIncrementDescriptionFormat,
			getRangeDescriptionFormat,
			getDescriptionFormat,
		)
	case doesExpressionContainRange:
		// Range
		description = d.generateRangeSegmentDescription(
			expression,
			getRangeDescriptionFormat,
			getSingleItemDescription,
		)
	}

	return description
}

// getWrappingIncrementDescription: a step which doesn't divide the field
// evenly (e.g. minute "*/50") matches the values 0 and 50, so the gaps
// alternate between 50 and 10 minutes instead of being a constant interval.
// In verbose mode, spell out the matched values after the increment
// description by reusing the description of the equivalent list expression
// ("0,50").
func (d *expressionDescriptor) getWrappingIncrementDescription(
	expression string,
	rng *segmentRange,
	allDescription string,
	getSingleItemDescription singleItemDescriptionFunc,
	getIncrementDescriptionFormat func(s string) string,
	getRangeDescriptionFormat func(s string) string,
	getDescriptionFormat func(s string) string,
) string {
	if !d.options.Verbose || rng == nil {
		return ""
	}

	parts := strings.Split(expression, "/")
	start := parts[0]
	stepValue := parts[1]
	step, stepOK := jsParseInt(stepValue)
	length := rng.max - rng.min + 1

	// Only a step over the whole field wraps around: a range with a step
	// ("2-30/7") already states its bounds, and an evenly dividing step
	// ("*/15") really is a constant interval.
	coversWholeField := start == "*"
	isUsableStep := stepOK && step > 1
	dividesEvenly := isUsableStep && length%step == 0

	if !coversWholeField || !isUsableStep || dividesEvenly {
		return ""
	}

	var values []int
	for value := rng.min; value <= rng.max; value += step {
		values = append(values, value)
	}

	if len(values) < 2 {
		return ""
	}

	strs := make([]string, len(values))
	for i, v := range values {
		strs[i] = strconv.Itoa(v)
	}

	valuesDescription := d.getSegmentDescription(
		strings.Join(strs, ","),
		allDescription,
		getSingleItemDescription,
		getIncrementDescriptionFormat,
		getRangeDescriptionFormat,
		getDescriptionFormat,
		nil,
	)

	if valuesDescription != "" {
		return ", " + valuesDescription
	}
	return ""
}

func (d *expressionDescriptor) generateRangeSegmentDescription(
	rangeExpression string,
	getRangeDescriptionFormat func(s string) string,
	getSingleItemDescription singleItemDescriptionFunc,
) string {
	rangeSegments := strings.Split(rangeExpression, "-")
	rangeSegment1Description := getSingleItemDescription(rangeSegments[0], 1)
	rangeSegment2Description := getSingleItemDescription(rangeSegments[1], 2)
	rangeDescriptionFormat := getRangeDescriptionFormat(rangeExpression)
	return format(rangeDescriptionFormat, rangeSegment1Description, rangeSegment2Description)
}

func (d *expressionDescriptor) formatTime(hourExpression, minuteExpression, secondExpression string) string {
	hourOffset := 0
	minuteOffset := 0

	hour, _ := jsParseInt(hourExpression)
	minute, _ := jsParseInt(minuteExpression)
	hour += hourOffset
	minute += minuteOffset

	if minute >= 60 {
		minute -= 60
		hour++
	} else if minute < 0 {
		minute += 60
		hour--
	}

	if hour >= 24 {
		hour -= 24
	} else if hour < 0 {
		hour = 24 + hour
	}

	period := ""
	setPeriodBeforeTime := false
	if !d.options.Use24HourTimeFormat {
		setPeriodBeforeTime = d.i18n.setPeriodBeforeTime()
		if setPeriodBeforeTime {
			period = d.getPeriod(hour) + " "
		} else {
			period = " " + d.getPeriod(hour)
		}
		if hour > 12 {
			hour -= 12
		}
		if hour == 0 {
			hour = 12
		}
	}

	second := ""
	if secondExpression != "" {
		second = ":" + padTwo(secondExpression)
	}

	hourStr := strconv.Itoa(hour)
	paddedHour := padTwo(hourStr)
	minuteStr := strconv.Itoa(minute)
	paddedMinute := padTwo(minuteStr)
	displayHour := paddedHour
	if d.options.TrimHoursLeadingZero {
		displayHour = hourStr
	}

	if setPeriodBeforeTime {
		return period + displayHour + ":" + paddedMinute + second
	}
	return displayHour + ":" + paddedMinute + second + period
}

var trailingCommaSpaceRe = regexp.MustCompile(`, ?$`)

func (d *expressionDescriptor) transformVerbosity(description string, useVerboseFormat bool) string {
	if !useVerboseFormat {
		description = strings.ReplaceAll(description, ", "+d.i18n.everyMinute(), "")
		description = strings.ReplaceAll(description, ", "+d.i18n.everyHour(), "")
		description = strings.ReplaceAll(description, d.i18n.commaEveryDay(), "")
		description = trailingCommaSpaceRe.ReplaceAllString(description, "")

		if replacements, ok := d.i18n.conciseVerbosityReplacements(); ok {
			for key, value := range replacements {
				description = strings.ReplaceAll(description, key, value)
			}
		}
	}
	return description
}

func (d *expressionDescriptor) getPeriod(hour int) string {
	if hour >= 12 {
		if p, ok := d.i18n.pm(); ok {
			return p
		}
		return "PM"
	}
	if a, ok := d.i18n.am(); ok {
		return a
	}
	return "AM"
}
