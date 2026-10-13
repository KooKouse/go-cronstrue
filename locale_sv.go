package cronstrue

// Code generated from cronstrue/src/i18n/locales/sv.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("sv", SvLocale{enLocale{}})
}

// SvLocale is the sv locale.
type SvLocale struct {
	enLocale
}

func (SvLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (SvLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Ett fel inträffade vid generering av uttryckets beskrivning. Kontrollera cron-uttryckets syntax."
}

func (SvLocale) everyMinute() string {
	return "varje minut"
}

func (SvLocale) everyHour() string {
	return "varje timme"
}

func (SvLocale) atSpace() string {
	return "Kl "
}

func (SvLocale) everyMinuteBetweenX0AndX1() string {
	return "Varje minut mellan %s och %s"
}

func (SvLocale) at() string {
	return "Kl"
}

func (SvLocale) spaceAnd() string {
	return " och"
}

func (SvLocale) everySecond() string {
	return "varje sekund"
}

func (SvLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekunderna från %s till och med %s efter minuten"
}

func (SvLocale) minutesX0ThroughX1PastTheHour() string {
	return "minuterna från %s till och med %s efter timmen"
}

func (SvLocale) betweenX0AndX1() string {
	return "mellan %s och %s"
}

func (SvLocale) atX0() string {
	return "kl %s"
}

func (SvLocale) commaEveryDay() string {
	return ", varje dag"
}

func (SvLocale) spaceX0OfTheMonth() string {
	return " %sen av månaden"
}

func (SvLocale) lastDay() string {
	return "den sista dagen"
}

func (SvLocale) commaAndOnX0() string {
	return ", och på %s"
}

func (SvLocale) commaOnlyInX0() string {
	return ", bara på %s"
}

func (SvLocale) commaOnTheLastDayOfTheMonth() string {
	return ", på sista dagen av månaden"
}

func (SvLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", på sista veckodag av månaden"
}

func (SvLocale) firstWeekday() string {
	return "första veckodag"
}

func (SvLocale) weekdayNearestDayX0() string {
	return "veckodagen närmast dag %s"
}

func (SvLocale) commaOnTheX0OfTheMonth() string {
	return ", på den %s av månaden"
}

func (SvLocale) commaStartingX0() string {
	return ", startar %s"
}

func (SvLocale) atReboot() (string, bool) {
	return "Kör en gång, vid uppstart", true
}

func (SvLocale) onTheHour() (string, bool) {
	return "på heltimmen", true
}

func (SvLocale) everyX0Seconds(string) string {
	return "varje %s sekund"
}

func (SvLocale) everyX0Minutes(string) string {
	return "var %s minut"
}

func (SvLocale) everyX0Hours(string) string {
	return "var %s timme"
}

func (SvLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", var %s dag i veckan"
}

func (SvLocale) commaEveryX0Months(string) string {
	return ", var %s månad"
}

func (SvLocale) commaEveryX0Days(string) string {
	return ", var %s dag"
}

func (SvLocale) commaEveryX0Years(string) string {
	return ", var %s år"
}

func (SvLocale) atX0SecondsPastTheMinute(string) string {
	return "på %s sekunder efter minuten"
}

func (SvLocale) atX0MinutesPastTheHour(string) string {
	return "på %s minuten efter timmen"
}

func (SvLocale) commaX0ThroughX1(string) string {
	return ", %s till %s"
}

func (SvLocale) commaAndX0ThroughX1(string) string {
	return ", och %s till %s"
}

func (SvLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", på sista %s av månaden"
}

func (SvLocale) commaOnDayX0OfTheMonth(string) string {
	return ", på dag %s av månaden"
}

func (SvLocale) first(string) string {
	return "första"
}

func (SvLocale) second(string) string {
	return "andra"
}

func (SvLocale) third(string) string {
	return "tredje"
}

func (SvLocale) fourth(string) string {
	return "fjärde"
}

func (SvLocale) fifth(string) string {
	return "femte"
}

func (SvLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", mellan dag %s och %s av månaden"
}

func (SvLocale) commaOnlyOnX0(string) string {
	return ", varje %s"
}

func (SvLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dagar före den sista dagen i månaden"
}

func (SvLocale) commaOnThe(string, string) string {
	return ", den "
}

func (SvLocale) daysOfTheWeek() []string {
	return []string{"söndag", "måndag", "tisdag", "onsdag", "torsdag", "fredag", "lördag"}
}

func (SvLocale) monthsOfTheYear() []string {
	return []string{"januari", "februari", "mars", "april", "maj", "juni", "juli", "augusti", "september", "oktober", "november", "december"}
}
