package cronstrue

// Code generated from cronstrue/src/i18n/locales/hr.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("hr", HrLocale{enLocale{}})
}

// HrLocale is the hr locale.
type HrLocale struct {
	enLocale
}

func (HrLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (HrLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Došlo je do pogreške pri generiranju izraza. Provjerite sintaksu cron izraza."
}

func (HrLocale) everyMinute() string {
	return "svaku minutu"
}

func (HrLocale) everyHour() string {
	return "svaki sat"
}

func (HrLocale) atSpace() string {
	return "U "
}

func (HrLocale) everyMinuteBetweenX0AndX1() string {
	return "Svaku minutu između %s i %s"
}

func (HrLocale) at() string {
	return "U"
}

func (HrLocale) spaceAnd() string {
	return " i"
}

func (HrLocale) everySecond() string {
	return "svaku sekundu"
}

func (HrLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekunde od %s do %s nakon pune minute"
}

func (HrLocale) minutesX0ThroughX1PastTheHour() string {
	return "minute od %s do %s nakon punog sata"
}

func (HrLocale) betweenX0AndX1() string {
	return "između %s i %s"
}

func (HrLocale) atX0() string {
	return "u %s"
}

func (HrLocale) commaEveryDay() string {
	return ", svaki dan"
}

func (HrLocale) spaceX0OfTheMonth() string {
	return " %s u mjesecu"
}

func (HrLocale) lastDay() string {
	return "zadnji dan"
}

func (HrLocale) commaAndOnX0() string {
	return ", i %s"
}

func (HrLocale) commaOnlyInX0() string {
	return ", samo u %s"
}

func (HrLocale) commaOnTheLastDayOfTheMonth() string {
	return ", zadnji dan u mjesecu"
}

func (HrLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", zadnji radni dan u mjesecu"
}

func (HrLocale) firstWeekday() string {
	return "prvi radni dan"
}

func (HrLocale) weekdayNearestDayX0() string {
	return "radni dan najbliži %s. danu"
}

func (HrLocale) commaOnTheX0OfTheMonth() string {
	return ", %s u mjesecu"
}

func (HrLocale) commaStartingX0() string {
	return ", počevši od %s"
}

func (HrLocale) atReboot() (string, bool) {
	return "Pokreni jednom, pri pokretanju", true
}

func (HrLocale) onTheHour() (string, bool) {
	return "u puni sat", true
}

func (HrLocale) everyX0Seconds(string) string {
	return "svakih %s sekundi"
}

func (HrLocale) everyX0Minutes(string) string {
	return "svakih %s minuta"
}

func (HrLocale) everyX0Hours(string) string {
	return "svakih %s sati"
}

func (HrLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", svakih %s dana u tjednu"
}

func (HrLocale) commaEveryX0Months(string) string {
	return ", svakih %s mjeseci"
}

func (HrLocale) commaEveryX0Days(string) string {
	return ", svakih %s dana"
}

func (HrLocale) commaEveryX0Years(string) string {
	return ", svakih %s godina"
}

func (HrLocale) atX0SecondsPastTheMinute(string) string {
	return "u %s sekundi nakon pune minute"
}

func (HrLocale) atX0MinutesPastTheHour(string) string {
	return "u %s minuta nakon punog sata"
}

func (HrLocale) commaX0ThroughX1(string) string {
	return ", od %s do %s"
}

func (HrLocale) commaAndX0ThroughX1(string) string {
	return ", i od %s do %s"
}

func (HrLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", zadnji %s u mjesecu"
}

func (HrLocale) commaOnDayX0OfTheMonth(string) string {
	return ", %s. dan u mjesecu"
}

func (HrLocale) first(string) string {
	return "prvi"
}

func (HrLocale) second(string) string {
	return "drugi"
}

func (HrLocale) third(string) string {
	return "treći"
}

func (HrLocale) fourth(string) string {
	return "četvrti"
}

func (HrLocale) fifth(string) string {
	return "peti"
}

func (HrLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", između %s. i %s. dana u mjesecu"
}

func (HrLocale) commaOnlyOnX0(string) string {
	return ", samo %s"
}

func (HrLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dana prije kraja mjeseca"
}

func (HrLocale) commaOnThe(string, string) string {
	return ", "
}

func (HrLocale) daysOfTheWeek() []string {
	return []string{"Nedjelja", "Ponedjeljak", "Utorak", "Srijeda", "Četvrtak", "Petak", "Subota"}
}

func (HrLocale) monthsOfTheYear() []string {
	return []string{"siječanj", "veljača", "ožujak", "travanj", "svibanj", "lipanj", "srpanj", "kolovoz", "rujan", "listopad", "studeni", "prosinac"}
}
