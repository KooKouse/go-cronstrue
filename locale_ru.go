package cronstrue

import (
	"math"
	"strings"
)

func init() {
	registerLocale("ru", ruLocale{enLocale{}})
}

// ruLocale is the Russian locale, hand-ported from
// cronstrue/src/i18n/locales/ru.ts. Russian plural and case forms are
// selected by helper functions mirroring the upstream logic exactly.
type ruLocale struct {
	enLocale
}

// ruPhraseByNumber mirrors getPhraseByNumber: Slavic plural form selection
// (one / few / many) based on the number modulo 10 and 100.
func ruPhraseByNumber(s string, words [3]string) string {
	n, ok := jsNumberRU(s)
	if !ok || math.IsNaN(n) {
		return words[2]
	}
	ni := int(math.Abs(n))
	if ni%100 > 4 && ni%100 < 20 {
		return words[2]
	}
	switch d := ni % 10; {
	case d == 1:
		return words[0]
	case d >= 2 && d <= 4:
		return words[1]
	default: // 0, 5-9
		return words[2]
	}
}

// ruPhraseByDayOfWeek mirrors getPhraseByDayOfWeek: gender form selection
// based on the day of week number.
func ruPhraseByDayOfWeek(s string, words [3]string) string {
	n, ok := jsNumberRU(s)
	if !ok || math.IsNaN(n) {
		return words[2]
	}
	switch int(n) {
	case 0:
		return words[0]
	case 1, 2, 4:
		return words[1]
	default:
		return words[2]
	}
}

// jsNumberRU mimics Number(str): empty and blank strings convert to 0
// (unlike plain numbers, which yield NaN for empty input).
func jsNumberRU(s string) (float64, bool) {
	if strings.TrimSpace(s) == "" {
		return 0, true
	}
	return jsNumber(s)
}

// ruStartsWith2BeforeDash mirrors s.substring(0, s.indexOf("-")) == "2";
// a missing "-" yields the empty string (never "2").
func ruStartsWith2BeforeDash(s string) bool {
	if s == "" {
		return false
	}
	i := strings.Index(s, "-")
	if i < 0 {
		return false
	}
	return s[:i] == "2"
}

func (ruLocale) use24HourTimeFormatByDefault() bool { return true }

func (ruLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Произошла ошибка во время генерации описания выражения. Проверьте синтаксис крон-выражения."
}
func (ruLocale) everyMinute() string                     { return "каждую минуту" }
func (ruLocale) everyHour() string                       { return "каждый час" }
func (ruLocale) atSpace() string                         { return "В " }
func (ruLocale) everyMinuteBetweenX0AndX1() string       { return "Каждую минуту с %s по %s" }
func (ruLocale) at() string                              { return "В" }
func (ruLocale) spaceAnd() string                        { return " и" }
func (ruLocale) everySecond() string                     { return "каждую секунду" }
func (ruLocale) secondsX0ThroughX1PastTheMinute() string { return "секунды с %s по %s" }
func (ruLocale) minutesX0ThroughX1PastTheHour() string   { return "минуты с %s по %s" }
func (ruLocale) betweenX0AndX1() string                  { return "с %s по %s" }
func (ruLocale) atX0() string                            { return "в %s" }
func (ruLocale) commaEveryDay() string                   { return ", каждый день" }
func (ruLocale) commaEveryX0DaysOfTheWeek(s string) string {
	return ruPhraseByNumber(s, [3]string{", каждый %s день недели", ", каждые %s дня недели", ", каждые %s дней недели"})
}
func (ruLocale) commaX0ThroughX1(s string) string {
	if s != "" && (s[0] == '2' || s[0] == '3') {
		return ", со %s по %s"
	}
	return ", с %s по %s"
}
func (ruLocale) commaAndX0ThroughX1(s string) string {
	if s != "" && (s[0] == '2' || s[0] == '3') {
		return " и со %s по %s"
	}
	return " и с %s по %s"
}
func (ruLocale) first(s string) string {
	return ruPhraseByDayOfWeek(s, [3]string{"первое", "первый", "первую"})
}
func (ruLocale) second(s string) string {
	return ruPhraseByDayOfWeek(s, [3]string{"второе", "второй", "вторую"})
}
func (ruLocale) third(s string) string {
	return ruPhraseByDayOfWeek(s, [3]string{"третье", "третий", "третью"})
}
func (ruLocale) fourth(s string) string {
	return ruPhraseByDayOfWeek(s, [3]string{"четвертое", "четвертый", "четвертую"})
}
func (ruLocale) fifth(s string) string {
	return ruPhraseByDayOfWeek(s, [3]string{"пятое", "пятый", "пятую"})
}
func (ruLocale) commaOnThe(s, day string) string {
	if s == "2" {
		return ", во "
	}
	return ", в "
}
func (ruLocale) spaceX0OfTheMonth() string          { return " %s месяца" }
func (ruLocale) lastDay() string                    { return "последний день" }
func (ruLocale) commaOnlyInMonthX0() (string, bool) { return ", только %s", true }

func (ruLocale) commaOnTheLastX0OfTheMonth(s string) string {
	return ruPhraseByDayOfWeek(s, [3]string{", в последнее %s месяца", ", в последний %s месяца", ", в последнюю %s месяца"})
}

func (ruLocale) commaOnlyOnX0(s string) string {
	if s != "" && s[0] == '2' {
		return ", только во %s"
	}
	return ", только в %s"
}
func (ruLocale) commaAndOnX0() string { return ", и %s" }
func (ruLocale) commaEveryX0Months(s string) string {
	return ruPhraseByNumber(s, [3]string{"", " каждые %s месяца", " каждые %s месяцев"})
}
func (ruLocale) commaOnlyInX0() string { return ", только в %s" }

func (ruLocale) commaOnTheLastDayOfTheMonth() string {
	return ", в последний день месяца"
}
func (ruLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", в последний будний день месяца"
}
func (ruLocale) commaDaysBeforeTheLastDayOfTheMonth(s string) string {
	return ruPhraseByNumber(s, [3]string{
		", за %s день до конца месяца",
		", за %s дня до конца месяца",
		", за %s дней до конца месяца",
	})
}
func (ruLocale) firstWeekday() string { return "первый будний день" }
func (ruLocale) weekdayNearestDayX0() string {
	return "ближайший будний день к %s числу"
}
func (ruLocale) commaOnTheX0OfTheMonth() string {
	return ", в %s месяца"
}
func (ruLocale) commaEveryX0Days(s string) string {
	return ruPhraseByNumber(s, [3]string{", каждый %s день", ", каждые %s дня", ", каждые %s дней"})
}
func (ruLocale) commaBetweenDayX0AndX1OfTheMonth(s string) string {
	if ruStartsWith2BeforeDash(s) {
		return ", со %s по %s число месяца"
	}
	return ", с %s по %s число месяца"
}
func (ruLocale) commaOnDayX0OfTheMonth(s string) string {
	if s != "" && s[0] == '2' {
		return ", во %s число месяца"
	}
	return ", в %s число месяца"
}
func (ruLocale) commaEveryX0Years(s string) string {
	return ruPhraseByNumber(s, [3]string{", каждый %s год", ", каждые %s года", ", каждые %s лет"})
}
func (ruLocale) commaStartingX0() string { return ", начало %s" }

func (ruLocale) everyX0Seconds(s string) string {
	return ruPhraseByNumber(s, [3]string{"каждую %s секунду", "каждые %s секунды", "каждые %s секунд"})
}
func (ruLocale) everyX0Minutes(s string) string {
	return ruPhraseByNumber(s, [3]string{"каждую %s минуту", "каждые %s минуты", "каждые %s минут"})
}
func (ruLocale) everyX0Hours(s string) string {
	return ruPhraseByNumber(s, [3]string{"каждый %s час", "каждые %s часа", "каждые %s часов"})
}
func (ruLocale) atX0SecondsPastTheMinute(s string) string {
	return ruPhraseByNumber(s, [3]string{"в %s секунду", "в %s секунды", "в %s секунд"})
}
func (ruLocale) atX0MinutesPastTheHour(s string) string {
	return ruPhraseByNumber(s, [3]string{"в %s минуту", "в %s минуты", "в %s минут"})
}

func (ruLocale) daysOfTheWeek() []string {
	return []string{"воскресенье", "понедельник", "вторник", "среда", "четверг", "пятница", "суббота"}
}

func (ruLocale) daysOfTheWeekInCase(form int) ([]string, bool) {
	if form == 1 {
		return []string{"воскресенья", "понедельника", "вторника", "среды", "четверга", "пятницы", "субботы"}, true
	}
	return []string{"воскресенье", "понедельник", "вторник", "среду", "четверг", "пятницу", "субботу"}, true
}

func (ruLocale) monthsOfTheYear() []string {
	return []string{
		"январь", "февраль", "март", "апрель", "май", "июнь",
		"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь",
	}
}

func (ruLocale) monthsOfTheYearInCase(form int) ([]string, bool) {
	if form == 1 {
		return []string{
			"января", "февраля", "марта", "апреля", "мая", "июня",
			"июля", "августа", "сентября", "октября", "ноября", "декабря",
		}, true
	}
	return ruLocale{}.monthsOfTheYear(), true
}

func (ruLocale) atReboot() (string, bool) {
	return "Выполнить один раз, при запуске", true
}
func (ruLocale) onTheHour() (string, bool) {
	return "ровно в час", true
}
