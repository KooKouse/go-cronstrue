package cronstrue

// Code generated from cronstrue/src/i18n/locales/my.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("my", MyLocale{enLocale{}})
}

// MyLocale is the my locale.
type MyLocale struct {
	enLocale
}

func (MyLocale) use24HourTimeFormatByDefault() bool {
	return false
}

func (MyLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Terdapat ralat semasa menjana penerangan ungkapan. Sila periksa sintaks ungkapan cron."
}

func (MyLocale) everyMinute() string {
	return "setiap minit"
}

func (MyLocale) everyHour() string {
	return "setiap jam"
}

func (MyLocale) atSpace() string {
	return "Pada "
}

func (MyLocale) everyMinuteBetweenX0AndX1() string {
	return "Setiap minit antara %s dan %s"
}

func (MyLocale) at() string {
	return "Pada"
}

func (MyLocale) spaceAnd() string {
	return " dan"
}

func (MyLocale) everySecond() string {
	return "setiap saat"
}

func (MyLocale) secondsX0ThroughX1PastTheMinute() string {
	return "saat ke %s hingga %s selepas minit"
}

func (MyLocale) minutesX0ThroughX1PastTheHour() string {
	return "minit ke %s hingga %s selepas jam"
}

func (MyLocale) betweenX0AndX1() string {
	return "antara %s dan %s"
}

func (MyLocale) atX0() string {
	return "pada %s"
}

func (MyLocale) commaEveryDay() string {
	return ", setiap hari"
}

func (MyLocale) spaceX0OfTheMonth() string {
	return " %s pada bulan"
}

func (MyLocale) lastDay() string {
	return "hari terakhir"
}

func (MyLocale) commaAndOnX0() string {
	return ", dan pada %s"
}

func (MyLocale) commaOnlyInX0() string {
	return ", hanya pada %s"
}

func (MyLocale) commaOnTheLastDayOfTheMonth() string {
	return ", pada hari terakhir bulan"
}

func (MyLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", pada minggu terakhir bulan"
}

func (MyLocale) firstWeekday() string {
	return "hari pertama minggu bekerja"
}

func (MyLocale) weekdayNearestDayX0() string {
	return "hari bekerja yang terdekat dengan %s"
}

func (MyLocale) commaOnTheX0OfTheMonth() string {
	return ", pada %s bulan"
}

func (MyLocale) commaStartingX0() string {
	return ", bermula %s"
}

func (MyLocale) atReboot() (string, bool) {
	return "Jalankan sekali, semasa permulaan", true
}

func (MyLocale) onTheHour() (string, bool) {
	return "pada waktu yang tepat", true
}

func (MyLocale) everyX0Seconds(string) string {
	return "setiap %s saat"
}

func (MyLocale) everyX0Minutes(string) string {
	return "setiap %s minit"
}

func (MyLocale) everyX0Hours(string) string {
	return "setiap %s jam"
}

func (MyLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", setiap %s hari dalam minggu"
}

func (MyLocale) commaEveryX0Months(string) string {
	return ", setiap bulan %s"
}

func (MyLocale) commaEveryX0Days(string) string {
	return ", setiap %s hari"
}

func (MyLocale) commaEveryX0Years(string) string {
	return ", setiap %s tahun"
}

func (MyLocale) atX0SecondsPastTheMinute(string) string {
	return "pada %s saat selepas minit"
}

func (MyLocale) atX0MinutesPastTheHour(string) string {
	return "pada %s minit selepas jam"
}

func (MyLocale) commaX0ThroughX1(string) string {
	return ", %s hingga %s"
}

func (MyLocale) commaAndX0ThroughX1(string) string {
	return ", dan %s hingga %s"
}

func (MyLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", pada %s terakhir bulan"
}

func (MyLocale) commaOnDayX0OfTheMonth(string) string {
	return ", pada hari %s dalam bulan"
}

func (MyLocale) first(string) string {
	return "pertama"
}

func (MyLocale) second(string) string {
	return "kedua"
}

func (MyLocale) third(string) string {
	return "ketiga"
}

func (MyLocale) fourth(string) string {
	return "keempat"
}

func (MyLocale) fifth(string) string {
	return "kelima"
}

func (MyLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", antara hari %s dan %s dalam bulan"
}

func (MyLocale) commaOnlyOnX0(string) string {
	return ", hanya pada %s"
}

func (MyLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s hari sebelum hari terakhir bulan"
}

func (MyLocale) commaOnThe(string, string) string {
	return ", pada "
}

func (MyLocale) daysOfTheWeek() []string {
	return []string{"Ahad", "Isnin", "Selasa", "Rabu", "Khamis", "Jumaat", "Sabtu"}
}

func (MyLocale) monthsOfTheYear() []string {
	return []string{"Januari", "Februari", "Mac", "April", "Mei", "Jun", "Julai", "Ogos", "September", "Oktober", "November", "Disember"}
}
