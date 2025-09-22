package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Logger wraps the standard log package with additional functionality
type Logger struct {
	*log.Logger
	mainLogFile   *os.File
	errorLogFile  *os.File
	debugLogFile  *os.File
	timingLogFile *os.File
}

// NewLogger creates a new logger with organized output
func NewLogger() *Logger {
	// Ensure output directory exists
	outputDir := "output/logs"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Create timestamp for unique log files
	timestamp := time.Now().Format("2006-01-02_15-04-05")

	// Create main log file
	mainLogFile, err := os.OpenFile(
		filepath.Join(outputDir, fmt.Sprintf("active-goference_%s.log", timestamp)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Create error log file
	errorLogFile, err := os.OpenFile(
		filepath.Join(outputDir, fmt.Sprintf("errors_%s.log", timestamp)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Create debug log file
	debugLogFile, err := os.OpenFile(
		filepath.Join(outputDir, fmt.Sprintf("debug_%s.log", timestamp)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Create timing log file
	timingLogFile, err := os.OpenFile(
		filepath.Join(outputDir, fmt.Sprintf("timing_%s.log", timestamp)),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Multi-writer: main file and stdout
	multiWriter := io.MultiWriter(mainLogFile, os.Stdout)

	logger := log.New(multiWriter, "ACTIVE-GOFERENCE: ", log.LstdFlags|log.Lshortfile)

	return &Logger{
		Logger:        logger,
		mainLogFile:   mainLogFile,
		errorLogFile:  errorLogFile,
		debugLogFile:  debugLogFile,
		timingLogFile: timingLogFile,
	}
}

// Info logs an info message
func (l *Logger) Info(args ...interface{}) {
	l.Logger.SetPrefix("INFO: ")
	l.Logger.Println(args...)
}

// Error logs an error message
func (l *Logger) Error(args ...interface{}) {
	l.Logger.SetPrefix("ERROR: ")
	l.Logger.Println(args...)

	// Also write to error log file
	errorLogger := log.New(l.errorLogFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
	errorLogger.Println(args...)
}

// Debug logs a debug message
func (l *Logger) Debug(args ...interface{}) {
	l.Logger.SetPrefix("DEBUG: ")
	l.Logger.Println(args...)

	// Also write to debug log file
	debugLogger := log.New(l.debugLogFile, "DEBUG: ", log.LstdFlags|log.Lshortfile)
	debugLogger.Println(args...)
}

// Warn logs a warning message
func (l *Logger) Warn(args ...interface{}) {
	l.Logger.SetPrefix("WARN: ")
	l.Logger.Println(args...)

	// Also write to error log file
	errorLogger := log.New(l.errorLogFile, "WARN: ", log.LstdFlags|log.Lshortfile)
	errorLogger.Println(args...)
}

// Timing logs the duration of an operation
func (l *Logger) Timing(operation string, start time.Time) {
	duration := time.Since(start)
	l.Info(operation, "completed in", duration)

	// Also write to timing log file
	timingLogger := log.New(l.timingLogFile, "TIMING: ", log.LstdFlags|log.Lshortfile)
	timingLogger.Println(fmt.Sprintf("%s completed in %v", operation, duration))
}

// Trace logs detailed trace information
func (l *Logger) Trace(args ...interface{}) {
	// Create a new slice with "TRACE:" prepended
	traceArgs := append([]interface{}{"TRACE:"}, args...)
	l.Debug(traceArgs...)
}

// Close closes all log files
func (l *Logger) Close() error {
	l.mainLogFile.Close()
	l.errorLogFile.Close()
	l.debugLogFile.Close()
	l.timingLogFile.Close()
	return nil
}

// Global logger instance
var Log = NewLogger()
