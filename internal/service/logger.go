package service

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const defaultTimeFormat = time.RFC3339Nano

// Capacity of writers depends on writers slice, and once on stdout write (3+1=4)
const writersCapacity = 4

const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Cyan = "\033[36m"
const DarkGray = "\033[91m"
const colorReset = "\033[0m"

type FilteredWriter struct {
	w      zerolog.LevelWriter
	levels []zerolog.Level
}

func (w *FilteredWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w *FilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	for _, filteredLevel := range w.levels {
		if level == filteredLevel {
			return w.w.WriteLevel(level, p)
		}
	}
	return len(p), nil
}

type LoggerConfig struct {
	LoggerLogPath  string
	LoggerJsonPath string
	OutOnly        bool
}

func NewLogger(keyModule string, config LoggerConfig) (zerolog.Logger, error) {
	defaultWriters := make([]io.Writer, 0, writersCapacity)

	if config.LoggerJsonPath != "" && !config.OutOnly {
		logWriter, err := configureLoggerFile(config.LoggerJsonPath)
		if err != nil {
			return zerolog.Logger{}, err
		}

		defaultWriters = append(defaultWriters, logWriter)
	}

	if config.LoggerLogPath != "" && !config.OutOnly {
		logWriter, err := configureLoggerFile(config.LoggerLogPath)
		if err != nil {
			return zerolog.Logger{}, err
		}

		// zerolog.ConsoleWriter needs to make readable logs in terminal
		logWriter = configureOutputMessage(zerolog.ConsoleWriter{Out: logWriter, TimeFormat: defaultTimeFormat})

		defaultWriters = append(defaultWriters, logWriter)
	}

	outWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: defaultTimeFormat}
	outWriter = configureOutputMessage(outWriter)
	defaultWriters = append(defaultWriters, outWriter)

	if !config.OutOnly {
		errWriter := zerolog.MultiLevelWriter(os.Stderr)
		filteredLevels := []zerolog.Level{zerolog.ErrorLevel, zerolog.PanicLevel, zerolog.FatalLevel}
		errWriter = &FilteredWriter{errWriter, filteredLevels}
		defaultWriters = append(defaultWriters, errWriter)
	}

	w := zerolog.MultiLevelWriter(defaultWriters...)

	logger := zerolog.New(w).With().Str("module", keyModule).Caller().Timestamp().Logger()

	return logger, nil
}

func configureOutputMessage(writer zerolog.ConsoleWriter) zerolog.ConsoleWriter {
	// configure time format
	zerolog.TimeFieldFormat = time.RFC3339Nano
	writer.TimeFormat = "15:04:05.000000"

	// level color configure
	writer.FormatLevel = configureFormatLevel

	// cut message if its to long or add space in its to short
	writer.FormatMessage = configureFormatMessage

	writer.FormatFieldValue = func(i interface{}) string {
		if str, ok := i.(string); ok {
			return strings.ReplaceAll(str, `\"`, `"`)
		}

		return fmt.Sprintf("%s", i)
	}

	writer.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf(Red+"%s=", i)
	}

	// formatting caller string
	zerolog.CallerMarshalFunc = configureCaller

	return writer
}

func configureFormatLevel(i interface{}) string {
	inputMessage, ok := i.(string)
	if !ok {
		fmt.Printf("WARNING! Cant convert interface to string")
	}
	level, _ := zerolog.ParseLevel(inputMessage)
	switch level {
	case zerolog.DebugLevel:
		return fmt.Sprintf(Cyan+"%-5s ➙"+colorReset, i)
	case zerolog.InfoLevel:
		return fmt.Sprintf(Green+"%-5s ➙"+colorReset, i)
	case zerolog.WarnLevel:
		return fmt.Sprintf(Yellow+"%-5s ➙"+colorReset, i)
	case zerolog.ErrorLevel, zerolog.FatalLevel, zerolog.PanicLevel:
		return fmt.Sprintf(Red+"%-5s ➙"+colorReset, i)
	case zerolog.NoLevel:
		return fmt.Sprintf(DarkGray+"%-5s ➙"+colorReset, i)
	default:
		return fmt.Sprintf("%-5s ➙", i)
	}
}

const messagePostfix = "..."

func configureFormatMessage(i interface{}) string {
	// lengthRequired means how many symbols would be in message between code pointer and params
	// this was done for greater readability
	const lengthRequired = 80
	messageLength := len(fmt.Sprintf("%s", i))
	if messageLength <= lengthRequired {
		return fmt.Sprintf("%-*s", lengthRequired, i)
	} else {
		message := fmt.Sprintf("%s", i)
		// lengthRequired-3 it is 80 symbols string with ellipsis (80-3=77  len"..."=3)
		message = message[:lengthRequired-len(messagePostfix)] + messagePostfix
		return message
	}
}

const maxFileNameLen = 8
const maxLinePointerLen = 4
const delimiter = ":"
const callerPostfix = "*"

func configureCaller(_ uintptr, file string, line int) string {
	// lengthRequired means how many symbols can contain caller info
	fileName := filepath.Base(file)
	linePointer := strconv.Itoa(line)

	if len(linePointer) > maxLinePointerLen {
		linePointer = linePointer[:maxLinePointerLen-len(callerPostfix)] + callerPostfix
	}

	restAvailableLen := (maxLinePointerLen - len(linePointer)) + maxFileNameLen
	if len(fileName) > restAvailableLen {
		fileName = fileName[:restAvailableLen-len(callerPostfix)] + callerPostfix
	}

	return fmt.Sprintf("%*s", maxFileNameLen+maxLinePointerLen+len(delimiter), fileName+delimiter+linePointer)
}

func configureLoggerFile(path string) (io.Writer, error) {
	if len(path) > 0 {
		// Ensure target directory exists
		targetDirPath := filepath.Dir(path)
		if _, err := os.Stat(targetDirPath); os.IsNotExist(err) {
			// Logs directory does not exist so create it
			err = os.MkdirAll(targetDirPath, os.ModeDir)
			if err != nil {
				fmt.Printf("Error MkdirAll: %s", err.Error())
				os.Exit(1)
			}
		}
		flag, perm := os.O_RDWR|os.O_APPEND|os.O_CREATE, os.FileMode(0644)
		if file, err := os.OpenFile(path, flag, perm); err == nil {
			return file, nil
		} else {
			fmt.Printf("WARNING: Unable to create file for logs at path: %s, error: %s\n", path, err.Error())
			return nil, err
		}
	}
	err := fmt.Errorf("cannot create file with path %s", path)
	return nil, err
}
