package api

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Asset struct {
	id          uint64 `json:"id"`
	assettype   string `json:"assettype"`
	ip          string `json:"ip"`
	environment string `json:"environment"`
	pci         bool   `json:"pci"`
	sox         bool   `json:"sox"`
	gdpr        bool   `json:"gdpr"`
	datacenter  string `json:"datacenter"`
	owner       int    `json:"owner"`
}

type Owner struct {
	teamid   uint64 `json:"teamid"`
	teamname string `json:"teamname"`
}

func DatabaseConfiguration() mysql.Config {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "database-eeveentory.cwr4vgen2iib.us-east-1.rds.amazonaws.com:3306",
		DBName:               "eeveentory",
		AllowNativePasswords: true,
	}

	return cfg
}

func ConnectToDb(cfg mysql.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
