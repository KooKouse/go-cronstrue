package cronstrue

// Code generated from cronstrue/src/i18n/locales/pt_BR.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("pt_BR", PtBRLocale{enLocale{}})
}

// PtBRLocale is the pt_BR locale.
type PtBRLocale struct {
	enLocale
}

func (PtBRLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (PtBRLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Ocorreu um erro ao gerar a descrição da expressão Cron."
}

func (PtBRLocale) everyMinute() string {
	return "a cada minuto"
}

func (PtBRLocale) everyHour() string {
	return "a cada hora"
}

func (PtBRLocale) atSpace() string {
	return "às "
}

func (PtBRLocale) everyMinuteBetweenX0AndX1() string {
	return "a cada minuto entre %s e %s"
}

func (PtBRLocale) at() string {
	return "às"
}

func (PtBRLocale) spaceAnd() string {
	return " e"
}

func (PtBRLocale) everySecond() string {
	return "a cada segundo"
}

func (PtBRLocale) secondsX0ThroughX1PastTheMinute() string {
	return "No segundo %s até %s de cada minuto"
}

func (PtBRLocale) minutesX0ThroughX1PastTheHour() string {
	return "do minuto %s até %s de cada hora"
}

func (PtBRLocale) betweenX0AndX1() string {
	return "entre %s e %s"
}

func (PtBRLocale) atX0() string {
	return "Às %s"
}

func (PtBRLocale) commaEveryDay() string {
	return ", a cada dia"
}

func (PtBRLocale) spaceX0OfTheMonth() string {
	return " %s do mês"
}

func (PtBRLocale) lastDay() string {
	return "o último dia"
}

func (PtBRLocale) commaAndOnX0() string {
	return ", e de %s"
}

func (PtBRLocale) commaOnlyInX0() string {
	return ", somente em %s"
}

func (PtBRLocale) commaOnTheLastDayOfTheMonth() string {
	return ", no último dia do mês"
}

func (PtBRLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", no último dia da semana do mês"
}

func (PtBRLocale) firstWeekday() string {
	return "primeiro dia da semana"
}

func (PtBRLocale) weekdayNearestDayX0() string {
	return "dia da semana mais próximo do dia %s"
}

func (PtBRLocale) commaOnTheX0OfTheMonth() string {
	return ", no %s do mês"
}

func (PtBRLocale) commaStartingX0() string {
	return ", iniciando %s"
}

func (PtBRLocale) atReboot() (string, bool) {
	return "Executar uma vez, na inicialização", true
}

func (PtBRLocale) onTheHour() (string, bool) {
	return "na hora certa", true
}

func (PtBRLocale) everyX0Seconds(string) string {
	return "a cada %s segundos"
}

func (PtBRLocale) everyX0Minutes(string) string {
	return "a cada %s minutos"
}

func (PtBRLocale) everyX0Hours(string) string {
	return "a cada %s horas"
}

func (PtBRLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", a cada %s dias de semana"
}

func (PtBRLocale) commaEveryX0Months(string) string {
	return ", a cada %s meses"
}

func (PtBRLocale) commaEveryX0Days(string) string {
	return ", a cada %s dias"
}

func (PtBRLocale) commaEveryX0Years(string) string {
	return ", a cada %s anos"
}

func (PtBRLocale) atX0SecondsPastTheMinute(string) string {
	return "aos %s segundos do minuto"
}

func (PtBRLocale) atX0MinutesPastTheHour(string) string {
	return "aos %s minutos da hora"
}

func (PtBRLocale) commaX0ThroughX1(string) string {
	return ", de %s a %s"
}

func (PtBRLocale) commaAndX0ThroughX1(string) string {
	return ", e de %s a %s"
}

func (PtBRLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", na última %s do mês"
}

func (PtBRLocale) commaOnDayX0OfTheMonth(string) string {
	return ", no dia %s do mês"
}

func (PtBRLocale) first(s string) string {
	switch s {
	case "0", "6":
		return "primeiro"
	default:
		return "primeira"
	}
}

func (PtBRLocale) second(s string) string {
	switch s {
	case "0", "6":
		return "segundo"
	default:
		return "segunda"
	}
}

func (PtBRLocale) third(s string) string {
	switch s {
	case "0", "6":
		return "terceiro"
	default:
		return "terceira"
	}
}

func (PtBRLocale) fourth(s string) string {
	switch s {
	case "0", "6":
		return "quarto"
	default:
		return "quarta"
	}
}

func (PtBRLocale) fifth(s string) string {
	switch s {
	case "0", "6":
		return "quinto"
	default:
		return "quinta"
	}
}

func (PtBRLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", entre os dias %s e %s do mês"
}

func (PtBRLocale) commaOnlyOnX0(s string) string {
	if s != "" && len(s) > 1 && s[1] == '-' {
		return ", somente %s"
	}
	return ", somente de %s"
}

func (PtBRLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s dias antes do último dia do mês"
}

func (PtBRLocale) commaOnThe(s, day string) string {
	switch day {
	case "0", "6":
		return ", no"
	default:
		return ", na "
	}
}

func (PtBRLocale) daysOfTheWeek() []string {
	return []string{"domingo", "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado"}
}

func (PtBRLocale) monthsOfTheYear() []string {
	return []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
}
