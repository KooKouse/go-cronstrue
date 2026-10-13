package cronstrue

// Code generated from cronstrue/src/i18n/locales/ko.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("ko", KoLocale{enLocale{}})
}

// KoLocale is the ko locale.
type KoLocale struct {
	enLocale
}

func (KoLocale) setPeriodBeforeTime() bool {
	return true
}

func (KoLocale) use24HourTimeFormatByDefault() bool {
	return false
}

func (KoLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "표현식 설명을 생성하는 중 오류가 발생했습니다. cron 표현식 구문을 확인하십시오."
}

func (KoLocale) everyMinute() string {
	return "1분마다"
}

func (KoLocale) everyHour() string {
	return "1시간마다"
}

func (KoLocale) atSpace() string {
	return ""
}

func (KoLocale) everyMinuteBetweenX0AndX1() string {
	return "매분 %s~%s"
}

func (KoLocale) at() string {
	return ""
}

func (KoLocale) spaceAnd() string {
	return ","
}

func (KoLocale) everySecond() string {
	return "1초마다"
}

func (KoLocale) secondsX0ThroughX1PastTheMinute() string {
	return "%s~%s초"
}

func (KoLocale) minutesX0ThroughX1PastTheHour() string {
	return "%s~%s분"
}

func (KoLocale) betweenX0AndX1() string {
	return "%s부터 %s까지"
}

func (KoLocale) atX0() string {
	return "%s에서"
}

func (KoLocale) commaEveryDay() string {
	return ", 매일"
}

func (KoLocale) spaceX0OfTheMonth() string {
	return " %s"
}

func (KoLocale) lastDay() string {
	return "마지막 날"
}

func (KoLocale) commaAndOnX0() string {
	return ", %s에"
}

func (KoLocale) commaOnlyInX0() string {
	return ", %s에만"
}

func (KoLocale) commaOnTheLastDayOfTheMonth() string {
	return ", 해당 월의 마지막 날에"
}

func (KoLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", 매월 마지막 평일"
}

func (KoLocale) firstWeekday() string {
	return "첫 번째 평일"
}

func (KoLocale) weekdayNearestDayX0() string {
	return "%s일과 가장 가까운 평일"
}

func (KoLocale) commaOnTheX0OfTheMonth() string {
	return ", 매월 %s"
}

func (KoLocale) commaStartingX0() string {
	return ", %s부터"
}

func (KoLocale) atReboot() (string, bool) {
	return "시작 시 한 번 실행", true
}

func (KoLocale) onTheHour() (string, bool) {
	return "정각", true
}

func (KoLocale) pm() (string, bool) {
	return "오후", true
}

func (KoLocale) am() (string, bool) {
	return "오전", true
}

func (KoLocale) everyX0Seconds(string) string {
	return "%s초마다"
}

func (KoLocale) everyX0Minutes(string) string {
	return "%s분마다"
}

func (KoLocale) everyX0Hours(string) string {
	return "%s시간마다"
}

func (KoLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", 주 중 %s일마다"
}

func (KoLocale) commaEveryX0Months(string) string {
	return ", %s개월마다"
}

func (KoLocale) commaEveryX0Days(string) string {
	return ", %s일마다"
}

func (KoLocale) commaEveryX0Years(string) string {
	return ", %s년마다"
}

func (KoLocale) atX0SecondsPastTheMinute(string) string {
	return "%s초"
}

func (KoLocale) atX0MinutesPastTheHour(string) string {
	return "%s분"
}

func (KoLocale) commaX0ThroughX1(string) string {
	return ", %s부터 %s까지"
}

func (KoLocale) commaAndX0ThroughX1(string) string {
	return ", %s부터 %s까지"
}

func (KoLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", 해당 월의 마지막 %s"
}

func (KoLocale) commaOnDayX0OfTheMonth(string) string {
	return ", 매월 %s일"
}

func (KoLocale) first(string) string {
	return "첫 번째"
}

func (KoLocale) second(string) string {
	return "두 번째"
}

func (KoLocale) third(string) string {
	return "세 번째"
}

func (KoLocale) fourth(string) string {
	return "네 번째"
}

func (KoLocale) fifth(string) string {
	return "다섯 번째"
}

func (KoLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", 해당 월의 %s일에서 %s일까지"
}

func (KoLocale) commaOnlyOnX0(string) string {
	return ", %s에만"
}

func (KoLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", 해당 월의 마지막 날 %s일 전"
}

func (KoLocale) commaOnThe(string, string) string {
	return ", 매월 "
}

func (KoLocale) daysOfTheWeek() []string {
	return []string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"}
}

func (KoLocale) monthsOfTheYear() []string {
	return []string{"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"}
}
