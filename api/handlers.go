package api

import (
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
