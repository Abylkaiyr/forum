package configAddr

type Config struct {
	Addr string
	Port int
}

func SetConfig() *Config {
	return &Config{
		Addr: "localhost",
		Port: 8080,
	}
}
