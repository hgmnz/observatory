package main

import (
	"database/sql"
	"fmt"
	"github.com/bmizerany/pq"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	urls := [...]string{
		"postgres://localhost/hgmnz?sslmode=disable",
	}

	for _, url := range urls {
		go feel(url)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	<-sigs
	fmt.Println("Exiting program")
	os.Exit(0)
}

func feel(databaseUrl string) {
	dataSource, err := pq.ParseURL(databaseUrl)
	if err != nil {
		fmt.Printf("Unable to parse database url (%s)", databaseUrl)
		panic(err)
	}
	u, err := url.Parse(databaseUrl)
	if err != nil {
		panic(err)
	}
	databaseName := u.Path[1:]
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		fmt.Println("Unable to connect to database")
	}

	for {
		row := db.QueryRow("SELECT count(*) from pg_stat_activity")
		var count uint16
		err = row.Scan(&count)
		if err != nil {
			panic(err)
		}
		row = db.QueryRow("SELECT pg_database_size($1)", databaseName)
		var bytes uint64
		err = row.Scan(&bytes)
		if err != nil {
			panic(err)
		}
		o := observation{connections: count, databaseUrl: databaseUrl, bytes: bytes, serviceAvailable: true, updatedAt: time.Now()}
		o.persist()
	}
}
