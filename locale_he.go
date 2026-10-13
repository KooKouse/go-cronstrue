package cronstrue

// Code generated from cronstrue/src/i18n/locales/he.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("he", HeLocale{enLocale{}})
}

// HeLocale is the he locale.
type HeLocale struct {
	enLocale
}

func (HeLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (HeLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "אירעה שגיאה בעת יצירת תיאור הביטוי. בדוק את תחביר הביטוי cron."
}

func (HeLocale) everyMinute() string {
	return "כל דקה"
}

func (HeLocale) everyHour() string {
	return "כל שעה"
}

func (HeLocale) atSpace() string {
	return "ב "
}

func (HeLocale) everyMinuteBetweenX0AndX1() string {
	return "כל דקה %s עד %s"
}

func (HeLocale) at() string {
	return "ב"
}

func (HeLocale) spaceAnd() string {
	return " ו"
}

func (HeLocale) everySecond() string {
	return "כל שניה"
}

func (HeLocale) secondsX0ThroughX1PastTheMinute() string {
	return "%s עד %s שניות של הדקה"
}

func (HeLocale) minutesX0ThroughX1PastTheHour() string {
	return "%s עד %s דקות של השעה"
}

func (HeLocale) betweenX0AndX1() string {
	return "%s עד %s"
}

func (HeLocale) atX0() string {
	return "ב %s"
}

func (HeLocale) commaEveryDay() string {
	return ", כל יום"
}

func (HeLocale) spaceX0OfTheMonth() string {
	return " %s של החודש"
}

func (HeLocale) lastDay() string {
	return "היום האחרון"
}

func (HeLocale) commaAndOnX0() string {
	return ", וב %s"
}

func (HeLocale) commaOnlyInX0() string {
	return ", רק ב %s"
}

func (HeLocale) commaOnTheLastDayOfTheMonth() string {
	return ", ביום האחרון של החודש"
}

func (HeLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", ביום החול האחרון של החודש"
}

func (HeLocale) firstWeekday() string {
	return "יום החול הראשון"
}

func (HeLocale) weekdayNearestDayX0() string {
	return "יום החול הראשון הקרוב אל %s"
}

func (HeLocale) commaOnTheX0OfTheMonth() string {
	return ", ביום ה%s של החודש"
}

func (HeLocale) commaStartingX0() string {
	return ", החל מ %s"
}

func (HeLocale) atReboot() (string, bool) {
	return "הרצה אחת, בעת ההפעלה", true
}

func (HeLocale) onTheHour() (string, bool) {
	return "בשעה עגולה", true
}

func (HeLocale) everyX0Seconds(string) string {
	return "כל %s שניות"
}

func (HeLocale) everyX0Minutes(string) string {
	return "כל %s דקות"
}

func (HeLocale) everyX0Hours(string) string {
	return "כל %s שעות"
}

func (HeLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", כל %s ימים בשבוע"
}

func (HeLocale) commaEveryX0Months(string) string {
	return ", כל %s חודשים"
}

func (HeLocale) commaEveryX0Days(string) string {
	return ", כל %s ימים"
}

func (HeLocale) commaEveryX0Years(string) string {
	return ", כל %s שנים"
}

func (HeLocale) atX0SecondsPastTheMinute(string) string {
	return "ב %s שניות של הדקה"
}

func (HeLocale) atX0MinutesPastTheHour(string) string {
	return "ב %s דקות של השעה"
}

func (HeLocale) commaX0ThroughX1(string) string {
	return ", %s עד %s"
}

func (HeLocale) commaAndX0ThroughX1(string) string {
	return ", ו %s עד %s"
}

func (HeLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", רק ב %s של החודש"
}

func (HeLocale) commaOnDayX0OfTheMonth(string) string {
	return ", ביום ה%s של החודש"
}

func (HeLocale) first(string) string {
	return "ראשון"
}

func (HeLocale) second(string) string {
	return "שני"
}

func (HeLocale) third(string) string {
	return "שלישי"
}

func (HeLocale) fourth(string) string {
	return "רביעי"
}

func (HeLocale) fifth(string) string {
	return "חמישי"
}

func (HeLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", בין היום ה%s וה%s של החודש"
}

func (HeLocale) commaOnlyOnX0(string) string {
	return ", רק ב %s"
}

func (HeLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s ימים לפני היום האחרון בחודש"
}

func (HeLocale) commaOnThe(string, string) string {
	return ", ב "
}

func (HeLocale) daysOfTheWeek() []string {
	return []string{"יום ראשון", "יום שני", "יום שלישי", "יום רביעי", "יום חמישי", "יום שישי", "יום שבת"}
}

func (HeLocale) monthsOfTheYear() []string {
	return []string{"ינואר", "פברואר", "מרץ", "אפריל", "מאי", "יוני", "יולי", "אוגוסט", "ספטמבר", "אוקטובר", "נובמבר", "דצמבר"}
}
