package db

import (
	"database/sql"
	"fmt"
)

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_url TEXT NOT NULL,
		origin_url TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func StoreUrl(db *sql.DB, shortUrl string, originUrl string) error {
	query := `INSERT INTO urls (short_url, origin_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortUrl, originUrl)
	return err
}
func GetOriginalUrl(db *sql.DB, shortUrl string) (string, error) {
	query := `SELECT origin_url FROM urls WHERE short_url = ?`
	row := db.QueryRow(query, shortUrl)
	var originUrl string
	err := row.Scan(&originUrl)
	if err == sql.ErrNoRows {
		return "", nil
	}
	fmt.Println("Original", originUrl)
	return originUrl, err
}
func GetAllUrls(db *sql.DB) error {
	query := `SELECT * FROM urls`
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var shortUrl string
		var originUrl string
		err := rows.Scan(&shortUrl, &originUrl)
		if err != nil {
			return err
		}
		fmt.Println(originUrl, shortUrl)
	}
	return nil
}
