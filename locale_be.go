package cronstrue

// Code generated from cronstrue/src/i18n/locales/be.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("be", BeLocale{enLocale{}})
}

// BeLocale is the be locale.
type BeLocale struct {
	enLocale
}

func (BeLocale) use24HourTimeFormatByDefault() bool {
	return false
}

func (BeLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Адбылася памылка падчас генерацыі апісання выразы. Праверце сінтаксіс крон-выразы."
}

func (BeLocale) everyMinute() string {
	return "кожную хвіліну"
}

func (BeLocale) everyHour() string {
	return "кожную гадзіну"
}

func (BeLocale) atSpace() string {
	return "У "
}

func (BeLocale) everyMinuteBetweenX0AndX1() string {
	return "Кожную хвіліну з %s да %s"
}

func (BeLocale) at() string {
	return "У"
}

func (BeLocale) spaceAnd() string {
	return " і"
}

func (BeLocale) everySecond() string {
	return "кожную секунду"
}

func (BeLocale) secondsX0ThroughX1PastTheMinute() string {
	return "секунды з %s па %s"
}

func (BeLocale) minutesX0ThroughX1PastTheHour() string {
	return "хвіліны з %s па %s"
}

func (BeLocale) betweenX0AndX1() string {
	return "з %s па %s"
}

func (BeLocale) atX0() string {
	return "у %s"
}

func (BeLocale) commaEveryDay() string {
	return ", кожны дзень"
}

func (BeLocale) spaceX0OfTheMonth() string {
	return " %s месяца"
}

func (BeLocale) lastDay() string {
	return "апошні дзень"
}

func (BeLocale) commaAndOnX0() string {
	return ", і ў %s"
}

func (BeLocale) commaOnlyInX0() string {
	return ", толькі ў %s"
}

func (BeLocale) commaOnTheLastDayOfTheMonth() string {
	return ", у апошні дзень месяца"
}

func (BeLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", у апошні будні дзень месяца"
}

func (BeLocale) firstWeekday() string {
	return "першы будны дзень"
}

func (BeLocale) weekdayNearestDayX0() string {
	return "найбліжэйшы будны дзень да %s"
}

func (BeLocale) commaOnTheX0OfTheMonth() string {
	return ", у %s месяцы"
}

func (BeLocale) commaStartingX0() string {
	return ", пачатак %s"
}

func (BeLocale) atReboot() (string, bool) {
	return "Выканаць адзін раз, пры запуску", true
}

func (BeLocale) onTheHour() (string, bool) {
	return "роўна ў гадзіну", true
}

func (BeLocale) everyX0Seconds(string) string {
	return "кожныя %s секунд"
}

func (BeLocale) everyX0Minutes(string) string {
	return "кожныя %s хвілін"
}

func (BeLocale) everyX0Hours(string) string {
	return "кожныя %s гадзін"
}

func (BeLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", кожныя %s дзён тыдня"
}

func (BeLocale) commaEveryX0Months(string) string {
	return ", кожныя %s месяцаў"
}

func (BeLocale) commaEveryX0Days(string) string {
	return ", кожныя %s дзён"
}

func (BeLocale) commaEveryX0Years(string) string {
	return ", кожныя %s гадоў"
}

func (BeLocale) atX0SecondsPastTheMinute(string) string {
	return "у %s секунд"
}

func (BeLocale) atX0MinutesPastTheHour(string) string {
	return "у %s хвілін"
}

func (BeLocale) commaX0ThroughX1(string) string {
	return ", %s па %s"
}

func (BeLocale) commaAndX0ThroughX1(string) string {
	return ", і %s па %s"
}

func (BeLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", у апошні %s месяца"
}

func (BeLocale) commaOnDayX0OfTheMonth(string) string {
	return ", у %s лік месяца"
}

func (BeLocale) first(string) string {
	return "першы"
}

func (BeLocale) second(string) string {
	return "другі"
}

func (BeLocale) third(string) string {
	return "трэці"
}

func (BeLocale) fourth(string) string {
	return "чацвёрты"
}

func (BeLocale) fifth(string) string {
	return "пяты"
}

func (BeLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", з %s па %s лік месяца"
}

func (BeLocale) commaOnlyOnX0(string) string {
	return ", толькі ў %s"
}

func (BeLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s дзён да апошняга дня месяца"
}

func (BeLocale) commaOnThe(string, string) string {
	return ", у "
}

func (BeLocale) daysOfTheWeek() []string {
	return []string{"нядзеля", "панядзелак", "аўторак", "серада", "чацвер", "пятніца", "субота"}
}

func (BeLocale) monthsOfTheYear() []string {
	return []string{"студзень", "люты", "сакавік", "красавік", "травень", "чэрвень", "ліпень", "жнівень", "верасень", "кастрычнік", "лістапад", "снежань"}
}
