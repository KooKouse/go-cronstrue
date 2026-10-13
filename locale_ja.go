package cronstrue

// Code generated from cronstrue/src/i18n/locales/ja.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("ja", JaLocale{enLocale{}})
}

// JaLocale is the ja locale.
type JaLocale struct {
	enLocale
}

func (JaLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (JaLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "式の記述を生成する際にエラーが発生しました。Cron 式の構文を確認してください。"
}

func (JaLocale) everyMinute() string {
	return "毎分"
}

func (JaLocale) everyHour() string {
	return "毎時"
}

func (JaLocale) atSpace() string {
	return "次において実施"
}

func (JaLocale) everyMinuteBetweenX0AndX1() string {
	return "%s から %s まで毎分"
}

func (JaLocale) at() string {
	return "次において実施"
}

func (JaLocale) spaceAnd() string {
	return "と"
}

func (JaLocale) everySecond() string {
	return "毎秒"
}

func (JaLocale) secondsX0ThroughX1PastTheMinute() string {
	return "毎分 %s 秒から %s 秒まで"
}

func (JaLocale) minutesX0ThroughX1PastTheHour() string {
	return "毎時 %s 分から %s 分まで"
}

func (JaLocale) betweenX0AndX1() string {
	return "%s と %s の間"
}

func (JaLocale) atX0() string {
	return "次において実施 %s"
}

func (JaLocale) commaEveryDay() string {
	return "、毎日"
}

func (JaLocale) spaceX0OfTheMonth() string {
	return "月のうち %s"
}

func (JaLocale) lastDay() string {
	return "最終日"
}

func (JaLocale) commaAndOnX0() string {
	return "、〜と %s"
}

func (JaLocale) commaOnlyInX0() string {
	return "%s でのみ"
}

func (JaLocale) commaOnTheLastDayOfTheMonth() string {
	return "次の最終日に"
}

func (JaLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return "月の最後の平日に"
}

func (JaLocale) firstWeekday() string {
	return "最初の平日"
}

func (JaLocale) weekdayNearestDayX0() string {
	return "%s 日の直近の平日"
}

func (JaLocale) commaOnTheX0OfTheMonth() string {
	return "月の %s に"
}

func (JaLocale) commaStartingX0() string {
	return "、%s に開始"
}

func (JaLocale) atReboot() (string, bool) {
	return "起動時に1回実行", true
}

func (JaLocale) onTheHour() (string, bool) {
	return "時ちょうど", true
}

func (JaLocale) everyX0Seconds(string) string {
	return "%s 秒ごと"
}

func (JaLocale) everyX0Minutes(string) string {
	return "%s 分ごと"
}

func (JaLocale) everyX0Hours(string) string {
	return "%s 時間ごと"
}

func (JaLocale) commaEveryX0DaysOfTheWeek(string) string {
	return "、週のうち %s 日ごと"
}

func (JaLocale) commaEveryX0Months(string) string {
	return "、%s か月ごと"
}

func (JaLocale) commaEveryX0Days(string) string {
	return "、月のうち %s 日ごと"
}

func (JaLocale) commaEveryX0Years(string) string {
	return "、%s 年ごと"
}

func (JaLocale) atX0SecondsPastTheMinute(string) string {
	return "毎分 %s 秒過ぎ"
}

func (JaLocale) atX0MinutesPastTheHour(string) string {
	return "毎時 %s 分過ぎ"
}

func (JaLocale) commaX0ThroughX1(string) string {
	return "、%s から %s まで"
}

func (JaLocale) commaAndX0ThroughX1(string) string {
	return "、%s から %s まで"
}

func (JaLocale) commaOnTheLastX0OfTheMonth(string) string {
	return "月の最後の %s に"
}

func (JaLocale) commaOnDayX0OfTheMonth(string) string {
	return "、月の %s 日目"
}

func (JaLocale) first(string) string {
	return "1 番目"
}

func (JaLocale) second(string) string {
	return "2 番目"
}

func (JaLocale) third(string) string {
	return "3 番目"
}

func (JaLocale) fourth(string) string {
	return "4 番目"
}

func (JaLocale) fifth(string) string {
	return "5 番目"
}

func (JaLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return "、月の %s 日から %s 日の間"
}

func (JaLocale) commaOnlyOnX0(string) string {
	return "%s にのみ"
}

func (JaLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return "月の最終日の %s 日前"
}

func (JaLocale) commaOnThe(string, string) string {
	return "次に"
}

func (JaLocale) daysOfTheWeek() []string {
	return []string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}
}

func (JaLocale) monthsOfTheYear() []string {
	return []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}
}
