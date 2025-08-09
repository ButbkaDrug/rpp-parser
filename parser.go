package rppparser

import (
	"errors"
	"log/slog"
	"os"
	"rpp-parser/pkg/lexer"
	"rpp-parser/pkg/parser"
)

type options struct {
	logger *slog.Logger
}

type Options func(options *options) error

func WithLogger(l *slog.Logger) Options {
	return func(options *options) error {
		if l == nil {
			return errors.New("nil pointer not allowed")
		}

		options.logger = l

		return nil

	}
}

func Parse(path string, opts ...Options) (parser.Project, error) {
	var options options

	for _, opt := range opts {
		err := opt(&options)

		if err != nil {
			return parser.Project{}, err
		}
	}

	if options.logger == nil {
		logfile, err := os.Create("/tmp/reaper-parser.log")

		if err != nil {
			slog.Error("prr-parser: failed to creat a log file", "error", err.Error())
			logfile = os.Stdout
		}
		options.logger = slog.New(slog.NewTextHandler(logfile, &slog.HandlerOptions{
			Level: slog.LevelError,
		}))
	}

	slog.SetDefault(options.logger)

	fileBytes, err := os.ReadFile(path)

	if err != nil {
		slog.Error("rpp-parser.Parse() failed:", "error", err.Error())
		return parser.Project{}, err
	}

	fileString := string(fileBytes)

	l := lexer.NewLexer(fileString)
	p := parser.NewParser(l)
	proj := &parser.Project{}

	p.ParseProject(proj)
	proj.Path = path

	return *proj, nil
}

func ParseBytesIntoStruct(data []byte, proj *parser.Project) error {

	logfile, err := os.Create("/tmp/reaper-parser.log")

	if err != nil {
		slog.Error("prr-parser: failed to creat a log file", "error", err.Error())
		logfile = os.Stdout
	}

	_ = slog.New(slog.NewTextHandler(logfile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	fileString := string(data)

	l := lexer.NewLexer(fileString)
	p := parser.NewParser(l)

	p.ParseProject(proj)

	return nil
}
