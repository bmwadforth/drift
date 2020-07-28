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

func getMigrations() map[string]DriftMigration {
	db := Connect()

	rows, err := db.Query("SELECT name, checksum, applied FROM drift_migrations"); if err != nil {
		log.Fatal(err)
	}

	migrations := make(map[string]DriftMigration, 0)
	for rows.Next() {
		migration := DriftMigration{}
		_ = rows.Scan(&migration.Name, &migration.Checksum, &migration.Applied)
		migrations[migration.Name] = migration
	}

	return migrations
}

func runMigrations(appliedMigrations map[string]DriftMigration, migrationsToApply map[string][]byte) {
	db := Connect()
	for fileName, fileBytes := range migrationsToApply {
		checkSum := fmt.Sprintf("% x", getChecksumFromBytes(fileBytes))
		appliedMigration, foundInDB := appliedMigrations[fileName]

		if foundInDB && string(appliedMigration.Checksum) != checkSum {
			log.Println(fmt.Sprintf("ERROR: Checksum difference between database and migration folder for migration: %s", fileName))
			continue
		}

		if foundInDB {
			log.Println(fmt.Sprintf("skipping migration: %s - already applied", fileName))
			continue
		}

		_, err := db.Exec("INSERT INTO drift_migrations (name, checksum) VALUES ($1, $2);", fileName, checkSum)
		if err != nil {
			log.Fatal(err)
		}
	}
}