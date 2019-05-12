# gologger-native

[![Build Status](https://travis-ci.org/Nivl/gologger-native.svg)](https://travis-ci.org/Nivl/gologger-native)
[![Go Report Card](https://goreportcard.com/badge/github.com/nivl/gologger-native)](https://goreportcard.com/report/github.com/nivl/gologger-native)
[![GoDoc](https://godoc.org/github.com/Nivl/gologger-native?status.svg)](https://godoc.org/github.com/Nivl/gologger-native)

gologger-native contains native logger implementations for go-logger

## Implementations

### Mac OS (requires cgo)

On MacOS, gologger-native will use [ASL](https://developer.apple.com/documentation/os/logging) on macOS 10.12+
and will be a noop for all previous versions. This lib targets MacOS 10.10 and above only.

### Windows

On Windows, the Event Logger will be used.

### Linux, BSD, Solaris ...

On all other systems, the logs are sent to syslog.

**Note that this package is not compatible with nacl and plan9**
