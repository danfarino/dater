package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/justincampbell/timeago"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "")
	}

	for _, arg := range os.Args[1:] {
		showDate(arg)
	}
}

func showDate(input string) {
	tm := parseInput(input)

	if tm.IsZero() {
		_, _ = fmt.Fprintln(os.Stderr, "ERROR: unable to parse input")
		os.Exit(1)
	}

	fromNow := time.Duration(math.Abs(float64(time.Since(tm))))

	fmt.Printf("Local: %s\nUTC:   %s\nEpoch sec: %d\n%s; %v; %.2f days total\n",
		humanFormatTime(tm.Local()),
		humanFormatTime(tm.UTC()),
		tm.Unix(),
		timeago.FromTime(tm),
		fromNow.Round(time.Second),
		fromNow.Hours()/24,
	)
}

var isDigitsRegexp = regexp.MustCompile(`^\d+$`)

func parseInput(input string) time.Time {
	if input == "" {
		return time.Now()
	}

	// Try as epoch seconds
	if isDigitsRegexp.MatchString(input) {
		val, err := strconv.Atoi(input)
		if err != nil {
			return time.Time{}
		}

		return time.Unix(int64(val), 0)
	}

	parsed, err := time.Parse(time.RFC3339, input)
	if err == nil {
		return parsed
	}

	parsed, err = time.Parse(time.RFC1123Z, input)
	if err == nil {
		return parsed
	}

	// OpenSSL cert output
	parsed, err = time.Parse("Jan _2 15:04:05 2006 MST", input)
	if err == nil {
		return parsed
	}

	parsed, err = time.Parse("Mon Jan _2 15:04:05 MST 2006", input)
	if err == nil {
		return parsed
	}

	parsed, err = time.Parse("2006-01-02 15:04:05 MST", input)
	if err == nil {
		return parsed
	}

	parsed, err = time.ParseInLocation("2006-01-02 15:04:05", input, time.Local)
	if err == nil {
		return parsed
	}

	parsed, err = time.ParseInLocation("3:04PM", strings.ToUpper(input), time.Local)
	if err == nil {
		return parsed
	}

	// Try as a duration
	dur, err := time.ParseDuration(input)
	if err == nil {
		return time.Now().Add(dur)
	}

	return time.Time{}
}

func humanFormatTime(t time.Time) string {
	return t.Format(time.RFC1123Z)
}
