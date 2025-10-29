package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func New() zerolog.Logger {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "15:04:05.000",
	}

	writer.FormatLevel = func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case zerolog.LevelInfoValue: l = "✅"
			case zerolog.LevelWarnValue: l = "⚠️"
			case zerolog.LevelErrorValue: l = "❌"
			case zerolog.LevelFatalValue: l = "💀"
			case zerolog.LevelDebugValue: l = "🐞"
			default: l = "➡️"
			}
		}
		return fmt.Sprintf("| %s |", l)
	}

	writer.FormatCaller = func(i interface{}) string {
		var caller string
		if c, ok := i.(string); ok {
			parts := strings.Split(c, "/")
			if len(parts) > 2 {
				caller = strings.Join(parts[len(parts)-2:], "/")
			} else {
				caller = c
			}
		}
		return fmt.Sprintf("%-25s >", caller)
	}

	return zerolog.New(writer).
		With().
		Timestamp().
		Caller().
		Logger()
}