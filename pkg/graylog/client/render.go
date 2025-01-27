package client

import (
	"fmt"
	"strings"

	"github.com/jeehoon/graylog-cli/pkg/timeutil"
)

const (
	Reset = "\033[0m"

	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"

	LightRed     = "\033[91m"
	LightGreen   = "\033[92m"
	LightYellow  = "\033[93m"
	LightBlue    = "\033[94m"
	LightMagenta = "\033[95m"
	LightCyan    = "\033[96m"

	White    = "\033[97m"
	DarkGray = "\033[90m"
)

func Render(dec *Decoder, useColor bool, msg map[string]any) string {
	keys, values := dec.Fields(msg)
	var fields []string
	for idx, key := range keys {
		value := values[idx]

		if useColor {
			key = Cyan + key + Reset
			value = Blue + value + Reset
		}

		fields = append(fields, fmt.Sprintf("%v:%v", key, value))
	}

	hostname := dec.Hostname(msg)

	if useColor {
		hostname = LightMagenta + hostname + Reset
	}

	lv := dec.Level(msg)
	level := lv.String()

	if useColor {
		switch lv {
		case LevelEmergency, LevelAlert, LevelCritical, LevelError:
			level = Red + level + Reset
		case LevelWarning:
			level = Yellow + level + Reset
		case LevelNotice, LevelInformational:
			level = White + level + Reset
		case LevelDebug:
			level = DarkGray + level + Reset
		}
	}

	timestamp := timeutil.Format(dec.Timestamp(msg))

	text := dec.Text(msg)

	output := fmt.Sprintln(hostname, timestamp, level, text, strings.Join(fields, " "))

	return strings.TrimSpace(output)
}
