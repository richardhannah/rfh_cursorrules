package config

type DBConfig struct {
	Username string
	Password string
	Host     string
	Database string
}

func GetDBConfig() DBConfig {

	return DBConfig{
		Username: "richard",
		Password: "Onlyone1",
		Host:     "localhost",
		Database: "richard",
	}
}
