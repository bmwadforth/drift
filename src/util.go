package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var workingDir string
var migrationDir string

func SetWD() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	workingDir = path
}

func SetMigrationPath() {
	migrationDir = fmt.Sprintf("%s/%s", workingDir, "migration")
}

func dirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func createDir(path string) bool {
	mkdirErr := os.Mkdir(path, 0770)
	if mkdirErr != nil {
		log.Fatal(mkdirErr)
	}

	return true
}

func writeConfigTemplate(path string) bool {
	bytes, err := json.Marshal(DriftConfig{})
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(path, bytes, 0770)
	if err != nil {
		log.Fatal(err)
	}

	return true
}

func writeFile(path string) bool {
	err := ioutil.WriteFile(path, nil, 0770)
	if err != nil {
		log.Fatal(err)
	}

	return true
}
