package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST = "localhost"
	PORT = 5432
)

// type Url struct {
// 	UUId      string `json:"uuid"`
// 	ShortUrl  string `json:"short_url"`
// 	LongUrl   string `json:"long_url"`
// 	ExpiresAt string `json:"expires_at`
// }

type Database struct {
	Conn *sql.DB
}

var (
	dbConn = &Database{}
)

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()

	if err != nil {
		return db, err
	}

	dbConn.Conn = conn

	return db, nil
}

func SaveUrlMapping(uuid string, shortUrl string, originalUrl string, expiresAt string) error {
	query, err := dbConn.Conn.Prepare("INSERT INTO urls (uuid, url, short_url, expires_at) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return errors.New("Failed while preparing query to insert")
	}

	_, err = query.Exec(uuid, originalUrl, shortUrl, NewNullString(expiresAt))
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func RetrieveInitialUrl(uuid string) (string, string, error) {
	var (
		url        string
		expires_at *string
	)

	row := dbConn.Conn.QueryRow("SELECT url, expires_at FROM urls WHERE uuid = $1", uuid)
	switch err := row.Scan(&url, &expires_at); err {
	case sql.ErrNoRows:
		return "", "", errors.New(fmt.Sprintf("No rows were returned for uuid: %v", uuid))
	case nil:
	default:
		return "", "", errors.New(err.Error())
	}

	if expires_at == nil {
		return url, "", nil
	} else {
		return url, *expires_at, nil
	}
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
