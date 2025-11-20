package config

type LimiterConfig struct {
	LoginRateLimitMaxAttempts int
	LoginRateLimitWindowSec   int
	LoginLockoutDurationMin   int
}

func loadLimiterConfig() LimiterConfig {
	return LimiterConfig{
		LoginRateLimitMaxAttempts: getInt("LOGIN_RATE_LIMIT_MAX_ATTEMPTS", 10),
		LoginRateLimitWindowSec:   getInt("LOGIN_RATE_LIMIT_WINDOW_SEC", 3600),
		LoginLockoutDurationMin:   getInt("LOGIN_LOCKOUT_DURATION_MIN", 300),
	}
}