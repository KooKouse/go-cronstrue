package cronstrue

// Code generated from cronstrue/src/i18n/locales/th.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("th", ThLocale{enLocale{}})
}

// ThLocale is the th locale.
type ThLocale struct {
	enLocale
}

func (ThLocale) use24HourTimeFormatByDefault() bool {
	return false
}

func (ThLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "เกิดข้อผิดพลาดขณะสร้างคำอธิบายนิพจน์ ตรวจสอบไวยากรณ์นิพจน์ครอน"
}

func (ThLocale) everyMinute() string {
	return "ทุกๆ นาที"
}

func (ThLocale) everyHour() string {
	return "ทุกๆ ชั่วโมง"
}

func (ThLocale) atSpace() string {
	return "เมื่อ "
}

func (ThLocale) everyMinuteBetweenX0AndX1() string {
	return "ทุกๆ นาที %s และ %s"
}

func (ThLocale) at() string {
	return "เมื่อ"
}

func (ThLocale) spaceAnd() string {
	return " และ"
}

func (ThLocale) everySecond() string {
	return "ทุกๆ วินาที"
}

func (ThLocale) secondsX0ThroughX1PastTheMinute() string {
	return "วินาที %s ถึง %s นาทีที่ผ่านมา"
}

func (ThLocale) minutesX0ThroughX1PastTheHour() string {
	return "นาที %s ถึง %s ชั่วโมงที่ผ่านมา"
}

func (ThLocale) betweenX0AndX1() string {
	return "ระหว่าง %s ถึง %s"
}

func (ThLocale) atX0() string {
	return "เมื่อ %s"
}

func (ThLocale) commaEveryDay() string {
	return ", ทุกๆ วัน"
}

func (ThLocale) spaceX0OfTheMonth() string {
	return " %s ของเดือน"
}

func (ThLocale) lastDay() string {
	return "วันสุดท้าย"
}

func (ThLocale) commaAndOnX0() string {
	return ", และใน %s"
}

func (ThLocale) commaOnlyInX0() string {
	return ", เท่านั้น %s"
}

func (ThLocale) commaOnTheLastDayOfTheMonth() string {
	return ", ในวันสิ้นเดือน"
}

func (ThLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", ในวันธรรมดาสุดท้ายของเดือน"
}

func (ThLocale) firstWeekday() string {
	return "วันธรรมดาวันแรก"
}

func (ThLocale) weekdayNearestDayX0() string {
	return "วันธรรมดาที่ใกล้ที่สุด %s"
}

func (ThLocale) commaOnTheX0OfTheMonth() string {
	return ", ในวัน %s ของเดือน"
}

func (ThLocale) commaStartingX0() string {
	return ", เริ่ม %s"
}

func (ThLocale) atReboot() (string, bool) {
	return "ทำงานครั้งเดียว เมื่อเริ่มระบบ", true
}

func (ThLocale) onTheHour() (string, bool) {
	return "ตรงชั่วโมง", true
}

func (ThLocale) everyX0Seconds(string) string {
	return "ทุกๆ %s วินาที"
}

func (ThLocale) everyX0Minutes(string) string {
	return "ทุกๆ %s นาที"
}

func (ThLocale) everyX0Hours(string) string {
	return "ทุกๆ %s ชั่วโมง"
}

func (ThLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", ทุกๆ %s วันของสัปดาห์"
}

func (ThLocale) commaEveryX0Months(string) string {
	return ", ทุกๆ %s เดือน"
}

func (ThLocale) commaEveryX0Days(string) string {
	return ", ทุกๆ %s วัน"
}

func (ThLocale) commaEveryX0Years(string) string {
	return ", ทุกๆ %s ปี"
}

func (ThLocale) atX0SecondsPastTheMinute(string) string {
	return "เมื่อ %s วินาที นาทีที่ผ่านมา"
}

func (ThLocale) atX0MinutesPastTheHour(string) string {
	return "เมื่อ %s นาที ชั่วโมงที่ผ่านมา"
}

func (ThLocale) commaX0ThroughX1(string) string {
	return ", %s ถึง %s"
}

func (ThLocale) commaAndX0ThroughX1(string) string {
	return ", %s ถึง %s"
}

func (ThLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", ณ สุดท้าย %s ของเดือน"
}

func (ThLocale) commaOnDayX0OfTheMonth(string) string {
	return ", ในวัน %s ของเดือน"
}

func (ThLocale) first(string) string {
	return "แรก"
}

func (ThLocale) second(string) string {
	return "ที่สอง"
}

func (ThLocale) third(string) string {
	return "ที่สาม"
}

func (ThLocale) fourth(string) string {
	return "ที่สี่"
}

func (ThLocale) fifth(string) string {
	return "ที่ห้า"
}

func (ThLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", ระหว่างวัน %s และ %s ของเดือน"
}

func (ThLocale) commaOnlyOnX0(string) string {
	return ", เท่านั้น %s"
}

func (ThLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", %s วันก่อนวันสุดท้ายของเดือน"
}

func (ThLocale) commaOnThe(string, string) string {
	return ", ในวัน "
}

func (ThLocale) daysOfTheWeek() []string {
	return []string{"วันอาทิตย์", "วันจันทร์", "วันอังคาร", "วันพุธ", "วันพฤหัสบดี", "วันศุกร์", "วันเสาร์"}
}

func (ThLocale) monthsOfTheYear() []string {
	return []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}
}
