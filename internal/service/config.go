package service

var Config struct {
	Server struct {
		Port int    `env:"PORT" envDefault:"8080"`
		Name string `env:"SERVICE_NAME" envDefault:"service-name"`
	} `envPrefix:"SERVER_"`
}
