package cronstrue

// Code generated from cronstrue/src/i18n/locales/cs.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("cs", CsLocale{enLocale{}})
}

// CsLocale is the cs locale.
type CsLocale struct {
	enLocale
}

func (CsLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (CsLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Při vytváření popisu došlo k chybě. Zkontrolujte prosím správnost syntaxe cronu."
}

func (CsLocale) everyMinute() string {
	return "každou minutu"
}

func (CsLocale) everyHour() string {
	return "každou hodinu"
}

func (CsLocale) atSpace() string {
	return "V "
}

func (CsLocale) everyMinuteBetweenX0AndX1() string {
	return "Každou minutu mezi %s a %s"
}

func (CsLocale) at() string {
	return "V"
}

func (CsLocale) spaceAnd() string {
	return " a"
}

func (CsLocale) everySecond() string {
	return "každou sekundu"
}

func (CsLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekundy od %s do %s"
}

func (CsLocale) minutesX0ThroughX1PastTheHour() string {
	return "minuty od %s do %s"
}

func (CsLocale) betweenX0AndX1() string {
	return "mezi %s a %s"
}

func (CsLocale) atX0() string {
	return "v %s"
}

func (CsLocale) commaEveryDay() string {
	return ", každý den"
}

func (CsLocale) spaceX0OfTheMonth() string {
	return " %s v měsíci"
}

func (CsLocale) lastDay() string {
	return "poslední den"
}

func (CsLocale) commaAndOnX0() string {
	return ", a v %s"
}

func (CsLocale) commaOnlyInX0() string {
	return ", pouze v %s"
}

func (CsLocale) commaOnTheLastDayOfTheMonth() string {
	return ", poslední den v měsíci"
}

func (CsLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", poslední pracovní den v měsíci"
}

func (CsLocale) firstWeekday() string {
	return "první pracovní den"
}

func (CsLocale) weekdayNearestDayX0() string {
	return "pracovní den nejblíže %s. dni"
}

func (CsLocale) commaOnTheX0OfTheMonth() string {
	return ", v %s v měsíci"
}

func (CsLocale) commaStartingX0() string {
	return ", začínající %s"
}

func (CsLocale) atReboot() (string, bool) {
	return "Spustit jednou, při startu", true
}

func (CsLocale) onTheHour() (string, bool) {
	return "v celou hodinu", true
}

func (CsLocale) everyX0Seconds(string) string {
	return "každých %s sekund"
}

func (CsLocale) everyX0Minutes(string) string {
	return "každých %s minut"
}

func (CsLocale) everyX0Hours(string) string {
	return "každých %s hodin"
}

func (CsLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", každých %s dní v týdnu"
}

func (CsLocale) commaEveryX0Months(string) string {
	return ", každých %s měsíců"
}

func (CsLocale) commaEveryX0Days(string) string {
	return ", každých %s dnů"
}

func (CsLocale) commaEveryX0Years(string) string {
	return ", každých %s roků"
}

func (CsLocale) atX0SecondsPastTheMinute(string) string {
	return "v %s sekund"
}

func (CsLocale) atX0MinutesPastTheHour(string) string {
	return "v %s minut"
}

func (CsLocale) commaX0ThroughX1(string) string {
	return ", od %s do %s"
}

func (CsLocale) commaAndX0ThroughX1(string) string {
	return ", a od %s do %s"
}

func (CsLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", poslední %s v měsíci"
}

func (CsLocale) commaOnDayX0OfTheMonth(string) string {
	return ", %s. den v měsíci"
}

func (CsLocale) first(string) string {
	return "první"
}

func (CsLocale) second(string) string {
	return "druhý"
}

func (CsLocale) third(string) string {
	return "třetí"
}

func (CsLocale) fourth(string) string {
	return "čtvrtý"
}

func (CsLocale) fifth(string) string {
	return "pátý"
}

func (CsLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", mezi dny %s a %s v měsíci"
}

func (CsLocale) commaOnlyOnX0(string) string {
	return ", pouze v %s"
}

func (CsLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dní před posledním dnem v měsíci"
}

func (CsLocale) commaOnThe(string, string) string {
	return ", "
}

func (CsLocale) daysOfTheWeek() []string {
	return []string{"Neděle", "Pondělí", "Úterý", "Středa", "Čtvrtek", "Pátek", "Sobota"}
}

func (CsLocale) monthsOfTheYear() []string {
	return []string{"Leden", "Únor", "Březen", "Duben", "Květen", "Červen", "Červenec", "Srpen", "Září", "Říjen", "Listopad", "Prosinec"}
}
