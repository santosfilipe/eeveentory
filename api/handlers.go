package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAssets(c *gin.Context) {
	var assets []Asset

	cfg := DatabaseConfiguration()
	db := ConnectToDb(cfg)

	rows, err := db.Query("SELECT * FROM asset")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var asset Asset
		err := rows.Scan(&asset.Id, &asset.Assettype, &asset.Ip, &asset.Environment, &asset.Pci, &asset.Sox, &asset.Gdpr, &asset.Datacenter, &asset.Owner)
		if err != nil {
			log.Fatal(err)
		}
		assets = append(assets, asset)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, assets)
}

func GetAssetsByIp(c *gin.Context) {
	var asset Asset
	ip := c.Param("ip")

	cfg := DatabaseConfiguration()
	db := ConnectToDb(cfg)

	rows, err := db.Query("SELECT * FROM asset WHERE ip = ?", ip)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	err = rows.Scan(&asset.Id, &asset.Assettype, &asset.Ip, &asset.Environment, &asset.Pci, &asset.Sox, &asset.Gdpr, &asset.Datacenter, &asset.Owner)
	if err == sql.ErrNoRows {
		log.Println(err)
	} else if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}

	c.IndentedJSON(http.StatusOK, asset)
}
