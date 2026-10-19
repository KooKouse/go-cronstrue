package cronstrue

// Code generated from cronstrue/src/i18n/locales/zh_TW.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("zh-Hant", ZhTWLocale{enLocale{}})
}

// ZhTWLocale is the zh_TW locale.
type ZhTWLocale struct {
	enLocale
}

func (ZhTWLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (ZhTWLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "產生表達式描述時發生了錯誤，請檢查 cron 表達式語法。"
}

func (ZhTWLocale) everyMinute() string {
	return "每分鐘"
}

func (ZhTWLocale) everyHour() string {
	return "每小時"
}

func (ZhTWLocale) atSpace() string {
	return "在 "
}

func (ZhTWLocale) everyMinuteBetweenX0AndX1() string {
	return "在 %s 和 %s 之間的每分鐘"
}

func (ZhTWLocale) at() string {
	return "在"
}

func (ZhTWLocale) spaceAnd() string {
	return " 和"
}

func (ZhTWLocale) everySecond() string {
	return "每秒"
}

func (ZhTWLocale) secondsX0ThroughX1PastTheMinute() string {
	return "在一分鐘後的 %s 到 %s 秒"
}

func (ZhTWLocale) minutesX0ThroughX1PastTheHour() string {
	return "在整點後的 %s 到 %s 分鐘"
}

func (ZhTWLocale) betweenX0AndX1() string {
	return "在 %s 和 %s 之間"
}

func (ZhTWLocale) atX0() string {
	return "在 %s"
}

func (ZhTWLocale) commaEveryDay() string {
	return ", 每天"
}

func (ZhTWLocale) spaceX0OfTheMonth() string {
	return "%s "
}

func (ZhTWLocale) lastDay() string {
	return "最後一天"
}

func (ZhTWLocale) commaAndOnX0() string {
	return ", 或 %s"
}

func (ZhTWLocale) commaOnlyInX0() string {
	return ", 僅在 %s"
}

func (ZhTWLocale) commaOnTheLastDayOfTheMonth() string {
	return ", 每月的最後一天"
}

func (ZhTWLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", 每月的最後一個工作日"
}

func (ZhTWLocale) firstWeekday() string {
	return "第一個工作日"
}

func (ZhTWLocale) weekdayNearestDayX0() string {
	return "最接近 %s 號的工作日"
}

func (ZhTWLocale) commaOnTheX0OfTheMonth() string {
	return ", 每月的 %s "
}

func (ZhTWLocale) commaStartingX0() string {
	return ", %s 開始"
}

func (ZhTWLocale) commaYearX0ThroughYearX1() (string, bool) {
	return ", 從 %s 年至 %s 年", true
}

func (ZhTWLocale) commaOnlyInMonthX0() (string, bool) {
	return ", 僅在 %s", true
}

func (ZhTWLocale) commaOnlyInYearX0() (string, bool) {
	return ", 僅在 %s 年", true
}

func (ZhTWLocale) atReboot() (string, bool) {
	return "開機時執行一次", true
}

func (ZhTWLocale) onTheHour() (string, bool) {
	return "整點", true
}

func (ZhTWLocale) everyX0Seconds(string) string {
	return "每 %s 秒"
}

func (ZhTWLocale) everyX0Minutes(string) string {
	return "每 %s 分鐘"
}

func (ZhTWLocale) everyX0Hours(string) string {
	return "每 %s 小時"
}

func (ZhTWLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", 每週的每 %s 天"
}

func (ZhTWLocale) commaEveryX0Months(string) string {
	return ", 每 %s 月"
}

func (ZhTWLocale) commaEveryX0Days(string) string {
	return ", 每 %s 天"
}

func (ZhTWLocale) commaEveryX0Years(string) string {
	return ", 每 %s 年"
}

func (ZhTWLocale) atX0SecondsPastTheMinute(string) string {
	return "在一分鐘後的 %s 秒"
}

func (ZhTWLocale) atX0MinutesPastTheHour(string) string {
	return "在整點後的 %s 分"
}

func (ZhTWLocale) commaX0ThroughX1(string) string {
	return ", %s 到 %s"
}

func (ZhTWLocale) commaAndX0ThroughX1(string) string {
	return ", 和 %s 到 %s"
}

func (ZhTWLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", 每月的最後一個 %s "
}

func (ZhTWLocale) dayX0() (string, bool) {
	return " %s 號", true
}

func (ZhTWLocale) commaOnDayX0OfTheMonth(string) string {
	return ", 每月的 %s"
}

func (ZhTWLocale) first(string) string {
	return "第一個"
}

func (ZhTWLocale) second(string) string {
	return "第二個"
}

func (ZhTWLocale) third(string) string {
	return "第三個"
}

func (ZhTWLocale) fourth(string) string {
	return "第四個"
}

func (ZhTWLocale) fifth(string) string {
	return "第五個"
}

func (ZhTWLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", 在每月的 %s 和 %s 之間"
}

func (ZhTWLocale) commaOnlyOnX0(string) string {
	return ", 僅在 %s"
}

func (ZhTWLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s 這個月的最後一天的前幾天"
}

func (ZhTWLocale) commaOnThe(string, string) string {
	return ", 在每月 "
}

func (ZhTWLocale) daysOfTheWeek() []string {
	return []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
}

func (ZhTWLocale) monthsOfTheYear() []string {
	return []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
}
