package config

type DBConfig struct {
	Username         string
	Password         string
	Host             string
	Database         string
	ConnectionString *string
}

var dbConfig *DBConfig

func GetDBConfig() *DBConfig {

	return dbConfig
}

func SetDBConfig(connectionString *string) {
	dbConfig = &DBConfig{
		Username:         "",
		Password:         "",
		Host:             "",
		Database:         "",
		ConnectionString: connectionString,
	}
}
