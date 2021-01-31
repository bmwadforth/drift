package src

import (
	"fmt"
	"log"
	"time"
)

func Add(name string) (bool, error) {
	path := migrationDir

	//TODO: Update file extension based on provider, also check config.json is set - otherwise throw error
	unixTime := time.Now().Unix()
	migrationNameUp := fmt.Sprintf("%d_%s_up.sql", unixTime, name)
	migrationNameDown := fmt.Sprintf("%d_%s_down.sql", unixTime, name)

	//TODO: ensure patch folder exists
	log.Println(fmt.Sprintf("%s:%s", "Adding migration", migrationNameUp))
	writeFile(fmt.Sprintf("%s/%s/%s", path, "patch", migrationNameUp))

	log.Println(fmt.Sprintf("%s:%s", "Adding migration", migrationNameDown))
	writeFile(fmt.Sprintf("%s/%s/%s", path, "patch", migrationNameDown))

	return true, nil
}
