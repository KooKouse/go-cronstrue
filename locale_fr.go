package cronstrue

// Code generated from cronstrue/src/i18n/locales/fr.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("fr", FrLocale{enLocale{}})
}

// FrLocale is the fr locale.
type FrLocale struct {
	enLocale
}

func (FrLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (FrLocale) conciseVerbosityReplacements() (map[string]string, bool) {
	return map[string]string{"de le": "du"}, true
}

func (FrLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Une erreur est survenue en générant la description de l'expression cron. Vérifiez sa syntaxe."
}

func (FrLocale) everyMinute() string {
	return "toutes les minutes"
}

func (FrLocale) everyHour() string {
	return "toutes les heures"
}

func (FrLocale) atSpace() string {
	return "À "
}

func (FrLocale) everyMinuteBetweenX0AndX1() string {
	return "Toutes les minutes entre %s et %s"
}

func (FrLocale) at() string {
	return "À"
}

func (FrLocale) spaceAnd() string {
	return " et"
}

func (FrLocale) everySecond() string {
	return "toutes les secondes"
}

func (FrLocale) secondsX0ThroughX1PastTheMinute() string {
	return "les secondes entre %s et %s après la minute"
}

func (FrLocale) minutesX0ThroughX1PastTheHour() string {
	return "les minutes entre %s et %s après l'heure"
}

func (FrLocale) betweenX0AndX1() string {
	return "de %s à %s"
}

func (FrLocale) atX0() string {
	return "%s"
}

func (FrLocale) commaEveryDay() string {
	return ", tous les jours"
}

func (FrLocale) spaceX0OfTheMonth() string {
	return " %s du mois"
}

func (FrLocale) lastDay() string {
	return "le dernier jour"
}

func (FrLocale) commaAndOnX0() string {
	return ", et %s"
}

func (FrLocale) commaOnlyInX0() string {
	return ", uniquement en %s"
}

func (FrLocale) commaOnTheLastDayOfTheMonth() string {
	return ", le dernier jour du mois"
}

func (FrLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", le dernier jour ouvrable du mois"
}

func (FrLocale) firstWeekday() string {
	return "premier jour ouvrable"
}

func (FrLocale) weekdayNearestDayX0() string {
	return "jour ouvrable le plus proche du %s"
}

func (FrLocale) commaOnTheX0OfTheMonth() string {
	return ", le %s du mois"
}

func (FrLocale) commaStartingX0() string {
	return ", à partir de %s"
}

func (FrLocale) atReboot() (string, bool) {
	return "Exécuter une fois, au démarrage", true
}

func (FrLocale) onTheHour() (string, bool) {
	return "à l'heure pile", true
}

func (FrLocale) everyX0Seconds(string) string {
	return "toutes les %s secondes"
}

func (FrLocale) everyX0Minutes(string) string {
	return "toutes les %s minutes"
}

func (FrLocale) everyX0Hours(string) string {
	return "toutes les %s heures"
}

func (FrLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", tous les %s jours de la semaine"
}

func (FrLocale) commaEveryX0Months(string) string {
	return ", tous les %s mois"
}

func (FrLocale) commaEveryX0Days(string) string {
	return ", tous les %s jours"
}

func (FrLocale) commaEveryX0Years(string) string {
	return ", tous les %s ans"
}

func (FrLocale) atX0SecondsPastTheMinute(string) string {
	return "%s secondes après la minute"
}

func (FrLocale) atX0MinutesPastTheHour(string) string {
	return "%s minutes après l'heure"
}

func (FrLocale) commaX0ThroughX1(string) string {
	return ", de %s à %s"
}

func (FrLocale) commaAndX0ThroughX1(string) string {
	return ", et de %s à %s"
}

func (FrLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", le dernier %s du mois"
}

func (FrLocale) commaOnDayX0OfTheMonth(string) string {
	return ", le %s du mois"
}

func (FrLocale) first(string) string {
	return "premier"
}

func (FrLocale) second(string) string {
	return "second"
}

func (FrLocale) third(string) string {
	return "troisième"
}

func (FrLocale) fourth(string) string {
	return "quatrième"
}

func (FrLocale) fifth(string) string {
	return "cinquième"
}

func (FrLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", du %s au %s du mois"
}

func (FrLocale) commaOnlyOnX0(string) string {
	return ", uniquement le %s"
}

func (FrLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s jours avant le dernier jour du mois"
}

func (FrLocale) commaOnThe(string, string) string {
	return ", le "
}

func (FrLocale) daysOfTheWeek() []string {
	return []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"}
}

func (FrLocale) monthsOfTheYear() []string {
	return []string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"}
}
