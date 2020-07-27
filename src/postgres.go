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

			//TODO: Calculate checksum for each migration, compare 'name' of migration to checksum by comparing filemap with db table
			rows, err := db.Query("SELECT name, checksum, applied FROM public.drift_migrations"); if err != nil {
				log.Fatal(err)
			}

			migrations := make([]DriftMigration, 0)
			for rows.Next() {
				migration := DriftMigration{}
				_ = rows.Scan(&migration.Name, &migration.Checksum, &migration.Applied)
				migrations = append(migrations, migration)
			}

			log.Println(migrations)
		}
	} else {
		createMigrationTable()
	}
}
