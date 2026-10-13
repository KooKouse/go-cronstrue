package cronstrue

// Code generated from cronstrue/src/i18n/locales/fi.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("fi", FiLocale{enLocale{}})
}

// FiLocale is the fi locale.
type FiLocale struct {
	enLocale
}

func (FiLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (FiLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Virhe kuvauksen generoinnissa. Tarkista cron-syntaksi."
}

func (FiLocale) everyMinute() string {
	return "joka minuutti"
}

func (FiLocale) everyHour() string {
	return "joka tunti"
}

func (FiLocale) atSpace() string {
	return "Klo "
}

func (FiLocale) everyMinuteBetweenX0AndX1() string {
	return "joka minuutti %s - %s välillä"
}

func (FiLocale) at() string {
	return "Klo"
}

func (FiLocale) spaceAnd() string {
	return " ja"
}

func (FiLocale) everySecond() string {
	return "joka sekunti"
}

func (FiLocale) secondsX0ThroughX1PastTheMinute() string {
	return "joka minuutti sekunttien %s - %s välillä"
}

func (FiLocale) minutesX0ThroughX1PastTheHour() string {
	return "joka tunti minuuttien %s - %s välillä"
}

func (FiLocale) betweenX0AndX1() string {
	return "%s - %s välillä"
}

func (FiLocale) atX0() string {
	return "klo %s"
}

func (FiLocale) commaEveryDay() string {
	return ", joka päivä"
}

func (FiLocale) spaceX0OfTheMonth() string {
	return " %s kuukaudessa"
}

func (FiLocale) lastDay() string {
	return "viimeinen päivä"
}

func (FiLocale) commaAndOnX0() string {
	return ", ja edelleen %s"
}

func (FiLocale) commaOnlyInX0() string {
	return ", vain %s"
}

func (FiLocale) commaOnTheLastDayOfTheMonth() string {
	return ", kuukauden viimeisenä päivänä"
}

func (FiLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", kuukauden viimeisenä viikonpäivänä"
}

func (FiLocale) firstWeekday() string {
	return "ensimmäinen viikonpäivä"
}

func (FiLocale) weekdayNearestDayX0() string {
	return "viikonpäivä lähintä %s päivää"
}

func (FiLocale) commaOnTheX0OfTheMonth() string {
	return ", kuukauden %s"
}

func (FiLocale) commaStartingX0() string {
	return ", alkaen %s"
}

func (FiLocale) atX0MinutesPastTheHourGt20() (string, bool) {
	return "%s minuuttia yli", true
}

func (FiLocale) atReboot() (string, bool) {
	return "Suorita kerran, käynnistyksen yhteydessä", true
}

func (FiLocale) onTheHour() (string, bool) {
	return "tasalta", true
}

func (FiLocale) everyX0Seconds(string) string {
	return "joka %s. sekunti"
}

func (FiLocale) everyX0Minutes(string) string {
	return "joka %s. minuutti"
}

func (FiLocale) everyX0Hours(string) string {
	return "joka %s. tunti"
}

func (FiLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", joka %s. viikonpäivä"
}

func (FiLocale) commaEveryX0Months(string) string {
	return ", joka %s. kuukausi"
}

func (FiLocale) commaEveryX0Days(string) string {
	return ", joka %s. päivä"
}

func (FiLocale) commaEveryX0Years(string) string {
	return ", joka %s. vuosi"
}

func (FiLocale) atX0SecondsPastTheMinute(string) string {
	return "%s sekunnnin jälkeen"
}

func (FiLocale) atX0MinutesPastTheHour(string) string {
	return "%s minuuttia yli"
}

func (FiLocale) commaX0ThroughX1(string) string {
	return ", %s - %s"
}

func (FiLocale) commaAndX0ThroughX1(string) string {
	return ", %s - %s"
}

func (FiLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", kuukauden viimeinen %s"
}

func (FiLocale) commaOnDayX0OfTheMonth(string) string {
	return ", kuukauden %s päivä"
}

func (FiLocale) first(string) string {
	return "ensimmäinen"
}

func (FiLocale) second(string) string {
	return "toinen"
}

func (FiLocale) third(string) string {
	return "kolmas"
}

func (FiLocale) fourth(string) string {
	return "neljäs"
}

func (FiLocale) fifth(string) string {
	return "viides"
}

func (FiLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", kuukauden päivien %s ja %s välillä"
}

func (FiLocale) commaOnlyOnX0(string) string {
	return ", vain %s"
}

func (FiLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s päivää ennen kuukauden viimeistä päivää"
}

func (FiLocale) commaOnThe(string, string) string {
	return ","
}

func (FiLocale) daysOfTheWeek() []string {
	return []string{"sunnuntai", "maanantai", "tiistai", "keskiviikko", "torstai", "perjantai", "lauantai"}
}

func (FiLocale) monthsOfTheYear() []string {
	return []string{"tammikuu", "helmikuu", "maaliskuu", "huhtikuu", "toukokuu", "kesäkuu", "heinäkuu", "elokuu", "syyskuu", "lokakuu", "marraskuu", "joulukuu"}
}
