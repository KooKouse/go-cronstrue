package cronstrue

func init() {
	registerLocale("zh-Hans", zhCNLocale{enLocale{}})
}

// zhCNLocale is the Simplified Chinese locale, ported verbatim from
// cronstrue/src/i18n/locales/zh_CN.ts. It embeds enLocale so that any
// optional method it does not provide keeps its English "not provided"
// semantics (ok == false), which triggers the descriptor fallbacks.
type zhCNLocale struct {
	enLocale
}

func (zhCNLocale) setPeriodBeforeTime() bool { return true }
func (zhCNLocale) pm() (string, bool)        { return "下午", true }
func (zhCNLocale) am() (string, bool)        { return "上午", true }
func (zhCNLocale) commaYearX0ThroughYearX1() (string, bool) {
	return ", 从%s年至%s年", true
}
func (zhCNLocale) commaOnlyInMonthX0() (string, bool) { return ", 仅于%s份", true }
func (zhCNLocale) commaOnlyInYearX0() (string, bool)  { return ", 仅于 %s 年", true }
func (zhCNLocale) dayX0() (string, bool)              { return " %s 号", true }
func (zhCNLocale) atReboot() (string, bool)           { return "开机时运行一次", true }
func (zhCNLocale) onTheHour() (string, bool)          { return "整点", true }

func (zhCNLocale) use24HourTimeFormatByDefault() bool { return false }
func (zhCNLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "生成表达式描述时发生了错误，请检查cron表达式语法。"
}
func (zhCNLocale) everyMinute() string               { return "每分钟" }
func (zhCNLocale) everyHour() string                 { return "每小时" }
func (zhCNLocale) atSpace() string                   { return "在" }
func (zhCNLocale) everyMinuteBetweenX0AndX1() string { return "在 %s 至 %s 之间的每分钟" }
func (zhCNLocale) at() string                        { return "在" }
func (zhCNLocale) spaceAnd() string                  { return " 和" }
func (zhCNLocale) everySecond() string               { return "每秒" }
func (zhCNLocale) everyX0Seconds(string) string      { return "每隔 %s 秒" }
func (zhCNLocale) secondsX0ThroughX1PastTheMinute() string {
	return "在一分钟后的第 %s 到 %s 秒"
}
func (zhCNLocale) atX0SecondsPastTheMinute(string) string { return "在一分钟后的第 %s 秒" }
func (zhCNLocale) everyX0Minutes(string) string           { return "每隔 %s 分钟" }
func (zhCNLocale) minutesX0ThroughX1PastTheHour() string {
	return "在整点后的第 %s 到 %s 分钟"
}
func (zhCNLocale) atX0MinutesPastTheHour(string) string    { return "在整点后的第 %s 分钟" }
func (zhCNLocale) everyX0Hours(string) string              { return "每隔 %s 小时" }
func (zhCNLocale) betweenX0AndX1() string                  { return "在 %s 和 %s 之间" }
func (zhCNLocale) atX0() string                            { return "在%s" }
func (zhCNLocale) commaEveryDay() string                   { return ", 每天" }
func (zhCNLocale) commaEveryX0DaysOfTheWeek(string) string { return ", 每周的每 %s 天" }
func (zhCNLocale) commaX0ThroughX1(string) string          { return ", %s至%s" }
func (zhCNLocale) commaAndX0ThroughX1(string) string       { return ", 和%s至%s" }
func (zhCNLocale) first(string) string                     { return "第一个" }
func (zhCNLocale) second(string) string                    { return "第二个" }
func (zhCNLocale) third(string) string                     { return "第三个" }
func (zhCNLocale) fourth(string) string                    { return "第四个" }
func (zhCNLocale) fifth(string) string                     { return "第五个" }
func (zhCNLocale) commaOnThe(string, string) string        { return ", 限每月的" }
func (zhCNLocale) spaceX0OfTheMonth() string               { return "%s" }
func (zhCNLocale) lastDay() string                         { return "本月最后一天" }
func (zhCNLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", 限每月的最后一个%s"
}
func (zhCNLocale) commaOnlyOnX0(string) string { return ", 仅%s" }
func (zhCNLocale) commaAndOnX0() string        { return ", 或者为%s" }
func (zhCNLocale) commaEveryX0Months(string) string {
	return ", 每隔 %s 个月"
}
func (zhCNLocale) commaOnlyInX0() string { return ", 仅限%s" }
func (zhCNLocale) commaOnTheLastDayOfTheMonth() string {
	return ", 限每月的最后一天"
}
func (zhCNLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", 限每月的最后一个工作日"
}
func (zhCNLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", 限每月最后%s天"
}
func (zhCNLocale) firstWeekday() string           { return "第一个工作日" }
func (zhCNLocale) weekdayNearestDayX0() string    { return "最接近 %s 号的工作日" }
func (zhCNLocale) commaOnTheX0OfTheMonth() string { return ", 限每月的%s" }
func (zhCNLocale) commaEveryX0Days(string) string { return ", 每隔 %s 天" }
func (zhCNLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", 限每月的 %s 至 %s 之间"
}
func (zhCNLocale) commaOnDayX0OfTheMonth(string) string { return ", 限每月%s" }
func (zhCNLocale) commaEveryX0Years(string) string      { return ", 每隔 %s 年" }
func (zhCNLocale) commaStartingX0() string              { return ", %s开始" }

func (zhCNLocale) daysOfTheWeek() []string {
	return []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
}

func (zhCNLocale) monthsOfTheYear() []string {
	return []string{
		"一月", "二月", "三月", "四月", "五月", "六月",
		"七月", "八月", "九月", "十月", "十一月", "十二月",
	}
}
