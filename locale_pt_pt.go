package cronstrue

// Code generated from cronstrue/src/i18n/locales/pt_PT.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("pt-PT", PtPTLocale{enLocale{}})
}

// PtPTLocale is the pt_PT locale.
type PtPTLocale struct {
	enLocale
}

func (PtPTLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (PtPTLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Ocorreu um erro ao gerar a descrição da expressão Cron."
}

func (PtPTLocale) everyMinute() string {
	return "a cada minuto"
}

func (PtPTLocale) everyHour() string {
	return "a cada hora"
}

func (PtPTLocale) atSpace() string {
	return "às "
}

func (PtPTLocale) everyMinuteBetweenX0AndX1() string {
	return "a cada minuto entre %s e %s"
}

func (PtPTLocale) at() string {
	return "às"
}

func (PtPTLocale) spaceAnd() string {
	return " e"
}

func (PtPTLocale) everySecond() string {
	return "a cada segundo"
}

func (PtPTLocale) secondsX0ThroughX1PastTheMinute() string {
	return "No segundo %s até %s de cada minuto"
}

func (PtPTLocale) minutesX0ThroughX1PastTheHour() string {
	return "do minuto %s até %s de cada hora"
}

func (PtPTLocale) betweenX0AndX1() string {
	return "entre %s e %s"
}

func (PtPTLocale) atX0() string {
	return "Às %s"
}

func (PtPTLocale) commaEveryDay() string {
	return ", a cada dia"
}

func (PtPTLocale) spaceX0OfTheMonth() string {
	return " %s do mês"
}

func (PtPTLocale) lastDay() string {
	return "o último dia"
}

func (PtPTLocale) commaAndOnX0() string {
	return ", e de %s"
}

func (PtPTLocale) commaOnlyInX0() string {
	return ", somente em %s"
}

func (PtPTLocale) commaOnTheLastDayOfTheMonth() string {
	return ", no último dia do mês"
}

func (PtPTLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", no último dia da semana do mês"
}

func (PtPTLocale) firstWeekday() string {
	return "primeiro dia da semana"
}

func (PtPTLocale) weekdayNearestDayX0() string {
	return "dia da semana mais próximo do dia %s"
}

func (PtPTLocale) commaOnTheX0OfTheMonth() string {
	return ", no %s do mês"
}

func (PtPTLocale) commaStartingX0() string {
	return ", iniciando %s"
}

func (PtPTLocale) atReboot() (string, bool) {
	return "Executar uma vez, no arranque", true
}

func (PtPTLocale) onTheHour() (string, bool) {
	return "à hora certa", true
}

func (PtPTLocale) everyX0Seconds(string) string {
	return "a cada %s segundos"
}

func (PtPTLocale) everyX0Minutes(string) string {
	return "a cada %s minutos"
}

func (PtPTLocale) everyX0Hours(string) string {
	return "a cada %s horas"
}

func (PtPTLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", a cada %s dias de semana"
}

func (PtPTLocale) commaEveryX0Months(string) string {
	return ", a cada %s meses"
}

func (PtPTLocale) commaEveryX0Days(string) string {
	return ", a cada %s dias"
}

func (PtPTLocale) commaEveryX0Years(string) string {
	return ", a cada %s anos"
}

func (PtPTLocale) atX0SecondsPastTheMinute(string) string {
	return "aos %s segundos do minuto"
}

func (PtPTLocale) atX0MinutesPastTheHour(string) string {
	return "aos %s minutos da hora"
}

func (PtPTLocale) commaX0ThroughX1(string) string {
	return ", de %s a %s"
}

func (PtPTLocale) commaAndX0ThroughX1(string) string {
	return ", e de %s a %s"
}

func (PtPTLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", na última %s do mês"
}

func (PtPTLocale) commaOnDayX0OfTheMonth(string) string {
	return ", no dia %s do mês"
}

func (PtPTLocale) first(s string) string {
	switch s {
	case "0", "6":
		return "primeiro"
	default:
		return "primeira"
	}
}

func (PtPTLocale) second(s string) string {
	switch s {
	case "0", "6":
		return "segundo"
	default:
		return "segunda"
	}
}

func (PtPTLocale) third(s string) string {
	switch s {
	case "0", "6":
		return "terceiro"
	default:
		return "terceira"
	}
}

func (PtPTLocale) fourth(s string) string {
	switch s {
	case "0", "6":
		return "quarto"
	default:
		return "quarta"
	}
}

func (PtPTLocale) fifth(s string) string {
	switch s {
	case "0", "6":
		return "quinto"
	default:
		return "quinta"
	}
}

func (PtPTLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", entre os dias %s e %s do mês"
}

func (PtPTLocale) commaOnlyOnX0(string) string {
	return ", somente de %s"
}

func (PtPTLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dias antes do último dia do mês"
}

func (PtPTLocale) commaOnThe(s, day string) string {
	switch day {
	case "0", "6":
		return ", no"
	default:
		return ", na "
	}
}

func (PtPTLocale) daysOfTheWeek() []string {
	return []string{"domingo", "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado"}
}

func (PtPTLocale) monthsOfTheYear() []string {
	return []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
}
