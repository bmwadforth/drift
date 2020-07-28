package src

import (
	_ "github.com/lib/pq"
)

func RunPG(migrationsToApply map[string][]byte) {
	tableExists := migrationTableExists()
	appliedMigrations := getMigrations()
	if tableExists {
		runMigrations(appliedMigrations, migrationsToApply)
	} else {
		createMigrationTable()
		runMigrations(appliedMigrations, migrationsToApply)
	}
}
