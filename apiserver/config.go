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
		DataBaseURL: "host=localhost user=postgres password=87770738522q dbname=myrest sslmode=disable",
		SecretKey:   "MySecretKey",
	}
}
