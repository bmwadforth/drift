package src

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func RunPG(fileMap map[string][]byte) {
	tableExists := migrationTableExists()
	if tableExists {
		db := Connect()

		for fileName, fileBytes := range fileMap {
			log.Println(fmt.Sprintf("file name: %s", fileName))
			log.Println(fmt.Sprintf("file bytes: %s", string(fileBytes)))
			fmt.Println()

			checkSum := fmt.Sprintf("% x", getChecksumFromBytes(fileBytes))

			//TODO: Calculate checksum for each migration, compare 'name' of migration to checksum by comparing filemap with db table
			_, err := db.Exec("INSERT INTO drift_migrations (name, checksum) VALUES ($1, $2);", fileName, checkSum); if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		createMigrationTable()
	}
}
