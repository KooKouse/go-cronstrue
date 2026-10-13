package cronstrue

// Code generated from cronstrue upstream test suite (test/cronstrue.ts,
// test/i18n.ts). golden outputs are the JavaScript library's outputs.

import "strings"
import "testing"

func TestToString(t *testing.T) {
	tests := []struct {
		name string
		expr string
		opts []Option
		want string
	}{
		{name: "* * * * * *", expr: "* * * * * *", want: "Every second"},
		{name: "* * * * *", expr: "* * * * *", want: "Every minute"},
		{name: "0 * * * * (verbose)", expr: "0 * * * *", opts: []Option{WithVerbose(true)}, want: "Every hour, every day"},
		{name: "* * * * * (verbose)", expr: "* * * * *", opts: []Option{WithVerbose(true)}, want: "Every minute, every hour, every day"},
		{name: "*/1 * * * *", expr: "*/1 * * * *", want: "Every minute"},
		{name: "*/5 * * * *", expr: "*/5 * * * *", want: "Every 5 minutes"},
		{name: "0 0/1 * * * ?", expr: "0 0/1 * * * ?", want: "Every minute"},
		{name: "0 0 * * * ?", expr: "0 0 * * * ?", want: "Every hour"},
		{name: "0 0 0/1 * * ?", expr: "0 0 0/1 * * ?", want: "Every hour"},
		{name: "* * * 3 *", expr: "* * * 3 *", want: "Every minute, only in March"},
		{name: "* * * 3,6 *", expr: "* * * 3,6 *", want: "Every minute, only in March and June"},
		{name: "* * * * * * 2013", expr: "* * * * * * 2013", want: "Every second, only in 2013"},
		{name: "* * * * * 2013", expr: "* * * * * 2013", want: "Every minute, only in 2013"},
		{name: "* * * * * 2013,2014", expr: "* * * * * 2013,2014", want: "Every minute, only in 2013 and 2014"},
		{name: "*/45 * * * * *", expr: "*/45 * * * * *", want: "Every 45 seconds"},
		{name: "*/5 * * * *", expr: "*/5 * * * *", want: "Every 5 minutes"},
		{name: "*/10 * * * *", expr: "*/10 * * * *", want: "Every 10 minutes"},
		{name: "0 */5 * * * *", expr: "0 */5 * * * *", want: "Every 5 minutes"},
		{name: "0 9-17 * * *", expr: "0 9-17 * * *", want: "Every hour, between 09:00 AM and 05:00 PM"},
		{name: "0 * ? * 2/1 *", expr: "0 * ? * 2/1 *", want: "Every hour, Tuesday through Saturday"},
		{name: "0 * ? * 1/1", expr: "0 * ? * 1/1", want: "Every hour, Monday through Saturday"},
		{name: "0 * ? * 2/1", expr: "0 * ? * 2/1", want: "Every hour, Tuesday through Saturday"},
		{name: "0 52 13 ? * 3/1", expr: "0 52 13 ? * 3/1", want: "At 01:52 PM, Wednesday through Saturday"},
		{name: "0 23 ? * MON-FRI", expr: "0 23 ? * MON-FRI", want: "At 11:00 PM, Monday through Friday"},
		{name: "30 11 * * 1-5", expr: "30 11 * * 1-5", want: "At 11:30 AM, Monday through Friday"},
		{name: "0-10 11 * * *", expr: "0-10 11 * * *", want: "Every minute between 11:00 AM and 11:10 AM"},
		{name: "23 12 * Jan-Mar *", expr: "23 12 * Jan-Mar *", want: "At 12:23 PM, January through March"},
		{name: "23 12 * JAN-FEB *", expr: "23 12 * JAN-FEB *", want: "At 12:23 PM, January through February"},
		{name: "1 1,3-4 * * *", expr: "1 1,3-4 * * *", want: "At 1 minutes past the hour, at 01:00 AM and 03:00 AM through 04:59 AM"},
		{name: "* 0 */4 * * *", expr: "* 0 */4 * * *", want: "Every second, on the hour, every 4 hours"},
		{name: "*/10 0 * * * *", expr: "*/10 0 * * * *", want: "Every 10 seconds, on the hour"},
		{name: "* 0 0 * * *", expr: "* 0 0 * * *", want: "Every second, on the hour, between 12:00 AM and 12:00 AM"},
		{name: "* 0 * * *", expr: "* 0 * * *", want: "Every minute, between 12:00 AM and 12:59 AM"},
		{name: "* 0 * * * *", expr: "* 0 * * * *", want: "Every second, on the hour"},
		{name: "30 11 * * *", expr: "30 11 * * *", opts: []Option{With24HourTimeFormat(true)}, want: "At 11:30"},
		{name: "31 10 * * *", expr: "31 10 * * *", opts: []Option{With24HourTimeFormat(true)}, want: "At 10:31"},
		{name: "29 10 * * *", expr: "29 10 * * *", want: "At 10:29 AM"},
		{name: "30 10 * * *", expr: "30 10 * * *", want: "At 10:30 AM"},
		{name: "31 10 * * *", expr: "31 10 * * *", want: "At 10:31 AM"},
		{name: "29 10 * * *", expr: "29 10 * * *", want: "At 10:29 AM"},
		{name: "30 10 * * *", expr: "30 10 * * *", want: "At 10:30 AM"},
		{name: "30 11 * * *", expr: "30 11 * * *", want: "At 11:30 AM"},
		{name: "30 3 * * *", expr: "30 3 * * *", want: "At 03:30 AM"},
		{name: "30 10 * * *", expr: "30 10 * * *", opts: []Option{WithTrimHoursLeadingZero(true)}, want: "At 10:30 AM"},
		{name: "31 10 * * *", expr: "31 10 * * *", opts: []Option{WithTrimHoursLeadingZero(true), With24HourTimeFormat(true)}, want: "At 10:31"},
		{name: "29 9 * * *", expr: "29 9 * * *", opts: []Option{WithTrimHoursLeadingZero(true)}, want: "At 9:29 AM"},
		{name: "30 9 * * *", expr: "30 9 * * *", opts: []Option{WithTrimHoursLeadingZero(true), With24HourTimeFormat(true)}, want: "At 9:30"},
		{name: "0 0 * * *", expr: "0 0 * * *", opts: []Option{WithTrimHoursLeadingZero(true), With24HourTimeFormat(true)}, want: "At 0:00"},
		{name: "10 11 * * *", expr: "10 11 * * *", opts: []Option{WithVerbose(true)}, want: "At 11:10 AM, every day"},
		{name: "30 22 * * *", expr: "30 22 * * *", opts: []Option{With24HourTimeFormat(true)}, want: "At 22:30"},
		{name: "5 1 * * 1", expr: "5 1 * * 1", want: "At 01:05 AM, only on Monday"},
		{name: "5 23 * * 1", expr: "5 23 * * 1", want: "At 11:05 PM, only on Monday"},
		{name: "30 11 * * *", expr: "30 11 * * *", want: "At 11:30 AM"},
		{name: "23 12 * * SUN", expr: "23 12 * * SUN", want: "At 12:23 PM, only on Sunday"},
		{name: "30 02 14 * * *", expr: "30 02 14 * * *", want: "At 02:02:30 PM"},
		{name: "0 0 6 1/1 * ?", expr: "0 0 6 1/1 * ?", want: "At 06:00 AM"},
		{name: "0 5 0/1 * * ?", expr: "0 5 0/1 * * ?", want: "At 5 minutes past the hour"},
		{name: "46 9 * * 1", expr: "46 9 * * 1", want: "At 09:46 AM, only on Monday"},
		{name: "23 12 15 * *", expr: "23 12 15 * *", want: "At 12:23 PM, on day 15 of the month"},
		{name: "23 12 * JAN *", expr: "23 12 * JAN *", want: "At 12:23 PM, only in January"},
		{name: "23 12 ? JAN *", expr: "23 12 ? JAN *", want: "At 12:23 PM, only in January"},
		{name: "30 14,16 * * *", expr: "30 14,16 * * *", want: "At 02:30 PM and 04:30 PM"},
		{name: "30 6,14,16 * * *", expr: "30 6,14,16 * * *", want: "At 06:30 AM, 02:30 PM and 04:30 PM"},
		{name: "0 * 31 * 1", expr: "0 * 31 * 1", want: "Every hour, on day 31 of the month, and on Monday"},
		{name: "0 * 31 * 2", expr: "0 * 31 * 2", opts: []Option{WithLogicalAndDayFields(true)}, want: "Every hour, on day 31 of the month, only on Tuesday"},
		{name: "0 0 0 1,2,3 * WED,FRI", expr: "0 0 0 1,2,3 * WED,FRI", opts: []Option{WithLogicalAndDayFields(true)}, want: "At 12:00 AM, on day 1, 2, and 3 of the month, only on Wednesday and Friday"},
		{name: "0 2,16 1,8,15,22 * 1,2", expr: "0 2,16 1,8,15,22 * 1,2", opts: []Option{WithLogicalAndDayFields(true)}, want: "At 02:00 AM and 04:00 PM, on day 1, 8, 15, and 22 of the month, only on Monday and Tuesday"},
		{name: "0 45 12 22,17,6,30,26 * *", expr: "0 45 12 22,17,6,30,26 * *", want: "At 12:45 PM, on day 6, 17, 22, 26, and 30 of the month"},
		{name: "* * LW * *", expr: "* * LW * *", want: "Every minute, on the last weekday of the month"},
		{name: "* * WL * *", expr: "* * WL * *", want: "Every minute, on the last weekday of the month"},
		{name: "* * 1W * *", expr: "* * 1W * *", want: "Every minute, on the first weekday of the month"},
		{name: "* * 13W * *", expr: "* * 13W * *", want: "Every minute, on the weekday nearest day 13 of the month"},
		{name: "* * W1 * *", expr: "* * W1 * *", want: "Every minute, on the first weekday of the month"},
		{name: "* * 5W * *", expr: "* * 5W * *", want: "Every minute, on the weekday nearest day 5 of the month"},
		{name: "* * W5 * *", expr: "* * W5 * *", want: "Every minute, on the weekday nearest day 5 of the month"},
		{name: "0 0 1W-2 * *", expr: "0 0 1W-2 * *", want: "At 12:00 AM, 2 days before the first weekday of the month"},
		{name: "0 0 15W+3 * *", expr: "0 0 15W+3 * *", want: "At 12:00 AM, 3 days after the weekday nearest day 15 of the month"},
		{name: "* * * * 4L", expr: "* * * * 4L", want: "Every minute, on the last Thursday of the month"},
		{name: "*/5 * L JAN *", expr: "*/5 * L JAN *", want: "Every 5 minutes, on the last day of the month, only in January"},
		{name: "0 20 15,L * *", expr: "0 20 15,L * *", want: "At 08:00 PM, on day 15 and the last day of the month"},
		{name: "0 20 1-10,20-L * *", expr: "0 20 1-10,20-L * *", want: "At 08:00 PM, on day 1 through 10 and 20 through the last day of the month"},
		{name: "0 15 10 * * L", expr: "0 15 10 * * L", want: "At 10:15 AM, only on Saturday"},
		{name: "0 15 10 L * *", expr: "0 15 10 L * *", want: "At 10:15 AM, on the last day of the month"},
		{name: "0 0 0 L-5 * ?", expr: "0 0 0 L-5 * ?", want: "At 12:00 AM, 5 days before the last day of the month"},
		{name: "23 12 * * 1#2", expr: "23 12 * * 1#2", opts: []Option{WithDayOfWeekStartIndexZero(false)}, want: "At 12:23 PM, on the second Sunday of the month"},
		{name: "* * * ? * 2-6/2", expr: "* * * ? * 2-6/2", opts: []Option{WithDayOfWeekStartIndexZero(false)}, want: "Every second, every 2 days of the week, Monday through Friday"},
		{name: "* * * ? * 7", expr: "* * * ? * 7", opts: []Option{WithDayOfWeekStartIndexZero(false)}, want: "Every second, only on Saturday"},
		{name: "* * * ? * 1,2,3,4,5", expr: "* * * ? * 1,2,3,4,5", opts: []Option{WithDayOfWeekStartIndexZero(false)}, want: "Every second, only on Sunday, Monday, Tuesday, Wednesday, and Thursday"},
		{name: "0 * ? * 1/1", expr: "0 * ? * 1/1", opts: []Option{WithDayOfWeekStartIndexZero(false)}, want: "Every hour, Sunday through Saturday"},
		{name: "0 * ? * 2/1", expr: "0 * ? * 2/1", opts: []Option{WithDayOfWeekStartIndexZero(false)}, want: "Every hour, Monday through Saturday"},
		{name: "* * * 7 *", expr: "* * * 7 *", opts: []Option{WithMonthStartIndexZero(true)}, want: "Every minute, only in August"},
		{name: "30 * * 6-8 *", expr: "30 * * 6-8 *", opts: []Option{WithMonthStartIndexZero(true)}, want: "At 30 minutes past the hour, July through September"},
		{name: "30 * * 1-10/2 *", expr: "30 * * 1-10/2 *", opts: []Option{WithMonthStartIndexZero(true)}, want: "At 30 minutes past the hour, every 2 months, February through November"},
		{name: "30 * * 4,5,6 *", expr: "30 * * 4,5,6 *", opts: []Option{WithMonthStartIndexZero(true)}, want: "At 30 minutes past the hour, only in May, June, and July"},
		{name: "30 * * */10 *", expr: "30 * * */10 *", opts: []Option{WithMonthStartIndexZero(true)}, want: "At 30 minutes past the hour, every 10 months"},
		{name: "30 * * 2/11 *", expr: "30 * * 2/11 *", opts: []Option{WithMonthStartIndexZero(true)}, want: "At 30 minutes past the hour, every 11 months, March through December"},
		{name: "30 * * 1/2 *", expr: "30 * * 1/2 *", opts: []Option{WithMonthStartIndexZero(true)}, want: "At 30 minutes past the hour, every 2 months, February through December"},
		{name: "30 * * JAN *", expr: "30 * * JAN *", opts: []Option{WithMonthStartIndexZero(true)}, want: "At 30 minutes past the hour, only in January"},
		{name: "*/5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Every 5 minutes, between 03:00 PM and 03:59 PM, Monday through Friday"},
		{name: "* * * * MON#3", expr: "* * * * MON#3", want: "Every minute, on the third Monday of the month"},
		{name: "* * * * MON#3,THU#1", expr: "* * * * MON#3,THU#1", want: "Every minute, on the third Monday and first Thursday of the month"},
		{name: "5-10 * * * * *", expr: "5-10 * * * * *", want: "Seconds 5 through 10 past the minute"},
		{name: "5-10 30-35 10-12 * * *", expr: "5-10 30-35 10-12 * * *", want: "Seconds 5 through 10 past the minute, minutes 30 through 35 past the hour, between 10:00 AM and 12:59 PM"},
		{name: "30 */5 * * * *", expr: "30 */5 * * * *", want: "At 30 seconds past the minute, every 5 minutes"},
		{name: "10 0/5 * * * ?", expr: "10 0/5 * * * ?", want: "At 10 seconds past the minute, every 5 minutes"},
		{name: "2-59/3 1,9,22 11-26 1-6 ?", expr: "2-59/3 1,9,22 11-26 1-6 ?", want: "Every 3 minutes, minutes 2 through 59 past the hour, at 01:00 AM, 09:00 AM, and 10:00 PM, between day 11 and 26 of the month, January through June"},
		{name: "23 12 * JAN-FEB * 2013-2014", expr: "23 12 * JAN-FEB * 2013-2014", want: "At 12:23 PM, January through February, 2013 through 2014"},
		{name: "23 12 * JAN-MAR * 2013-2015", expr: "23 12 * JAN-MAR * 2013-2015", want: "At 12:23 PM, January through March, 2013 through 2015"},
		{name: "12-50 0-10 6 * * * 2022", expr: "12-50 0-10 6 * * * 2022", want: "Seconds 12 through 50 past the minute, minutes 0 through 10 past the hour, at 06:00 AM, only in 2022"},
		{name: "0 0/30 8-9 5,20 * ?", expr: "0 0/30 8-9 5,20 * ?", want: "Every 30 minutes, between 08:00 AM and 09:59 AM, on day 5 and 20 of the month"},
		{name: "23 12 * * SUN#2", expr: "23 12 * * SUN#2", want: "At 12:23 PM, on the second Sunday of the month"},
		{name: "0 25 7-19/8 ? * *", expr: "0 25 7-19/8 ? * *", want: "At 25 minutes past the hour, every 8 hours, between 07:00 AM and 07:59 PM"},
		{name: "0 25 7-20/13 ? * *", expr: "0 25 7-20/13 ? * *", want: "At 25 minutes past the hour, every 13 hours, between 07:00 AM and 08:59 PM"},
		{name: "0 0 8 1/3 * ? *", expr: "0 0 8 1/3 * ? *", want: "At 08:00 AM, every 3 days in a month"},
		{name: "0 0 14 */30 * *", expr: "0 0 14 */30 * *", want: "At 02:00 PM, every 30 days in a month"},
		{name: "0 15 10 ? * */3", expr: "0 15 10 ? * */3", want: "At 10:15 AM, every 3 days of the week"},
		{name: "* * * ? * 1-5/2", expr: "* * * ? * 1-5/2", want: "Every second, every 2 days of the week, Monday through Friday"},
		{name: "0 5 7 2 1/3 ? *", expr: "0 5 7 2 1/3 ? *", want: "At 07:05 AM, on day 2 of the month, every 3 months"},
		{name: "0 15 6 1 1 ? 1/2", expr: "0 15 6 1 1 ? 1/2", want: "At 06:15 AM, on day 1 of the month, only in January, every 2 years"},
		{name: "2,4-5 1 * * *", expr: "2,4-5 1 * * *", want: "At 2 and 4 through 5 minutes past the hour, at 01:00 AM"},
		{name: "2,26-28 18 * * *", expr: "2,26-28 18 * * *", want: "At 2 and 26 through 28 minutes past the hour, at 06:00 PM"},
		{name: "5/30 * * * * ?", expr: "5/30 * * * * ?", want: "Every 30 seconds, starting at 5 seconds past the minute"},
		{name: "0 5/30 * * * ?", expr: "0 5/30 * * * ?", want: "Every 30 minutes, starting at 5 minutes past the hour"},
		{name: "* * 5/8 * * ?", expr: "* * 5/8 * * ?", want: "Every second, every 8 hours, starting at 05:00 AM"},
		{name: "0 5 7 2/3 * ? *", expr: "0 5 7 2/3 * ? *", want: "At 07:05 AM, every 3 days in a month, starting on day 2 of the month"},
		{name: "0 5 7 ? 3/2 ? *", expr: "0 5 7 ? 3/2 ? *", want: "At 07:05 AM, every 2 months, March through December"},
		{name: "0 5 7 ? * 2/3 *", expr: "0 5 7 ? * 2/3 *", want: "At 07:05 AM, every 3 days of the week, Tuesday through Saturday"},
		{name: "0 5 7 ? * ? 2016/4", expr: "0 5 7 ? * ? 2016/4", want: "At 07:05 AM, every 4 years, 2016 through 9999"},
		{name: "0 30 10-13 ? * wed,FRI", expr: "0 30 10-13 ? * wed,FRI", want: "At 30 minutes past the hour, between 10:00 AM and 01:59 PM, only on Wednesday and Friday"},
		{name: "0 00 10 ? * MON-THU,SUN *", expr: "0 00 10 ? * MON-THU,SUN *", want: "At 10:00 AM, only on Monday through Thursday and Sunday"},
		{name: "0 0 0 1,2,3 * WED,FRI", expr: "0 0 0 1,2,3 * WED,FRI", want: "At 12:00 AM, on day 1, 2, and 3 of the month, and on Wednesday and Friday"},
		{name: "0 2,16 1,8,15,22 * 1,2", expr: "0 2,16 1,8,15,22 * 1,2", want: "At 02:00 AM and 04:00 PM, on day 1, 8, 15, and 22 of the month, and on Monday and Tuesday"},
		{name: "0 */4 * * *", expr: "0 */4 * * *", want: "On the hour, every 4 hours"},
		{name: "0 */4,6 * * * ", expr: "0 */4,6 * * * ", want: "On the hour, every 4 hours and at 06:00 AM"},
		{name: "5 30 6,14,16 5 * *", expr: "5 30 6,14,16 5 * *", want: "At 5 seconds past the minute, at 30 minutes past the hour, at 06:00 AM, 02:00 PM, and 04:00 PM, on day 5 of the month"},
		{name: "0-20/3 9 * * *", expr: "0-20/3 9 * * *", want: "Every 3 minutes, minutes 0 through 20 past the hour, between 09:00 AM and 09:59 AM"},
		{name: "5-45/10,*/5,9 * * * *", expr: "5-45/10,*/5,9 * * * *", want: "Every 10 minutes, minutes 5 through 45 past the hour, every 5 minutes, and at 9 minutes past the hour"},
		{name: "*/5 0-4,6-8 * * *", expr: "*/5 0-4,6-8 * * *", want: "Every 5 minutes, at 12:00 AM through 04:59 AM and 06:00 AM through 08:59 AM"},
		{name: "0 9 * * 0-4/2,5,6", expr: "0 9 * * 0-4/2,5,6", want: "At 09:00 AM, every 2 days of the week, Sunday through Thursday, only on Friday, and only on Saturday"},
		{name: "* 1-2,22-23 * * *", expr: "* 1-2,22-23 * * *", want: "Every minute, at 01:00 AM through 02:59 AM and 10:00 PM through 11:59 PM"},
		{name: "@yearly", expr: "@yearly", want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@annually", expr: "@annually", want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@monthly", expr: "@monthly", want: "At 12:00 AM, on day 1 of the month"},
		{name: "@weekly", expr: "@weekly", want: "At 12:00 AM, only on Sunday"},
		{name: "@daily", expr: "@daily", want: "At 12:00 AM"},
		{name: "@midnight", expr: "@midnight", want: "At 12:00 AM"},
		{name: "@hourly", expr: "@hourly", want: "Every hour"},
		{name: "@reboot", expr: "@reboot", want: "Run once, at startup"},
		{name: "30 4 1 * *", expr: "30 4 1 * *", opts: []Option{WithVerbose(true)}, want: "At 04:30 AM, on day 1 of the month"},
		{name: "0 13 * * 1", expr: "0 13 * * 1", opts: []Option{WithVerbose(true)}, want: "At 01:00 PM, only on Monday"},
		{name: "*/7 * * * *", expr: "*/7 * * * *", opts: []Option{WithVerbose(true)}, want: "Every 7 minutes, at 0, 7, 14, 21, 28, 35, 42, 49, and 56 minutes past the hour, every hour, every day"},
		{name: "*/59 * * * *", expr: "*/59 * * * *", opts: []Option{WithVerbose(true)}, want: "Every 59 minutes, at 0 and 59 minutes past the hour, every hour, every day"},
		{name: "0 */7 * * *", expr: "0 */7 * * *", opts: []Option{WithVerbose(true)}, want: "On the hour, every 7 hours, at 12:00 AM, 07:00 AM, 02:00 PM, and 09:00 PM, every day"},
		{name: "*/50 * * * * *", expr: "*/50 * * * * *", opts: []Option{WithVerbose(true)}, want: "Every 50 seconds, at 0 and 50 seconds past the minute, every minute, every hour, every day"},
		{name: "*/15 * * * *", expr: "*/15 * * * *", opts: []Option{WithVerbose(true)}, want: "Every 15 minutes, every hour, every day"},
		{name: "*/30 * * * *", expr: "*/30 * * * *", opts: []Option{WithVerbose(true)}, want: "Every 30 minutes, every hour, every day"},
		{name: "5-30/7 * * * *", expr: "5-30/7 * * * *", opts: []Option{WithVerbose(true)}, want: "Every 7 minutes, minutes 5 through 30 past the hour, every hour, every day"},
		{name: "garbage expression with option (throwExceptionOnParseError = false)", expr: "garbage", opts: []Option{WithThrowExceptionOnParseError(false)}, want: "An error occurred when generating the expression description. Check the cron expression syntax."},
		{name: "* * * * *", expr: "* * * * *", opts: []Option{WithLocale("zh_CN")}, want: "每分钟"},
		{name: "*/5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", opts: []Option{WithLocale("zh_CN")}, want: "每隔 5 分钟, 在 下午 03:00 和 下午 03:59 之间, 星期一至星期五"},
		{name: "@yearly dow=true mon=false", expr: "@yearly", opts: []Option{WithDayOfWeekStartIndexZero(true), WithMonthStartIndexZero(false)}, want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@yearly dow=true mon=true", expr: "@yearly", opts: []Option{WithDayOfWeekStartIndexZero(true), WithMonthStartIndexZero(true)}, want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@yearly dow=false mon=false", expr: "@yearly", opts: []Option{WithDayOfWeekStartIndexZero(false), WithMonthStartIndexZero(false)}, want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@yearly dow=false mon=true", expr: "@yearly", opts: []Option{WithDayOfWeekStartIndexZero(false), WithMonthStartIndexZero(true)}, want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@annually dow=true mon=false", expr: "@annually", opts: []Option{WithDayOfWeekStartIndexZero(true), WithMonthStartIndexZero(false)}, want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@annually dow=true mon=true", expr: "@annually", opts: []Option{WithDayOfWeekStartIndexZero(true), WithMonthStartIndexZero(true)}, want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@annually dow=false mon=false", expr: "@annually", opts: []Option{WithDayOfWeekStartIndexZero(false), WithMonthStartIndexZero(false)}, want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@annually dow=false mon=true", expr: "@annually", opts: []Option{WithDayOfWeekStartIndexZero(false), WithMonthStartIndexZero(true)}, want: "At 12:00 AM, on day 1 of the month, only in January"},
		{name: "@weekly dow=true mon=false", expr: "@weekly", opts: []Option{WithDayOfWeekStartIndexZero(true), WithMonthStartIndexZero(false)}, want: "At 12:00 AM, only on Sunday"},
		{name: "@weekly dow=true mon=true", expr: "@weekly", opts: []Option{WithDayOfWeekStartIndexZero(true), WithMonthStartIndexZero(true)}, want: "At 12:00 AM, only on Sunday"},
		{name: "@weekly dow=false mon=false", expr: "@weekly", opts: []Option{WithDayOfWeekStartIndexZero(false), WithMonthStartIndexZero(false)}, want: "At 12:00 AM, only on Sunday"},
		{name: "@weekly dow=false mon=true", expr: "@weekly", opts: []Option{WithDayOfWeekStartIndexZero(false), WithMonthStartIndexZero(true)}, want: "At 12:00 AM, only on Sunday"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.expr, tt.opts...)
			if err != nil {
				t.Fatalf("ToString(%q) returned error: %v", tt.expr, err)
			}
			if got != tt.want {
				t.Errorf("ToString(%q) = %q, want %q", tt.expr, got, tt.want)
			}
		})
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		opts    []Option
		wantMsg string
	}{
		{name: "second out of range", expr: "61 * * * * *", wantMsg: "seconds part must be >= 0 and <= 59"},
		{name: "minute out of range", expr: "0 -1 * * * *", wantMsg: "minutes part must be >= 0 and <= 59"},
		{name: "hour out of range", expr: "0 0 24 * * *", wantMsg: "hours part must be >= 0 and <= 23"},
		{name: "dayOfMonth out of range", expr: "0 0 0 32 * *", wantMsg: "DOM part must be >= 1 and <= 31"},
		{name: "garbage expression", expr: "sdlksldksldksd", wantMsg: "Expression has only 1 part. At least 5 parts are required."},
		{name: "too many parts", expr: "* * * * * * * *", wantMsg: "Expression has 8 parts; too many!"},
		{name: "empty expression", expr: "", wantMsg: "cron expression is empty"},
		{name: "undefined expression", expr: "", wantMsg: "cron expression is empty"},
		{name: "'W' list is invalid", expr: "0 30 14 1W,15W * ? *", wantMsg: "The 'W' character can be specified only when the day-of-month is a single day, not a range or list of days."},
		{name: "W offset invalid 1-5W", expr: "0 0 1-5W * *", wantMsg: "The 'W' character can be specified only when the day-of-month is a single day, not a range or list of days."},
		{name: "W offset invalid LW-2", expr: "0 0 LW-2 * *", wantMsg: "The 'W' character can be specified only when the day-of-month is a single day, not a range or list of days."},
		{name: "W offset invalid 1W+2+3", expr: "0 0 1W+2+3 * *", wantMsg: "The 'W' character can be specified only when the day-of-month is a single day, not a range or list of days."},
		{name: "W offset invalid 1W+-2", expr: "0 0 1W+-2 * *", wantMsg: "The 'W' character can be specified only when the day-of-month is a single day, not a range or list of days."},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := ToString(tt.expr, tt.opts...)
			if err == nil {
				t.Fatalf("ToString(%q) expected error containing %q, got nil", tt.expr, tt.wantMsg)
			}
			if !strings.Contains(err.Error(), tt.wantMsg) {
				t.Errorf("ToString(%q) error = %q, want containing %q", tt.expr, err.Error(), tt.wantMsg)
			}
		})
	}
}
