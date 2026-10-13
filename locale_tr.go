package cronstrue

// Code generated from cronstrue/src/i18n/locales/tr.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("tr", TrLocale{enLocale{}})
}

// TrLocale is the tr locale.
type TrLocale struct {
	enLocale
}

func (TrLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (TrLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "İfade açıklamasını oluştururken bir hata oluştu. Cron ifadesini gözden geçirin."
}

func (TrLocale) everyMinute() string {
	return "her dakika"
}

func (TrLocale) everyHour() string {
	return "her saat"
}

func (TrLocale) atSpace() string {
	return "Saat "
}

func (TrLocale) everyMinuteBetweenX0AndX1() string {
	return "Saat %s ve %s arasındaki her dakika"
}

func (TrLocale) at() string {
	return "Saat"
}

func (TrLocale) spaceAnd() string {
	return " ve"
}

func (TrLocale) everySecond() string {
	return "her saniye"
}

func (TrLocale) secondsX0ThroughX1PastTheMinute() string {
	return "dakikaların %s. ve %s. saniyeleri arası"
}

func (TrLocale) minutesX0ThroughX1PastTheHour() string {
	return "saatlerin %s. ve %s. dakikaları arası"
}

func (TrLocale) betweenX0AndX1() string {
	return "%s ile %s arasında"
}

func (TrLocale) atX0() string {
	return "saat %s"
}

func (TrLocale) commaEveryDay() string {
	return ", her gün"
}

func (TrLocale) spaceX0OfTheMonth() string {
	return " %s günü"
}

func (TrLocale) lastDay() string {
	return "son gün"
}

func (TrLocale) commaAndOnX0() string {
	return ", ve %s"
}

func (TrLocale) commaOnlyInX0() string {
	return ", sadece %s için"
}

func (TrLocale) commaOnTheLastDayOfTheMonth() string {
	return ", ayın son günü"
}

func (TrLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", ayın son iş günü"
}

func (TrLocale) firstWeekday() string {
	return "ilk iş günü"
}

func (TrLocale) weekdayNearestDayX0() string {
	return "%s. günü sonrasındaki ilk iş günü"
}

func (TrLocale) commaOnTheX0OfTheMonth() string {
	return ", ayın %s"
}

func (TrLocale) commaStartingX0() string {
	return ", başlangıç %s"
}

func (TrLocale) atReboot() (string, bool) {
	return "Başlangıçta bir kez çalıştır", true
}

func (TrLocale) onTheHour() (string, bool) {
	return "saatin başında", true
}

func (TrLocale) everyX0Seconds(string) string {
	return "her %s saniyede bir"
}

func (TrLocale) everyX0Minutes(string) string {
	return "her %s dakikada bir"
}

func (TrLocale) everyX0Hours(string) string {
	return "her %s saatte"
}

func (TrLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", ayın her %s günü"
}

func (TrLocale) commaEveryX0Months(string) string {
	return ", %s ayda bir"
}

func (TrLocale) commaEveryX0Days(string) string {
	return ", %s günde bir"
}

func (TrLocale) commaEveryX0Years(string) string {
	return ", %s yılda bir"
}

func (TrLocale) atX0SecondsPastTheMinute(string) string {
	return "dakikaların %s. saniyesinde"
}

func (TrLocale) atX0MinutesPastTheHour(string) string {
	return "saatlerin %s. dakikasında"
}

func (TrLocale) commaX0ThroughX1(string) string {
	return ", %s ile %s arasında"
}

func (TrLocale) commaAndX0ThroughX1(string) string {
	return ", ve %s ile %s arasında"
}

func (TrLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", ayın son %s günü"
}

func (TrLocale) commaOnDayX0OfTheMonth(string) string {
	return ", ayın %s. günü"
}

func (TrLocale) first(string) string {
	return "ilk"
}

func (TrLocale) second(string) string {
	return "ikinci"
}

func (TrLocale) third(string) string {
	return "üçüncü"
}

func (TrLocale) fourth(string) string {
	return "dördüncü"
}

func (TrLocale) fifth(string) string {
	return "beşinci"
}

func (TrLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", ayın %s. ve %s. günleri arası"
}

func (TrLocale) commaOnlyOnX0(string) string {
	return ", sadece %s günü"
}

func (TrLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s ayın son gününden önceki günler"
}

func (TrLocale) commaOnThe(string, string) string {
	return ", ayın "
}

func (TrLocale) daysOfTheWeek() []string {
	return []string{"Pazar", "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi"}
}

func (TrLocale) monthsOfTheYear() []string {
	return []string{"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran", "Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık"}
}
