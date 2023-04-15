package app

type Config struct {
	HttpAddr string
}

func NewConfig(httpAddr string) Config {
	return Config{
		HttpAddr: httpAddr,
	}
}
