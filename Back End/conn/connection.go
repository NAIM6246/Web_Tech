package conn

import (
	"fmt"
	"log"
	"sync"
	"webTech/configs"

	"github.com/jinzhu/gorm"
	//mssql
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type DB struct {
	*gorm.DB
}

const dbtype = "mssql"

func NewDB() *DB {
	return &DB{}
}

var dbInstance *DB
var connDBOnce sync.Once

func ConnectDB(config *configs.DBConfig) *DB {
	connDBOnce.Do(func() {
		_ = connectDB(config)
	})
	return dbInstance
}

func connectDB(config *configs.DBConfig) error {
	connString := fmt.Sprintf("server=%s; port=%d; database=%s;", config.Server, config.Port, config.DbName)
	conn, err := gorm.Open(dbtype, connString)
	if err != nil {
		log.Fatal("Open connection failed : ", err.Error())
		panic(err)
	}
	fmt.Println("Database connected succesfully")
	dbInstance = &DB{conn}
	return nil
}

func (db *DB) Migration() {

}
