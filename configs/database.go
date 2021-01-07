package configs

import "time"

var (
	DBHost        = GetString("DB_HOST", "postgres")
	DBPort        = GetInt64("DB_PORT", 5432)
	DBUser        = GetString("DB_USER", "admin")
	DBPassword    = GetString("DB_PASSWORD", "admin")
	DBDBName      = GetString("DB_NAME", "cubicasa_test")
	DBMaxOpenConn = GetInt64("MAX_OPEN_CONN", 10)
	DBMaxIdleConn = GetInt64("MAX_IDLE_CONN", 5)
	DBMaxLifeTime = time.Duration(GetInt64("MAX_LIFE_TIME", 10))
	DBShowQuery   = GetBool("SHOW_QUERY", true)
)
