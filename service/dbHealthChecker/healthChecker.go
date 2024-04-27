package dbhealthchecker

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/lib/pq"
	"github.com/s29papi/wag3r-bot/worker/db"
	"github.com/s29papi/wag3r-bot/worker/env"
)

type DBHealthChecker struct {
	*db.DB
	stopHealthChecker chan struct{}
	wg                *sync.WaitGroup
	done              chan struct{}
}

func NewDBHealthChecker() *DBHealthChecker {
	var wg sync.WaitGroup
	c := make(chan struct{})
	d := make(chan struct{})
	psqlInfo, err := pq.ParseURL(env.RENDER_POSTGRES_URL)
	if err != nil {
		log.Fatalln(err)
	}
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	db := db.NewDB(sdb)
	return &DBHealthChecker{db, c, &wg, d}
}

func (dbHC *DBHealthChecker) Start() {
	dur := time.Duration(12) * time.Second
	t := time.NewTicker(dur)
	dbHC.wg.Add(1)
	go healthCheckerLoop(dbHC.DB, t.C, dbHC.stopHealthChecker, dbHC.wg)
	dbHC.wg.Wait()
	t.Stop()
	dbHC.done <- struct{}{}
}

func (dbHC *DBHealthChecker) Stop() {
	dbHC.stopHealthChecker <- struct{}{}
	<-dbHC.done
}

func healthCheckerLoop(db *db.DB, t <-chan time.Time, c <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	queryStatement := `SELECT is_active FROM db_health_checker`
	for {
		select {
		case <-t:
			log.Println("Tick go-routine:, new request health checker initiated")

			rows, err := db.Query(queryStatement)
			if err != nil {
				log.Printf("Error: db_health_checker Failed %v\n", err)
				return
			}

			var is_actives []bool
			for rows.Next() {
				var is_active bool
				if err := rows.Scan(&is_active); err != nil {
					log.Printf("Error: db_health_checker Failed %v\n", err)
					return
				}
				is_actives = append(is_actives, is_active)
			}
			log.Printf("Health Checker Status: Asking is db_health_checker active, Responce: %v\n", is_actives[0])
		case <-c:
			return
		}
	}
}
