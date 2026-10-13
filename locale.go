package cronstrue

// locale is implemented by every supported language. It is the Go
// counterpart of the Locale interface in the original JavaScript library
// (src/i18n/locale.ts).
//
// Methods whose JavaScript counterparts are optional (they may return null
// or be undefined) return an additional boolean here: ok == false means "not
// provided" and triggers the documented fallback behavior in the descriptor.
//
// The form parameter passed by the descriptor follows the original library:
// 0 means "not applicable" (single item), 1 means "from" and 2 means
// "through" for range expressions. It is only consumed by locales that need
// grammatical cases for ranges (e.g. Russian).
type locale interface {
	// conciseVerbosityReplacements optionally returns replacements applied
	// to make the non-verbose description more concise.
	conciseVerbosityReplacements() (map[string]string, bool)
	// setPeriodBeforeTime optionally places AM/PM before the time (e.g. Chinese).
	setPeriodBeforeTime() bool
	// pm optionally localizes "PM".
	pm() (string, bool)
	// am optionally localizes "AM".
	am() (string, bool)

	use24HourTimeFormatByDefault() bool
	anErrorOccuredWhenGeneratingTheExpressionD() string
	everyMinute() string
	everyHour() string
	atSpace() string
	everyMinuteBetweenX0AndX1() string
	at() string
	spaceAnd() string
	everySecond() string
	everyX0Seconds(s string) string
	secondsX0ThroughX1PastTheMinute() string
	atX0SecondsPastTheMinute(s string) string
	atX0SecondsPastTheMinuteGt20() (string, bool)
	everyX0Minutes(s string) string
	minutesX0ThroughX1PastTheHour() string
	atX0MinutesPastTheHour(s string) string
	atX0MinutesPastTheHourGt20() (string, bool)
	everyX0Hours(s string) string
	betweenX0AndX1() string
	atX0() string
	commaEveryDay() string
	commaEveryX0DaysOfTheWeek(s string) string
	commaX0ThroughX1(s string) string
	commaAndX0ThroughX1(s string) string
	commaMonthX0ThroughMonthX1() (string, bool)
	commaYearX0ThroughYearX1() (string, bool)
	first(s string) string
	second(s string) string
	third(s string) string
	fourth(s string) string
	fifth(s string) string
	commaOnThe(s string, day string) string
	spaceX0OfTheMonth() string
	lastDay() string
	commaOnTheLastX0OfTheMonth(s string) string
	commaOnlyOnX0(s string) string
	commaAndOnX0() string
	commaEveryX0Months(s string) string
	commaOnlyInX0() string
	commaOnlyInMonthX0() (string, bool)
	commaOnlyInYearX0() (string, bool)
	commaOnTheLastDayOfTheMonth() string
	commaOnTheLastWeekdayOfTheMonth() string
	commaDaysBeforeTheLastDayOfTheMonth(s string) string
	firstWeekday() string
	weekdayNearestDayX0() string
	commaOnTheX0OfTheMonth() string
	commaDaysBeforeTheX0OfTheMonth(s string) (string, bool)
	commaDaysAfterTheX0OfTheMonth(s string) (string, bool)
	commaEveryX0Days(s string) string
	commaBetweenDayX0AndX1OfTheMonth(s string) string
	commaOnDayX0OfTheMonth(s string) string
	commaEveryX0Years(s string) string
	commaStartingX0() string
	dayX0() (string, bool)
	daysOfTheWeek() []string
	daysOfTheWeekInCase(form int) ([]string, bool)
	monthsOfTheYear() []string
	monthsOfTheYearInCase(form int) ([]string, bool)
	atReboot() (string, bool)
	onTheHour() (string, bool)
}

// locales holds all registered locales by name.
var locales = map[string]locale{}

// defaultLocaleName is used when Options.LocaleName is empty or unknown.
const defaultLocaleName = "en"

// registerLocale registers a locale under the given name (e.g. "en",
// "zh_CN"). It is called from init functions of the locale files.
func registerLocale(name string, l locale) {
	locales[name] = l
}

// LocaleNames returns the sorted names of all registered locales.
func LocaleNames() []string {
	names := make([]string, 0, len(locales))
	for name := range locales {
		names = append(names, name)
	}
	sortStrings(names)
	return names
}
