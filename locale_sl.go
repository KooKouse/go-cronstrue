package cronstrue

// Code generated from cronstrue/src/i18n/locales/sl.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("sl", SlLocale{enLocale{}})
}

// SlLocale is the sl locale.
type SlLocale struct {
	enLocale
}

func (SlLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (SlLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Pri generiranju opisa izraza je prišlo do napake. Preverite sintakso izraza cron."
}

func (SlLocale) everyMinute() string {
	return "vsako minuto"
}

func (SlLocale) everyHour() string {
	return "vsako uro"
}

func (SlLocale) atSpace() string {
	return "Ob "
}

func (SlLocale) everyMinuteBetweenX0AndX1() string {
	return "Vsako minuto od %s do %s"
}

func (SlLocale) at() string {
	return "Ob"
}

func (SlLocale) spaceAnd() string {
	return " in"
}

func (SlLocale) everySecond() string {
	return "vsako sekundo"
}

func (SlLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekunde od %s do %s"
}

func (SlLocale) minutesX0ThroughX1PastTheHour() string {
	return "minute od %s do %s"
}

func (SlLocale) betweenX0AndX1() string {
	return "od %s do %s"
}

func (SlLocale) atX0() string {
	return "ob %s"
}

func (SlLocale) commaEveryDay() string {
	return ", vsak dan"
}

func (SlLocale) spaceX0OfTheMonth() string {
	return " %s v mesecu"
}

func (SlLocale) lastDay() string {
	return "zadnjič"
}

func (SlLocale) commaAndOnX0() string {
	return "in naprej %s"
}

func (SlLocale) commaOnlyInX0() string {
	return ", samo v %s"
}

func (SlLocale) commaOnTheLastDayOfTheMonth() string {
	return ", zadnji dan v mesecu"
}

func (SlLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", zadnji delovni dan v mesecu"
}

func (SlLocale) firstWeekday() string {
	return "prvi delovni dan"
}

func (SlLocale) weekdayNearestDayX0() string {
	return "delovni dan, najbližji %s. dnevu"
}

func (SlLocale) commaOnTheX0OfTheMonth() string {
	return ", %s v mesecu"
}

func (SlLocale) commaStartingX0() string {
	return ", začenši %s"
}

func (SlLocale) atReboot() (string, bool) {
	return "Zaženi enkrat, ob zagonu", true
}

func (SlLocale) onTheHour() (string, bool) {
	return "ob točni uri", true
}

func (SlLocale) everyX0Seconds(string) string {
	return "vsakih %s sekund"
}

func (SlLocale) everyX0Minutes(string) string {
	return "vsakih %s minut"
}

func (SlLocale) everyX0Hours(string) string {
	return "vsakih %s ur"
}

func (SlLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", vsakih %s dni v tednu"
}

func (SlLocale) commaEveryX0Months(string) string {
	return ", vsakih %s mesecev"
}

func (SlLocale) commaEveryX0Days(string) string {
	return ", vsakih %s dni"
}

func (SlLocale) commaEveryX0Years(string) string {
	return ", vsakih %s let"
}

func (SlLocale) atX0SecondsPastTheMinute(string) string {
	return "ob %s."
}

func (SlLocale) atX0MinutesPastTheHour(string) string {
	return "ob %s."
}

func (SlLocale) commaX0ThroughX1(string) string {
	return ", od %s do %s"
}

func (SlLocale) commaAndX0ThroughX1(string) string {
	return ", in od %s do %s"
}

func (SlLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", zadnji %s v mesecu"
}

func (SlLocale) commaOnDayX0OfTheMonth(string) string {
	return ", %s. dan v mesecu"
}

func (SlLocale) first(string) string {
	return "prvi"
}

func (SlLocale) second(string) string {
	return "drugi"
}

func (SlLocale) third(string) string {
	return "tretji"
}

func (SlLocale) fourth(string) string {
	return "četrti"
}

func (SlLocale) fifth(string) string {
	return "peti"
}

func (SlLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", od %s. do %s. dne v mesecu"
}

func (SlLocale) commaOnlyOnX0(string) string {
	return ", samo v %s"
}

func (SlLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dni pred koncem meseca"
}

func (SlLocale) commaOnThe(string, string) string {
	return ", "
}

func (SlLocale) daysOfTheWeek() []string {
	return []string{"Nedelja", "Ponedeljek", "Torek", "Sreda", "Četrtek", "Petek", "Sobota"}
}

func (SlLocale) monthsOfTheYear() []string {
	return []string{"januar", "februar", "marec", "april", "maj", "junij", "julij", "avgust", "september", "oktober", "november", "december"}
}
