// Package cronstrue converts cron expressions into human readable
// descriptions. It is a Go port of the JavaScript library cronstrue
// (https://github.com/bradymholt/cronstrue, MIT licensed, based on upstream
// v3.32.0), keeping its output and option semantics bug-for-bug compatible
// where practical.
//
// Basic usage:
//
//	cronstrue.ToString("* * * * *")           // "Every minute"
//	cronstrue.ToString("*/5 * * * *")         // "Every 5 minutes"
//	cronstrue.ToString("30 6 * * *")          // "At 06:30 AM"
//	cronstrue.ToString("0 12 1 * *", cronstrue.WithLocale("zh_CN")) // "在下午 12:00, 限每月 1 号"
//
// See the Options type for all available settings.
package cronstrue

import "fmt"

// ToString converts a cron expression into a description a human can read.
//
// By default a parse error is returned as err (mirroring the exception the
// JavaScript library throws). Pass WithThrowExceptionOnParseError(false) to
// instead receive a localized "an error occurred" description with a nil
// error.
func ToString(expression string, opts ...Option) (string, error) {
	options := newOptions(opts...)
	descriptor := newExpressionDescriptor(expression, options)

	description, err := dToString(descriptor)
	if err != nil {
		if options.ThrowExceptionOnParseError {
			return "", err
		}
		return descriptor.i18n.anErrorOccuredWhenGeneratingTheExpressionD(), nil
	}
	return description, nil
}

// dToString runs getFullDescription, converting an internal panic into an
// error the same way the JavaScript original catches exceptions.
func dToString(d *expressionDescriptor) (description string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return d.getFullDescription()
}
