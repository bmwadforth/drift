package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func dirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func writeConfigTemplate(path string) bool {
	bytes, err := json.Marshal(DriftConfig{}); if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(path, bytes, 0770); if err != nil {
		log.Fatal(err)
	}

	return true
}

func Initialise() (bool, error) {
	log.Println("Initialising")

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("%s: %s", "Current working directory", path))
	/*dirExist := dirExists(path); if dirExist == true {
		//Check files/folders under dir are correct and not corrupted
		log.Println(fmt.Sprintf("%s: %s", "Migration folder exists under", path))
	} else {

	}*/

	log.Println(fmt.Sprintf("%s: %s", "Creating migration folder under", path))
	migrationPath := fmt.Sprintf("%s/%s", path, "migration")

	mkdirErr := os.Mkdir(migrationPath, 0770)
	if mkdirErr != nil {
		log.Fatal(mkdirErr)
	}

	_ = writeConfigTemplate(fmt.Sprintf("%s/config.json", migrationPath))

	return true, nil
}
