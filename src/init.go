package src

import (
	"fmt"
	"log"
)

func Initialise() (bool, error) {
	log.Println("Initialising")
	path := workingDir

	//TODO: Handle errors a little better
	log.Println(fmt.Sprintf("%s: %s", "Current working directory", path))
	/*dirExist := dirExists(path); if dirExist == true {
		//Check files/folders under dir are correct and not corrupted
		log.Println(fmt.Sprintf("%s: %s", "Migration folder exists under", path))
	} else {

	}*/

	// Create migration dir
	log.Println(fmt.Sprintf("%s: %s", "Creating migration folder under", path))
	migrationPath := fmt.Sprintf("%s/%s", path, "migration")
	_ = createDir(migrationPath)

	// Write Config.json file
	log.Println(fmt.Sprintf("%s: %s", "Creating Config.json file under", migrationPath))
	_ = writeConfigTemplate(fmt.Sprintf("%s/Config.json", migrationPath))

	// Create patches dir under migration dir
	log.Println(fmt.Sprintf("%s: %s", "Creating patches folder under", migrationPath))
	patchDir := fmt.Sprintf("%s/%s", migrationPath, "patch")
	_ = createDir(patchDir)

	// Create seed dir under migration dir
	log.Println(fmt.Sprintf("%s: %s", "Creating seed folder under", migrationPath))
	seedDir := fmt.Sprintf("%s/%s", migrationPath, "seed")
	_ = createDir(seedDir)

	return true, nil
}
