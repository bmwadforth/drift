package src

import (
	"fmt"
	"log"
)

func Up() (bool, error) {
	path := migrationDir
	patchDir := fmt.Sprintf("%s/%s", path, "patch")

	files := readFilesInDir(patchDir)
	fileMap := make(map[string][]byte)
	for _, file := range files {
		if file.Size() <= 0 {
			log.Println(fmt.Sprintf("skipping migration: %s - no data inside migration file", file.Name()))
			continue
		}
		fileMap[file.Name()] = readFileInDir(fmt.Sprintf("%s/%s", patchDir, file.Name()))
	}

	if len(fileMap) > 0 {
		switch Config.Provider {
		case POSTGRES:
			fallthrough
		case SQLSERVER:
			fallthrough
		case MYSQL:
			tableExists := migrationTableExists()
			appliedMigrations := getMigrations()
			if tableExists {
				runMigrations(appliedMigrations, fileMap)
			} else {
				createMigrationTable()
				runMigrations(appliedMigrations, fileMap)
			}
		default:
			log.Fatal("provider not supported")
		}
	}

	return true, nil
}
