package cronstrue

// Code generated from cronstrue/src/i18n/locales/hu.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("hu", HuLocale{enLocale{}})
}

// HuLocale is the hu locale.
type HuLocale struct {
	enLocale
}

func (HuLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (HuLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Hiba történt a kifejezésleírás generálásakor. Ellenőrizze a cron kifejezés szintaxisát."
}

func (HuLocale) everyMinute() string {
	return "minden percben"
}

func (HuLocale) everyHour() string {
	return "minden órában"
}

func (HuLocale) atSpace() string {
	return "Ekkor: "
}

func (HuLocale) everyMinuteBetweenX0AndX1() string {
	return "percenként %s és %s között"
}

func (HuLocale) at() string {
	return "Ekkor:"
}

func (HuLocale) spaceAnd() string {
	return " és"
}

func (HuLocale) everySecond() string {
	return "minden másodpercben"
}

func (HuLocale) secondsX0ThroughX1PastTheMinute() string {
	return "%s. másodpercben %s perc után"
}

func (HuLocale) minutesX0ThroughX1PastTheHour() string {
	return "%s. percben %s óra után"
}

func (HuLocale) betweenX0AndX1() string {
	return "%s és %s között"
}

func (HuLocale) atX0() string {
	return "ekkor %s"
}

func (HuLocale) commaEveryDay() string {
	return ", minden nap"
}

func (HuLocale) spaceX0OfTheMonth() string {
	return " %s a hónapban"
}

func (HuLocale) lastDay() string {
	return "az utolsó nap"
}

func (HuLocale) commaAndOnX0() string {
	return ", és %s"
}

func (HuLocale) commaOnlyInX0() string {
	return ", csak ekkor: %s"
}

func (HuLocale) commaOnTheLastDayOfTheMonth() string {
	return ", a hónap utolsó napján"
}

func (HuLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", a hónap utolsó hétköznapján"
}

func (HuLocale) firstWeekday() string {
	return "első hétköznap"
}

func (HuLocale) weekdayNearestDayX0() string {
	return "hétköznap legközelebbi nap %s"
}

func (HuLocale) commaOnTheX0OfTheMonth() string {
	return ", a hónap %s"
}

func (HuLocale) commaStartingX0() string {
	return ", %s kezdettel"
}

func (HuLocale) atReboot() (string, bool) {
	return "Egyszer fut le, rendszerindításkor", true
}

func (HuLocale) onTheHour() (string, bool) {
	return "órakor", true
}

func (HuLocale) everyX0Seconds(string) string {
	return "%s másodpercenként"
}

func (HuLocale) everyX0Minutes(string) string {
	return "minden %s. percben"
}

func (HuLocale) everyX0Hours(string) string {
	return "minden %s órában"
}

func (HuLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", a hét minden %s napján"
}

func (HuLocale) commaEveryX0Months(string) string {
	return ", minden %s hónapban"
}

func (HuLocale) commaEveryX0Days(string) string {
	return ", %s naponként"
}

func (HuLocale) commaEveryX0Years(string) string {
	return ", %s évente"
}

func (HuLocale) atX0SecondsPastTheMinute(string) string {
	return "%s. másodpercben"
}

func (HuLocale) atX0MinutesPastTheHour(string) string {
	return "%s. percben"
}

func (HuLocale) commaX0ThroughX1(string) string {
	return ", %s - %s"
}

func (HuLocale) commaAndX0ThroughX1(string) string {
	return ", és %s - %s"
}

func (HuLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", a hónap utolsó %s"
}

func (HuLocale) commaOnDayX0OfTheMonth(string) string {
	return ", a hónap %s napján"
}

func (HuLocale) first(string) string {
	return "első"
}

func (HuLocale) second(string) string {
	return "második"
}

func (HuLocale) third(string) string {
	return "harmadik"
}

func (HuLocale) fourth(string) string {
	return "negyedik"
}

func (HuLocale) fifth(string) string {
	return "ötödik"
}

func (HuLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", a hónap %s és %s napja között"
}

func (HuLocale) commaOnlyOnX0(string) string {
	return ", csak ekkor: %s"
}

func (HuLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s nappal a hónap utolsó napja előtt"
}

func (HuLocale) commaOnThe(string, string) string {
	return ", "
}

func (HuLocale) daysOfTheWeek() []string {
	return []string{"vasárnap", "hétfő", "kedd", "szerda", "csütörtök", "péntek", "szombat"}
}

func (HuLocale) monthsOfTheYear() []string {
	return []string{"január", "február", "március", "április", "május", "június", "július", "augusztus", "szeptember", "október", "november", "december"}
}
