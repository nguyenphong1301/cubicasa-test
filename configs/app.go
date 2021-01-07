package configs

var (
	Environment = GetString("ENVIRONMENT", "staging")
	LogLevel    = GetInt64("LOG_LEVEL", 7)
	Port        = GetInt64("APP_PORT", 80)
)
