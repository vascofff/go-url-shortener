package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "localhost"
	PORT = 5432
)

type Url struct {
	UUId     string `json:"uuid"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
	//ExpiresOn uint   `json:"expires_on`
}

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

func SaveUrlMapping(uuid string, shortUrl string, originalUrl string, expiresOn string) {
	// query, err := dbConn.Conn.Prepare("INSERT INTO urls (uuid, long_url, short_url) VALUES (?, ?, ?)")

	// if err != nil {
	// 	panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - UUId: %s\n", err, query))
	// }

	// _, err := query.Exec(uuid, shortUrl, "https://stackoverflow.com/questions/31622052/how-to-serve-up-a-json-response-using-go")

	// log.Fatalf("uuid: %v, shortUrl: %v, originalUrl: %v, expiresOn: %v", uuid, shortUrl, originalUrl, expiresOn)

	_, err := dbConn.Conn.Exec(
		"INSERT INTO urls (uuid, url, short_url, expires_on) VALUES ($1, $2, $3, $4)",
		uuid, originalUrl, shortUrl, expiresOn)

	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

func RetrieveInitialUrl(uuid string) (string, string) {
	var (
		url        string
		expires_on string
	)

	row := dbConn.Conn.QueryRow("SELECT url, expires_on FROM urls WHERE uuid = $1", uuid)

	switch err := row.Scan(&url, &expires_on); err {
	case sql.ErrNoRows:
		log.Fatal("No rows were returned for uuid: %v", uuid)
	case nil:
		fmt.Println(url, expires_on)
	default:
		panic(err)
	}
	// err := row.Scan(&url, &expires_on)

	// if err != nil {
	// 	panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - UUId: %s\n", err, uuid))
	// }
	// log.Fatal("%v %v", url, expires_on)

	return url, expires_on
}
