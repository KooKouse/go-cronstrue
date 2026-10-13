package cronstrue

// Code generated from cronstrue/src/i18n/locales/ca.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("ca", CaLocale{enLocale{}})
}

// CaLocale is the ca locale.
type CaLocale struct {
	enLocale
}

func (CaLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (CaLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "S'ha produït un error mentres es generava la descripció de l'expressió. Revisi la sintaxi de la expressió de cron."
}

func (CaLocale) everyMinute() string {
	return "cada minut"
}

func (CaLocale) everyHour() string {
	return "cada hora"
}

func (CaLocale) atSpace() string {
	return "A les "
}

func (CaLocale) everyMinuteBetweenX0AndX1() string {
	return "cada minut entre les %s i les %s"
}

func (CaLocale) at() string {
	return "A les"
}

func (CaLocale) spaceAnd() string {
	return " i"
}

func (CaLocale) everySecond() string {
	return "cada segon"
}

func (CaLocale) secondsX0ThroughX1PastTheMinute() string {
	return "En els segons %s al %s de cada minut"
}

func (CaLocale) minutesX0ThroughX1PastTheHour() string {
	return "del minut %s al %s passada l'hora"
}

func (CaLocale) betweenX0AndX1() string {
	return "entre les %s i les %s"
}

func (CaLocale) atX0() string {
	return "a les %s"
}

func (CaLocale) commaEveryDay() string {
	return ", cada dia"
}

func (CaLocale) spaceX0OfTheMonth() string {
	return " %s del mes"
}

func (CaLocale) lastDay() string {
	return "l'últim dia"
}

func (CaLocale) commaAndOnX0() string {
	return ", i el %s"
}

func (CaLocale) commaOnlyInX0() string {
	return ", sólo en %s"
}

func (CaLocale) commaOnTheLastDayOfTheMonth() string {
	return ", en l'últim dia del mes"
}

func (CaLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", en l'últim dia de la setmana del mes"
}

func (CaLocale) firstWeekday() string {
	return "primer dia de la setmana"
}

func (CaLocale) weekdayNearestDayX0() string {
	return "dia de la setmana més proper al %s"
}

func (CaLocale) commaOnTheX0OfTheMonth() string {
	return ", en el %s del mes"
}

func (CaLocale) commaStartingX0() string {
	return ", començant %s"
}

func (CaLocale) atReboot() (string, bool) {
	return "Executa una vegada, en iniciar", true
}

func (CaLocale) onTheHour() (string, bool) {
	return "en punt", true
}

func (CaLocale) everyX0Seconds(string) string {
	return "cada %s segons"
}

func (CaLocale) everyX0Minutes(string) string {
	return "cada %s minuts"
}

func (CaLocale) everyX0Hours(string) string {
	return "cada %s hores"
}

func (CaLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", cada %s dies de la setmana"
}

func (CaLocale) commaEveryX0Months(string) string {
	return ", cada %s mesos"
}

func (CaLocale) commaEveryX0Days(string) string {
	return ", cada %s dies"
}

func (CaLocale) commaEveryX0Years(string) string {
	return ", cada %s anys"
}

func (CaLocale) atX0SecondsPastTheMinute(string) string {
	return "als %s segonds del minut"
}

func (CaLocale) atX0MinutesPastTheHour(string) string {
	return "als %s minuts de l'hora"
}

func (CaLocale) commaX0ThroughX1(string) string {
	return ", de %s a %s"
}

func (CaLocale) commaAndX0ThroughX1(string) string {
	return ", i de %s a %s"
}

func (CaLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", en l'últim %s del mes"
}

func (CaLocale) commaOnDayX0OfTheMonth(string) string {
	return ", el dia %s del mes"
}

func (CaLocale) first(string) string {
	return "primer"
}

func (CaLocale) second(string) string {
	return "segon"
}

func (CaLocale) third(string) string {
	return "tercer"
}

func (CaLocale) fourth(string) string {
	return "quart"
}

func (CaLocale) fifth(string) string {
	return "cinquè"
}

func (CaLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", entre els dies %s i %s del mes"
}

func (CaLocale) commaOnlyOnX0(string) string {
	return ", només el %s"
}

func (CaLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dies abans de l'últim dia del mes"
}

func (CaLocale) commaOnThe(string, string) string {
	return ", en el "
}

func (CaLocale) daysOfTheWeek() []string {
	return []string{"diumenge", "dilluns", "dimarts", "dimecres", "dijous", "divendres", "dissabte"}
}

func (CaLocale) monthsOfTheYear() []string {
	return []string{"gener", "febrer", "març", "abril", "maig", "juny", "juliol", "agost", "setembre", "octubre", "novembre", "desembre"}
}
