package cronstrue

// Code generated from cronstrue/src/i18n/locales/vi.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("vi", ViLocale{enLocale{}})
}

// ViLocale is the vi locale.
type ViLocale struct {
	enLocale
}

func (ViLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (ViLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "Đã xảy ra lỗi khi tạo mô tả biểu thức. Vui lòng kiểm tra cú pháp biểu thức cron."
}

func (ViLocale) everyMinute() string {
	return "mỗi phút"
}

func (ViLocale) everyHour() string {
	return "mỗi giờ"
}

func (ViLocale) atSpace() string {
	return "Vào "
}

func (ViLocale) everyMinuteBetweenX0AndX1() string {
	return "Mỗi phút giữa %s and %s"
}

func (ViLocale) at() string {
	return "Vào"
}

func (ViLocale) spaceAnd() string {
	return " và"
}

func (ViLocale) everySecond() string {
	return "mỗi giây"
}

func (ViLocale) secondsX0ThroughX1PastTheMinute() string {
	return "giây thứ %s qua phút thứ %s"
}

func (ViLocale) minutesX0ThroughX1PastTheHour() string {
	return "phút thứ %s qua %s tiếng"
}

func (ViLocale) betweenX0AndX1() string {
	return "giữa %s và %s"
}

func (ViLocale) atX0() string {
	return "vào %s"
}

func (ViLocale) commaEveryDay() string {
	return ", mỗi ngày"
}

func (ViLocale) spaceX0OfTheMonth() string {
	return " %s của tháng"
}

func (ViLocale) lastDay() string {
	return "ngày cuối cùng"
}

func (ViLocale) commaAndOnX0() string {
	return ", và hơn %s"
}

func (ViLocale) commaOnlyInX0() string {
	return ", chỉ trong %s"
}

func (ViLocale) commaOnTheLastDayOfTheMonth() string {
	return ", vào ngày cuối cùng của tháng"
}

func (ViLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", vào ngày cuối tuần của tháng"
}

func (ViLocale) firstWeekday() string {
	return "ngày đầu tuần"
}

func (ViLocale) weekdayNearestDayX0() string {
	return "ngày trong tuần ngày gần nhất %s"
}

func (ViLocale) commaOnTheX0OfTheMonth() string {
	return ", vào ngày %s của tháng"
}

func (ViLocale) commaStartingX0() string {
	return ", bắt đầu %s"
}

func (ViLocale) atReboot() (string, bool) {
	return "Chạy một lần, khi khởi động", true
}

func (ViLocale) onTheHour() (string, bool) {
	return "đúng giờ", true
}

func (ViLocale) everyX0Seconds(string) string {
	return "mỗi %s giây"
}

func (ViLocale) everyX0Minutes(string) string {
	return "mỗi %s phút"
}

func (ViLocale) everyX0Hours(string) string {
	return "mỗi %s tiếng"
}

func (ViLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", mỗi %s ngày trong tuần"
}

func (ViLocale) commaEveryX0Months(string) string {
	return ", mỗi ngày %s tháng"
}

func (ViLocale) commaEveryX0Days(string) string {
	return ", mỗi %s ngày"
}

func (ViLocale) commaEveryX0Years(string) string {
	return ", mỗi %s năm"
}

func (ViLocale) atX0SecondsPastTheMinute(string) string {
	return "tại giây thứ %s của mỗi phút"
}

func (ViLocale) atX0MinutesPastTheHour(string) string {
	return "vào %s phút của mỗi tiếng"
}

func (ViLocale) commaX0ThroughX1(string) string {
	return ", %s đến %s"
}

func (ViLocale) commaAndX0ThroughX1(string) string {
	return ", %s đến %s"
}

func (ViLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", vào ngày %s cuối cùng của tháng"
}

func (ViLocale) commaOnDayX0OfTheMonth(string) string {
	return ", vào %s ngày trong tháng"
}

func (ViLocale) first(string) string {
	return "đầu tiên"
}

func (ViLocale) second(string) string {
	return "thứ 2"
}

func (ViLocale) third(string) string {
	return "thứ 3"
}

func (ViLocale) fourth(string) string {
	return "thứ 4"
}

func (ViLocale) fifth(string) string {
	return "thứ 5"
}

func (ViLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", giữa ngày %s và %s trong tháng"
}

func (ViLocale) commaOnlyOnX0(string) string {
	return ", chỉ trên %s"
}

func (ViLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s ngày trước ngày cuối cùng của tháng"
}

func (ViLocale) commaOnThe(string, string) string {
	return ", trên "
}

func (ViLocale) daysOfTheWeek() []string {
	return []string{"Chủ nhật", "Thứ 2", "Thứ 3", "Thứ 4", "Thứ 5", "Thứ 6", "Thứ 7"}
}

func (ViLocale) monthsOfTheYear() []string {
	return []string{"Tháng 1", "Tháng 2", "Tháng 3", "Tháng 4", "Tháng 5", "Tháng 6", "Tháng 7", "Tháng 8", "Tháng 9", "Tháng 10", "Tháng 11", "Tháng 12"}
}
