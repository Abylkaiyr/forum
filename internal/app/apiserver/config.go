package apiserver

// Config ...
type Config struct {
	BindAddr int
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		BindAddr: 8080,
	}
}
