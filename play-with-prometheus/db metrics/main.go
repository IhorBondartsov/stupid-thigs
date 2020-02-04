package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/common/log"

	"github.com/dlmiddlecote/sqlstats"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//POSTGRES_USER: 'user'
//POSTGRES_PASSWORD: 'password'
//POSTGRESS_DB: 'db_amex01'

func main() {
	time.Sleep(time.Second * 10)
	dbinfo := fmt.Sprintf("user=postgres password=postgres host=postgres port=5432 dbname=postgres sslmode=disable")
	fmt.Println(dbinfo)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Error(err)

	}

    // Create a lot of connections to db.
	go func() {
		for {
			conn, err := db.Conn(context.Background())
			if err != nil {
				log.Error(err)
				break
			}
			err = conn.PingContext(context.Background())
			if err != nil {
				log.Error(err)
			}
			time.Sleep(time.Second * 1)
		}
	}()

	if err := run(db); err != nil {
		log.Error(err)
		panic(err)
	}
}

func run(db *sql.DB) error {
	// Create a new collector, the name will be used as a label on the metrics
	collector := sqlstats.NewStatsCollector("db_name", db)

	// Register it with Prometheus
	prometheus.MustRegister(collector)

	// Register the metrics handler
	http.Handle("/metrics", promhttp.Handler())

	// Run the web server
	return http.ListenAndServe(":3329", nil)
}
