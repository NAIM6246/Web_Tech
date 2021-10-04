package configs

import "sync"

type DBConfig struct {
	Server   string
	Port     int
	User     string
	Password string
	DbName   string
}

func NewDBConfig() *DBConfig {
	var loadDBOnce sync.Once
	loadDBOnce.Do(mapConfig)
	return dbConfig
}

var dbConfig *DBConfig

func mapConfig() {
	dbConfig = &DBConfig{
		Server:   "localhost",
		Port:     1433,
		User:     "sa",
		Password: "golangdb123456.",
		DbName:   "WebTech",
	}
}
