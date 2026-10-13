package cronstrue

// Code generated from cronstrue/src/i18n/locales/af.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("af", AfLocale{enLocale{}})
}

// AfLocale is the af locale.
type AfLocale struct {
	enLocale
}

func (AfLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (AfLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Daar was 'n fout om die tydsuitdrukking the genereer. Raadpleeg asb die uitdrukking formaat."
}

func (AfLocale) everyMinute() string {
	return "elke minuut"
}

func (AfLocale) everyHour() string {
	return "elke uur"
}

func (AfLocale) atSpace() string {
	return "Teen "
}

func (AfLocale) everyMinuteBetweenX0AndX1() string {
	return "Elke minuut tussen %s en %s"
}

func (AfLocale) at() string {
	return "Teen"
}

func (AfLocale) spaceAnd() string {
	return " en"
}

func (AfLocale) everySecond() string {
	return "elke sekonde"
}

func (AfLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekonde %s deur na %s na die minuut"
}

func (AfLocale) minutesX0ThroughX1PastTheHour() string {
	return "minute %s deur na %s na die uur"
}

func (AfLocale) betweenX0AndX1() string {
	return "tussen %s en %s"
}

func (AfLocale) atX0() string {
	return "teen %s"
}

func (AfLocale) commaEveryDay() string {
	return ", elke dag"
}

func (AfLocale) spaceX0OfTheMonth() string {
	return " %s van die maand"
}

func (AfLocale) lastDay() string {
	return "die laaste dag"
}

func (AfLocale) commaAndOnX0() string {
	return ", en op %s"
}

func (AfLocale) commaOnlyInX0() string {
	return ", net in %s"
}

func (AfLocale) commaOnTheLastDayOfTheMonth() string {
	return ", op die laaste dag van die maand"
}

func (AfLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", op die laaste weeksdag van die maand"
}

func (AfLocale) firstWeekday() string {
	return "eerste weeksdag"
}

func (AfLocale) weekdayNearestDayX0() string {
	return "weeksdag naaste aan dag %s"
}

func (AfLocale) commaOnTheX0OfTheMonth() string {
	return ", op die %s van die maande"
}

func (AfLocale) commaStartingX0() string {
	return ", beginnende %s"
}

func (AfLocale) commaYearX0ThroughYearX1() (string, bool) {
	return ", jaar %s na %s", true
}

func (AfLocale) atReboot() (string, bool) {
	return "Loop een keer, met opstart", true
}

func (AfLocale) onTheHour() (string, bool) {
	return "op die uur", true
}

func (AfLocale) everyX0Seconds(string) string {
	return "elke %s sekonde"
}

func (AfLocale) everyX0Minutes(string) string {
	return "elke %s minute"
}

func (AfLocale) everyX0Hours(string) string {
	return "elke %s ure"
}

func (AfLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", elke %s dae van die week"
}

func (AfLocale) commaEveryX0Months(string) string {
	return ", elke %s maande"
}

func (AfLocale) commaEveryX0Days(string) string {
	return ", elke %s dae"
}

func (AfLocale) commaEveryX0Years(string) string {
	return ", elke %s jare"
}

func (AfLocale) atX0SecondsPastTheMinute(string) string {
	return "teen %s sekondes na die minuut"
}

func (AfLocale) atX0MinutesPastTheHour(string) string {
	return "teen %s minute na die uur"
}

func (AfLocale) commaX0ThroughX1(string) string {
	return ", %s deur na %s"
}

func (AfLocale) commaAndX0ThroughX1(string) string {
	return ", en %s deur na %s"
}

func (AfLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", op die laaste %s van die maand"
}

func (AfLocale) commaOnDayX0OfTheMonth(string) string {
	return ", op dag %s van die maand"
}

func (AfLocale) first(string) string {
	return "eerste"
}

func (AfLocale) second(string) string {
	return "tweede"
}

func (AfLocale) third(string) string {
	return "derde"
}

func (AfLocale) fourth(string) string {
	return "vierde"
}

func (AfLocale) fifth(string) string {
	return "vyfde"
}

func (AfLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", tussen dag %s en %s van die maand"
}

func (AfLocale) commaOnlyOnX0(string) string {
	return ", net op %s"
}

func (AfLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dae voor die laaste dag van die maand"
}

func (AfLocale) commaOnThe(string, string) string {
	return ", op die "
}

func (AfLocale) daysOfTheWeek() []string {
	return []string{"Sondag", "Maandag", "Dinsdag", "Woensdag", "Donderdag", "Vrydag", "Saterdag"}
}

func (AfLocale) monthsOfTheYear() []string {
	return []string{"Januarie", "Februarie", "Maart", "April", "Mei", "Junie", "Julie", "Augustus", "September", "Oktober", "November", "Desember"}
}
