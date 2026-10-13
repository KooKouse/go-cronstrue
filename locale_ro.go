package cronstrue

import "math"

func init() {
	registerLocale("ro", roLocale{enLocale{}})
}

// roLocale is the Romanian locale, hand-ported from
// cronstrue/src/i18n/locales/ro.ts. Romanian inserts "de" between a
// number and its noun when the number is 20 or more and ends in 00 or
// 20 to 99: "5 minute", "19 minute", "20 de minute".
type roLocale struct {
	enLocale
}

func roPhraseByNumber(s string, words [2]string) string {
	n, ok := jsNumber(s)
	if !ok || math.IsNaN(n) {
		return words[0]
	}
	if n >= 20 && (math.Mod(n, 100) == 0 || math.Mod(n, 100) >= 20) {
		return words[1]
	}
	return words[0]
}

func (roLocale) use24HourTimeFormatByDefault() bool { return true }

func (roLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Eroare la generarea descrierii. Verificați sintaxa."
}
func (roLocale) at() string      { return "La" }
func (roLocale) atSpace() string { return "La " }
func (roLocale) atX0() string    { return "la %s" }
func (roLocale) commaMonthX0ThroughMonthX1() (string, bool) {
	return ", din %s până în %s", true
}
func (roLocale) commaYearX0ThroughYearX1() (string, bool) {
	return ", din %s până în %s", true
}
func (roLocale) atX0MinutesPastTheHourGt20() (string, bool) {
	return "la și %s de minute", true
}
func (roLocale) atX0SecondsPastTheMinuteGt20() (string, bool) {
	return "la și %s de secunde", true
}
func (roLocale) everyMinute() string               { return "în fiecare minut" }
func (roLocale) everyHour() string                 { return "în fiecare oră" }
func (roLocale) everyMinuteBetweenX0AndX1() string { return "În fiecare minut între %s și %s" }
func (roLocale) everySecond() string               { return "în fiecare secundă" }
func (roLocale) everyX0Seconds(s string) string {
	return roPhraseByNumber(s, [2]string{"la fiecare %s secunde", "la fiecare %s de secunde"})
}
func (roLocale) secondsX0ThroughX1PastTheMinute() string { return "între secunda %s și secunda %s" }
func (roLocale) atX0SecondsPastTheMinute(string) string  { return "la și %s secunde" }
func (roLocale) everyX0Minutes(s string) string {
	return roPhraseByNumber(s, [2]string{"la fiecare %s minute", "la fiecare %s de minute"})
}
func (roLocale) minutesX0ThroughX1PastTheHour() string { return "între minutele %s și %s" }
func (roLocale) atX0MinutesPastTheHour(string) string  { return "la și %s minute" }
func (roLocale) everyX0Hours(s string) string {
	return roPhraseByNumber(s, [2]string{"la fiecare %s ore", "la fiecare %s de ore"})
}
func (roLocale) betweenX0AndX1() string { return "între %s și %s" }
func (roLocale) commaEveryDay() string  { return ", în fiecare zi" }
func (roLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", la fiecare a %s-a zi a săptămânii"
}
func (roLocale) commaX0ThroughX1(string) string    { return ", de %s până %s" }
func (roLocale) commaAndX0ThroughX1(string) string { return ", și de %s până %s" }
func (roLocale) first(string) string               { return "prima" }
func (roLocale) second(string) string              { return "a doua" }
func (roLocale) third(string) string               { return "a treia" }
func (roLocale) fourth(string) string              { return "a patra" }
func (roLocale) fifth(string) string               { return "a cincea" }
func (roLocale) commaOnThe(string, string) string  { return ", în " }
func (roLocale) spaceAnd() string                  { return " și" }
func (roLocale) spaceX0OfTheMonth() string         { return " %s a lunii" }
func (roLocale) lastDay() string                   { return "ultima zi" }

func (roLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", în ultima %s a lunii"
}
func (roLocale) commaOnlyOnX0(string) string { return ", doar %s" }
func (roLocale) commaAndOnX0() string        { return ", și %s" }
func (roLocale) commaEveryX0Months(s string) string {
	return roPhraseByNumber(s, [2]string{", la fiecare %s luni", ", la fiecare %s de luni"})
}
func (roLocale) commaOnlyInX0() string { return ", doar în %s" }

func (roLocale) commaOnTheLastDayOfTheMonth() string {
	return ", în ultima zi a lunii"
}
func (roLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", în ultima zi lucrătoare a lunii"
}
func (roLocale) commaDaysBeforeTheLastDayOfTheMonth(s string) string {
	return roPhraseByNumber(s, [2]string{
		", %s zile înainte de ultima zi a lunii",
		", %s de zile înainte de ultima zi a lunii",
	})
}
func (roLocale) firstWeekday() string        { return "prima zi lucrătoare" }
func (roLocale) weekdayNearestDayX0() string { return "cea mai apropiată zi lucrătoare de ziua %s" }
func (roLocale) commaOnTheX0OfTheMonth() string {
	return ", în %s a lunii"
}
func (roLocale) commaEveryX0Days(s string) string {
	return roPhraseByNumber(s, [2]string{", la fiecare %s zile", ", la fiecare %s de zile"})
}
func (roLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", între zilele %s și %s ale lunii"
}
func (roLocale) commaOnDayX0OfTheMonth(string) string { return ", în ziua %s a lunii" }
func (roLocale) commaEveryX0Years(s string) string {
	return roPhraseByNumber(s, [2]string{", o dată la %s ani", ", o dată la %s de ani"})
}
func (roLocale) commaStartingX0() string { return ", pornire %s" }

func (roLocale) daysOfTheWeek() []string {
	return []string{"duminică", "luni", "marți", "miercuri", "joi", "vineri", "sâmbătă"}
}

func (roLocale) monthsOfTheYear() []string {
	return []string{
		"ianuarie", "februarie", "martie", "aprilie", "mai", "iunie",
		"iulie", "august", "septembrie", "octombrie", "noiembrie", "decembrie",
	}
}

func (roLocale) atReboot() (string, bool) { return "Rulează o dată, la pornire", true }
func (roLocale) onTheHour() (string, bool) {
	return "fix la oră", true
}
