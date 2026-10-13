package cronstrue

// Code generated from cronstrue/src/i18n/locales/fa.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("fa", FaLocale{enLocale{}})
}

// FaLocale is the fa locale.
type FaLocale struct {
	enLocale
}

func (FaLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (FaLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "خطایی در نمایش توضیحات این وظیفه رخ داد. لطفا ساختار آن را بررسی کنید."
}

func (FaLocale) everyMinute() string {
	return "هر دقیقه"
}

func (FaLocale) everyHour() string {
	return "هر ساعت"
}

func (FaLocale) atSpace() string {
	return "در "
}

func (FaLocale) everyMinuteBetweenX0AndX1() string {
	return "هر دقیقه بین %s و %s"
}

func (FaLocale) at() string {
	return "در"
}

func (FaLocale) spaceAnd() string {
	return " و"
}

func (FaLocale) everySecond() string {
	return "هر ثانیه"
}

func (FaLocale) secondsX0ThroughX1PastTheMinute() string {
	return "ثانیه %s تا %s دقیقه گذشته"
}

func (FaLocale) minutesX0ThroughX1PastTheHour() string {
	return "دقیقه %s تا %s ساعت گذشته"
}

func (FaLocale) betweenX0AndX1() string {
	return "بین %s و %s"
}

func (FaLocale) atX0() string {
	return "در %s"
}

func (FaLocale) commaEveryDay() string {
	return ", هر روز"
}

func (FaLocale) spaceX0OfTheMonth() string {
	return " %s ماه"
}

func (FaLocale) lastDay() string {
	return "آخرین روز"
}

func (FaLocale) commaAndOnX0() string {
	return ", و در %s"
}

func (FaLocale) commaOnlyInX0() string {
	return ", فقط در %s"
}

func (FaLocale) commaOnTheLastDayOfTheMonth() string {
	return ", در آخرین روز ماه"
}

func (FaLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", در آخرین روز ماه"
}

func (FaLocale) firstWeekday() string {
	return "اولین روز"
}

func (FaLocale) weekdayNearestDayX0() string {
	return "روز نزدیک به روز %s"
}

func (FaLocale) commaOnTheX0OfTheMonth() string {
	return ", در %s ماه"
}

func (FaLocale) commaStartingX0() string {
	return ", آغاز %s"
}

func (FaLocale) atReboot() (string, bool) {
	return "یک بار اجرا شود، هنگام راه‌اندازی", true
}

func (FaLocale) onTheHour() (string, bool) {
	return "سر ساعت", true
}

func (FaLocale) everyX0Seconds(string) string {
	return "هر %s ثانیه"
}

func (FaLocale) everyX0Minutes(string) string {
	return "هر %s دقیقه"
}

func (FaLocale) everyX0Hours(string) string {
	return "هر %s ساعت"
}

func (FaLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", هر %s روز از هفته"
}

func (FaLocale) commaEveryX0Months(string) string {
	return ", هر %s ماه"
}

func (FaLocale) commaEveryX0Days(string) string {
	return ", هر %s روز"
}

func (FaLocale) commaEveryX0Years(string) string {
	return ", هر %s سال"
}

func (FaLocale) atX0SecondsPastTheMinute(string) string {
	return "در %s قانیه از دقیقه گذشته"
}

func (FaLocale) atX0MinutesPastTheHour(string) string {
	return "در %s دقیقه پس از ساعت"
}

func (FaLocale) commaX0ThroughX1(string) string {
	return ", %s تا %s"
}

func (FaLocale) commaAndX0ThroughX1(string) string {
	return ", و %s تا %s"
}

func (FaLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", در %s ماه"
}

func (FaLocale) commaOnDayX0OfTheMonth(string) string {
	return ", در %s ماه"
}

func (FaLocale) first(string) string {
	return "اول"
}

func (FaLocale) second(string) string {
	return "دوم"
}

func (FaLocale) third(string) string {
	return "سوم"
}

func (FaLocale) fourth(string) string {
	return "چهارم"
}

func (FaLocale) fifth(string) string {
	return "پنجم"
}

func (FaLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", بین روز %s و %s ماه"
}

func (FaLocale) commaOnlyOnX0(string) string {
	return ", فقط در %s"
}

func (FaLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s روز قبل از آخرین روز ماه"
}

func (FaLocale) commaOnThe(string, string) string {
	return ", در "
}

func (FaLocale) daysOfTheWeek() []string {
	return []string{"یک‌شنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه", "شنبه"}
}

func (FaLocale) monthsOfTheYear() []string {
	return []string{"ژانویه", "فوریه", "مارس", "آپریل", "مه", "ژوئن", "ژوئیه", "آگوست", "سپتامبر", "اکتبر", "نوامبر", "دسامبر"}
}
