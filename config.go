package indexmaker

// Config holds the fileserver config
type Config struct {
	recursionDepth    int
	indexName         string
	overwriteExisting bool
	skipOnError       bool
	indexRenderer     IndexRenderer
	logger  Logger
}

// DefaultConfig returns a default configuration
func DefaultConfig() Config {
	return Config{
		recursionDepth:    0,
		indexName:         "index.html",
		overwriteExisting: false,
		indexRenderer:     DefaultRenderer,
		logger: &NullLogger{},
	}
}

// Option is influencing the config
type Option func(*Config)

// RecursionDepth sets the recursion depth
func RecursionDepth(d int) Option {
	return func(c *Config) {
		c.recursionDepth = d
	}
}

// IndexName sets the name of the index file
func IndexName(name string) Option {
	return func(c *Config) {
		c.indexName = name
	}
}

// OverwriteExisting defines whether an index should be overwritten if it exists
func OverwriteExisting(overwriteExisting bool) Option {
	return func(c *Config) {
		c.overwriteExisting = overwriteExisting
	}
}

// SkipOnError makes Make() either crash or skip on error
func SkipOnError(skipOnError bool) Option {
	return func(c *Config) {
		c.skipOnError = skipOnError
	}
}

// WithRenderer changes the renderer function
func WithRenderer(ir IndexRenderer) Option {
	return func(c *Config) {
		c.indexRenderer = ir
	}
}

// WithLogger changes the renderer function
func WithLogger(l Logger) Option {
	return func(c *Config) {
		c.logger = l
	}
}
