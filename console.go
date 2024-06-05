package console_go

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// Define color constants
const (
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Blue   = "\033[34m"
	White  = "\033[37m"
	Reset  = "\033[0m" // Reset color
)

// Logger interface
type Logger interface {
	Log(format string, args ...interface{})
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Error(format string, args ...interface{})
	Warn(format string, args ...interface{})
}

// Package-level logger settings structured as a singleton
var loggerSettings = struct {
	EnableLog   bool `json:"enable_log"`
	EnableInfo  bool `json:"enable_info"`
	EnableDebug bool `json:"enable_debug"`
	EnableError bool `json:"enable_error"`
	EnableWarn  bool `json:"enable_warn"`
	Output      io.Writer
}{
	EnableLog:   true,
	EnableInfo:  true,
	EnableDebug: true,
	EnableError: true,
	EnableWarn:  true,
	Output:      os.Stdout,
}

var loadOnce sync.Once

// loadConfigLoads the configuration from the file
func loadConfig() {
	// fmt.Println("loading config.json file")
	configFile, err := os.Open("config.json")
	if err != nil {
		// fmt.Fprintln(os.Stderr, "Failed to open config file:", err)
		return
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&loggerSettings); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to decode config file:", err)
	}
}

// getCallerInfo retrieves the file and line number of the caller
func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown:0"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

// log prints a log message with color if the corresponding log level is enabled
func log(color, caller, format string, args ...interface{}) {
	loadOnce.Do(loadConfig)

	// Print current time
	if _, err := fmt.Fprintf(loggerSettings.Output, "%s %s: ", time.Now().Format("15:04:05"), caller); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write log time:", err)
		return
	}
	// Print log color
	if _, err := fmt.Fprint(loggerSettings.Output, color); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write log color:", err)
		return
	}
	// Print log message
	if _, err := fmt.Fprintf(loggerSettings.Output, format+"\n", args...); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write log message:", err)
		return
	}
	// Reset log color
	if _, err := fmt.Fprint(loggerSettings.Output, Reset); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to reset log color:", err)
	}
}

// Log logs a message if logging is enabled
func Log(format string, args ...interface{}) {
	if loggerSettings.EnableLog {
		log(Green, getCallerInfo(), format, args...)
	}
}

// Info logs an info message if info logging is enabled
func Info(format string, args ...interface{}) {
	if loggerSettings.EnableInfo {
		log(Blue, getCallerInfo(), format, args...)
	}
}

// Debug logs a debug message if debug logging is enabled
func Debug(format string, args ...interface{}) {
	if loggerSettings.EnableDebug {
		log(White, getCallerInfo(), format, args...)
	}
}

// Warn logs a warning message if warning logging is enabled
func Warn(format string, args ...interface{}) {
	if loggerSettings.EnableWarn {
		log(Yellow, getCallerInfo(), format, args...)
	}
}

// Error logs an error message if error logging is enabled
func Error(format string, args ...interface{}) {
	if loggerSettings.EnableError {
		log(Red, getCallerInfo(), format, args...)
	}
}

// Fatal panics always
func Fatal(err error) {
	panic(err)
}
