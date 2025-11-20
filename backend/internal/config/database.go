package config

type DatabaseConfig struct {
	DatabaseURL string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
}

func loadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DatabaseURL: getString("DATABASE_URL", ""),
		DBHost:      getString("DB_HOST", "postgres"),
		DBPort:      getString("DB_PORT", "5432"),
		DBUser:      getString("DB_USER", "aegis"),
		DBPassword:  getString("DB_PASSWORD", "aegis"),
		DBName:      getString("DB_NAME", "aegisdb"),
	}
}