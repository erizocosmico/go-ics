# go-ics [![GoDoc](https://godoc.org/github.com/erizocosmico/go-ics?status.svg)](http://godoc.org/github.com/erizocosmico/go-ics) [![Build Status](https://travis-ci.org/erizocosmico/go-ics.svg?branch=master)](https://travis-ci.org/erizocosmico/go-ics)
This library provides a way of parsing ics calendar files. Supports repetition patterns, organizer and attendees. It also supports both local and remote files to be parsed.

### Status

This is a work in progress. It needs a lot more tests and a small refactor on some parts as well as CI and gopkg.in versions.

## Install

`go get https://github.com/erizocosmico/go-ics`

## How to use it

```go
import "github.com/erizocosmico/go-ics"

// ...
calendar, err := ics.ParseCalendar("local file URL or remote URL", 0, nil)
```

### TODO's

* [ ] Urgently rewrite the whole parser
* [ ] Explicitly handle all errors.
* [ ] trimField should NOT be a regex compiled on runtime.
* [ ] func names improvement
* [ ] divide parseEvents in smaller, testable functions
* [ ] test individual functions

## LICENSE

MIT License, see [LICENSE](/LICENSE)

Based on the work of [PuloV](https://github.com/PuloV/ics-golang).
