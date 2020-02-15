dtf
===

dtf is a command line tool for datetime format converting which is inspired by https://github.com/mwaterfall/alfred-datetime-format-converter

dtf can convert any format of datetime string to second timetamp, millisecond timestamp,  microsecond timestamp and nanosecond timestamp.
And dtf also can convert any timestamp to datetime string.

# Background

日常开发中经常要处理时间戳，每次想要看看时间戳对应的时间或者时间对应的时间戳都比较麻烦，需要去网上打开那种小工具网站来转换。

这种网站虽然方便，但是偶尔也不能满足我的需求，比如网站工具几乎时间戳都是最小只支持到毫秒，微秒、纳秒级别的都要我手动去处理下，而且字符串时间的显示也无法自定义格式，还有就是在家办公期间网速不好网站还打不开或者没网的时候就用不了，有的网站请求频繁还限制公司IP，所以我想自己写一个本地的命令行工具，能够便捷无脑转换出结果。

除了以上说的网站工具，有的mac用户用alfred时间戳转换插件，甚至有的大神直接打开终端敲代码，都各有优缺点，但是既然都要打开终端了，那为啥不来试试我这个dtf？

只要给我参数我就直接告诉你结果，任意时间戳我告诉你字符串结果，任意字符串我告诉你时间戳结果。

# Install

    go get -u github.com/axiaoxin/dtf

or download the binary file at <https://github.com/axiaoxin/dtf/releases>

# Usage

    dtf <any datetime string or timestamp>

# Example

## now:

    dtf now
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

    dtf 1581749145
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

    dtf 1581749145336
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

    dtf 1581749145336928
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

## nanosecond timestamp:

    dtf 1581749145336928000
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

## any datetime string:

    dtf 2020-02-15 14:45:45
    1581749145
    1581749145000
    1581749145000000
    1581749145000000000

    dtf 2020/02/15 14:45:45.336928
    1581749145
    1581749145336
    1581749145336928
    1581749145336928000

    dtf Sat Feb 15 14:45:45 2020
    1581749145
    1581749145000
    1581749145000000
    1581749145000000000

    dtf 2020-02-15T14:45:45.336928+08:00
    1581749145
    1581749145336
    1581749145336928
    1581749145336928000

    dtf 2020
    1577808000
    1577808000000
    1577808000000000
    1577808000000000000

    dtf 2020-02-02
    1580572800
    1580572800000
    1580572800000000
    1580572800000000000

    dtf 20200202
    1580572800
    1580572800000
    1580572800000000
    1580572800000000000

    dtf 2020年02月02日
    1580572800
    1580572800000
    1580572800000000
    1580572800000000000

    dtf 2020年02月02日 02:02
    1580580120
    1580580120000
    1580580120000000
    1580580120000000000

## how long ago:

    dtf 10 minutes ago
    2020-02-16 00:06:00.551004
    2020/02/16 00:06:00.551004
    Sun Feb 16 00:06:00 2020
    Sun Feb 16 00:06:00 CST 2020
    Sun Feb 16 00:06:00 +0800 2020
    16 Feb 20 00:06 CST
    16 Feb 20 00:06 +0800
    Sunday, 16-Feb-20 00:06:00 CST
    Sun, 16 Feb 2020 00:06:00 CST
    Sun, 16 Feb 2020 00:06:00 +0800
    2020-02-16T00:06:00+08:00
    2020-02-16T00:06:00.551004+08:00
    12:06AM
    Feb 16 00:06:00
    Feb 16 00:06:00.551
    Feb 16 00:06:00.551004
    Feb 16 00:06:00.551004000
    1581782760
    1581782760551
    1581782760551004
    1581782760551004000

    dtf 1 hours ago
    2020-02-15 23:16:50.251579
    2020/02/15 23:16:50.251579
    Sat Feb 15 23:16:50 2020
    Sat Feb 15 23:16:50 CST 2020
    Sat Feb 15 23:16:50 +0800 2020
    15 Feb 20 23:16 CST
    15 Feb 20 23:16 +0800
    Saturday, 15-Feb-20 23:16:50 CST
    Sat, 15 Feb 2020 23:16:50 CST
    Sat, 15 Feb 2020 23:16:50 +0800
    2020-02-15T23:16:50+08:00
    2020-02-15T23:16:50.251579+08:00
    11:16PM
    Feb 15 23:16:50
    Feb 15 23:16:50.251
    Feb 15 23:16:50.251579
    Feb 15 23:16:50.251579000
    1581779810
    1581779810251
    1581779810251579
    1581779810251579000

    dtf 1 days ago
    2020-02-15 00:17:36.787268
    2020/02/15 00:17:36.787268
    Sat Feb 15 00:17:36 2020
    Sat Feb 15 00:17:36 CST 2020
    Sat Feb 15 00:17:36 +0800 2020
    15 Feb 20 00:17 CST
    15 Feb 20 00:17 +0800
    Saturday, 15-Feb-20 00:17:36 CST
    Sat, 15 Feb 2020 00:17:36 CST
    Sat, 15 Feb 2020 00:17:36 +0800
    2020-02-15T00:17:36+08:00
    2020-02-15T00:17:36.787268+08:00
    12:17AM
    Feb 15 00:17:36
    Feb 15 00:17:36.787
    Feb 15 00:17:36.787268
    Feb 15 00:17:36.787268000
    1581697056
    1581697056787
    1581697056787268
    1581697056787268000
