package loggernative

/*
#cgo CFLAGS: -x objective-c -mmacosx-version-min=10.10
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>
#ifdef AVAILABLE_MAC_OS_X_VERSION_10_12_AND_LATER
	#import <os/log.h>
#endif
void logDefault(const char *message) {
	if (@available(macOS 10.12, *)) {
		os_log(OS_LOG_DEFAULT, "%{public}s", message);
	}
}
void logInfo(const char *message) {
	if (@available(macOS 10.12, *)) {
		os_log_info(OS_LOG_DEFAULT, "%{public}s", message);
	}
}
void logError(const char *message) {
	if (@available(macOS 10.12, *)) {
		os_log_error(OS_LOG_DEFAULT, "%{public}s", message);
	}
}
void logDebug(const char *message) {
	if (@available(macOS 10.12, *)) {
		os_log_debug(OS_LOG_DEFAULT, "%{public}s", message);
	}
}
*/
import (
	"C"
)
import (
	"sync"
	"unsafe"

	logger "github.com/Nivl/go-logger"
)

// we make sure Logger implements logger.Logger
var _ logger.Logger = (*Logger)(nil)

// New returns a logger using ASL
func New() (logger.Logger, error) {
	return &Logger{}, nil
}

// Logger represents a logger using ASL on macOS 10.12+, will be a noop
// on all previous versions.
// macOS < 10.10 are not supported
type Logger struct {
	sync.RWMutex
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
	return nil
}

// IsClosed returns wether the logger is closed or not
func (l *Logger) IsClosed() bool {
	l.RLock()
	defer l.RUnlock()
	return l.isClose
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	C.logError(cMsg)
}

// Info logs a message that may be helpful, but isn’t essential,
// for troubleshooting
func (l *Logger) Info(msg string) {
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	C.logInfo(cMsg)
}

// Debug logs a message that is intended for use in a development
func (l *Logger) Debug(msg string) {
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	C.logDebug(cMsg)
}

// Log logs a message that might result a failure
func (l *Logger) Log(msg string) {
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	C.logDefault(cMsg)
}
