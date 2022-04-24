package api

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Asset struct {
	Id          uint64 `json:"id"`
	Assettype   string `json:"assettype"`
	Ip          string `json:"ip"`
	Environment string `json:"environment"`
	Pci         bool   `json:"pci"`
	Sox         bool   `json:"sox"`
	Gdpr        bool   `json:"gdpr"`
	Datacenter  string `json:"datacenter"`
	Owner       int    `json:"owner"`
}

type Owner struct {
	Teamid   uint64 `json:"teamid"`
	Teamname string `json:"teamname"`
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
