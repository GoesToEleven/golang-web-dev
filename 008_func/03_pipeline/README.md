# Formatting with time

Since MST is GMT-0700, the reference time can be thought of as:

01/02 03:04:05PM '06 -0700

(January 02, 2006)

[godoc.org/time](https://godoc.org/time#pkg-constants)

***

## func (Time) Format

```go
func (t Time) Format(layout string) string
```

Format returns a textual representation of the time value formatted according to layout, which defines the format by showing how the reference time, defined to be

```go
Mon Jan 2 15:04:05 -0700 MST 2006
```

would be displayed if it were the value; it serves as an example of the desired output. The same display rules will then be applied to the time value.

A fractional second is represented by adding a period and zeros to the end of the seconds section of layout string, as in "15:04:05.000" to format a time stamp with millisecond precision.

Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard and convenient representations of the reference time. For more information about the formats and the definition of the reference time, see the documentation for ANSIC and the other constants defined by this package.

# MST

The Mountain Time Zone of North America keeps time by subtracting seven hours from Coordinated Universal Time (UTC), during the shortest days of autumn and winter (UTC−7), and by subtracting six hours during daylight saving time in the spring, summer, and early autumn (UTC−6). The clock time in this zone is based on the mean solar time at the 105th meridian west of the Greenwich Observatory. In the United States, the exact specification for the location of time zones and the dividing lines between zones is set forth in the Code of Federal Regulations at 49 CFR 71.[a]

In the United States and Canada, this time zone is generically called Mountain Time (MT). Specifically, it is Mountain Standard Time (MST) when observing standard time (fall and winter), and Mountain Daylight Time (MDT) when observing daylight saving time (spring and summer). The term refers to the fact that the Rocky Mountains, which range from northwestern Canada to the US state of New Mexico, are located almost entirely in the time zone. 

# UTC

Coordinated Universal Time, abbreviated as UTC, is the primary time standard by which the world regulates clocks and time. It is within about 1 second of mean solar time at 0° longitude; it does not observe daylight saving time. It is one of several closely related successors to Greenwich Mean Time (GMT). For most purposes, UTC is considered interchangeable with GMT, but GMT is no longer precisely defined by the scientific community.

# GMT

Greenwich Mean Time (GMT) is the mean solar time at the Royal Observatory in Greenwich, London. 

GMT was formerly used as the international civil time standard, now superseded in that function by Coordinated Universal Time (UTC). 

Today GMT is considered equivalent to UTC for UK civil purposes (but this is not formalised) and for navigation is considered equivalent to UT1 (the modern form of mean solar time at 0° longitude); these two meanings can differ by up to 0.9 s. Consequently, the term GMT should not be used for precise purposes.[1]

Because of Earth's uneven speed in its elliptical orbit and its axial tilt, noon (12:00:00) GMT is rarely the exact moment the sun crosses the Greenwich meridian and reaches its highest point in the sky there. This event may occur up to 16 minutes before or after noon GMT, a discrepancy calculated by the equation of time. Noon GMT is the annual average (i.e., "mean") moment of this event, which accounts for the word "mean" in "Greenwich Mean Time".