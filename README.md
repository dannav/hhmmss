# HH:MM:SS

Parse HH:MM:SS strings into a Go time.Duration type.

## Installation

Install this package with:

```
go get github.com/dannav/hhmmss
```

## Usage

Given the following string `03:30:26` call the package's `Parse` method:

```
dur, err := hhmmss.Parse("03:30:26")
```

`Parse` returns a `time.Duration` value that allows you to retrieve the value's time
in seconds, hours, minutes, etc.

If an integer string is passed to `Parse` the argument will be treated as seconds.

## Examples

View the tests for more examples on using `hhmmss`.