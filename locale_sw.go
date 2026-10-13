package cronstrue

// Code generated from cronstrue/src/i18n/locales/sw.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("sw", SwLocale{enLocale{}})
}

// SwLocale is the sw locale.
type SwLocale struct {
	enLocale
}

func (SwLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (SwLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Kuna tatizo wakati wa kutunga msemo. Angalia cron expression syntax."
}

func (SwLocale) everyMinute() string {
	return "kila dakika"
}

func (SwLocale) everyHour() string {
	return "kila saa"
}

func (SwLocale) atSpace() string {
	return "Kwa "
}

func (SwLocale) everyMinuteBetweenX0AndX1() string {
	return "Kila dakika kwanzia %s hadi %s"
}

func (SwLocale) at() string {
	return "Kwa"
}

func (SwLocale) spaceAnd() string {
	return " na"
}

func (SwLocale) everySecond() string {
	return "kila sekunde"
}

func (SwLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekunde ya %s hadi %s baada ya dakika"
}

func (SwLocale) minutesX0ThroughX1PastTheHour() string {
	return "minutes %s through %s past the hour"
}

func (SwLocale) betweenX0AndX1() string {
	return "kati ya %s na %s"
}

func (SwLocale) atX0() string {
	return "kwenye %s"
}

func (SwLocale) commaEveryDay() string {
	return ", kila siku"
}

func (SwLocale) spaceX0OfTheMonth() string {
	return " siku %s ya mwezi"
}

func (SwLocale) lastDay() string {
	return "siku ya mwisho"
}

func (SwLocale) commaAndOnX0() string {
	return ", na pia %s"
}

func (SwLocale) commaOnlyInX0() string {
	return ", kwa %s tu"
}

func (SwLocale) commaOnTheLastDayOfTheMonth() string {
	return ", siku ya mwisho wa mwezi"
}

func (SwLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", wikendi ya mwisho wa mwezi"
}

func (SwLocale) firstWeekday() string {
	return "siku za kazi ya kwanza"
}

func (SwLocale) weekdayNearestDayX0() string {
	return "siku ya kazi karibu na siku ya %s"
}

func (SwLocale) commaOnTheX0OfTheMonth() string {
	return ", siku ya %s ya mwezi"
}

func (SwLocale) commaStartingX0() string {
	return ", kwanzia %s"
}

func (SwLocale) atReboot() (string, bool) {
	return "Endesha mara moja, wakati wa kuanza", true
}

func (SwLocale) onTheHour() (string, bool) {
	return "saa kamili", true
}

func (SwLocale) everyX0Seconds(string) string {
	return "kila sekunde %s"
}

func (SwLocale) everyX0Minutes(string) string {
	return "kila dakika %s"
}

func (SwLocale) everyX0Hours(string) string {
	return "every %s hours"
}

func (SwLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", kila siku %s ya wiki"
}

func (SwLocale) commaEveryX0Months(string) string {
	return ", kila mwezi wa %s"
}

func (SwLocale) commaEveryX0Days(string) string {
	return ", kila siku %s"
}

func (SwLocale) commaEveryX0Years(string) string {
	return ", kila miaka %s"
}

func (SwLocale) atX0SecondsPastTheMinute(string) string {
	return "sekunde %s baada ya dakika"
}

func (SwLocale) atX0MinutesPastTheHour(string) string {
	return "at %s minutes past the hour"
}

func (SwLocale) commaX0ThroughX1(string) string {
	return ", %s hadi %s"
}

func (SwLocale) commaAndX0ThroughX1(string) string {
	return ", na %s hadi %s"
}

func (SwLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", siku ya %s ya mwezi"
}

func (SwLocale) commaOnDayX0OfTheMonth(string) string {
	return ", siku ya %s ya mwezi"
}

func (SwLocale) first(string) string {
	return "ya kwanza"
}

func (SwLocale) second(string) string {
	return "ya pili"
}

func (SwLocale) third(string) string {
	return "ya tatu"
}

func (SwLocale) fourth(string) string {
	return "ya nne"
}

func (SwLocale) fifth(string) string {
	return "ya tano"
}

func (SwLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", kati ya siku %s na %s ya mwezi"
}

func (SwLocale) commaOnlyOnX0(string) string {
	return ", kwa %s tu"
}

func (SwLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", siku ya %s kabla ya siku ya mwisho wa mwezi"
}

func (SwLocale) commaOnThe(string, string) string {
	return ", kwenye "
}

func (SwLocale) daysOfTheWeek() []string {
	return []string{"Jumapili", "Jumatatu", "Jumanne", "Jumatano", "Alhamisi", "Ijumaa", "Jumamosi"}
}

func (SwLocale) monthsOfTheYear() []string {
	return []string{"Januari", "Februari", "Machi", "Aprili", "Mei", "Juni", "Julai", "Agosti", "Septemba", "Oktoba", "Novemba", "Desemba"}
}
