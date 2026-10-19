package cronstrue

// Options controls how a cron expression is converted into a human readable
// description. Option values mirror the options accepted by the original
// cronstrue JavaScript library (https://github.com/bradymholt/cronstrue).
type Options struct {
	// ThrowExceptionOnParseError indicates whether a parse error should be
	// returned as an error (true, the default) or replaced by a localized
	// "an error occurred" description string (false).
	ThrowExceptionOnParseError bool

	// Verbose, when true, produces a longer description that spells out
	// implied fields (e.g. "Every minute, every hour, every day").
	Verbose bool

	// DayOfWeekStartIndexZero determines whether Sunday is represented as 0
	// (standard cron, the default) or as 1 (Quartz style).
	DayOfWeekStartIndexZero bool

	// MonthStartIndexZero determines whether January is represented as 1
	// (standard cron, the default) or as 0.
	MonthStartIndexZero bool

	// Use24HourTimeFormat forces 24 hour time output (e.g. "13:00" instead
	// of "1:00 PM"). When Use24HourTimeFormatSet is false the locale default
	// (see locale.use24HourTimeFormatByDefault) is used instead.
	Use24HourTimeFormat    bool
	Use24HourTimeFormatSet bool

	// TrimHoursLeadingZero removes the leading zero of single digit hours
	// (e.g. "9:00 AM" instead of "09:00 AM").
	TrimHoursLeadingZero bool

	// LocaleName selects the output language (e.g. "en", "zh-Hans"). Empty
	// means the default locale ("en").
	LocaleName string

	// LogicalAndDayFields causes the day-of-month and day-of-week fields to
	// be described as a logical AND rather than the cron standard OR, e.g.
	// "At 12:00 AM, on day 1, 2, and 3 of the month, only on Wednesday and
	// Friday" rather than "... and on Wednesday and Friday".
	LogicalAndDayFields bool
}

// Option is a functional option used by ToString.
type Option func(*Options)

func newOptions(opts ...Option) *Options {
	o := &Options{
		ThrowExceptionOnParseError: true,
		DayOfWeekStartIndexZero:    true,
		MonthStartIndexZero:        false,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
	return o
}

// WithVerbose enables or disables the verbose description format.
func WithVerbose(verbose bool) Option {
	return func(o *Options) { o.Verbose = verbose }
}

// WithThrowExceptionOnParseError controls the behavior on parse errors:
// when true the error is returned; when false a localized error description
// string is returned instead.
func WithThrowExceptionOnParseError(throw bool) Option {
	return func(o *Options) { o.ThrowExceptionOnParseError = throw }
}

// WithDayOfWeekStartIndexZero sets whether Sunday is day 0 (true) or day 1
// (Quartz convention).
func WithDayOfWeekStartIndexZero(zero bool) Option {
	return func(o *Options) { o.DayOfWeekStartIndexZero = zero }
}

// WithMonthStartIndexZero sets whether January is month 0 (true) or month 1
// (standard cron).
func WithMonthStartIndexZero(zero bool) Option {
	return func(o *Options) { o.MonthStartIndexZero = zero }
}

// With24HourTimeFormat forces 12/24 hour time output.
func With24HourTimeFormat(use bool) Option {
	return func(o *Options) {
		o.Use24HourTimeFormat = use
		o.Use24HourTimeFormatSet = true
	}
}

// WithTrimHoursLeadingZero removes the leading zero of single digit hours.
func WithTrimHoursLeadingZero(trim bool) Option {
	return func(o *Options) { o.TrimHoursLeadingZero = trim }
}

// WithLocale selects the output language by name, e.g. "en" or "zh-Hans".
func WithLocale(name string) Option {
	return func(o *Options) { o.LocaleName = name }
}

// WithLogicalAndDayFields describes the day-of-month and day-of-week fields
// as a logical AND (see Options.LogicalAndDayFields).
func WithLogicalAndDayFields(and bool) Option {
	return func(o *Options) { o.LogicalAndDayFields = and }
}
