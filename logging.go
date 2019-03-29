package indexmaker

// Logger is used to make logging pluggable, works best with github.com/sirupsen/logrus
type Logger interface {
	Info (...interface{})
	Infof (string,  ...interface{})
	Warn (...interface{})
	Warnf (string,  ...interface{})
	Debug (...interface{})
	Debugf (string,  ...interface{})
}

// NullLogger empty
type NullLogger struct {}

// Info ...
func (n *NullLogger) Info(args ...interface{}){}

// Infof ...
func (n *NullLogger) Infof(format string, args ...interface{}){}

// Warn ...
func (n *NullLogger) Warn(args ...interface{}){}

// Warnf ...
func (n *NullLogger) Warnf(format string, args ...interface{}){}

// Debug ...
func (n *NullLogger) Debug(args ...interface{}){}

// Debugf ...
func (n *NullLogger) Debugf(format string, args ...interface{}){}
