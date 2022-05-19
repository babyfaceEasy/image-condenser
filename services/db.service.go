package services

import (
	"database/sql"
	"fmt"
	"thumbnail-generator/models"
)

/*
Handles the interaction with the database
*/

type DB struct {
	dbConn *sql.DB
}

func NewDBService(dbConn *sql.DB) DB {
	return DB{dbConn: dbConn}
}

// tutorials : https://learningprogramming.net/golang/golang-and-mysql/update-entity-in-golang-and-mysql-database/

func (db DB) UpdateStockThumbnailPath(path string, prdtID int) error {
	_, err := db.dbConn.Exec("UPDATE stocks SET thumbnail_path = ? WHERE id = ?", path, prdtID)
	if err != nil {
		return err
	}

	return nil
}

func (db DB) GetAllProduct(columns string) ([]models.Stock, error) {

	stocks := make([]models.Stock, 0)
	// https://stackoverflow.com/questions/12939690/mysql-query-for-empty-and-null-value-together
	results, err := db.dbConn.Query("SELECT id, picture_path, thumbnail_path FROM stocks WHERE deleted_at IS NULL ORDER BY id LIMIT 1")
	if err != nil {
		return stocks, err
	}

	for results.Next() {
		var stock models.Stock

		err = results.Scan(&stock.ID, &stock.PicturePath, &stock.ThumbnailPath)
		if err != nil {
			fmt.Println(err.Error())
		}

		//log.Printf("%v\n", prdt)
		stocks = append(stocks, stock)
	}

	return stocks, nil
}
