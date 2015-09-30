# go-ics [![GoDoc](https://godoc.org/github.com/mvader/go-ics?status.svg)](http://godoc.org/github.com/mvader/go-ics) [![Build Status](https://travis-ci.org/mvader/go-ics.svg?branch=master)](https://travis-ci.org/mvader/go-ics)
This library provides a way of parsing ics calendar files. Supports repetition patterns, organizer and attendees. It also supports both local and remote files to be parsed.

### Status
This is a work in progress. It needs a lot more tests and a small refactor on some parts as well as CI and gopkg.in versions.

##Install
`go get https://github.com/mvader/go-ics`

##How to use it
```go
import "github.com/mvader/go-ics"

// ...
calendar, err := ics.ParseCalendar("local file URL or remote URL", 0, nil)
```

### TODO's

* [ ] Explicitly handle all errors.
* [ ] trimField should NOT be a regex compiled on runtime.
* [ ] func names improvement
* [ ] divide parseEvents in smaller, testable functions
* [ ] test individual functions

## LICENSE
The MIT License (MIT)

Copyright (c) 2015 Miguel Molina

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Based on the work of [PuloV](https://github.com/PuloV/ics-golang).
