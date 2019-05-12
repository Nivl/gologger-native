package loggernative

import (
	"sync"

	"github.com/Nivl/go-logger"
	"golang.org/x/sys/windows/svc/eventlog"
)

// we make sure Logger implements logger.Logger
var _ logger.Logger = (*Logger)(nil)

// New returns a logger using Windows' Events logger
func New() (logger.Logger, error) {
	log, err := eventlog.Open("Application")
	if err != nil {
		return nil, err
	}

	return &Logger{
		eventLogger: log,
	}, nil
}

// Logger represents a logger using Windows' Events logger
type Logger struct {
	sync.RWMutex
	eventLogger *eventlog.Log
	isClose     bool
}

// ID returns the logger's unique ID
func (l *Logger) ID() string {
	return "native-logger"
}

// Close frees any resource allocated by the logger
// the logger may not be reusable after being closed
func (l *Logger) Close() error {
	l.Lock()
	defer l.Unlock()

	l.isClose = true
	return l.eventLogger.Close()
}

// IsClosed returns wether the logger is closed or not
func (l *Logger) IsClosed() bool {
	l.RLock()
	defer l.RUnlock()
	return l.isClose
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.eventLogger.Error(1, msg)
}

// Info logs a message that may be helpful, but isn’t essential,
// for troubleshooting
func (l *Logger) Info(msg string) {
	l.eventLogger.Info(2, msg)
}

// Debug logs a message that is intended for use in a development
func (l *Logger) Debug(msg string) {
	l.eventLogger.Warning(3, msg)
}

// Log logs a message that might result a failure
func (l *Logger) Log(msg string) {
	l.eventLogger.Warning(4, msg)
}
