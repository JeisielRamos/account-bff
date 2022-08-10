package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBRepository struct {
	DB *sql.DB
}

var instance *DBRepository

func newDBRepository() *DBRepository {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PWD")+"@tcp("+os.Getenv("DB_ADDR")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return &DBRepository{
		DB: db,
	}
}

func InitDBRepository() *DBRepository {
	if instance == nil {
		instance = newDBRepository()
		fmt.Println("New config")
	}

	fmt.Println("conected")
	return instance
}
