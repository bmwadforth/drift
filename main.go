package main

import (
	"errors"
	"fmt"
	"github.com/bmwadforth/drift/src"
	"log"
	"os"
)

func main() {
	src.SetWorkingPath()
	src.SetMigrationPath()
	src.SetConfig()
	src.SetSQLPath()

	args := os.Args[1:]

	if len(args) > 0 {
		fmt.Println(args)
		switch args[0] {
		case "init": {
			_, err := src.Initialise(); if err != nil {
				log.Fatal(err)
			}
		}
		case "add": {
			//TODO: Ensure args[1] exists
			_, err := src.Add(args[1]); if err != nil {
				log.Fatal(err)
			}
		}
		case "remove": {
			//TODO: Ensure args[1] exists
			_, err := src.Remove(args[1]); if err != nil {
				log.Fatal(err)
			}
		}
		case "up": {
			_, err := src.Up(); if err != nil {
				log.Fatal(err)
			}
		}
		case "down": {
			_, err := src.Down(); if err != nil {
				log.Fatal(err)
			}
		}
		default:
			log.Fatal(errors.New("invalid argument supplied"))
		}
	} else {
		//TODO: Binary called with no arguments, process.exit
		//TODO: Print CLI HELP
		log.Fatal(errors.New("an argument must be supplied"))
	}
}
