package db

import (
	"database/sql"
	"fmt"
	"github.com/golobby/dotenv"
	_ "github.com/lib/pq"
	"os"
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
	file, err := os.Open(env)
	if err != nil {
		panic(err)
	}

	err = dotenv.NewDecoder(file).Decode(&dbConf)
	if err != nil {
		return nil, err
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
