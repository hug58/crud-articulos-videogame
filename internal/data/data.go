package data

import (
	"database/sql"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *sql.DB
}

func initDB() {
	log.Println("Initializing database connection pool.")
	db, err := getConnection()
	if err != nil {
		log.Println("Error connecting to database: ", err)
		log.Panic(err)
	}

	err = MakeMigration(db)
	if err != nil {
		log.Println("Error creating database: ", err)
		//log.Panic(err)
	}

	data = &Data{
		DB: db,
	}
}

func New() *Data {
	once.Do(initDB)
	return data
}

// Close closes the resources used by data.
func Close() error {
	if data == nil {
		return nil
	}

	return data.DB.Close()
}
