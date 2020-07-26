package src

type DriftProvider string

const (
	POSTGRES  DriftProvider = "POSTGRES"
	SQLSERVER DriftProvider = "SQL_SERVER"
	MYSQL     DriftProvider = "MY_SQL"
)

type DriftConfig struct {
	Provider DriftProvider `json:"provider"`
	Host     string        `json:"host"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Port     int           `json:"port"`
}
