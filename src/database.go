package src

import (
	"database/sql"
	"fmt"
	"log"
)

var ConnectionString string
var Database *sql.DB

func Connect() *sql.DB {
	//TODO: Load connection string here based on configuration - connection string style *might* change
	ConnectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", Config.Host, Config.Username, Config.Password, Config.DatabaseName, Config.Port)

	if Database == nil {
		driver := getDriver()
		db, err := sql.Open(driver, ConnectionString)
		if err != nil {
			log.Fatal(err)
		}
		Database = db
	}

	return Database
}

func migrationTableExists() bool {
	db := Connect()

	existsQuery := loadExistsSQL()
	qResult, err := db.Query(existsQuery); if err != nil {
		log.Fatal(err)
	}

	var migrationTableExists bool
	qResult.Next()
	err = qResult.Scan(&migrationTableExists); if err != nil {
		log.Fatal(err)
	}
	if migrationTableExists {
		return true
	}

	return false
}

func createMigrationTable() bool {
	db := Connect()

	createQuery := loadCreateSQL()
	_, err := db.Exec(createQuery); if err != nil {
		log.Fatal(err)
	}

	return true
}