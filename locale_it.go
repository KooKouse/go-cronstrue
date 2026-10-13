package cronstrue

// Code generated from cronstrue/src/i18n/locales/it.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("it", ItLocale{enLocale{}})
}

// ItLocale is the it locale.
type ItLocale struct {
	enLocale
}

func (ItLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (ItLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "È verificato un errore durante la generazione la descrizione espressione. Controllare la sintassi delle espressioni cron."
}

func (ItLocale) everyMinute() string {
	return "ogni minuto"
}

func (ItLocale) everyHour() string {
	return "ogni ora"
}

func (ItLocale) atSpace() string {
	return "Alle "
}

func (ItLocale) everyMinuteBetweenX0AndX1() string {
	return "Ogni minuto tra le %s e le %s"
}

func (ItLocale) at() string {
	return "Alle"
}

func (ItLocale) spaceAnd() string {
	return " e"
}

func (ItLocale) everySecond() string {
	return "ogni secondo"
}

func (ItLocale) secondsX0ThroughX1PastTheMinute() string {
	return "secondi %s al %s oltre il minuto"
}

func (ItLocale) minutesX0ThroughX1PastTheHour() string {
	return "minuti %s al %s dopo l'ora"
}

func (ItLocale) betweenX0AndX1() string {
	return "tra le %s e le %s"
}

func (ItLocale) atX0() string {
	return "alle %s"
}

func (ItLocale) commaEveryDay() string {
	return ", ogni giorno"
}

func (ItLocale) spaceX0OfTheMonth() string {
	return " %s del mese"
}

func (ItLocale) lastDay() string {
	return "l'ultimo giorno"
}

func (ItLocale) commaAndOnX0() string {
	return ", e il %s"
}

func (ItLocale) commaOnlyInX0() string {
	return ", solo in %s"
}

func (ItLocale) commaOnTheLastDayOfTheMonth() string {
	return ", l'ultimo giorno del mese"
}

func (ItLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", nell'ultima settimana del mese"
}

func (ItLocale) firstWeekday() string {
	return "primo giorno della settimana"
}

func (ItLocale) weekdayNearestDayX0() string {
	return "giorno della settimana più vicino al %s"
}

func (ItLocale) commaOnTheX0OfTheMonth() string {
	return ", il %s del mese"
}

func (ItLocale) commaStartingX0() string {
	return ", a partire %s"
}

func (ItLocale) atReboot() (string, bool) {
	return "Esegui una volta, all'avvio", true
}

func (ItLocale) onTheHour() (string, bool) {
	return "all'ora esatta", true
}

func (ItLocale) everyX0Seconds(string) string {
	return "ogni %s secondi"
}

func (ItLocale) everyX0Minutes(string) string {
	return "ogni %s minuti"
}

func (ItLocale) everyX0Hours(string) string {
	return "ogni %s ore"
}

func (ItLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", ogni %s giorni della settimana"
}

func (ItLocale) commaEveryX0Months(string) string {
	return ", ogni %s mesi"
}

func (ItLocale) commaEveryX0Days(string) string {
	return ", ogni %s giorni"
}

func (ItLocale) commaEveryX0Years(string) string {
	return ", ogni %s anni"
}

func (ItLocale) atX0SecondsPastTheMinute(string) string {
	return "al %s secondo passato il minuto"
}

func (ItLocale) atX0MinutesPastTheHour(string) string {
	return "al %s minuto passata l'ora"
}

func (ItLocale) commaX0ThroughX1(string) string {
	return ", %s al %s"
}

func (ItLocale) commaAndX0ThroughX1(string) string {
	return ", e %s al %s"
}

func (ItLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", l'ultimo %s del mese"
}

func (ItLocale) commaOnDayX0OfTheMonth(string) string {
	return ", il giorno %s del mese"
}

func (ItLocale) first(string) string {
	return "primo"
}

func (ItLocale) second(string) string {
	return "secondo"
}

func (ItLocale) third(string) string {
	return "terzo"
}

func (ItLocale) fourth(string) string {
	return "quarto"
}

func (ItLocale) fifth(string) string {
	return "quinto"
}

func (ItLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", tra il giorno %s e %s del mese"
}

func (ItLocale) commaOnlyOnX0(string) string {
	return ", solo il %s"
}

func (ItLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s giorni prima dell'ultimo giorno del mese"
}

func (ItLocale) commaOnThe(string, string) string {
	return ", il "
}

func (ItLocale) daysOfTheWeek() []string {
	return []string{"domenica", "lunedì", "martedì", "mercoledì", "giovedì", "venerdì", "sabato"}
}

func (ItLocale) monthsOfTheYear() []string {
	return []string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"}
}
