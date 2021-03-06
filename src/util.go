package src

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	if workingDir == "" {
		SetWorkingPath()
	}
	migrationDir = fmt.Sprintf("%s/%s", workingDir, "migration")
}

func SetConfig() {
	if workingDir == "" {
		SetWorkingPath()
	}
	configBytes := readFileInDir(fmt.Sprintf("%s/%s", migrationDir, "config.json"))
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
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		//If we end up here, the path might or might not exist, thus we need to inspect the error (permission denied, for e.g.)
		log.Fatal(err)
	}

	return false
}

func readFilesInDir(path string, pattern string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	filesToReturn := make([]os.FileInfo, 0, 0)
	for _, file := range files {
		if strings.Contains(file.Name(), pattern) {
			filesToReturn = append(filesToReturn, file)
		}
	}

	return filesToReturn
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
	return `SELECT EXISTS (
   SELECT FROM information_schema.tables
   WHERE  table_schema = 'public'
   AND    table_name   = 'drift_migrations'
);`
}

func loadCreateSQL() string {
	return `CREATE TABLE DRIFT_MIGRATIONS (
    id serial not null primary key,
    name varchar(128) unique not null,
    checksum bytea unique not null,
    applied timestamptz default now()
);`
}

func getDriver() string {
	switch Config.Provider {
	case POSTGRES:
		return "postgres"
	case MYSQL:
		return "mysql"
	case SQLSERVER:
		return "sqlserver"
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