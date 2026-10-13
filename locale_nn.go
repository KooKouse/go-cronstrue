package cronstrue

// Code generated from cronstrue/src/i18n/locales/nn.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("nn", NnLocale{enLocale{}})
}

// NnLocale is the nn locale.
type NnLocale struct {
	enLocale
}

func (NnLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (NnLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Ein feil oppstod ved generering av uttrykksskildring. Sjekk cron-syntaksen."
}

func (NnLocale) everyMinute() string {
	return "kvart minutt"
}

func (NnLocale) everyHour() string {
	return "kvar time"
}

func (NnLocale) atSpace() string {
	return "Kl."
}

func (NnLocale) everyMinuteBetweenX0AndX1() string {
	return "Kvart minutt mellom %s og %s"
}

func (NnLocale) at() string {
	return "Kl."
}

func (NnLocale) spaceAnd() string {
	return " og"
}

func (NnLocale) everySecond() string {
	return "kvart sekund"
}

func (NnLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekunda frå %s til og med %s etter minuttet"
}

func (NnLocale) minutesX0ThroughX1PastTheHour() string {
	return "minutta frå %s til og med %s etter timen"
}

func (NnLocale) betweenX0AndX1() string {
	return "mellom %s og %s"
}

func (NnLocale) atX0() string {
	return "på %s"
}

func (NnLocale) commaEveryDay() string {
	return ", kvar dag"
}

func (NnLocale) spaceX0OfTheMonth() string {
	return " %s i månaden"
}

func (NnLocale) lastDay() string {
	return "den siste dagen"
}

func (NnLocale) commaAndOnX0() string {
	return ", og på %s"
}

func (NnLocale) commaOnlyInX0() string {
	return ", berre i %s"
}

func (NnLocale) commaOnTheLastDayOfTheMonth() string {
	return ", på den siste dagen i månaden"
}

func (NnLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", den siste vekedagen i månaden"
}

func (NnLocale) firstWeekday() string {
	return "første vekedag"
}

func (NnLocale) weekdayNearestDayX0() string {
	return "vekedag nærmast dag %s"
}

func (NnLocale) commaOnTheX0OfTheMonth() string {
	return ", på den %s av månaden"
}

func (NnLocale) commaStartingX0() string {
	return ", startar %s"
}

func (NnLocale) atReboot() (string, bool) {
	return "Køyr éin gong, ved oppstart", true
}

func (NnLocale) onTheHour() (string, bool) {
	return "på timen", true
}

func (NnLocale) everyX0Seconds(string) string {
	return "kvart %s sekund"
}

func (NnLocale) everyX0Minutes(string) string {
	return "kvart %s minutt"
}

func (NnLocale) everyX0Hours(string) string {
	return "kvar %s time"
}

func (NnLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", kvar %s vekedag"
}

func (NnLocale) commaEveryX0Months(string) string {
	return ", kvar %s månad"
}

func (NnLocale) commaEveryX0Days(string) string {
	return ", kvar %s dag i månaden"
}

func (NnLocale) commaEveryX0Years(string) string {
	return ", kvart %s år"
}

func (NnLocale) atX0SecondsPastTheMinute(string) string {
	return "på %s sekund etter minuttet"
}

func (NnLocale) atX0MinutesPastTheHour(string) string {
	return "på %s minutt etter timen"
}

func (NnLocale) commaX0ThroughX1(string) string {
	return ", %s til og med %s"
}

func (NnLocale) commaAndX0ThroughX1(string) string {
	return ", og %s til og med %s"
}

func (NnLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", på den siste %s av månaden"
}

func (NnLocale) commaOnDayX0OfTheMonth(string) string {
	return ", på dag %s av månaden"
}

func (NnLocale) first(string) string {
	return "første"
}

func (NnLocale) second(string) string {
	return "andre"
}

func (NnLocale) third(string) string {
	return "tredje"
}

func (NnLocale) fourth(string) string {
	return "fjerde"
}

func (NnLocale) fifth(string) string {
	return "femte"
}

func (NnLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", mellom dag %s og %s av månaden"
}

func (NnLocale) commaOnlyOnX0(string) string {
	return ", på %s"
}

func (NnLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dagar før den siste dagen i månaden"
}

func (NnLocale) commaOnThe(string, string) string {
	return ", på "
}

func (NnLocale) daysOfTheWeek() []string {
	return []string{"sundag", "måndag", "tysdag", "onsdag", "torsdag", "fredag", "laurdag"}
}

func (NnLocale) monthsOfTheYear() []string {
	return []string{"januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"}
}
