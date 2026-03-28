package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

// Time represents a time value in nanoseconds since the Unix epoch.
type Time int64

// DateFormat represents a date format preference.
type DateFormat int

// Date format preferences.
const (
	DATE_FORMAT_YYYYMMDD DateFormat = C.SDL_DATE_FORMAT_YYYYMMDD
	DATE_FORMAT_DDMMYYYY DateFormat = C.SDL_DATE_FORMAT_DDMMYYYY
	DATE_FORMAT_MMDDYYYY DateFormat = C.SDL_DATE_FORMAT_MMDDYYYY
)

// TimeFormat represents a time format preference.
type TimeFormat int

// Time format preferences.
const (
	TIME_FORMAT_24HR TimeFormat = C.SDL_TIME_FORMAT_24HR
	TIME_FORMAT_12HR TimeFormat = C.SDL_TIME_FORMAT_12HR
)

// DateTime represents a calendar date and time.
type DateTime struct {
	Year       int32
	Month      int32
	Day        int32
	Hour       int32
	Minute     int32
	Second     int32
	Nanosecond int32
	DayOfWeek  int32
	UTCOffset  int32
}

// GetCurrentTime returns the current time.
func GetCurrentTime() (Time, error) {
	var t C.SDL_Time
	if !C.SDL_GetCurrentTime(&t) {
		return 0, getError()
	}
	return Time(t), nil
}

// TimeToDateTime converts a time value to a calendar date and time.
func TimeToDateTime(ticks Time, localTime bool) (DateTime, error) {
	var cdt C.SDL_DateTime
	if !C.SDL_TimeToDateTime(C.SDL_Time(ticks), &cdt, C.bool(localTime)) {
		return DateTime{}, getError()
	}
	return DateTime{
		Year:       int32(cdt.year),
		Month:      int32(cdt.month),
		Day:        int32(cdt.day),
		Hour:       int32(cdt.hour),
		Minute:     int32(cdt.minute),
		Second:     int32(cdt.second),
		Nanosecond: int32(cdt.nanosecond),
		DayOfWeek:  int32(cdt.day_of_week),
		UTCOffset:  int32(cdt.utc_offset),
	}, nil
}

// GetDaysInMonth returns the number of days in a month.
func GetDaysInMonth(year, month int) int {
	return int(C.SDL_GetDaysInMonth(C.int(year), C.int(month)))
}

// GetDayOfYear returns the day of year for a date.
func GetDayOfYear(year, month, day int) int {
	return int(C.SDL_GetDayOfYear(C.int(year), C.int(month), C.int(day)))
}

// GetDayOfWeek returns the day of week for a date.
func GetDayOfWeek(year, month, day int) int {
	return int(C.SDL_GetDayOfWeek(C.int(year), C.int(month), C.int(day)))
}

// GetDateTimeLocalePreferences returns the user's preferred date and time formats.
func GetDateTimeLocalePreferences() (DateFormat, TimeFormat, error) {
	var df C.SDL_DateFormat
	var tf C.SDL_TimeFormat
	if !C.SDL_GetDateTimeLocalePreferences(&df, &tf) {
		return 0, 0, getError()
	}
	return DateFormat(df), TimeFormat(tf), nil
}

// DateTimeToTime converts a DateTime to a Time value.
func DateTimeToTime(dt DateTime) (Time, error) {
	cdt := C.SDL_DateTime{
		year: C.int(dt.Year), month: C.int(dt.Month), day: C.int(dt.Day),
		hour: C.int(dt.Hour), minute: C.int(dt.Minute), second: C.int(dt.Second),
		nanosecond: C.int(dt.Nanosecond), day_of_week: C.int(dt.DayOfWeek), utc_offset: C.int(dt.UTCOffset),
	}
	var t C.SDL_Time
	if !C.SDL_DateTimeToTime(&cdt, &t) {
		return 0, getError()
	}
	return Time(t), nil
}

// TimeToWindows converts a Time value to Windows FILETIME components.
func TimeToWindows(ticks Time) (dwLow, dwHigh uint32) {
	var lo, hi C.Uint32
	C.SDL_TimeToWindows(C.SDL_Time(ticks), &lo, &hi)
	return uint32(lo), uint32(hi)
}

// TimeFromWindows converts Windows FILETIME components to a Time value.
func TimeFromWindows(dwLow, dwHigh uint32) Time {
	return Time(C.SDL_TimeFromWindows(C.Uint32(dwLow), C.Uint32(dwHigh)))
}
