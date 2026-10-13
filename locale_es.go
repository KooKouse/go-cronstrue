package cronstrue

// Code generated from cronstrue/src/i18n/locales/es.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("es", EsLocale{enLocale{}})
}

// EsLocale is the es locale.
type EsLocale struct {
	enLocale
}

func (EsLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (EsLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Ocurrió un error mientras se generaba la descripción de la expresión. Revise la sintaxis de la expresión de cron."
}

func (EsLocale) everyMinute() string {
	return "cada minuto"
}

func (EsLocale) everyHour() string {
	return "cada hora"
}

func (EsLocale) atSpace() string {
	return "A las "
}

func (EsLocale) everyMinuteBetweenX0AndX1() string {
	return "cada minuto entre las %s y las %s"
}

func (EsLocale) at() string {
	return "A las"
}

func (EsLocale) spaceAnd() string {
	return " y"
}

func (EsLocale) everySecond() string {
	return "cada segundo"
}

func (EsLocale) secondsX0ThroughX1PastTheMinute() string {
	return "En los segundos %s al %s de cada minuto"
}

func (EsLocale) minutesX0ThroughX1PastTheHour() string {
	return "del minuto %s al %s pasada la hora"
}

func (EsLocale) betweenX0AndX1() string {
	return "entre las %s y las %s"
}

func (EsLocale) atX0() string {
	return "a las %s"
}

func (EsLocale) commaEveryDay() string {
	return ", cada día"
}

func (EsLocale) spaceX0OfTheMonth() string {
	return " %s del mes"
}

func (EsLocale) lastDay() string {
	return "el último día"
}

func (EsLocale) commaAndOnX0() string {
	return ", y el %s"
}

func (EsLocale) commaOnlyInX0() string {
	return ", sólo en %s"
}

func (EsLocale) commaOnTheLastDayOfTheMonth() string {
	return ", en el último día del mes"
}

func (EsLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", en el último día de la semana del mes"
}

func (EsLocale) firstWeekday() string {
	return "primer día de la semana"
}

func (EsLocale) weekdayNearestDayX0() string {
	return "día de la semana más próximo al %s"
}

func (EsLocale) commaOnTheX0OfTheMonth() string {
	return ", en el %s del mes"
}

func (EsLocale) commaStartingX0() string {
	return ", comenzando %s"
}

func (EsLocale) atReboot() (string, bool) {
	return "Ejecutar una vez, al iniciar", true
}

func (EsLocale) onTheHour() (string, bool) {
	return "en punto", true
}

func (EsLocale) everyX0Seconds(string) string {
	return "cada %s segundos"
}

func (EsLocale) everyX0Minutes(string) string {
	return "cada %s minutos"
}

func (EsLocale) everyX0Hours(string) string {
	return "cada %s horas"
}

func (EsLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", cada %s días de la semana"
}

func (EsLocale) commaEveryX0Months(string) string {
	return ", cada %s meses"
}

func (EsLocale) commaEveryX0Days(string) string {
	return ", cada %s días"
}

func (EsLocale) commaEveryX0Years(string) string {
	return ", cada %s años"
}

func (EsLocale) atX0SecondsPastTheMinute(string) string {
	return "a los %s segundos del minuto"
}

func (EsLocale) atX0MinutesPastTheHour(string) string {
	return "a los %s minutos de la hora"
}

func (EsLocale) commaX0ThroughX1(string) string {
	return ", de %s a %s"
}

func (EsLocale) commaAndX0ThroughX1(string) string {
	return ", y de %s a %s"
}

func (EsLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", en el último %s del mes"
}

func (EsLocale) commaOnDayX0OfTheMonth(string) string {
	return ", el día %s del mes"
}

func (EsLocale) first(string) string {
	return "primero"
}

func (EsLocale) second(string) string {
	return "segundo"
}

func (EsLocale) third(string) string {
	return "tercer"
}

func (EsLocale) fourth(string) string {
	return "cuarto"
}

func (EsLocale) fifth(string) string {
	return "quinto"
}

func (EsLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", entre los días %s y %s del mes"
}

func (EsLocale) commaOnlyOnX0(string) string {
	return ", sólo el %s"
}

func (EsLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s días antes del último día del mes"
}

func (EsLocale) commaOnThe(string, string) string {
	return ", en el "
}

func (EsLocale) daysOfTheWeek() []string {
	return []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"}
}

func (EsLocale) monthsOfTheYear() []string {
	return []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}
}
