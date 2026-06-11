package logging

import (
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/DeRuina/timberjack"
	"github.com/sirupsen/logrus"
)

type LogSettings struct {
	LogFile    string  `yaml:"log_file"`
	MaxSize    int     `yaml:"max_size"`
	MaxBackups int     `yaml:"max_backups"`
	MaxAge     int     `yaml:"max_age"`
	LogLevel   *string `yaml:"log_level"`
	LogOutput  *string `yaml:"log_output,omitempty"` // "stdout" or "stderr"
}

// LoggerOption defines a functional option for configuring a logger.
type LoggerOption func(*logrus.Logger)

// WithOutput sets the output destination for the logger.
func WithOutput(w io.Writer) LoggerOption {
	return func(l *logrus.Logger) {
		l.SetOutput(w)
	}
}

// NewLogger creates and configures a new logrus.Logger based on the provided configuration and options.
func NewLogger(cfg *LogSettings, opts ...LoggerOption) (*logrus.Logger, error) {
	logger := logrus.New()

	// Set Log Level
	logLevel := logrus.InfoLevel
	if cfg.LogLevel != nil && *cfg.LogLevel != "" {
		if lv, err := logrus.ParseLevel(strings.ToLower(*cfg.LogLevel)); err == nil {
			logLevel = lv
		}
	}
	logger.SetLevel(logLevel)

	// Setup Output
	var output io.Writer = os.Stdout
	if cfg.LogOutput != nil {
		switch *cfg.LogOutput {
		case "stderr":
			output = os.Stderr
		case "stdout":
			output = os.Stdout
		}
	}

	// If file logging is enabled, create a multi-writer to log to both the configured output and the file.
	if cfg.LogFile != "" {
		fileLogger := &timberjack.Logger{
			Filename:   cfg.LogFile,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
		}
		output = io.MultiWriter(output, fileLogger)
		logrus.New().Infof("File logging enabled, writing to %s", cfg.LogFile)
	}
	logger.SetOutput(output)

	// Apply functional options
	for _, opt := range opts {
		opt(logger)
	}

	// Set Formatter
	textFormatter := &logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", ""
		},
		ForceColors: true,
	}

	// Wrap with our custom source formatter
	logger.SetFormatter(&SourceFormatter{
		Underlying: textFormatter,
		AddSpace:   true,
	})

	// Set Caller Reporting
	logger.SetReportCaller(true)

	return logger, nil
}
