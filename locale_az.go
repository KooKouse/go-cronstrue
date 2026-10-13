package cronstrue

// Code generated from cronstrue/src/i18n/locales/az.ts (upstream v3.32.0)
// and verified against the JavaScript implementation. Edit with care.

func init() {
	registerLocale("az", AzLocale{enLocale{}})
}

// AzLocale is the az locale.
type AzLocale struct {
	enLocale
}

func (AzLocale) use24HourTimeFormatByDefault() bool {
	return true
}

func (AzLocale) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "İfadənin təsviri yaradılarkən xəta baş verdi. Cron ifadəsinin sintaksisini yoxlayın."
}

func (AzLocale) everyMinute() string {
	return "hər dəqiqə"
}

func (AzLocale) everyHour() string {
	return "hər saat"
}

func (AzLocale) atSpace() string {
	return "Saat "
}

func (AzLocale) everyMinuteBetweenX0AndX1() string {
	return "Saat %s ilə %s arasında hər dəqiqə"
}

func (AzLocale) at() string {
	return "Saat"
}

func (AzLocale) spaceAnd() string {
	return " və"
}

func (AzLocale) everySecond() string {
	return "hər saniyə"
}

func (AzLocale) secondsX0ThroughX1PastTheMinute() string {
	return "dəqiqənin %s ilə %s saniyələri arasında"
}

func (AzLocale) minutesX0ThroughX1PastTheHour() string {
	return "saatın %s ilə %s dəqiqələri arasında"
}

func (AzLocale) betweenX0AndX1() string {
	return "saat %s ilə %s arasında"
}

func (AzLocale) atX0() string {
	return "saat %s"
}

func (AzLocale) commaEveryDay() string {
	return ", hər gün"
}

func (AzLocale) spaceX0OfTheMonth() string {
	return " %s günü"
}

func (AzLocale) lastDay() string {
	return "son gün"
}

func (AzLocale) commaAndOnX0() string {
	return ", həmçinin %s"
}

func (AzLocale) commaOnlyInX0() string {
	return ", yalnız %s"
}

func (AzLocale) commaOnTheLastDayOfTheMonth() string {
	return ", ayın son günü"
}

func (AzLocale) commaOnTheLastWeekdayOfTheMonth() string {
	return ", ayın son iş günü"
}

func (AzLocale) firstWeekday() string {
	return "ilk iş günü"
}

func (AzLocale) weekdayNearestDayX0() string {
	return "%s tarixinə ən yaxın iş günü"
}

func (AzLocale) commaOnTheX0OfTheMonth() string {
	return ", ayın %s"
}

func (AzLocale) commaStartingX0() string {
	return ", %s başlamaqla"
}

func (AzLocale) commaMonthX0ThroughMonthX1() (string, bool) {
	return ", %s %s qədər", true
}

func (AzLocale) commaYearX0ThroughYearX1() (string, bool) {
	return ", %s ilindən %s ilinə qədər", true
}

func (AzLocale) commaOnlyInYearX0() (string, bool) {
	return ", yalnız %s ili", true
}

func (AzLocale) atReboot() (string, bool) {
	return "Başlanğıcda bir dəfə icra olunur", true
}

func (AzLocale) onTheHour() (string, bool) {
	return "saat başında", true
}

func (AzLocale) everyX0Seconds(string) string {
	return "hər %s saniyədən bir"
}

func (AzLocale) everyX0Minutes(string) string {
	return "hər %s dəqiqədən bir"
}

func (AzLocale) everyX0Hours(string) string {
	return "hər %s saatdan bir"
}

func (AzLocale) commaEveryX0DaysOfTheWeek(string) string {
	return ", həftənin hər %s günündən bir"
}

func (AzLocale) commaEveryX0Months(string) string {
	return ", hər %s aydan bir"
}

func (AzLocale) commaEveryX0Days(string) string {
	return ", ayda hər %s gündən bir"
}

func (AzLocale) commaEveryX0Years(string) string {
	return ", hər %s ildən bir"
}

func (AzLocale) atX0SecondsPastTheMinute(s string) string {
	switch s {
	case "0", "6", "16", "26", "36", "40", "46", "56", "60", "66", "70", "76", "80", "86", "90", "96":
		return "dəqiqənin %s-cı saniyəsində"
	case "1", "2", "5", "7", "8", "11", "12", "15", "17", "18", "20", "21", "22", "25", "27", "28", "31", "32", "35", "37", "38", "41", "42", "45", "47", "48", "50", "51", "52", "55", "57", "58", "61", "62", "65", "67", "68", "71", "72", "75", "77", "78", "81", "82", "85", "87", "88", "91", "92", "95", "97", "98":
		return "dəqiqənin %s-ci saniyəsində"
	case "3", "4", "13", "14", "23", "24", "33", "34", "43", "44", "53", "54", "63", "64", "73", "74", "83", "84", "93", "94":
		return "dəqiqənin %s-cü saniyəsində"
	case "9", "10", "19", "29", "30", "39", "49", "59", "69", "79", "89", "99":
		return "dəqiqənin %s-cu saniyəsində"
	default:
		return "dəqiqənin %s saniyələrində"
	}
}

func (AzLocale) atX0MinutesPastTheHour(s string) string {
	switch s {
	case "0", "6", "16", "26", "36", "40", "46", "56", "60", "66", "70", "76", "80", "86", "90", "96":
		return "saatın %s-cı dəqiqəsində"
	case "1", "2", "5", "7", "8", "11", "12", "15", "17", "18", "20", "21", "22", "25", "27", "28", "31", "32", "35", "37", "38", "41", "42", "45", "47", "48", "50", "51", "52", "55", "57", "58", "61", "62", "65", "67", "68", "71", "72", "75", "77", "78", "81", "82", "85", "87", "88", "91", "92", "95", "97", "98":
		return "saatın %s-ci dəqiqəsində"
	case "3", "4", "13", "14", "23", "24", "33", "34", "43", "44", "53", "54", "63", "64", "73", "74", "83", "84", "93", "94":
		return "saatın %s-cü dəqiqəsində"
	case "9", "10", "19", "29", "30", "39", "49", "59", "69", "79", "89", "99":
		return "saatın %s-cu dəqiqəsində"
	default:
		return "saatın %s dəqiqələrində"
	}
}

func (AzLocale) commaX0ThroughX1(string) string {
	return ", %s %s qədər"
}

func (AzLocale) commaAndX0ThroughX1(string) string {
	return ", %s %s qədər"
}

func (AzLocale) commaOnTheLastX0OfTheMonth(string) string {
	return ", ayın son %s günü"
}

func (AzLocale) commaOnDayX0OfTheMonth(s string) string {
	switch s {
	case "0", "6", "16", "26", "36", "40", "46", "56", "60", "66", "70", "76", "80", "86", "90", "96":
		return ", ayın %s-cı günü"
	case "1", "2", "5", "7", "8", "11", "12", "15", "17", "18", "20", "21", "22", "25", "27", "28", "31", "32", "35", "37", "38", "41", "42", "45", "47", "48", "50", "51", "52", "55", "57", "58", "61", "62", "65", "67", "68", "71", "72", "75", "77", "78", "81", "82", "85", "87", "88", "91", "92", "95", "97", "98":
		return ", ayın %s-ci günü"
	case "3", "4", "13", "14", "23", "24", "33", "34", "43", "44", "53", "54", "63", "64", "73", "74", "83", "84", "93", "94":
		return ", ayın %s-cü günü"
	case "9", "10", "19", "29", "30", "39", "49", "59", "69", "79", "89", "99":
		return ", ayın %s-cu günü"
	default:
		return ", ayın %s günləri"
	}
}

func (AzLocale) first(string) string {
	return "birinci"
}

func (AzLocale) second(string) string {
	return "ikinci"
}

func (AzLocale) third(string) string {
	return "üçüncü"
}

func (AzLocale) fourth(string) string {
	return "dördüncü"
}

func (AzLocale) fifth(string) string {
	return "beşinci"
}

func (AzLocale) commaBetweenDayX0AndX1OfTheMonth(string) string {
	return ", ayın %s ilə %s günləri arasında"
}

func (AzLocale) commaOnlyOnX0(string) string {
	return ", yalnız %s"
}

func (AzLocale) commaDaysBeforeTheLastDayOfTheMonth(string) string {
	return ", ayın son günündən %s gün əvvəl"
}

func (AzLocale) commaOnThe(string, string) string {
	return ", ayın "
}

func (AzLocale) daysOfTheWeekInCase(form int) ([]string, bool) {
	switch form {
	case 1:
		return []string{"bazardan", "bazar ertəsindən", "çərşənbə axşamından", "çərşənbədən", "cümə axşamından", "cümədən", "şənbədən"}, true
	case 2:
		return []string{"bazara", "bazar ertəsinə", "çərşənbə axşamına", "çərşənbəyə", "cümə axşamına", "cüməyə", "şənbəyə"}, true
	default:
		return []string{"bazar", "bazar ertəsi", "çərşənbə axşamı", "çərşənbə", "cümə axşamı", "cümə", "şənbə"}, true
	}
}

func (AzLocale) monthsOfTheYearInCase(form int) ([]string, bool) {
	switch form {
	case 1:
		return []string{"yanvardan", "fevraldan", "martdan", "apreldən", "maydan", "iyundan", "iyuldan", "avqustdan", "sentyabrdan", "oktyabrdan", "noyabrdan", "dekabrdan"}, true
	case 2:
		return []string{"yanvara", "fevrala", "marta", "aprelə", "maya", "iyuna", "iyula", "avqusta", "sentyabra", "oktyabra", "noyabra", "dekabra"}, true
	default:
		return []string{"yanvara", "fevrala", "marta", "aprelə", "maya", "iyuna", "iyula", "avqusta", "sentyabra", "oktyabra", "noyabra", "dekabra"}, true
	}
}

func (AzLocale) daysOfTheWeek() []string {
	return []string{"bazar", "bazar ertəsi", "çərşənbə axşamı", "çərşənbə", "cümə axşamı", "cümə", "şənbə"}
}

func (AzLocale) monthsOfTheYear() []string {
	return []string{"yanvarda", "fevralda", "martda", "apreldə", "mayda", "iyunda", "iyulda", "avqustda", "sentyabrda", "oktyabrda", "noyabrda", "dekabrda"}
}
