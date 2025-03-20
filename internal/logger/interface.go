package logger

// BaseLogger used logger interface definition.
type BaseLogger interface {
	// Infof posts message on log 'info' Level.
	Infof(msg string, fields ...interface{})
	// Warnf posts message on log 'warn' Level.
	Warnf(msg string, fields ...interface{})
	// Errorf posts message on log 'error' Level.
	Errorf(msg string, fields ...interface{})
	// Fatalf posts message on log 'fatal' Level.
	Fatalf(msg string, fields ...interface{})
}
