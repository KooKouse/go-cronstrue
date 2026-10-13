package cronstrue

// Code generated from cronstrue/src/i18n/locales/da.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("da", DaLocale{enLocale{}})
}

// DaLocale is the da locale.
type DaLocale struct {
	enLocale
}

func (DaLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (DaLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Der opstod en fejl ved generering af udtryksbeskrivelsen. Tjek cron-ekspressionssyntaxen."
}

func (DaLocale) everyMinute() string {
	return "hvert minut"
}

func (DaLocale) everyHour() string {
	return "hver time"
}

func (DaLocale) atSpace() string {
	return "kl "
}

func (DaLocale) everyMinuteBetweenX0AndX1() string {
	return "hvert minut mellem %s og %s"
}

func (DaLocale) at() string {
	return "kl"
}

func (DaLocale) spaceAnd() string {
	return " og"
}

func (DaLocale) everySecond() string {
	return "hvert sekund"
}

func (DaLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekunderne fra %s til og med %s hvert minut"
}

func (DaLocale) minutesX0ThroughX1PastTheHour() string {
	return "minutterne fra %s til og med %s hver time"
}

func (DaLocale) betweenX0AndX1() string {
	return "mellem %s og %s"
}

func (DaLocale) atX0() string {
	return "kl %s"
}

func (DaLocale) commaEveryDay() string {
	return ", hver dag"
}

func (DaLocale) spaceX0OfTheMonth() string {
	return " %s i måneden"
}

func (DaLocale) lastDay() string {
	return "sidste dag"
}

func (DaLocale) commaAndOnX0() string {
	return ", og på %s"
}

func (DaLocale) commaOnlyInX0() string {
	return ", kun i %s"
}

func (DaLocale) commaOnTheLastDayOfTheMonth() string {
	return ", på den sidste dag i måneden"
}

func (DaLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", på den sidste hverdag i måneden"
}

func (DaLocale) firstWeekday() string {
	return "første hverdag"
}

func (DaLocale) weekdayNearestDayX0() string {
	return "hverdag nærmest dag %s"
}

func (DaLocale) commaOnTheX0OfTheMonth() string {
	return ", på den %s i måneden"
}

func (DaLocale) commaStartingX0() string {
	return ", startende %s"
}

func (DaLocale) atReboot() (string, bool) {
	return "Kør én gang, ved opstart", true
}

func (DaLocale) onTheHour() (string, bool) {
	return "på timen", true
}

func (DaLocale) everyX0Seconds(string) string {
	return "hvert %s. sekund"
}

func (DaLocale) everyX0Minutes(string) string {
	return "hvert %s. minut"
}

func (DaLocale) everyX0Hours(string) string {
	return "hver %s. time"
}

func (DaLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", hver %s. ugedag"
}

func (DaLocale) commaEveryX0Months(string) string {
	return ", hver %s. måned"
}

func (DaLocale) commaEveryX0Days(string) string {
	return ", hver %s. dag"
}

func (DaLocale) commaEveryX0Years(string) string {
	return ", hvert %s. år"
}

func (DaLocale) atX0SecondsPastTheMinute(string) string {
	return "%s sekunder efter minutskift"
}

func (DaLocale) atX0MinutesPastTheHour(string) string {
	return "%s minutter efter timeskift"
}

func (DaLocale) commaX0ThroughX1(string) string {
	return ", %s til og med %s"
}

func (DaLocale) commaAndX0ThroughX1(string) string {
	return ", og %s til og med %s"
}

func (DaLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", på den sidste %s i måneden"
}

func (DaLocale) commaOnDayX0OfTheMonth(string) string {
	return ", på dag %s i måneden"
}

func (DaLocale) first(string) string {
	return "første"
}

func (DaLocale) second(string) string {
	return "anden"
}

func (DaLocale) third(string) string {
	return "tredje"
}

func (DaLocale) fourth(string) string {
	return "fjerde"
}

func (DaLocale) fifth(string) string {
	return "femte"
}

func (DaLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", mellem dag %s og %s i måneden"
}

func (DaLocale) commaOnlyOnX0(string) string {
	return ", på enhver %s"
}

func (DaLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dage før den sidste dag i måneden"
}

func (DaLocale) commaOnThe(string, string) string {
	return ", på den "
}

func (DaLocale) daysOfTheWeek() []string {
	return []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"}
}

func (DaLocale) monthsOfTheYear() []string {
	return []string{"januar", "februar", "marts", "april", "maj", "juni", "juli", "august", "september", "oktober", "november", "december"}
}
