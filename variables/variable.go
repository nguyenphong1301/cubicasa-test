package variables

/**
this package to storage all global variables for convenience
*/

import (
	"cubicasa/configs"
	"cubicasa/libs/database"
	"cubicasa/libs/logger"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	PostgresDB *gorm.DB
	Logger     logger.LogInterface
)

func Init() error {

	// init global logger
	Logger = &logger.LogRus{}
	Logger.Init(logger.LogConfig{
		Level: int(configs.LogLevel),
	})

	var err error = nil

	// init postgres connection
	if PostgresDB, err = InitPostgres(); err != nil {
		Logger.Error("... postgres connect failed with error ", err.Error())
		return err
	}
	log.Println("... postgres connected")

	return nil
}
func DeInit() {
	if PostgresDB != nil {
		PostgresDB.Close()
		log.Println("close postgres connection")
	}

}
func InitPostgres() (*gorm.DB, error) {
	connStr := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable binary_parameters=yes`,
		configs.DBHost,
		configs.DBPort,
		configs.DBUser,
		configs.DBPassword,
		configs.DBDBName)
	log.Println("connecting to ", connStr, " ...")
	return database.NewDatabase(
		"postgres",
		connStr,
		int(configs.DBMaxOpenConn),
		int(configs.DBMaxIdleConn),
		configs.DBMaxLifeTime,
		configs.DBShowQuery,
	)
}
