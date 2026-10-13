package cronstrue

// Code generated from cronstrue/src/i18n/locales/ar.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("ar", ArLocale{enLocale{}})
}

// ArLocale is the ar locale.
type ArLocale struct {
	enLocale
}

func (ArLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (ArLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "حدث خطأ في إنشاء وصف المصطلح٠ تأكد من تركيب مصطلح الكرون"
}

func (ArLocale) everyMinute() string {
	return "كل دقيقة"
}

func (ArLocale) everyHour() string {
	return "كل ساعة"
}

func (ArLocale) atSpace() string {
	return " "
}

func (ArLocale) everyMinuteBetweenX0AndX1() string {
	return "كل دقيقة بين %s و %s"
}

func (ArLocale) at() string {
	return ""
}

func (ArLocale) spaceAnd() string {
	return " و"
}

func (ArLocale) everySecond() string {
	return "كل ثانية"
}

func (ArLocale) secondsX0ThroughX1PastTheMinute() string {
	return "الثواني %s حتى %s من بداية الدقيقة"
}

func (ArLocale) minutesX0ThroughX1PastTheHour() string {
	return "الدقائق %s حتى %s من بداية الساعة"
}

func (ArLocale) betweenX0AndX1() string {
	return "بين %s و %s"
}

func (ArLocale) atX0() string {
	return "%s"
}

func (ArLocale) commaEveryDay() string {
	return "، كل يوم"
}

func (ArLocale) spaceX0OfTheMonth() string {
	return " %s من الشهر"
}

func (ArLocale) lastDay() string {
	return "اليوم الأخير"
}

func (ArLocale) commaAndOnX0() string {
	return "، وفي %s"
}

func (ArLocale) commaOnlyInX0() string {
	return "، %s فقط"
}

func (ArLocale) commaOnTheLastDayOfTheMonth() string {
	return "، في اخر يوم من الشهر"
}

func (ArLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return "، في اخر يوم أسبوع من الشهر"
}

func (ArLocale) firstWeekday() string {
	return "اول ايام الأسبوع"
}

func (ArLocale) weekdayNearestDayX0() string {
	return "يوم الأسبوع الأقرب ليوم %s"
}

func (ArLocale) commaOnTheX0OfTheMonth() string {
	return "، في %s من الشهر"
}

func (ArLocale) commaStartingX0() string {
	return "، بداية من %s"
}

func (ArLocale) atReboot() (string, bool) {
	return "تشغيل مرة واحدة، عند بدء التشغيل", true
}

func (ArLocale) onTheHour() (string, bool) {
	return "في تمام الساعة", true
}

func (ArLocale) everyX0Seconds(string) string {
	return "كل %s ثواني"
}

func (ArLocale) everyX0Minutes(string) string {
	return "كل %s دقائق"
}

func (ArLocale) everyX0Hours(string) string {
	return "كل %s ساعات"
}

func (ArLocale) commaEveryX0DaysOfTheWeek(string) string {
	return "، كل %s من أيام الأسبوع"
}

func (ArLocale) commaEveryX0Months(string) string {
	return "، كل %s أشهر"
}

func (ArLocale) commaEveryX0Days(string) string {
	return "، كل %s أيام"
}

func (ArLocale) commaEveryX0Years(string) string {
	return "، كل %s سنوات"
}

func (ArLocale) atX0SecondsPastTheMinute(string) string {
	return "الثانية %s من بداية الدقيقة"
}

func (ArLocale) atX0MinutesPastTheHour(string) string {
	return "الدقيقة %s من بداية الساعة"
}

func (ArLocale) commaX0ThroughX1(string) string {
	return "، %s حتى %s"
}

func (ArLocale) commaAndX0ThroughX1(string) string {
	return "، و %s حتى %s"
}

func (ArLocale) commaOnTheLastX0OfTheMonth(string) string {
	return "، في اخر %s من الشهر"
}

func (ArLocale) commaOnDayX0OfTheMonth(string) string {
	return "، في اليوم %s من الشهر"
}

func (ArLocale) first(string) string {
	return "أول"
}

func (ArLocale) second(string) string {
	return "ثاني"
}

func (ArLocale) third(string) string {
	return "ثالث"
}

func (ArLocale) fourth(string) string {
	return "رابع"
}

func (ArLocale) fifth(string) string {
	return "خامس"
}

func (ArLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return "، بين يوم %s و %s من الشهر"
}

func (ArLocale) commaOnlyOnX0(string) string {
	return "، %s فقط"
}

func (ArLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return "، %s أيام قبل اخر يوم من الشهر"
}

func (ArLocale) commaOnThe(string, string) string {
	return "، في ال"
}

func (ArLocale) daysOfTheWeek() []string {
	return []string{"الأحد", "الإثنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت"}
}

func (ArLocale) monthsOfTheYear() []string {
	return []string{"يناير", "فبراير", "مارس", "ابريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"}
}
