package cronstrue

// Code generated from cronstrue/src/i18n/locales/id.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("id", IdLocale{enLocale{}})
}

// IdLocale is the id locale.
type IdLocale struct {
	enLocale
}

func (IdLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (IdLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Terjadi kesalahan saat membuat deskripsi ekspresi. Periksa sintaks ekspresi cron."
}

func (IdLocale) everyMinute() string {
	return "setiap menit"
}

func (IdLocale) everyHour() string {
	return "setiap jam"
}

func (IdLocale) atSpace() string {
	return "Pada "
}

func (IdLocale) everyMinuteBetweenX0AndX1() string {
	return "Setiap menit diantara %s dan %s"
}

func (IdLocale) at() string {
	return "Pada"
}

func (IdLocale) spaceAnd() string {
	return " dan"
}

func (IdLocale) everySecond() string {
	return "setiap detik"
}

func (IdLocale) secondsX0ThroughX1PastTheMinute() string {
	return "detik ke %s sampai %s melewati menit"
}

func (IdLocale) minutesX0ThroughX1PastTheHour() string {
	return "menit ke %s sampai %s melewati jam"
}

func (IdLocale) betweenX0AndX1() string {
	return "diantara %s dan %s"
}

func (IdLocale) atX0() string {
	return "pada %s"
}

func (IdLocale) commaEveryDay() string {
	return ", setiap hari"
}

func (IdLocale) spaceX0OfTheMonth() string {
	return " %s pada bulan"
}

func (IdLocale) lastDay() string {
	return "hari terakhir"
}

func (IdLocale) commaAndOnX0() string {
	return ", dan pada %s"
}

func (IdLocale) commaOnlyInX0() string {
	return ", hanya pada %s"
}

func (IdLocale) commaOnTheLastDayOfTheMonth() string {
	return ", pada hari terakhir bulan ini"
}

func (IdLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", pada hari kerja terakhir setiap bulan"
}

func (IdLocale) firstWeekday() string {
	return "hari kerja pertama"
}

func (IdLocale) weekdayNearestDayX0() string {
	return "hari kerja terdekat %s"
}

func (IdLocale) commaOnTheX0OfTheMonth() string {
	return ", pada %s bulan ini"
}

func (IdLocale) commaStartingX0() string {
	return ", mulai pada %s"
}

func (IdLocale) atReboot() (string, bool) {
	return "Jalankan sekali, saat mulai", true
}

func (IdLocale) onTheHour() (string, bool) {
	return "tepat pada jam", true
}

func (IdLocale) everyX0Seconds(string) string {
	return "setiap %s detik"
}

func (IdLocale) everyX0Minutes(string) string {
	return "setiap %s menit"
}

func (IdLocale) everyX0Hours(string) string {
	return "setiap %s jam"
}

func (IdLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", setiap hari %s  dalam seminggu"
}

func (IdLocale) commaEveryX0Months(string) string {
	return ", setiap bulan %s "
}

func (IdLocale) commaEveryX0Days(string) string {
	return ", setiap %s hari"
}

func (IdLocale) commaEveryX0Years(string) string {
	return ", setiap %s tahun"
}

func (IdLocale) atX0SecondsPastTheMinute(string) string {
	return "pada %s detik lewat satu menit"
}

func (IdLocale) atX0MinutesPastTheHour(string) string {
	return "pada %s menit melewati jam"
}

func (IdLocale) commaX0ThroughX1(string) string {
	return ", %s sampai %s"
}

func (IdLocale) commaAndX0ThroughX1(string) string {
	return ", dan %s sampai %s"
}

func (IdLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", pada %s terakhir bulan ini"
}

func (IdLocale) commaOnDayX0OfTheMonth(string) string {
	return ", pada hari %s dalam sebulan"
}

func (IdLocale) first(string) string {
	return "pertama"
}

func (IdLocale) second(string) string {
	return "kedua"
}

func (IdLocale) third(string) string {
	return "ketiga"
}

func (IdLocale) fourth(string) string {
	return "keempat"
}

func (IdLocale) fifth(string) string {
	return "kelima"
}

func (IdLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", antara hari %s dan %s dalam sebulan"
}

func (IdLocale) commaOnlyOnX0(string) string {
	return ", hanya pada %s"
}

func (IdLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s hari sebelum hari terakhir setiap bulan"
}

func (IdLocale) commaOnThe(string, string) string {
	return ", di "
}

func (IdLocale) daysOfTheWeek() []string {
	return []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
}

func (IdLocale) monthsOfTheYear() []string {
	return []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
}
