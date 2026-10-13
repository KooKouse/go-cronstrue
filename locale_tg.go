package cronstrue

// Code generated from cronstrue/src/i18n/locales/tg.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("tg", TgLocale{enLocale{}})
}

// TgLocale is the tg locale.
type TgLocale struct {
	enLocale
}

func (TgLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (TgLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Ҳангоми тавсифи ифодаи cron хато рӯй дод. Синтаксисро санҷед."
}

func (TgLocale) everyMinute() string {
	return "ҳар дақиқа"
}

func (TgLocale) everyHour() string {
	return "ҳар соат"
}

func (TgLocale) atSpace() string {
	return "Дар "
}

func (TgLocale) everyMinuteBetweenX0AndX1() string {
	return "Ҳар дақиқа аз %s то %s"
}

func (TgLocale) at() string {
	return "Дар"
}

func (TgLocale) spaceAnd() string {
	return " ва"
}

func (TgLocale) everySecond() string {
	return "ҳар сония"
}

func (TgLocale) secondsX0ThroughX1PastTheMinute() string {
	return "сонияҳои аз %s то %s"
}

func (TgLocale) minutesX0ThroughX1PastTheHour() string {
	return "дақиқаҳои аз %s то %s"
}

func (TgLocale) betweenX0AndX1() string {
	return "аз %s то %s"
}

func (TgLocale) atX0() string {
	return "дар %s"
}

func (TgLocale) commaEveryDay() string {
	return ", ҳар рӯз"
}

func (TgLocale) spaceX0OfTheMonth() string {
	return " %s-и моҳ"
}

func (TgLocale) lastDay() string {
	return "рӯзи охирин"
}

func (TgLocale) commaAndOnX0() string {
	return ", ва %s"
}

func (TgLocale) commaOnlyInX0() string {
	return ", танҳо дар %s"
}

func (TgLocale) commaOnTheLastDayOfTheMonth() string {
	return ", дар рӯзи охирини моҳ"
}

func (TgLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", дар охирин рӯзи кории моҳ"
}

func (TgLocale) firstWeekday() string {
	return "якумин рӯзи корӣ"
}

func (TgLocale) weekdayNearestDayX0() string {
	return "наздиктарин рӯзи корӣ ба %s"
}

func (TgLocale) commaOnTheX0OfTheMonth() string {
	return ", дар %s-и моҳ"
}

func (TgLocale) commaStartingX0() string {
	return ", оғоз %s"
}

func (TgLocale) commaOnlyInMonthX0() (string, bool) {
	return ", танҳо %s", true
}

func (TgLocale) atReboot() (string, bool) {
	return "Як бор иҷро шавад, ҳангоми оғоз", true
}

func (TgLocale) onTheHour() (string, bool) {
	return "дақиқан дар соат", true
}

func (TgLocale) everyX0Seconds(string) string {
	return "ҳар %s сония"
}

func (TgLocale) everyX0Minutes(string) string {
	return "ҳар %s дақиқа"
}

func (TgLocale) everyX0Hours(string) string {
	return "ҳар %s соат"
}

func (TgLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", ҳар %s рӯзи ҳафта"
}

func (TgLocale) commaEveryX0Months(string) string {
	return " ҳар %s моҳ"
}

func (TgLocale) commaEveryX0Days(string) string {
	return ", ҳар %s рӯз"
}

func (TgLocale) commaEveryX0Years(string) string {
	return ", ҳар %s сол"
}

func (TgLocale) atX0SecondsPastTheMinute(string) string {
	return "дар %s сония"
}

func (TgLocale) atX0MinutesPastTheHour(string) string {
	return "дар %s дақиқа"
}

func (TgLocale) commaX0ThroughX1(string) string {
	return ", аз %s то %s"
}

func (TgLocale) commaAndX0ThroughX1(string) string {
	return " ва аз %s то %s"
}

func (TgLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", дар охирин %s-и моҳ"
}

func (TgLocale) commaOnDayX0OfTheMonth(string) string {
	return ", дар рӯзи %s-и моҳ"
}

func (TgLocale) first(string) string {
	return "якум"
}

func (TgLocale) second(string) string {
	return "дуюм"
}

func (TgLocale) third(string) string {
	return "сеюм"
}

func (TgLocale) fourth(string) string {
	return "чорум"
}

func (TgLocale) fifth(string) string {
	return "панҷум"
}

func (TgLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", аз рӯзи %s то %s-и моҳ"
}

func (TgLocale) commaOnlyOnX0(string) string {
	return ", танҳо дар %s"
}

func (TgLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s рӯз пеш аз охири моҳ"
}

func (TgLocale) commaOnThe(string, string) string {
	return ", дар "
}

func (TgLocale) daysOfTheWeekInCase(int) ([]string, bool) {
	return []string{"якшанбе", "душанбе", "сешанбе", "чоршанбе", "панҷшанбе", "ҷумъа", "шанбе"}, true
}

func (TgLocale) monthsOfTheYearInCase(form int) ([]string, bool) {
	switch form {
	case 1:
		return []string{"январи", "феврали", "марти", "апрели", "майи", "июни", "июли", "августи", "сентябри", "октябри", "ноябри", "декабри"}, true
	case 2:
		return []string{"январ", "феврал", "март", "апрел", "май", "июн", "июл", "август", "сентябр", "октябр", "ноябр", "декабр"}, true
	default:
		return []string{"январ", "феврал", "март", "апрел", "май", "июн", "июл", "август", "сентябр", "октябр", "ноябр", "декабр"}, true
	}
}

func (TgLocale) daysOfTheWeek() []string {
	return []string{"якшанбе", "душанбе", "сешанбе", "чоршанбе", "панҷшанбе", "ҷумъа", "шанбе"}
}

func (TgLocale) monthsOfTheYear() []string {
	return []string{"январ", "феврал", "март", "апрел", "май", "июн", "июл", "август", "сентябр", "октябр", "ноябр", "декабр"}
}
