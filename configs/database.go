package configs

import (
	"fmt"
)

type Database struct {
    Connection string
    Host       string
    Port       int
    Database   string
    Username   string
    Password   string
}

func (d Database) MySQLDSN() string {
    return fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?parseTime=true",
        d.Username,
        d.Password,
        d.Host,
        d.Port,
        d.Database,
    )
}

func LoadDatabase() Database {
    return Database{
        Connection: getEnv("DB_CONNECTION", "mysql"),
        Host:       getEnv("DB_HOST", "127.0.0.1"),
        Port:       getEnvInt("DB_PORT", 3306),
        Database:   getEnv("DB_DATABASE", "erp_go"),
        Username:   getEnv("DB_USERNAME", "root"),
        Password:   getEnv("DB_PASSWORD", ""),
    }
}

