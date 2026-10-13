package cronstrue

// Code generated from cronstrue/src/i18n/locales/sr.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("sr", SrLocale{enLocale{}})
}

// SrLocale is the sr locale.
type SrLocale struct {
	enLocale
}

func (SrLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (SrLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Došlo je do greške pri generisanju izraza. Proverite sintaksu cron izraza."
}

func (SrLocale) everyMinute() string {
	return "svakog minuta"
}

func (SrLocale) everyHour() string {
	return "svaki sat"
}

func (SrLocale) atSpace() string {
	return "U "
}

func (SrLocale) everyMinuteBetweenX0AndX1() string {
	return "Svakog minuta između %s i %s"
}

func (SrLocale) at() string {
	return "U"
}

func (SrLocale) spaceAnd() string {
	return " i"
}

func (SrLocale) everySecond() string {
	return "svake sekunde"
}

func (SrLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekunde od %s do %s posle pune minute"
}

func (SrLocale) minutesX0ThroughX1PastTheHour() string {
	return "minute od %s do %s posle punog sata"
}

func (SrLocale) betweenX0AndX1() string {
	return "između %s i %s"
}

func (SrLocale) atX0() string {
	return "u %s"
}

func (SrLocale) commaEveryDay() string {
	return ", svaki dan"
}

func (SrLocale) spaceX0OfTheMonth() string {
	return " %s u mesecu"
}

func (SrLocale) lastDay() string {
	return "poslednji dan"
}

func (SrLocale) commaAndOnX0() string {
	return ", i %s"
}

func (SrLocale) commaOnlyInX0() string {
	return ", samo u %s"
}

func (SrLocale) commaOnTheLastDayOfTheMonth() string {
	return ", poslednjeg dana u mesecu"
}

func (SrLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", poslednjeg radnog dana u mesecu"
}

func (SrLocale) firstWeekday() string {
	return "prvi radni dan"
}

func (SrLocale) weekdayNearestDayX0() string {
	return "radni dan najbliži %s. danu"
}

func (SrLocale) commaOnTheX0OfTheMonth() string {
	return ", %s u mesecu"
}

func (SrLocale) commaStartingX0() string {
	return ", počevši od %s"
}

func (SrLocale) atReboot() (string, bool) {
	return "Pokreni jednom, pri pokretanju", true
}

func (SrLocale) onTheHour() (string, bool) {
	return "u pun sat", true
}

func (SrLocale) everyX0Seconds(string) string {
	return "svakih %s sekundi"
}

func (SrLocale) everyX0Minutes(string) string {
	return "svakih %s minuta"
}

func (SrLocale) everyX0Hours(string) string {
	return "svakih %s sati"
}

func (SrLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", svakih %s dana u nedelji"
}

func (SrLocale) commaEveryX0Months(string) string {
	return ", svakih %s meseci"
}

func (SrLocale) commaEveryX0Days(string) string {
	return ", svakih %s dana"
}

func (SrLocale) commaEveryX0Years(string) string {
	return ", svakih %s godina"
}

func (SrLocale) atX0SecondsPastTheMinute(string) string {
	return "u %s sekundi posle pune minute"
}

func (SrLocale) atX0MinutesPastTheHour(string) string {
	return "u %s minuta posle punog sata"
}

func (SrLocale) commaX0ThroughX1(string) string {
	return ", od %s do %s"
}

func (SrLocale) commaAndX0ThroughX1(string) string {
	return ", i od %s do %s"
}

func (SrLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", poslednji %s u mesecu"
}

func (SrLocale) commaOnDayX0OfTheMonth(string) string {
	return ", %s. dan u mesecu"
}

func (SrLocale) first(string) string {
	return "prvi"
}

func (SrLocale) second(string) string {
	return "drugi"
}

func (SrLocale) third(string) string {
	return "treći"
}

func (SrLocale) fourth(string) string {
	return "četvrti"
}

func (SrLocale) fifth(string) string {
	return "peti"
}

func (SrLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", između %s. i %s. dana u mesecu"
}

func (SrLocale) commaOnlyOnX0(string) string {
	return ", samo %s"
}

func (SrLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dana pre kraja meseca"
}

func (SrLocale) commaOnThe(string, string) string {
	return ", "
}

func (SrLocale) daysOfTheWeek() []string {
	return []string{"Nedelja", "Ponedeljak", "Utorak", "Sreda", "Četvrtak", "Petak", "Subota"}
}

func (SrLocale) monthsOfTheYear() []string {
	return []string{"januar", "februar", "mart", "april", "maj", "jun", "jul", "avgust", "septembar", "oktobar", "novembar", "decembar"}
}
