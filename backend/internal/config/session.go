package config

type SessionConfig struct {
	SessionSecret      string
	SessionMaxAge      int
	SessionIdleTimeout int
}

func loadSessionConfig() SessionConfig {
	return SessionConfig{
		SessionSecret:      getString("SESSION_SECRET", ""),
		SessionMaxAge:      getInt("SESSION_MAX_AGE", 1800),
		SessionIdleTimeout: getInt("SESSION_IDLE_TIMEOUT", 900),
	}
}