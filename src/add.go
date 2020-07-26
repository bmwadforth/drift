package src

import (
	"fmt"
	"log"
	"time"
)

func AddMigration(name string) (bool, error) {
	path := migrationDir

	//TODO: Update file extension based on provider, also check config.json is set - otherwise throw error
	unixTime := time.Now().Unix()
	migrationName := fmt.Sprintf("%d_%s.sql", unixTime, name)

	//TODO: ensure patch folder exists
	log.Println(fmt.Sprintf("%s:%s", "Adding migration", migrationName))
	writeFile(fmt.Sprintf("%s/%s/%s", path, "patch", migrationName))

	return true, nil
}
