package settings

import (
	"os"
)

func get_db_setting() DBSetting {
	// Чтение настроек из переменных окружения

	database_user, exists_user := os.LookupEnv("DATABASE_USER")
	database_password, exists_password := os.LookupEnv("DATABASE_PASSWORD")
	database_name, exists_name := os.LookupEnv("DATABASE_NAME")

	if !exists_user || !exists_password || !exists_name {
		panic("Key environment variables are missing")
	}
	return DBSetting{username: database_user, password: database_password, name: database_name}
}
