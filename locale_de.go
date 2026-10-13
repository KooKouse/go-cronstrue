package cronstrue

// Code generated from cronstrue/src/i18n/locales/de.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("de", DeLocale{enLocale{}})
}

// DeLocale is the de locale.
type DeLocale struct {
	enLocale
}

func (DeLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (DeLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Beim Generieren der Ausdrucksbeschreibung ist ein Fehler aufgetreten. Überprüfen Sie die Syntax des Cron-Ausdrucks."
}

func (DeLocale) everyMinute() string {
	return "jede Minute"
}

func (DeLocale) everyHour() string {
	return "jede Stunde"
}

func (DeLocale) atSpace() string {
	return "Um "
}

func (DeLocale) everyMinuteBetweenX0AndX1() string {
	return "Jede Minute zwischen %s und %s"
}

func (DeLocale) at() string {
	return "Um"
}

func (DeLocale) spaceAnd() string {
	return " und"
}

func (DeLocale) everySecond() string {
	return "Jede Sekunde"
}

func (DeLocale) secondsX0ThroughX1PastTheMinute() string {
	return "Sekunden %s bis %s"
}

func (DeLocale) minutesX0ThroughX1PastTheHour() string {
	return "Minuten %s bis %s"
}

func (DeLocale) betweenX0AndX1() string {
	return "zwischen %s und %s"
}

func (DeLocale) atX0() string {
	return "um %s"
}

func (DeLocale) commaEveryDay() string {
	return ", jeden Tag"
}

func (DeLocale) spaceX0OfTheMonth() string {
	return " %s des Monats"
}

func (DeLocale) lastDay() string {
	return "der letzte Tag"
}

func (DeLocale) commaAndOnX0() string {
	return ", und jeden %s"
}

func (DeLocale) commaOnlyInX0() string {
	return ", nur im %s"
}

func (DeLocale) commaOnTheLastDayOfTheMonth() string {
	return ", am letzten Tag des Monats"
}

func (DeLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", am letzten Werktag des Monats"
}

func (DeLocale) firstWeekday() string {
	return "ersten Werktag"
}

func (DeLocale) weekdayNearestDayX0() string {
	return "Werktag am nächsten zum %s Tag"
}

func (DeLocale) commaOnTheX0OfTheMonth() string {
	return ", am %s des Monats"
}

func (DeLocale) commaStartingX0() string {
	return ", beginnend %s"
}

func (DeLocale) atReboot() (string, bool) {
	return "Einmal ausführen, beim Start", true
}

func (DeLocale) onTheHour() (string, bool) {
	return "zur vollen Stunde", true
}

func (DeLocale) everyX0Seconds(string) string {
	return "alle %s Sekunden"
}

func (DeLocale) everyX0Minutes(string) string {
	return "alle %s Minuten"
}

func (DeLocale) everyX0Hours(string) string {
	return "alle %s Stunden"
}

func (DeLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", alle %s Tage der Woche"
}

func (DeLocale) commaEveryX0Months(string) string {
	return ", alle %s Monate"
}

func (DeLocale) commaEveryX0Days(string) string {
	return ", alle %s Tage"
}

func (DeLocale) commaEveryX0Years(string) string {
	return ", alle %s Jahre"
}

func (DeLocale) atX0SecondsPastTheMinute(string) string {
	return "bei Sekunde %s"
}

func (DeLocale) atX0MinutesPastTheHour(string) string {
	return "bei Minute %s"
}

func (DeLocale) commaX0ThroughX1(string) string {
	return ", %s bis %s"
}

func (DeLocale) commaAndX0ThroughX1(string) string {
	return ", und %s bis %s"
}

func (DeLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", am letzten %s des Monats"
}

func (DeLocale) commaOnDayX0OfTheMonth(string) string {
	return ", an Tag %s des Monats"
}

func (DeLocale) first(string) string {
	return "ersten"
}

func (DeLocale) second(string) string {
	return "zweiten"
}

func (DeLocale) third(string) string {
	return "dritten"
}

func (DeLocale) fourth(string) string {
	return "vierten"
}

func (DeLocale) fifth(string) string {
	return "fünften"
}

func (DeLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", zwischen Tag %s und %s des Monats"
}

func (DeLocale) commaOnlyOnX0(string) string {
	return ", nur jeden %s"
}

func (DeLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s tage vor dem letzten Tag des Monats"
}

func (DeLocale) commaOnThe(string, string) string {
	return ", am "
}

func (DeLocale) daysOfTheWeek() []string {
	return []string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"}
}

func (DeLocale) monthsOfTheYear() []string {
	return []string{"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"}
}
