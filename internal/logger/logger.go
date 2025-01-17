package logger

import (
	"log"
	"log/slog"
	"os"
	"time"

	slogformatter "github.com/samber/slog-formatter"
	sloghttp "github.com/samber/slog-http"
	"github.com/spf13/pflag"
)

var loggerLevels = map[string]slog.Level{
	"info":  slog.LevelInfo,
	"debug": slog.LevelDebug,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

var opts struct {
	http struct {
		port   int
		secure bool
	}
	featureToggles struct {
		configFile string
	}
	logger struct {
		style string
		level string
	}
}

func GetLogger() sloghttp.Config {
	pflag.StringVar(
		&opts.featureToggles.configFile,
		"feature-config",
		"featuretoggles.json",
		"Location of a feature toggle",
	)
	pflag.IntVar(&opts.http.port, "http-port", 8000, "The HTTP port to run the server on")
	pflag.BoolVar(&opts.http.secure, "secure", false, "Enable HTTPS on HTTP server")
	pflag.StringVar(
		&opts.logger.level,
		"logger-level",
		"info",
		"Set logger level ['debug', 'info', 'warn', 'error']",
	)
	pflag.StringVar(
		&opts.logger.style,
		"logger-style",
		"text",
		"Set logger style ['json', 'text']",
	)
	pflag.Parse()

	logLevel, ok := loggerLevels[opts.logger.level]
	if !ok {
		log.Fatalf(
			"invalid log level supplied: %s not in ['debug', 'info', 'warn', 'error']",
			logLevel,
		)
	}

	// my desired UTC logging with compiler-arg-defined style
	var handler slog.Handler
	switch opts.logger.style {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: &logLevel})
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: &logLevel})
	default:
		log.Fatalf("invalid log style supplied: %s", opts.logger.style)
	}
	logger := slog.New(
		slogformatter.NewFormatterHandler(
			slogformatter.TimezoneConverter(time.UTC),
			slogformatter.TimeFormatter(time.DateTime, nil),
		)(handler),
	)
	slog.SetDefault(logger)
	loggerConfig := sloghttp.Config{
		WithSpanID:  true,
		WithTraceID: true,
	}

	return loggerConfig
}
