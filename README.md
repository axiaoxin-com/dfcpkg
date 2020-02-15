dtf
===

dtf is a command line tool for datetime format converting which is inspired by https://github.com/mwaterfall/alfred-datetime-format-converter

dtf can convert any format of datetime string to second timetamp, millisecond timestamp,  microsecond timestamp and nanosecond timestamp.
And dtf also can convert any timestamp to datetime string.

# Install

    go get -u github.com/axiaoxin/dtf

or download the binary file at <https://github.com/axiaoxin/dtf/releases>

# Usage

    ./dtf <any datetime string or timestamp>

# Example

## now:

    % dtf now
    2020-02-15 14:45:45.336928
    2020/02/15 14:45:45.336928
    Sat Feb 15 14:45:45 2020
    Sat Feb 15 14:45:45 CST 2020
    Sat Feb 15 14:45:45 +0800 2020
    15 Feb 20 14:45 CST
    15 Feb 20 14:45 +0800
    Saturday, 15-Feb-20 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 +0800
    2020-02-15T14:45:45+08:00
    2020-02-15T14:45:45.336928+08:00
    2:45PM
    Feb 15 14:45:45
    Feb 15 14:45:45.336
    Feb 15 14:45:45.336928
    Feb 15 14:45:45.336928000
    1581749145
    1581749145336
    1581749145336928
    1581749145336928000

## second timestamp:

    % dtf 1581749145
    2020-02-15 14:45:45
    2020/02/15 14:45:45
    Sat Feb 15 14:45:45 2020
    Sat Feb 15 14:45:45 CST 2020
    Sat Feb 15 14:45:45 +0800 2020
    15 Feb 20 14:45 CST
    15 Feb 20 14:45 +0800
    Saturday, 15-Feb-20 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 +0800
    2020-02-15T14:45:45+08:00
    2020-02-15T14:45:45+08:00
    2:45PM
    Feb 15 14:45:45
    Feb 15 14:45:45.000
    Feb 15 14:45:45.000000
    Feb 15 14:45:45.000000000

## millisecond timestamp:

    % dtf 1581749145336
    2020-02-15 14:45:45.336
    2020/02/15 14:45:45.336
    Sat Feb 15 14:45:45 2020
    Sat Feb 15 14:45:45 CST 2020
    Sat Feb 15 14:45:45 +0800 2020
    15 Feb 20 14:45 CST
    15 Feb 20 14:45 +0800
    Saturday, 15-Feb-20 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 +0800
    2020-02-15T14:45:45+08:00
    2020-02-15T14:45:45.336+08:00
    2:45PM
    Feb 15 14:45:45
    Feb 15 14:45:45.336
    Feb 15 14:45:45.336000
    Feb 15 14:45:45.336000000

## microsecond timestamp:

    % dtf 1581749145336928
    2020-02-15 14:45:45.336928
    2020/02/15 14:45:45.336928
    Sat Feb 15 14:45:45 2020
    Sat Feb 15 14:45:45 CST 2020
    Sat Feb 15 14:45:45 +0800 2020
    15 Feb 20 14:45 CST
    15 Feb 20 14:45 +0800
    Saturday, 15-Feb-20 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 +0800
    2020-02-15T14:45:45+08:00
    2020-02-15T14:45:45.336928+08:00
    2:45PM
    Feb 15 14:45:45
    Feb 15 14:45:45.336
    Feb 15 14:45:45.336928
    Feb 15 14:45:45.336928000

# nanosecond timestamp

    % dtf 1581749145336928000
    2020-02-15 14:45:45.336928
    2020/02/15 14:45:45.336928
    Sat Feb 15 14:45:45 2020
    Sat Feb 15 14:45:45 CST 2020
    Sat Feb 15 14:45:45 +0800 2020
    15 Feb 20 14:45 CST
    15 Feb 20 14:45 +0800
    Saturday, 15-Feb-20 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 CST
    Sat, 15 Feb 2020 14:45:45 +0800
    2020-02-15T14:45:45+08:00
    2020-02-15T14:45:45.336928+08:00
    2:45PM
    Feb 15 14:45:45
    Feb 15 14:45:45.336
    Feb 15 14:45:45.336928
    Feb 15 14:45:45.336928000

# any datetime string:

    % dtf 2020-02-15 14:45:45
    1581777945
    1581777945000
    1581777945000000

    % dtf 2020/02/15 14:45:45.336928
    1581777945
    1581777945336
    1581777945336928
    1581777945336928000

    % dtf Sat Feb 15 14:45:45 2020
    1581777945
    1581777945000
    1581777945000000
    1581777945000000000

    % dtf 2020-02-15T14:45:45.336928+08:00
    1581749145
    1581749145336
    1581749145336928
    1581749145336928000
