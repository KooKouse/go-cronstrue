package cronstrue

// Code generated from cronstrue/src/i18n/locales/nb.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("nb", NbLocale{enLocale{}})
}

// NbLocale is the nb locale.
type NbLocale struct {
	enLocale
}

func (NbLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (NbLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "En feil inntraff ved generering av uttrykksbeskrivelse. Sjekk cron syntaks."
}

func (NbLocale) everyMinute() string {
	return "hvert minutt"
}

func (NbLocale) everyHour() string {
	return "hver time"
}

func (NbLocale) atSpace() string {
	return "Kl."
}

func (NbLocale) everyMinuteBetweenX0AndX1() string {
	return "Hvert minutt mellom %s og %s"
}

func (NbLocale) at() string {
	return "Kl."
}

func (NbLocale) spaceAnd() string {
	return " og"
}

func (NbLocale) everySecond() string {
	return "hvert sekund"
}

func (NbLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekundene fra %s til og med %s etter minuttet"
}

func (NbLocale) minutesX0ThroughX1PastTheHour() string {
	return "minuttene fra %s til og med %s etter timen"
}

func (NbLocale) betweenX0AndX1() string {
	return "mellom %s og %s"
}

func (NbLocale) atX0() string {
	return "på %s"
}

func (NbLocale) commaEveryDay() string {
	return ", hver dag"
}

func (NbLocale) spaceX0OfTheMonth() string {
	return " %s i måneden"
}

func (NbLocale) lastDay() string {
	return "den siste dagen"
}

func (NbLocale) commaAndOnX0() string {
	return ", og på %s"
}

func (NbLocale) commaOnlyInX0() string {
	return ", bare i %s"
}

func (NbLocale) commaOnTheLastDayOfTheMonth() string {
	return ", på den siste dagen i måneden"
}

func (NbLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", den siste ukedagen i måneden"
}

func (NbLocale) firstWeekday() string {
	return "første ukedag"
}

func (NbLocale) weekdayNearestDayX0() string {
	return "ukedag nærmest dag %s"
}

func (NbLocale) commaOnTheX0OfTheMonth() string {
	return ", på den %s av måneden"
}

func (NbLocale) commaStartingX0() string {
	return ", starter %s"
}

func (NbLocale) atReboot() (string, bool) {
	return "Kjør én gang, ved oppstart", true
}

func (NbLocale) onTheHour() (string, bool) {
	return "på timen", true
}

func (NbLocale) everyX0Seconds(string) string {
	return "hvert %s sekund"
}

func (NbLocale) everyX0Minutes(string) string {
	return "hvert %s minutt"
}

func (NbLocale) everyX0Hours(string) string {
	return "hver %s time"
}

func (NbLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", hver %s ukedag"
}

func (NbLocale) commaEveryX0Months(string) string {
	return ", hver %s måned"
}

func (NbLocale) commaEveryX0Days(string) string {
	return ", hver %s dag i måneden"
}

func (NbLocale) commaEveryX0Years(string) string {
	return ", hvert %s år"
}

func (NbLocale) atX0SecondsPastTheMinute(string) string {
	return "på %s sekunder etter minuttet"
}

func (NbLocale) atX0MinutesPastTheHour(string) string {
	return "på %s minutter etter timen"
}

func (NbLocale) commaX0ThroughX1(string) string {
	return ", %s til og med %s"
}

func (NbLocale) commaAndX0ThroughX1(string) string {
	return ", og %s til og med %s"
}

func (NbLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", på den siste %s av måneden"
}

func (NbLocale) commaOnDayX0OfTheMonth(string) string {
	return ", på dag %s av måneden"
}

func (NbLocale) first(string) string {
	return "første"
}

func (NbLocale) second(string) string {
	return "andre"
}

func (NbLocale) third(string) string {
	return "tredje"
}

func (NbLocale) fourth(string) string {
	return "fjerde"
}

func (NbLocale) fifth(string) string {
	return "femte"
}

func (NbLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", mellom dag %s og %s av måneden"
}

func (NbLocale) commaOnlyOnX0(string) string {
	return ", på %s"
}

func (NbLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dager før den siste dagen i måneden"
}

func (NbLocale) commaOnThe(string, string) string {
	return ", på "
}

func (NbLocale) daysOfTheWeek() []string {
	return []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"}
}

func (NbLocale) monthsOfTheYear() []string {
	return []string{"januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"}
}
