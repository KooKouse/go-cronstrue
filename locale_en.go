package cronstrue

import "sort"

func sortStrings(s []string) { sort.Strings(s) }

func init() {
	registerLocale("en", enLocale{})
}

// enLocale is the English locale, ported verbatim from
// cronstrue/src/i18n/locales/en.ts. It also acts as the fallback for every
// other locale: optional methods that a locale does not override return
// ok == false and the descriptor then falls back to its documented default.
type enLocale struct{}

func (enLocale) conciseVerbosityReplacements() (map[string]string, bool) { return nil, false }
func (enLocale) setPeriodBeforeTime() bool                               { return false }
func (enLocale) pm() (string, bool)                                      { return "", false }
func (enLocale) am() (string, bool)                                      { return "", false }
func (enLocale) atX0SecondsPastTheMinuteGt20() (string, bool)            { return "", false }
func (enLocale) atX0MinutesPastTheHourGt20() (string, bool)              { return "", false }
func (enLocale) commaMonthX0ThroughMonthX1() (string, bool)              { return "", false }
func (enLocale) commaYearX0ThroughYearX1() (string, bool)                { return "", false }
func (enLocale) commaOnlyInMonthX0() (string, bool)                      { return "", false }
func (enLocale) commaOnlyInYearX0() (string, bool)                       { return "", false }
func (enLocale) commaDaysBeforeTheX0OfTheMonth(string) (string, bool)    { return "", false }
func (enLocale) commaDaysAfterTheX0OfTheMonth(string) (string, bool)     { return "", false }
func (enLocale) dayX0() (string, bool)                                   { return "", false }
func (enLocale) daysOfTheWeekInCase(int) ([]string, bool)                { return nil, false }
func (enLocale) monthsOfTheYearInCase(int) ([]string, bool)              { return nil, false }
func (enLocale) atReboot() (string, bool)                                { return "", false }
func (enLocale) onTheHour() (string, bool)                               { return "on the hour", true }

func (enLocale) use24HourTimeFormatByDefault() bool { return false }
func (enLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "An error occurred when generating the expression description. Check the cron expression syntax."
}
func (enLocale) everyMinute() string               { return "every minute" }
func (enLocale) everyHour() string                 { return "every hour" }
func (enLocale) atSpace() string                   { return "At " }
func (enLocale) everyMinuteBetweenX0AndX1() string { return "Every minute between %s and %s" }
func (enLocale) at() string                        { return "At" }
func (enLocale) spaceAnd() string                  { return " and" }
func (enLocale) everySecond() string               { return "every second" }
func (enLocale) everyX0Seconds(string) string      { return "every %s seconds" }
func (enLocale) secondsX0ThroughX1PastTheMinute() string {
	return "seconds %s through %s past the minute"
}
func (enLocale) atX0SecondsPastTheMinute(string) string  { return "at %s seconds past the minute" }
func (enLocale) everyX0Minutes(string) string            { return "every %s minutes" }
func (enLocale) minutesX0ThroughX1PastTheHour() string   { return "minutes %s through %s past the hour" }
func (enLocale) atX0MinutesPastTheHour(string) string    { return "at %s minutes past the hour" }
func (enLocale) everyX0Hours(string) string              { return "every %s hours" }
func (enLocale) betweenX0AndX1() string                  { return "between %s and %s" }
func (enLocale) atX0() string                            { return "at %s" }
func (enLocale) commaEveryDay() string                   { return ", every day" }
func (enLocale) commaEveryX0DaysOfTheWeek(string) string { return ", every %s days of the week" }
func (enLocale) commaX0ThroughX1(string) string          { return ", %s through %s" }
func (enLocale) commaAndX0ThroughX1(string) string       { return ", %s through %s" }
func (enLocale) first(string) string                     { return "first" }
func (enLocale) second(string) string                    { return "second" }
func (enLocale) third(string) string                     { return "third" }
func (enLocale) fourth(string) string                    { return "fourth" }
func (enLocale) fifth(string) string                     { return "fifth" }
func (enLocale) commaOnThe(string, string) string        { return ", on the " }
func (enLocale) spaceX0OfTheMonth() string               { return " %s of the month" }
func (enLocale) lastDay() string                         { return "the last day" }
func (enLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", on the last %s of the month"
}
func (enLocale) commaOnlyOnX0(string) string { return ", only on %s" }
func (enLocale) commaAndOnX0() string        { return ", and on %s" }
func (enLocale) commaEveryX0Months(string) string {
	return ", every %s months"
}
func (enLocale) commaOnlyInX0() string { return ", only in %s" }
func (enLocale) commaOnTheLastDayOfTheMonth() string {
	return ", on the last day of the month"
}
func (enLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", on the last weekday of the month"
}
func (enLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s days before the last day of the month"
}
func (enLocale) firstWeekday() string           { return "first weekday" }
func (enLocale) weekdayNearestDayX0() string    { return "weekday nearest day %s" }
func (enLocale) commaOnTheX0OfTheMonth() string { return ", on the %s of the month" }
func (enLocale) commaEveryX0Days(string) string { return ", every %s days in a month" }
func (enLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", between day %s and %s of the month"
}
func (enLocale) commaOnDayX0OfTheMonth(string) string { return ", on day %s of the month" }
func (enLocale) commaEveryX0Years(string) string      { return ", every %s years" }
func (enLocale) commaStartingX0() string              { return ", starting %s" }

func (enLocale) daysOfTheWeek() []string {
	return []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
}

func (enLocale) monthsOfTheYear() []string {
	return []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}
}
