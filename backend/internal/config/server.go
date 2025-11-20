package config

type ServerConfig struct {
	Port           string
	Env            string
	AllowedOrigins string
}

func loadServerConfig() ServerConfig {
	return ServerConfig{
		Port:           getString("PORT", "8080"),
		Env:            getString("ENV", "development"),
		AllowedOrigins: getString("ALLOWED_ORIGINS", "http://localhost:5173"),
	}
}