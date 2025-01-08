package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	_ "github.com/lib/pq"
)

type databaseConfig struct {
	Host string `env:"DB_HOST"`
	Name string `env:"DB_NAME"`
	Port int16  `env:"DB_PORT"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
}

func Init(env string) (*sql.DB, error) {

	dbConf := databaseConfig{}
	if value := os.Getenv("DB_HOST"); value != "" {
		dbConf.Host = value
	} else {
		dbConf.Host = "localhost"
	}
	if value := os.Getenv("DB_NAME"); value != "" {
		dbConf.Name = value
	}
	if value := os.Getenv("DB_PORT"); value != "" {
		port, _ := strconv.Atoi(value)
		dbConf.Port = int16(port)
	} else {
		dbConf.Port = 5432
	}

	if value := os.Getenv("DB_USER"); value != "" {
		dbConf.User = value
	}

	if value := os.Getenv("DB_PASS"); value != "" {
		dbConf.Pass = value
	}
	

	dbdns := fmt.Sprintf("postgres://%s:%s@%s/%s",
		dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Name,
	)
	db, err := sql.Open("postgres", dbdns)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database ", dbdns)
	return db, nil
}
