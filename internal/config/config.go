package config

// Тут поля для настройки всего проекта
// TODO .env
const (

	///////////////////////////////////////////////
	// Data base and migration
	///////////////////////////////////////////////
	DB_HOST     = "localhost"
	DB_PORT     = "5430"
	DB_USER     = "postgres_user"
	DB_PASSWORD = "postgres_password"
	DB_NAME     = "postgres_db"

	MIGRATION_DIR = "migrations"

	//////////////////////////////////////////////
	// HTTP
	//////////////////////////////////////////////
	HTTP_HOST = "localhost"
	HTTP_PORT = "3000"
)
