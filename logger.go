package ronin

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"go.uber.org/fx/fxevent"
)

type fxLogger struct {
	l zerolog.Logger
}

func (f fxLogger) Write(p []byte) (n int, err error) {
	// from https://github.com/rs/zerolog/blob/a9a8199d2dd3578d37e459618515f34b5e917f8d/log.go#L435-L441
	n = len(p)
	if n > 0 && p[n-1] == '\n' {
		// Trim CR added by stdlog.
		p = p[0 : n-1]
	}
	p = []byte(strings.Replace(string(p), "\t", " ", -1))
	f.l.Debug().Msgf("%s", p)
	return
}

var _ io.Writer = (*fxLogger)(nil)

func Logger() fxevent.Logger {
	logger := fxLogger{
		l: log.Logger.
			With().
			Str("evt.name", "fx.init").
			Logger(),
	}
	return &fxevent.ConsoleLogger{
		W: logger,
	}
}

func InitLog(stage string) error {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.DurationFieldUnit = time.Nanosecond
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Equivalent of Lshortfile
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	var writer io.Writer
	if strings.EqualFold(stage, "development") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		writer = zerolog.ConsoleWriter{
			Out:        colorable.NewColorable(os.Stderr),
			TimeFormat: time.RFC3339Nano,
		}
	} else {
		writer = os.Stdout
	}

	log.Logger = zerolog.New(writer).With().Timestamp().Caller().Logger()
	return nil
}
