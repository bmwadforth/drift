package src

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var workingDir string
var migrationDir string
var sqlDir string
var Config DriftConfig

func SetWorkingPath() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	workingDir = path
}

func SetMigrationPath() {
	//TODO: make sure working dir has a value before doing the below
	migrationDir = fmt.Sprintf("%s/%s", workingDir, "migration")
}

func SetConfig() {
	configBytes := readFileInDir(fmt.Sprintf("%s/%s", migrationDir, "Config.json"))
	err := json.Unmarshal(configBytes, &Config)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to set configuration: %v", err))
	}
}

func SetSQLPath() {
	switch Config.Provider {
	case POSTGRES:
		sqlDir = fmt.Sprintf("%s/%s", workingDir, "sql/postgres")
	}
}

func dirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func readFilesInDir(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func readFileInDir(path string) []byte {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return fileBytes
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

func loadExistsSQL() string {
	fileBytes := readFileInDir(fmt.Sprintf("%s/%s", sqlDir, "exists.sql"))
	return string(fileBytes)
}

func loadCreateSQL() string {
	fileBytes := readFileInDir(fmt.Sprintf("%s/%s", sqlDir, "create.sql"))
	return string(fileBytes)
}

func getDriver() string {
	switch Config.Provider {
	case POSTGRES:
		return "postgres"
	}

	return ""
}

func getChecksumFromBytes(bytes []byte) []byte {
	h := sha1.New()
	_, err := io.WriteString(h, string(bytes)); if err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}