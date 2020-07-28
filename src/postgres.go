package src

import (
	_ "github.com/lib/pq"
)

func RunPG(fileMap map[string][]byte) {
	tableExists := migrationTableExists()
	migrations := getMigrations()
	if tableExists {
		runMigrations(migrations, fileMap)
	} else {
		createMigrationTable()
	}
}
