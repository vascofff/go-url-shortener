package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	HOST = "localhost"
	PORT = 5432
)

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

func SaveUrlMapping(uuid string, shortUrl string, originalUrl string, expiresAt *string) error {
	query, err := dbConn.Conn.Prepare("INSERT INTO urls (uuid, url, short_url, expires_at) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return errors.Wrap(err, "failed while preparing query to insert")
	}

	_, err = query.Exec(uuid, originalUrl, shortUrl, newNullString(*expiresAt))
	if err != nil {
		return errors.Wrap(err, "failed while executing query while insert")
	}

	return nil
}

func RetrieveInitialUrl(uuid string) (string, string, error) {
	var (
		url        string
		expires_at *string
	)

	row := dbConn.Conn.QueryRow("SELECT url, expires_at FROM urls WHERE uuid = $1", uuid)
	err := row.Scan(&url, &expires_at)
	if errors.Is(err, sql.ErrNoRows) {
		return "", "", errors.Wrap(err, fmt.Sprintf("no rows were returned for uuid: %v", uuid))
	}
	if err != nil {
		return "", "", errors.Wrap(err, "failed to execute query when getting initial url")
	}

	if expires_at == nil {
		return url, "", nil
	} else {
		return url, *expires_at, nil
	}
}

func newNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
