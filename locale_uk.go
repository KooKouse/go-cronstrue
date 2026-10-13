package cronstrue

// Code generated from cronstrue/src/i18n/locales/uk.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("uk", UkLocale{enLocale{}})
}

// UkLocale is the uk locale.
type UkLocale struct {
	enLocale
}

func (UkLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (UkLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "ВІдбулася помилка підчас генерації опису. Перевірта правильність написання cron виразу."
}

func (UkLocale) everyMinute() string {
	return "щохвилини"
}

func (UkLocale) everyHour() string {
	return "щогодини"
}

func (UkLocale) atSpace() string {
	return "О "
}

func (UkLocale) everyMinuteBetweenX0AndX1() string {
	return "Щохвилини між %s та %s"
}

func (UkLocale) at() string {
	return "О"
}

func (UkLocale) spaceAnd() string {
	return " та"
}

func (UkLocale) everySecond() string {
	return "Щосекунди"
}

func (UkLocale) secondsX0ThroughX1PastTheMinute() string {
	return "з %s по %s секунду"
}

func (UkLocale) minutesX0ThroughX1PastTheHour() string {
	return "з %s по %s хвилину"
}

func (UkLocale) betweenX0AndX1() string {
	return "між %s та %s"
}

func (UkLocale) atX0() string {
	return "о %s"
}

func (UkLocale) commaEveryDay() string {
	return ", щоденно"
}

func (UkLocale) spaceX0OfTheMonth() string {
	return " %s місяця"
}

func (UkLocale) lastDay() string {
	return "останній день"
}

func (UkLocale) commaAndOnX0() string {
	return ", і в %s"
}

func (UkLocale) commaOnlyInX0() string {
	return ", тільки в %s"
}

func (UkLocale) commaOnTheLastDayOfTheMonth() string {
	return ", в останній день місяця"
}

func (UkLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", в останній будень місяця"
}

func (UkLocale) firstWeekday() string {
	return "перший будень"
}

func (UkLocale) weekdayNearestDayX0() string {
	return "будень найближчий до %s дня"
}

func (UkLocale) commaOnTheX0OfTheMonth() string {
	return ", в %s місяця"
}

func (UkLocale) commaStartingX0() string {
	return ", початок %s"
}

func (UkLocale) atReboot() (string, bool) {
	return "Виконати один раз, при запуску", true
}

func (UkLocale) onTheHour() (string, bool) {
	return "рівно о годині", true
}

func (UkLocale) everyX0Seconds(string) string {
	return "кожні %s секунд"
}

func (UkLocale) everyX0Minutes(string) string {
	return "кожні %s хвилин"
}

func (UkLocale) everyX0Hours(string) string {
	return "кожні %s годин"
}

func (UkLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", кожен %s день тижня"
}

func (UkLocale) commaEveryX0Months(string) string {
	return ", кожен %s місяць"
}

func (UkLocale) commaEveryX0Days(string) string {
	return ", кожен %s день"
}

func (UkLocale) commaEveryX0Years(string) string {
	return ", кожні %s роки"
}

func (UkLocale) atX0SecondsPastTheMinute(string) string {
	return "о %s секунді"
}

func (UkLocale) atX0MinutesPastTheHour(string) string {
	return "о %s хвилині"
}

func (UkLocale) commaX0ThroughX1(string) string {
	return ", %s по %s"
}

func (UkLocale) commaAndX0ThroughX1(string) string {
	return ", та %s по %s"
}

func (UkLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", в останній %s місяця"
}

func (UkLocale) commaOnDayX0OfTheMonth(string) string {
	return ", на %s день місяця"
}

func (UkLocale) first(string) string {
	return "перший"
}

func (UkLocale) second(string) string {
	return "другий"
}

func (UkLocale) third(string) string {
	return "третій"
}

func (UkLocale) fourth(string) string {
	return "четвертий"
}

func (UkLocale) fifth(string) string {
	return "п'ятий"
}

func (UkLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", між %s та %s днями місяця"
}

func (UkLocale) commaOnlyOnX0(string) string {
	return ", тільки в %s"
}

func (UkLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s днів до останнього дня місяця"
}

func (UkLocale) commaOnThe(string, string) string {
	return ", в "
}

func (UkLocale) daysOfTheWeek() []string {
	return []string{"неділя", "понеділок", "вівторок", "середа", "четвер", "п'ятниця", "субота"}
}

func (UkLocale) monthsOfTheYear() []string {
	return []string{"січень", "лютий", "березень", "квітень", "травень", "червень", "липень", "серпень", "вересень", "жовтень", "листопад", "грудень"}
}
