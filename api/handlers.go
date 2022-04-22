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
		err := rows.Scan(&asset.id, &asset.assettype, &asset.ip, &asset.environment, &asset.pci, &asset.sox, &asset.gdpr, &asset.datacenter, &asset.owner)
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
