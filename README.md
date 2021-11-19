[![Build Status](https://travis-ci.org/gregoryv/worklog.svg?branch=main)](https://travis-ci.org/gregoryv/worklog)
[![codecov](https://codecov.io/gh/gregoryv/worklog/branch/main/graph/badge.svg)](https://codecov.io/gh/gregoryv/worklog)
[![Maintainability](https://api.codeclimate.com/v1/badges/83083a5e52d4ffad3288/maintainability)](https://codeclimate.com/github/gregoryv/worklog/maintainability)


[worklog](https://godoc.org/github.com/gregoryv/worklog) - package for working with the [timesheet fileformat](https://github.com/gregoryv/timesheet-file-format)

## Quick Start

    $ go get -u github.com/gregoryv/worklog/...
	$ gensheet -h

This package contains parser and commands for manipulating timesheet
files.

Let's face it, timesheets are extremely boring to fill out so I
decided to fix that. Check out
the
[timesheet fileformat](https://github.com/gregoryv/timesheet-file-format),
it's readable, lightweight and versatile enough to accommodate most
situations.

By default the command `gensheet` generates an 8 hour working day
timesheet already filled out, assuming you will be working each
weekday. Save it, put it under version control, do what you want with
it.

To summarize your timesheets use the `worklog` command.

## Summarize one timesheet

    $ worklog 201801.timesheet
    2018 January    179:30   (8:00 semester)

                   +179:30
                              8:00 semester
## Expected vs. actual worked hours

    worklog -origin expected/ 201801.timesheet

The output summarizes all tagged values and reported time and
calculates how much overtime you have put in if any. The `expected`
directory is assumed to contain a timesheet with the same name to
compare with.
