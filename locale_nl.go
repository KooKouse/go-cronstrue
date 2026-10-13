package cronstrue

// Code generated from cronstrue/src/i18n/locales/nl.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("nl", NlLocale{enLocale{}})
}

// NlLocale is the nl locale.
type NlLocale struct {
	enLocale
}

func (NlLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (NlLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Er is een fout opgetreden bij het vertalen van de gegevens. Controleer de gegevens."
}

func (NlLocale) everyMinute() string {
	return "elke minuut"
}

func (NlLocale) everyHour() string {
	return "elk uur"
}

func (NlLocale) atSpace() string {
	return "Om "
}

func (NlLocale) everyMinuteBetweenX0AndX1() string {
	return "Elke minuut tussen %s en %s"
}

func (NlLocale) at() string {
	return "Om"
}

func (NlLocale) spaceAnd() string {
	return " en"
}

func (NlLocale) everySecond() string {
	return "elke seconde"
}

func (NlLocale) secondsX0ThroughX1PastTheMinute() string {
	return "seconden %s t/m %s na de minuut"
}

func (NlLocale) minutesX0ThroughX1PastTheHour() string {
	return "minuut %s t/m %s na het uur"
}

func (NlLocale) betweenX0AndX1() string {
	return "tussen %s en %s"
}

func (NlLocale) atX0() string {
	return "om %s"
}

func (NlLocale) commaEveryDay() string {
	return ", elke dag"
}

func (NlLocale) spaceX0OfTheMonth() string {
	return " %s van de maand"
}

func (NlLocale) lastDay() string {
	return "de laatste dag"
}

func (NlLocale) commaAndOnX0() string {
	return ", en op %s"
}

func (NlLocale) commaOnlyInX0() string {
	return ", alleen in %s"
}

func (NlLocale) commaOnTheLastDayOfTheMonth() string {
	return ", op de laatste dag van de maand"
}

func (NlLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", op de laatste werkdag van de maand"
}

func (NlLocale) firstWeekday() string {
	return "eerste werkdag"
}

func (NlLocale) weekdayNearestDayX0() string {
	return "werkdag dichtst bij dag %s"
}

func (NlLocale) commaOnTheX0OfTheMonth() string {
	return ", op de %s van de maand"
}

func (NlLocale) commaStartingX0() string {
	return ", beginnend %s"
}

func (NlLocale) atReboot() (string, bool) {
	return "Eén keer uitvoeren, bij het opstarten", true
}

func (NlLocale) onTheHour() (string, bool) {
	return "op het hele uur", true
}

func (NlLocale) everyX0Seconds(string) string {
	return "elke %s seconden"
}

func (NlLocale) everyX0Minutes(string) string {
	return "elke %s minuten"
}

func (NlLocale) everyX0Hours(string) string {
	return "elke %s uur"
}

func (NlLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", elke %s dagen van de week"
}

func (NlLocale) commaEveryX0Months(string) string {
	return ", elke %s maanden"
}

func (NlLocale) commaEveryX0Days(string) string {
	return ", elke %s dagen"
}

func (NlLocale) commaEveryX0Years(string) string {
	return ", elke %s jaren"
}

func (NlLocale) atX0SecondsPastTheMinute(string) string {
	return "op %s seconden na de minuut"
}

func (NlLocale) atX0MinutesPastTheHour(string) string {
	return "op %s minuten na het uur"
}

func (NlLocale) commaX0ThroughX1(string) string {
	return ", %s t/m %s"
}

func (NlLocale) commaAndX0ThroughX1(string) string {
	return ", en %s t/m %s"
}

func (NlLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", op de laatste %s van de maand"
}

func (NlLocale) commaOnDayX0OfTheMonth(string) string {
	return ", op dag %s van de maand"
}

func (NlLocale) first(string) string {
	return "eerste"
}

func (NlLocale) second(string) string {
	return "tweede"
}

func (NlLocale) third(string) string {
	return "derde"
}

func (NlLocale) fourth(string) string {
	return "vierde"
}

func (NlLocale) fifth(string) string {
	return "vijfde"
}

func (NlLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", tussen dag %s en %s van de maand"
}

func (NlLocale) commaOnlyOnX0(string) string {
	return ", alleen op %s"
}

func (NlLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dagen vóór de laatste dag van de maand"
}

func (NlLocale) commaOnThe(string, string) string {
	return ", op de "
}

func (NlLocale) daysOfTheWeek() []string {
	return []string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"}
}

func (NlLocale) monthsOfTheYear() []string {
	return []string{"januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"}
}
