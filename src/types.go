package src

import (
	"crypto"
	"time"
)

type DriftProvider string

const (
	POSTGRES  DriftProvider = "POSTGRES"
	SQLSERVER DriftProvider = "SQL_SERVER"
	MYSQL     DriftProvider = "MY_SQL"
)

type DriftConfig struct {
	Provider     DriftProvider `json:"provider"`
	DatabaseName string        `json:"databaseName"`
	Host         string        `json:"host"`
	Username     string        `json:"username"`
	Password     string        `json:"password"`
	Port         int           `json:"port"`
}

type DriftMigration struct {
	Name     string
	Checksum crypto.Hash
	Applied  time.Time
}
