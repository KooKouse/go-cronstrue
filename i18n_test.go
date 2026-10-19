package cronstrue

// Code generated from cronstrue upstream test/i18n.ts; golden outputs
// are the JavaScript library's outputs for each locale.

import "testing"

func TestToStringI18n(t *testing.T) {
	tests := []struct {
		locale string
		name   string
		expr   string
		use24  bool
		want   string
	}{
		{locale: "de", name: "de: * * * * *", expr: "* * * * *", want: "Jede Minute"},
		{locale: "de", name: "de: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Alle 5 Minuten, zwischen 15:00 und 15:59, Montag bis Freitag"},
		{locale: "es", name: "es: * * * * *", expr: "* * * * *", want: "Cada minuto"},
		{locale: "es", name: "es: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Cada 5 minutos, entre las 15:00 y las 15:59, de lunes a viernes"},
		{locale: "fr", name: "fr: * * * * *", expr: "* * * * *", want: "Toutes les minutes"},
		{locale: "fr", name: "fr: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Toutes les 5 minutes, de 15:00 à 15:59, de lundi à vendredi"},
		{locale: "it", name: "it: * * * * *", expr: "* * * * *", want: "Ogni minuto"},
		{locale: "it", name: "it: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Ogni 5 minuti, tra le 15:00 e le 15:59, lunedì al venerdì"},
		{locale: "id", name: "id: * * * * *", expr: "* * * * *", want: "Setiap menit"},
		{locale: "id", name: "id: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Setiap 5 menit, diantara 15:00 dan 15:59, Senin sampai Jumat"},
		{locale: "nb", name: "nb: * * * * *", expr: "* * * * *", want: "Hvert minutt"},
		{locale: "nb", name: "nb: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", use24: true, want: "Hvert 5 minutt, mellom 15:00 og 15:59, mandag til og med fredag"},
		{locale: "nb", name: "nb: 0 5 1/1 * *", expr: "0 5 1/1 * *", use24: true, want: "Kl.05:00"},
		{locale: "nb", name: "nb: 15 11 * 1/1 MON#1", expr: "15 11 * 1/1 MON#1", use24: true, want: "Kl.11:15, på første mandag i måneden"},
		{locale: "nb", name: "nb: 15 11 * 1/5 MON#1", expr: "15 11 * 1/5 MON#1", use24: true, want: "Kl.11:15, på første mandag i måneden, hver 5 måned"},
		{locale: "nb", name: "nb: 0 7 * * MON,TUE,THU,FRI,SUN", expr: "0 7 * * MON,TUE,THU,FRI,SUN", use24: true, want: "Kl.07:00, på mandag, tirsdag, torsdag, fredag, og søndag"},
		{locale: "sv", name: "sv: * * * * *", expr: "* * * * *", want: "Varje minut"},
		{locale: "sv", name: "sv: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Var 5 minut, mellan 15:00 och 15:59, måndag till fredag"},
		{locale: "sv", name: "sv: 0 12 * * *", expr: "0 12 * * *", want: "Kl 12:00"},
		{locale: "sv", name: "sv: 0 15 10 ? * 6#3", expr: "0 15 10 ? * 6#3", want: "Kl 10:15, den tredje lördagen av månaden"},
		{locale: "sv", name: "sv: 0 0 15 ? * MON *", expr: "0 0 15 ? * MON *", want: "Kl 15:00, varje måndag"},
		{locale: "nl", name: "nl: * * * * *", expr: "* * * * *", want: "Elke minuut"},
		{locale: "nl", name: "nl: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Elke 5 minuten, tussen 15:00 en 15:59, maandag t/m vrijdag"},
		{locale: "pl", name: "pl: * * * * *", expr: "* * * * *", want: "Co minutę"},
		{locale: "pl", name: "pl: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Co 5 minut, od 15:00 do 15:59, od poniedziałek do piątek"},
		{locale: "pt-BR", name: "pt-BR: * * * * *", expr: "* * * * *", want: "A cada minuto"},
		{locale: "pt-BR", name: "pt-BR: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "A cada 5 minutos, entre 15:00 e 15:59, de segunda-feira a sexta-feira"},
		{locale: "pt-BR", name: "pt-BR: */5 15 * * MON-FRI,SUN", expr: "*/5 15 * * MON-FRI,SUN", want: "A cada 5 minutos, entre 15:00 e 15:59, somente de segunda-feira a sexta-feira e domingo"},
		{locale: "pt-BR", name: "pt-BR: 45 10 * * 6#2", expr: "45 10 * * 6#2", want: "Às 10:45, no segundo sábado do mês"},
		{locale: "pt-BR", name: "pt-BR: 45 10 * * 0#3", expr: "45 10 * * 0#3", want: "Às 10:45, no terceiro domingo do mês"},
		{locale: "pt-BR", name: "pt-BR: 45 10 * * 1#3", expr: "45 10 * * 1#3", want: "Às 10:45, na terceira segunda-feira do mês"},
		{locale: "pt-PT", name: "pt-PT: * * * * *", expr: "* * * * *", want: "A cada minuto"},
		{locale: "pt-PT", name: "pt-PT: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "A cada 5 minutos, entre 15:00 e 15:59, de segunda-feira a sexta-feira"},
		{locale: "pt-BR", name: "pt-BR: 45 10 * * 6#2 #2", expr: "45 10 * * 6#2", want: "Às 10:45, no segundo sábado do mês"},
		{locale: "pt-BR", name: "pt-BR: 45 10 * * 0#3 #2", expr: "45 10 * * 0#3", want: "Às 10:45, no terceiro domingo do mês"},
		{locale: "pt-BR", name: "pt-BR: 45 10 * * 1#3 #2", expr: "45 10 * * 1#3", want: "Às 10:45, na terceira segunda-feira do mês"},
		{locale: "ro", name: "ro: * * * * *", expr: "* * * * *", want: "În fiecare minut"},
		{locale: "ro", name: "ro: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "La fiecare 5 minute, între 15:00 și 15:59, de luni până vineri"},
		{locale: "ro", name: "ro: */30 * * * *", expr: "*/30 * * * *", want: "La fiecare 30 de minute"},
		{locale: "ro", name: "ro: */20 * * * * *", expr: "*/20 * * * * *", want: "La fiecare 20 de secunde"},
		{locale: "ro", name: "ro: 0 */20 * * *", expr: "0 */20 * * *", want: "Fix la oră, la fiecare 20 de ore"},
		{locale: "ro", name: "ro: 0 0 */20 * *", expr: "0 0 */20 * *", want: "La 00:00, la fiecare 20 de zile"},
		{locale: "ro", name: "ro: 0 0 L-20 * *", expr: "0 0 L-20 * *", want: "La 00:00, 20 de zile înainte de ultima zi a lunii"},
		{locale: "ro", name: "ro: 0 0 0 1 1 ? */100", expr: "0 0 0 1 1 ? */100", want: "La 00:00, în ziua 1 a lunii, doar în ianuarie, o dată la 100 de ani"},
		{locale: "ro", name: "ro: 0 0 0 1 1 ? */101", expr: "0 0 0 1 1 ? */101", want: "La 00:00, în ziua 1 a lunii, doar în ianuarie, o dată la 101 ani"},
		{locale: "ro", name: "ro: 0 9 1W * *", expr: "0 9 1W * *", want: "La 09:00, în prima zi lucrătoare a lunii"},
		{locale: "ro", name: "ro: 0 9 15W * *", expr: "0 9 15W * *", want: "La 09:00, în cea mai apropiată zi lucrătoare de ziua 15 a lunii"},
		{locale: "ru", name: "ru: * * * * *", expr: "* * * * *", want: "Каждую минуту"},
		{locale: "ru", name: "ru: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Каждые 5 минут, с 15:00 по 15:59, с понедельника по пятницу"},
		{locale: "ru", name: "ru: * * * * */31", expr: "* * * * */31", want: "Каждую минуту, каждый 31 день недели"},
		{locale: "tr", name: "tr: * * * * *", expr: "* * * * *", want: "Her dakika"},
		{locale: "tr", name: "tr: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Her 5 dakikada bir, 15:00 ile 15:59 arasında, Pazartesi ile Cuma arasında"},
		{locale: "tr", name: "tr: @reboot", expr: "@reboot", want: "Başlangıçta bir kez çalıştır"},
		{locale: "uk", name: "uk: * * * * *", expr: "* * * * *", want: "Щохвилини"},
		{locale: "uk", name: "uk: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Кожні 5 хвилин, між 15:00 та 15:59, понеділок по п'ятниця"},
		{locale: "zh-Hans", name: "zh-Hans: * * * * *", expr: "* * * * *", want: "每分钟"},
		{locale: "zh-Hans", name: "zh-Hans: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "每隔 5 分钟, 在 下午 03:00 和 下午 03:59 之间, 星期一至星期五"},
		{locale: "zh-Hant", name: "zh-Hant: * * * * *", expr: "* * * * *", want: "每分鐘"},
		{locale: "zh-Hant", name: "zh-Hant: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "每 5 分鐘, 在 15:00 和 15:59 之間, 星期一 到 星期五"},
		{locale: "he", name: "he: * * * * *", expr: "* * * * *", want: "כל דקה"},
		{locale: "he", name: "he: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "כל 5 דקות, 15:00 עד 15:59, יום שני עד יום שישי"},
		{locale: "da", name: "da: * * * * *", expr: "* * * * *", want: "Hvert minut"},
		{locale: "da", name: "da: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Hvert 5. minut, mellem 15:00 og 15:59, mandag til og med fredag"},
		{locale: "da", name: "da: 0 3 * * 4", expr: "0 3 * * 4", want: "Kl 03:00, på enhver torsdag"},
		{locale: "da", name: "da: 0 3 * * 4,5", expr: "0 3 * * 4,5", want: "Kl 03:00, på enhver torsdag og fredag"},
		{locale: "fa", name: "fa: * * * * *", expr: "* * * * *", want: "هر دقیقه"},
		{locale: "fa", name: "fa: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "هر 5 دقیقه, بین 15:00 و 15:59, دوشنبه تا جمعه"},
		{locale: "tg", name: "tg: * * * * *", expr: "* * * * *", want: "Ҳар дақиқа"},
		{locale: "tg", name: "tg: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Ҳар 5 дақиқа, аз 15:00 то 15:59, аз душанбе то ҷумъа"},
		{locale: "tg", name: "tg: 0 0 L * *", expr: "0 0 L * *", want: "Дар 00:00, дар рӯзи охирини моҳ"},
		{locale: "sl", name: "sl: * * * * *", expr: "* * * * *", want: "Vsako minuto"},
		{locale: "sl", name: "sl: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Vsakih 5 minut, od 15:00 do 15:59, od Ponedeljek do Petek"},
		{locale: "sl", name: "sl: */5 * L JAN *", expr: "*/5 * L JAN *", want: "Vsakih 5 minut, zadnji dan v mesecu, samo v januar"},
		{locale: "ca", name: "ca: * * * * *", expr: "* * * * *", want: "Cada minut"},
		{locale: "ca", name: "ca: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Cada 5 minuts, entre les 15:00 i les 15:59, de dilluns a divendres"},
		{locale: "ja", name: "ja: 0 0 14 */30 * *", expr: "0 0 14 */30 * *", want: "次において実施14:00、月のうち 30 日ごと"},
		{locale: "ko", name: "ko: * * * * *", expr: "* * * * *", want: "1분마다"},
		{locale: "ko", name: "ko: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "5분마다, 오후 03:00부터 오후 03:59까지, 월요일부터 금요일까지"},
		{locale: "ko", name: "ko: 0-10 11 * * *", expr: "0-10 11 * * *", want: "매분 오전 11:00~오전 11:10"},
		{locale: "ko", name: "ko: 23 12 15 * *", expr: "23 12 15 * *", want: "오후 12:23, 매월 15일"},
		{locale: "ko", name: "ko: 0 9 * * 1-5", expr: "0 9 * * 1-5", want: "오전 09:00, 월요일부터 금요일까지"},
		{locale: "ko", name: "ko: 0,22,24,26,28 * * * *", expr: "0,22,24,26,28 * * * *", want: "0, 22, 24, 26, 28분"},
		{locale: "ko", name: "ko: 30 6,14 * * *", expr: "30 6,14 * * *", want: "오전 06:30, 오후 02:30"},
		{locale: "ko", name: "ko: * * * * MON#3", expr: "* * * * MON#3", want: "1분마다, 매월 세 번째 월요일"},
		{locale: "ko", name: "ko: * * LW * *", expr: "* * LW * *", want: "1분마다, 매월 마지막 평일"},
		{locale: "hu", name: "hu: * * * * *", expr: "* * * * *", want: "Minden percben"},
		{locale: "hu", name: "hu: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Minden 5. percben, 15:00 és 15:59 között, hétfő - péntek"},
		{locale: "af", name: "af: * * * * *", expr: "* * * * *", want: "Elke minuut"},
		{locale: "af", name: "af: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Elke 5 minute, tussen 15:00 en 15:59, Maandag deur na Vrydag"},
		{locale: "th", name: "th: * * * * *", expr: "* * * * *", want: "ทุกๆ นาที"},
		{locale: "th", name: "th: 0 * * * *", expr: "0 * * * *", want: "ทุกๆ ชั่วโมง"},
		{locale: "th", name: "th: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "ทุกๆ 5 นาที, ระหว่าง 03:00 PM ถึง 03:59 PM, วันจันทร์ ถึง วันศุกร์"},
		{locale: "th", name: "th: 15 14 1 * *", expr: "15 14 1 * *", want: "เมื่อ 02:15 PM, ในวัน 1 ของเดือน"},
		{locale: "ar", name: "ar: * * * * *", expr: "* * * * *", want: "كل دقيقة"},
		{locale: "ar", name: "ar: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "كل 5 دقائق, بين 15:00 و 15:59، الإثنين حتى الجمعة"},
		{locale: "my", name: "my: * * * * *", expr: "* * * * *", want: "Setiap minit"},
		{locale: "my", name: "my: 0 * * * *", expr: "0 * * * *", use24: true, want: "Setiap jam"},
		{locale: "my", name: "my: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", use24: true, want: "Setiap 5 minit, antara 15:00 dan 15:59, Isnin hingga Jumaat"},
		{locale: "my", name: "my: 0 5 1/1 * *", expr: "0 5 1/1 * *", use24: true, want: "Pada 05:00"},
		{locale: "my", name: "my: 15 11 * 1/5 MON", expr: "15 11 * 1/5 MON", use24: true, want: "Pada 11:15, hanya pada Isnin, setiap bulan 5"},
		{locale: "my", name: "my: 0 7 * * MON,TUE,THU,FRI,SUN", expr: "0 7 * * MON,TUE,THU,FRI,SUN", use24: true, want: "Pada 07:00, hanya pada Isnin, Selasa, Khamis, Jumaat, dan Ahad"},
		{locale: "bg", name: "bg: * * * * *", expr: "* * * * *", want: "Всяка минута"},
		{locale: "bg", name: "bg: 0 * * * *", expr: "0 * * * *", want: "Всеки час"},
		{locale: "bg", name: "bg: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Всеки 5 минути, от 15:00 до 15:59, от понеделник до петък"},
		{locale: "be", name: "be: * * * * *", expr: "* * * * *", want: "Кожную хвіліну"},
		{locale: "be", name: "be: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Кожныя 5 хвілін, з 03:00 PM па 03:59 PM, панядзелак па пятніца"},
		{locale: "hr", name: "hr: * * * * *", expr: "* * * * *", want: "Svaku minutu"},
		{locale: "hr", name: "hr: 0 * * * *", expr: "0 * * * *", want: "Svaki sat"},
		{locale: "hr", name: "hr: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Svakih 5 minuta, između 15:00 i 15:59, od Ponedjeljak do Petak"},
		{locale: "sr", name: "sr: * * * * *", expr: "* * * * *", want: "Svakog minuta"},
		{locale: "sr", name: "sr: 0 * * * *", expr: "0 * * * *", want: "Svaki sat"},
		{locale: "sr", name: "sr: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Svakih 5 minuta, između 15:00 i 15:59, od Ponedeljak do Petak"},
		{locale: "az", name: "az: * * * * *", expr: "* * * * *", want: "Hər dəqiqə"},
		{locale: "az", name: "az: 0 * * * *", expr: "0 * * * *", want: "Hər saat"},
		{locale: "az", name: "az: */5 15 * * MON-FRI", expr: "*/5 15 * * MON-FRI", want: "Hər 5 dəqiqədən bir, saat 15:00 ilə 15:59 arasında, bazar ertəsindən cüməyə qədər"},
		{locale: "az", name: "az: 0 0 3 * *", expr: "0 0 3 * *", want: "Saat 00:00, ayın 3-cü günü"},
		{locale: "az", name: "az: 0 0 9 * *", expr: "0 0 9 * *", want: "Saat 00:00, ayın 9-cu günü"},
		{locale: "az", name: "az: 23 12 * JAN-MAR *", expr: "23 12 * JAN-MAR *", want: "Saat 12:23, yanvardan marta qədər"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			opts := []Option{WithLocale(tt.locale)}
			if tt.use24 {
				opts = append(opts, With24HourTimeFormat(true))
			}
			got, err := ToString(tt.expr, opts...)
			if err != nil {
				t.Fatalf("ToString(%q, locale=%q) returned error: %v", tt.expr, tt.locale, err)
			}
			if got != tt.want {
				t.Errorf("ToString(%q, locale=%q) = %q, want %q", tt.expr, tt.locale, got, tt.want)
			}
		})
	}
}
