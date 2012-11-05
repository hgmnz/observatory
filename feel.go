package main

import (
	"database/sql"
	"fmt"
	"github.com/bmizerany/pq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	urls := []string{
		"url1",
		"url2",
		"url3",
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
	} else {
		db, err := sql.Open("postgres", dataSource)
		if err != nil {
			fmt.Println("Unable to connect to database")
		}
		for {
			row := db.QueryRow("SELECT count(*) from pg_stat_activity")
			var count int
			row.Scan(&count)
			o := observation{connections: count, databaseUrl: databaseUrl, updatedAt: time.Now()}
			o.persist()
		}
	}
}
