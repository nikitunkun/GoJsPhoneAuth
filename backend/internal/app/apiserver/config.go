package apiserver

type Config struct {
	port        string
	DatabaseURL string
}

func NewConfig(port string, databaseURL string) *Config {
	return &Config{
		port:        port,
		DatabaseURL: databaseURL,
	}
}
