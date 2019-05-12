// +build !windows,!nacl,!plan9,!darwin

package loggernative

import (
	"log/syslog"
	"sync"

	logger "github.com/Nivl/go-logger"
)

// we make sure Logger implements logger.Logger
var _ logger.Logger = (*Logger)(nil)

// New returns a logger using syslog
func New() (logger.Logger, error) {
	log, err := syslog.New(syslog.LOG_ERR|syslog.LOG_INFO|syslog.LOG_DEBUG|syslog.LOG_NOTICE, "")
	if err != nil {
		return nil, err
	}

	return &Logger{
		syslog: log,
	}, nil
}

// Logger represents a logger using Windows' Events logger
type Logger struct {
	sync.RWMutex
	syslog  *syslog.Writer
	isClose bool
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
	return l.syslog.Close()
}

// IsClosed returns wether the logger is closed or not
func (l *Logger) IsClosed() bool {
	l.RLock()
	defer l.RUnlock()
	return l.isClose
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.Lock()
	defer l.Unlock()

	l.syslog.Err(msg)
}

// Info logs a message that may be helpful, but isn’t essential,
// for troubleshooting
func (l *Logger) Info(msg string) {
	l.Lock()
	defer l.Unlock()

	l.syslog.Info(msg)
}

// Debug logs a message that is intended for use in a development
func (l *Logger) Debug(msg string) {
	l.Lock()
	defer l.Unlock()

	l.syslog.Debug(msg)
}

// Log logs a message that might result a failure
func (l *Logger) Log(msg string) {
	l.Lock()
	defer l.Unlock()

	l.syslog.Notice(msg)
}
