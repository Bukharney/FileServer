package main

import (
	"errors"
	"log"
	"os"

	"github.com/bukharney/FileServer/configs"
	"github.com/bukharney/FileServer/databases"
	"github.com/bukharney/FileServer/server"
)

func main() {
	cfg := new(configs.Configs)

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(errors.New("failed to get hostname"))
	}

	log.Println(host)
	if host != "railway" {
		cfg.URL = "http://localhost:8081/"
		cfg.PostgreSQL.Host = "localhost"
		cfg.PostgreSQL.Port = "5432"
		cfg.PostgreSQL.Username = "postgres"
		cfg.PostgreSQL.Password = "postgres"
		cfg.PostgreSQL.Database = "ModX"
	} else {
		cfg.URL = os.Getenv("URL")
		cfg.PostgreSQL.Host = os.Getenv("PGHOST")
		cfg.PostgreSQL.Port = os.Getenv("PGPORT")
		cfg.PostgreSQL.Username = os.Getenv("POSTGRES_USER")
		cfg.PostgreSQL.Password = os.Getenv("POSTGRES_PASSWORD")
		cfg.PostgreSQL.Database = os.Getenv("POSTGRES_DB")
	}

	db, err := databases.NewPostgreSQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(db)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
