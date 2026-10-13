# go-cronstrue

[![Go Reference](https://pkg.go.dev/badge/github.com/KooKouse/go-cronstrue.svg)](https://pkg.go.dev/github.com/KooKouse/go-cronstrue)
[![Go Test](https://github.com/KooKouse/go-cronstrue/actions/workflows/go.yml/badge.svg)](https://github.com/KooKouse/go-cronstrue/actions/workflows/go.yml)

**English** | [中文说明](#中文说明)

A Go library that converts cron expressions into human-readable descriptions. It is a faithful port of the JavaScript library [cronstrue](https://github.com/bradymholt/cronstrue) (based on upstream v3.32.0), verified to produce **byte-identical output** on a 13,200-case differential test suite against the original library (including error messages).

```text
*/5 * * * *        ->  Every 5 minutes
30 6 * * *         ->  At 06:30 AM
0 12 1 * *         ->  At 12:00 PM, on day 1 of the month
0 0 * * MON-FRI    ->  At 12:00 AM, Monday through Friday
23 12 * * SUN#2    ->  At 12:23 PM, on the second Sunday of the month
```

## Features

- Supports 5-field (standard), 6-field (with seconds or year) and 7-field (Quartz style) cron expressions
- Supports `* / , - ? L W #` special characters, `@yearly`/`@monthly`/`@daily`/`...` aliases and `@reboot`
- All upstream options: `verbose`, `dayOfWeekStartIndexZero`, `monthStartIndexZero`, `use24HourTimeFormat`, `trimHoursLeadingZero`, `throwExceptionOnParseError`, `logicalAndDayFields`
- i18n with locale fallback; **English** and **Simplified Chinese** (`zh_CN`) are currently ported (more locales welcome — see [Adding a locale](#adding-a-locale))
- Zero dependencies, single pure-Go package
- Also ships a small CLI

## Installation

```sh
go get github.com/KooKouse/go-cronstrue
```

## Usage

```go
package main

import (
	"fmt"

	cronstrue "github.com/KooKouse/go-cronstrue"
)

func main() {
	desc, err := cronstrue.ToString("*/5 * * * *")
	fmt.Println(desc, err) // Every 5 minutes <nil>

	desc, _ = cronstrue.ToString("30 6 1,15 * *")
	fmt.Println(desc) // At 06:30 AM, on day 1 and 15 of the month

	desc, _ = cronstrue.ToString("0 12 1 * *", cronstrue.WithLocale("zh_CN"))
	fmt.Println(desc) // 在下午 12:00, 限每月 1 号

	desc, _ = cronstrue.ToString("0 * * * *", cronstrue.WithVerbose(true))
	fmt.Println(desc) // Every hour, every day
}
```

By default a parse error is returned as a Go `error` (mirroring the exception the JavaScript library throws). Pass `cronstrue.WithThrowExceptionOnParseError(false)` to instead receive a localized error description with a nil error, e.g. `An error occurred when generating the expression description. Check the cron expression syntax.`

## Options

| Option | Default | Description |
|---|---|---|
| `WithVerbose(v)` | `false` | Longer description spelling out implied fields ("Every minute, every hour, every day") |
| `WithLocale(name)` | `"en"` | Output language; currently `"en"`, `"zh_CN"` |
| `WithThrowExceptionOnParseError(v)` | `true` | `true`: return parse errors as `error`; `false`: return the localized error message string instead |
| `WithDayOfWeekStartIndexZero(v)` | `true` | `true`: Sunday is 0 (standard cron); `false`: Sunday is 1 (Quartz) |
| `WithMonthStartIndexZero(v)` | `false` | `true`: January is 0 |
| `With24HourTimeFormat(v)` | locale default | `true`: "18:30"; `false`: "06:30 PM" |
| `WithTrimHoursLeadingZero(v)` | `false` | `true`: "9:00 AM" instead of "09:00 AM" |
| `WithLogicalAndDayFields(v)` | `false` | Describe day-of-month and day-of-week as logical AND instead of cron's OR |

## CLI

```sh
go install github.com/KooKouse/go-cronstrue/cmd/cronstrue@latest

cronstrue "*/5 * * * *"
# Every 5 minutes

cronstrue -locale zh_CN "0 12 1 * *"
# 在下午 12:00, 限每月 1 号
```

## Compatibility with the original library

The goal of this port is **bug-for-bug compatible output**. Concretely:

- The full upstream test suite (`test/cronstrue.ts`, `test/i18n.ts`) was ported to Go table tests (187 cases) — all pass.
- A differential harness ran 13,200 expression × option combinations through both this library and the JavaScript original (cronstrue 3.32.0 on Node.js): outputs and error messages are byte-identical.
- A few JavaScript-specific quirks are intentionally preserved: numeric coercion inside comma lists (`"05"` → `"5"`), the `1900+n` mapping for two-digit years, and even the `"null"`/`"undefined"` concatenation artifacts of upstream edge cases.

Known intentional deviations (Go idioms):

- Errors are returned as `error` values instead of thrown exceptions.
- Unknown locales fall back to English silently (upstream logs a console warning).
- The deprecated `tzOffset` option is not ported.

## Adding a locale

Locales live in `locale_<name>.go` (e.g. `locale_zh_cn.go`) and implement the unexported `locale` interface. Optional upstream translations are modeled as `(string, bool)` returns — return `ok == false` to keep the fallback behavior. The easiest way to add one is to:

1. Copy `src/i18n/locales/<name>.ts` from [bradymholt/cronstrue](https://github.com/bradymholt/cronstrue/tree/master/src/i18n/locales),
2. create `locale_<name>.go` embedding `enLocale`, overriding every non-`null` method,
3. register it in an `init()` function with `registerLocale("<name>", ...)`.

PRs for the remaining ~40 upstream locales are very welcome!

## Development

```sh
go test ./...   # run the ported upstream test suite
go build ./...  # build library + CLI
```

## Related projects

- [bradymholt/cronstrue](https://github.com/bradymholt/cronstrue) — the original JavaScript/TypeScript library (MIT)
- [lnquy/cron](https://github.com/lnquy/cron) — an older Go port translated from the C# lineage of the same family (26 locales, no longer tracks upstream)

## License

[MIT](LICENSE) — same as the original library. Ported with credit to Brady Holt and the cronstrue contributors.

---

# 中文说明

[go-cronstrue](https://github.com/KooKouse/go-cronstrue) 是 [cronstrue](https://github.com/bradymholt/cronstrue)（基于上游 v3.32.0）的 Go 移植版，把 cron 表达式转换成人类可读的描述。已通过与原 JS 库 **13,200 条差分用例**的逐字节一致性验证（含错误信息）。

```text
*/5 * * * *        ->  Every 5 minutes
0 12 1 * *         ->  At 12:00 PM, on day 1 of the month
0 12 1 * * (zh_CN) ->  在下午 12:00, 限每月 1 号
0 0 * * MON-FRI (zh_CN) ->  在上午 12:00, 星期一至星期五
```

## 特性

- 支持 5 段（标准）、6 段（带秒或年份）、7 段（Quartz 风格）表达式
- 支持 `* / , - ? L W #` 特殊字符、`@daily` 等别名以及 `@reboot`
- 完整支持上游所有选项（verbose、24 小时制、星期/月份起始下标等）
- 内置 i18n：英文、简体中文（`zh_CN`），欢迎贡献其余 locale
- 零依赖、纯 Go；附带命令行工具

## 安装与使用

```sh
go get github.com/KooKouse/go-cronstrue
```

```go
desc, err := cronstrue.ToString("*/5 * * * *")            // Every 5 minutes
desc, _ = cronstrue.ToString("0 12 1 * *",
    cronstrue.WithLocale("zh_CN"))                        // 在下午 12:00, 限每月 1 号
```

解析失败时默认返回 `error`；也可以用 `cronstrue.WithThrowExceptionOnParseError(false)` 改为返回本地化的错误提示文案（与上游行为一致）。

完整选项说明见上方 [Options](#options)。

## 命令行

```sh
go install github.com/KooKouse/go-cronstrue/cmd/cronstrue@latest
cronstrue -locale zh_CN "0 12 1 * *"   # 在下午 12:00, 限每月 1 号
```

## 与原库的兼容性

- 移植了上游完整测试集（187 个 Go 用例），全部通过
- 13,200 条表达式 × 选项组合与 JS 原版输出逐字节一致
- 有意保留了原库的一些怪癖（如逗号列表中的 `"05"`→`"5"` 数值化、两位年份 +1900 等）

## 贡献新语言

参照 `locale_zh_cn.go`，从上游 `src/i18n/locales/` 复制对应 `.ts` 文件翻译为 Go，内嵌 `enLocale` 并覆盖非 `null` 的方法，最后在 `init()` 里 `registerLocale` 注册即可。欢迎 PR！

## 许可证

[MIT](LICENSE)，与原库一致。
