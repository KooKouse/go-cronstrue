package cronstrue

// Code generated from cronstrue/src/i18n/locales/sk.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("sk", SkLocale{enLocale{}})
}

// SkLocale is the sk locale.
type SkLocale struct {
	enLocale
}

func (SkLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (SkLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Pri vytváraní popisu došlo k chybe. Skontrolujte prosím správnosť syntaxe cronu."
}

func (SkLocale) everyMinute() string {
	return "každú minútu"
}

func (SkLocale) everyHour() string {
	return "každú hodinu"
}

func (SkLocale) atSpace() string {
	return "V "
}

func (SkLocale) everyMinuteBetweenX0AndX1() string {
	return "Každú minútu medzi %s a %s"
}

func (SkLocale) at() string {
	return "V"
}

func (SkLocale) spaceAnd() string {
	return " a"
}

func (SkLocale) everySecond() string {
	return "každú sekundu"
}

func (SkLocale) secondsX0ThroughX1PastTheMinute() string {
	return "sekundy od %s do %s"
}

func (SkLocale) minutesX0ThroughX1PastTheHour() string {
	return "minúty od %s do %s"
}

func (SkLocale) betweenX0AndX1() string {
	return "medzi %s a %s"
}

func (SkLocale) atX0() string {
	return "v %s"
}

func (SkLocale) commaEveryDay() string {
	return ", každý deň"
}

func (SkLocale) spaceX0OfTheMonth() string {
	return " %s v mesiaci"
}

func (SkLocale) lastDay() string {
	return "posledný deň"
}

func (SkLocale) commaAndOnX0() string {
	return ", a v %s"
}

func (SkLocale) commaOnlyInX0() string {
	return ", iba v %s"
}

func (SkLocale) commaOnTheLastDayOfTheMonth() string {
	return ", posledný deň v mesiaci"
}

func (SkLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", posledný pracovný deň v mesiaci"
}

func (SkLocale) firstWeekday() string {
	return "prvý pracovný deň"
}

func (SkLocale) weekdayNearestDayX0() string {
	return "pracovný deň najbližšie %s. dňu"
}

func (SkLocale) commaOnTheX0OfTheMonth() string {
	return ", v %s v mesiaci"
}

func (SkLocale) commaStartingX0() string {
	return ", začínajúcich %s"
}

func (SkLocale) atReboot() (string, bool) {
	return "Spustiť raz, pri štarte", true
}

func (SkLocale) onTheHour() (string, bool) {
	return "o celú hodinu", true
}

func (SkLocale) everyX0Seconds(string) string {
	return "každých %s sekúnd"
}

func (SkLocale) everyX0Minutes(string) string {
	return "každých %s minút"
}

func (SkLocale) everyX0Hours(string) string {
	return "každých %s hodín"
}

func (SkLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", každých %s dní v týždni"
}

func (SkLocale) commaEveryX0Months(string) string {
	return ", každých %s mesiacov"
}

func (SkLocale) commaEveryX0Days(string) string {
	return ", každých %s dní"
}

func (SkLocale) commaEveryX0Years(string) string {
	return ", každých %s rokov"
}

func (SkLocale) atX0SecondsPastTheMinute(string) string {
	return "v %s sekúnd"
}

func (SkLocale) atX0MinutesPastTheHour(string) string {
	return "v %s minút"
}

func (SkLocale) commaX0ThroughX1(string) string {
	return ", od %s do %s"
}

func (SkLocale) commaAndX0ThroughX1(string) string {
	return ", a od %s do %s"
}

func (SkLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", posledný %s v mesiaci"
}

func (SkLocale) commaOnDayX0OfTheMonth(string) string {
	return ", %s. deň v mesiaci"
}

func (SkLocale) first(string) string {
	return "prvý"
}

func (SkLocale) second(string) string {
	return "druhý"
}

func (SkLocale) third(string) string {
	return "tretí"
}

func (SkLocale) fourth(string) string {
	return "štvrtý"
}

func (SkLocale) fifth(string) string {
	return "piaty"
}

func (SkLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", medzi dňami %s a %s v mesiaci"
}

func (SkLocale) commaOnlyOnX0(string) string {
	return ", iba v %s"
}

func (SkLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dní pred posledným dňom v mesiaci"
}

func (SkLocale) commaOnThe(string, string) string {
	return ", "
}

func (SkLocale) daysOfTheWeek() []string {
	return []string{"Nedeľa", "Pondelok", "Utorok", "Streda", "Štvrtok", "Piatok", "Sobota"}
}

func (SkLocale) monthsOfTheYear() []string {
	return []string{"Január", "Február", "Marec", "Apríl", "Máj", "Jún", "Júl", "August", "September", "Október", "November", "December"}
}
