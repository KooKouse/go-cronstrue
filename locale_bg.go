package cronstrue

import (
	"math"
	"strings"
)

func init() {
	registerLocale("bg", bgLocale{enLocale{}})
}

// bgLocale is the Bulgarian locale, hand-ported from
// cronstrue/src/i18n/locales/bg.ts.
type bgLocale struct {
	enLocale
}

// bgPhraseByPlural mirrors getPhraseByPlural: number < 2 -> singular.
func bgPhraseByPlural(s string, words [2]string) string {
	n := bgNumber(s)
	if n < 2 {
		return words[0]
	}
	return words[1]
}

// bgPhraseByDayOfWeek mirrors getPhraseByDayOfWeek: gender form by weekday
// (0:неделя 1:понеделник 2:вторник 3:сряда 4:четвъртък 5:петък 6:събота;
// gender: ж м м ж м м ж).
func bgPhraseByDayOfWeek(s string, words [2]string) string {
	n, ok := bgNumberChecked(s)
	if !ok || math.IsNaN(n) {
		n = 0
	}
	idx := []int{1, 0, 0, 1, 0, 0, 1}
	i := int(n)
	if i < 0 || i >= len(idx) {
		return "undefined"
	}
	return words[idx[i]]
}

// bgNumberEnding mirrors getNumberEnding: ordinal ending (в/р/т + и/а/о)
// selected by the last significant digit and the grammatical gender.
// gender: 0=м, 1=ж, 2=ср.
func bgNumberEnding(s string, gender int) string {
	var n float64
	if s == "" {
		n = 1 // str == null / undefined default
	} else if v, ok := jsNumber(s); ok {
		n = v
	} else {
		n = math.NaN() // e.g. list expressions like "1,15"
	}
	last := 3.0
	if n < 10 || (n > 20 && math.Mod(n, 10) != 0) {
		last = math.Mod(n, 10)
	}
	last = math.Max(math.Min(last, 3), 1) - 1
	endings := []string{"в", "р", "т"}
	genders := []string{"и", "а", "о"}
	return endings[int(last)] + genders[gender]
}

// bgNumber mimics `str != null ? Number(str) : 0`; the empty string
// converts to 0 like JavaScript Number("").
func bgNumber(s string) float64 {
	n, _ := bgNumberChecked(s)
	return n
}

func bgNumberChecked(s string) (float64, bool) {
	if strings.TrimSpace(s) == "" {
		return 0, true
	}
	return jsNumber(s)
}

func (bgLocale) use24HourTimeFormatByDefault() bool { return true }

func (bgLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Възникна грешка при генериране на описанието на израза. Проверете синтаксиса на cron израза."
}
func (bgLocale) everyMinute() string                     { return "всяка минута" }
func (bgLocale) everyHour() string                       { return "всеки час" }
func (bgLocale) atSpace() string                         { return "В " }
func (bgLocale) everyMinuteBetweenX0AndX1() string       { return "Всяка минута от %s до %s" }
func (bgLocale) at() string                              { return "В" }
func (bgLocale) spaceAnd() string                        { return " и" }
func (bgLocale) everySecond() string                     { return "всяка секунда" }
func (bgLocale) everyX0Seconds(string) string            { return "всеки %s секунди" }
func (bgLocale) secondsX0ThroughX1PastTheMinute() string { return "секунди от %s до %s" }
func (bgLocale) atX0SecondsPastTheMinute(s string) string {
	return "%s-" + bgNumberEnding(s, 1) + " секунда"
}
func (bgLocale) everyX0Minutes(string) string          { return "всеки %s минути" }
func (bgLocale) minutesX0ThroughX1PastTheHour() string { return "минути от %s до %s" }
func (bgLocale) atX0MinutesPastTheHour(s string) string {
	return "%s-" + bgNumberEnding(s, 1) + " минутa"
}
func (bgLocale) everyX0Hours(string) string { return "всеки %s часа" }
func (bgLocale) betweenX0AndX1() string     { return "от %s до %s" }
func (bgLocale) atX0() string               { return "в %s" }
func (bgLocale) commaEveryDay() string      { return ", всеки ден" }
func (bgLocale) commaEveryX0DaysOfTheWeek(s string) string {
	return bgPhraseByPlural(s, [2]string{", всеки %s ден от седмицата", ", всеки %s дена от седмицата"})
}
func (bgLocale) commaX0ThroughX1(string) string    { return ", от %s до %s" }
func (bgLocale) commaAndX0ThroughX1(string) string { return " и от %s до %s" }
func (bgLocale) first(s string) string {
	return bgPhraseByDayOfWeek(s, [2]string{"первият", "первата"})
}
func (bgLocale) second(s string) string {
	return bgPhraseByDayOfWeek(s, [2]string{"вторият", "втората"})
}
func (bgLocale) third(s string) string {
	return bgPhraseByDayOfWeek(s, [2]string{"третият", "третата"})
}
func (bgLocale) fourth(s string) string {
	return bgPhraseByDayOfWeek(s, [2]string{"четвертият", "четвертата"})
}
func (bgLocale) fifth(s string) string {
	return bgPhraseByDayOfWeek(s, [2]string{"петият", "петата"})
}
func (bgLocale) commaOnThe(string, string) string { return ", " }
func (bgLocale) spaceX0OfTheMonth() string        { return " %s на месеца" }
func (bgLocale) lastDay() string                  { return "последният ден" }

func (bgLocale) commaOnTheLastX0OfTheMonth(s string) string {
	return bgPhraseByDayOfWeek(s, [2]string{", в последният %s от месеца", ", в последната %s отмесеца"})
}

func (bgLocale) commaOnlyOnX0(string) string        { return ", %s" }
func (bgLocale) commaAndOnX0() string               { return " и %s" }
func (bgLocale) commaEveryX0Months(string) string   { return " всеки %s месеца" }
func (bgLocale) commaOnlyInMonthX0() (string, bool) { return ", %s", true }
func (bgLocale) commaOnlyInX0() string              { return ", в %s" }

func (bgLocale) commaOnTheLastDayOfTheMonth() string {
	return ", в последният ден на месеца"
}
func (bgLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", в последния делничен ден от месеца"
}
func (bgLocale) commaDaysBeforeTheLastDayOfTheMonth(s string) string {
	return bgPhraseByPlural(s, [2]string{", %s ден преди края на месеца", ", %s дена преди края на месеца"})
}
func (bgLocale) firstWeekday() string { return "първият делничен ден" }
func (bgLocale) weekdayNearestDayX0() string {
	return "най-близкият делничен ден до %s число"
}
func (bgLocale) commaOnTheX0OfTheMonth() string {
	return ", на %s число от месеца"
}
func (bgLocale) commaEveryX0Days(s string) string {
	return bgPhraseByPlural(s, [2]string{", всеки %s ден", ", всеки %s дена"})
}
func (bgLocale) commaBetweenDayX0AndX1OfTheMonth(s string) string {
	parts := strings.Split(s, "-")
	b := ""
	if len(parts) == 2 {
		b = parts[1]
	}
	return ", от %s-" + bgNumberEnding(parts[0], 2) + " до %s-" + bgNumberEnding(b, 2) + " число на месеца"
}
func (bgLocale) commaOnDayX0OfTheMonth(s string) string {
	return ", на %s-" + bgNumberEnding(s, 2) + " число от месеца"
}
func (bgLocale) commaEveryX0Years(s string) string {
	return bgPhraseByPlural(s, [2]string{", всяка %s година", ", всеки %s години"})
}
func (bgLocale) commaStartingX0() string { return ", започвайки %s" }

func (bgLocale) daysOfTheWeek() []string {
	return []string{"неделя", "понеделник", "вторник", "сряда", "четвъртък", "петък", "събота"}
}

func (bgLocale) monthsOfTheYear() []string {
	return []string{
		"януари", "февруари", "март", "април", "май", "юни",
		"юли", "август", "септевмври", "октомври", "ноември", "декември",
	}
}

func (bgLocale) atReboot() (string, bool) {
	return "Изпълнение веднъж, при стартиране", true
}
func (bgLocale) onTheHour() (string, bool) {
	return "в началото на часа", true
}
