package apiserver

// Config ...
type Config struct {
	BindAddr    string
	DataBaseURL string
	SecretKey   string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		DataBaseURL: "host=db user=postgres password=giftexpass dbname=giftex sslmode=disable",
		SecretKey:   "MySecretKey",
	}
}
