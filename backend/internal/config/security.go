package config

type SecurityConfig struct {
	Argon2TimeCost    uint32
	Argon2MemoryCost  uint32
	Argon2Parallelism uint8
	Argon2KeyLength   uint32

	MinPasswordLength   int
	RequireUppercase    bool
	RequireLowercase    bool
	RequireNumbers      bool
	RequireSpecialChars bool
}

func loadSecurityConfig() SecurityConfig {
	return SecurityConfig{
		Argon2TimeCost:    uint32(getInt("ARGON2_TIME_COST", 4)),
		Argon2MemoryCost:  uint32(getInt("ARGON2_MEMORY_COST", 65536)),
		Argon2Parallelism: uint8(getInt("ARGON2_PARALLELISM", 4)),
		Argon2KeyLength:   uint32(getInt("ARGON2_KEY_LENGTH", 32)),

		MinPasswordLength:   getInt("MIN_PASSWORD_LENGTH", 14),
		RequireUppercase:    getBool("REQUIRE_UPPERCASE", true),
		RequireLowercase:    getBool("REQUIRE_LOWERCASE", true),
		RequireNumbers:      getBool("REQUIRE_NUMBERS", true),
		RequireSpecialChars: getBool("REQUIRE_SPECIAL_CHARS", true),
	}
}