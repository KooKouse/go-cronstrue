package cronstrue

// Code generated from cronstrue/src/i18n/locales/pl.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("pl", PlLocale{enLocale{}})
}

// PlLocale is the pl locale.
type PlLocale struct {
	enLocale
}

func (PlLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (PlLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Wystąpił błąd podczas generowania opisu wyrażenia cron. Sprawdź składnię wyrażenia cron."
}

func (PlLocale) everyMinute() string {
	return "co minutę"
}

func (PlLocale) everyHour() string {
	return "co godzinę"
}

func (PlLocale) atSpace() string {
	return "O "
}

func (PlLocale) everyMinuteBetweenX0AndX1() string {
	return "Co minutę od %s do %s"
}

func (PlLocale) at() string {
	return "O"
}

func (PlLocale) spaceAnd() string {
	return " i"
}

func (PlLocale) everySecond() string {
	return "co sekundę"
}

func (PlLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekundy od %s do %s"
}

func (PlLocale) minutesX0ThroughX1PastTheHour() string {
	return "minuty od %s do %s"
}

func (PlLocale) betweenX0AndX1() string {
	return "od %s do %s"
}

func (PlLocale) atX0() string {
	return "o %s"
}

func (PlLocale) commaEveryDay() string {
	return ", co dzień"
}

func (PlLocale) spaceX0OfTheMonth() string {
	return " %s miesiąca"
}

func (PlLocale) lastDay() string {
	return "ostatni dzień"
}

func (PlLocale) commaAndOnX0() string {
	return ", i %s"
}

func (PlLocale) commaOnlyInX0() string {
	return ", tylko %s"
}

func (PlLocale) commaOnTheLastDayOfTheMonth() string {
	return ", ostatni dzień miesiąca"
}

func (PlLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", ostatni dzień roboczy miesiąca"
}

func (PlLocale) firstWeekday() string {
	return "pierwszy dzień roboczy"
}

func (PlLocale) weekdayNearestDayX0() string {
	return "dzień roboczy najbliższy %s-ego dnia"
}

func (PlLocale) commaOnTheX0OfTheMonth() string {
	return ", %s miesiąca"
}

func (PlLocale) commaStartingX0() string {
	return ", startowy %s"
}

func (PlLocale) atReboot() (string, bool) {
	return "Uruchom raz, przy starcie", true
}

func (PlLocale) onTheHour() (string, bool) {
	return "o pełnej godzinie", true
}

func (PlLocale) everyX0Seconds(string) string {
	return "co %s sekund"
}

func (PlLocale) everyX0Minutes(string) string {
	return "co %s minut"
}

func (PlLocale) everyX0Hours(string) string {
	return "co %s godzin"
}

func (PlLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", co %s dni tygodnia"
}

func (PlLocale) commaEveryX0Months(string) string {
	return ", co %s miesięcy"
}

func (PlLocale) commaEveryX0Days(string) string {
	return ", co %s dni"
}

func (PlLocale) commaEveryX0Years(string) string {
	return ", co %s lat"
}

func (PlLocale) atX0SecondsPastTheMinute(string) string {
	return "w %s sekundzie"
}

func (PlLocale) atX0MinutesPastTheHour(string) string {
	return "w %s minucie"
}

func (PlLocale) commaX0ThroughX1(string) string {
	return ", od %s do %s"
}

func (PlLocale) commaAndX0ThroughX1(string) string {
	return ", i od %s do %s"
}

func (PlLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", ostatni %s miesiąca"
}

func (PlLocale) commaOnDayX0OfTheMonth(string) string {
	return ", %s-ego dnia miesiąca"
}

func (PlLocale) first(string) string {
	return "pierwszy"
}

func (PlLocale) second(string) string {
	return "drugi"
}

func (PlLocale) third(string) string {
	return "trzeci"
}

func (PlLocale) fourth(string) string {
	return "czwarty"
}

func (PlLocale) fifth(string) string {
	return "piąty"
}

func (PlLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", od %s-ego do %s-ego dnia miesiąca"
}

func (PlLocale) commaOnlyOnX0(string) string {
	return ", tylko %s"
}

func (PlLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dni przed ostatnim dniem miesiąca"
}

func (PlLocale) commaOnThe(string, string) string {
	return ", "
}

func (PlLocale) daysOfTheWeek() []string {
	return []string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"}
}

func (PlLocale) monthsOfTheYear() []string {
	return []string{"styczeń", "luty", "marzec", "kwiecień", "maj", "czerwiec", "lipiec", "sierpień", "wrzesień", "październik", "listopad", "grudzień"}
}
