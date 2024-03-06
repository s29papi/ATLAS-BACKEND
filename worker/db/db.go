package db

import (
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	*sql.DB
}

func NewDB(sdb *sql.DB) *DB {
	return &DB{sdb}
}

func (db *DB) GetLastProcReqTime() int64 {
	query := `SELECT last_proc_req_time FROM wager_bot_state WHERE id = 1000`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error: GetLastProcReqTime Failed %v\n", err)
		return 0
	}

	var lastProcReqTimes []int64
	for rows.Next() {
		var lastProcReqTime int64
		if err := rows.Scan(&lastProcReqTime); err != nil {
			log.Printf("Error: GetLastProcReqTime Failed %v\n", err)
			return 0
		}
		lastProcReqTimes = append(lastProcReqTimes, lastProcReqTime)
	}
	return lastProcReqTimes[0]
}

func (db *DB) UpdateLastProcReqTime(new_proc_req_time, last_proc_req_time int64) {
	queryFmt := `UPDATE wager_bot_state
	SET last_proc_req_time = %v
	WHERE last_proc_req_time = %v`
	query := fmt.Sprintf(queryFmt, new_proc_req_time, last_proc_req_time)
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
func (db *DB) UpdateLastProcReqGameId(last_Proc_Req_GameId int) {
	queryFmt := `UPDATE wager_bot_state
	SET last_Proc_Req_GameId = %v
	WHERE id = 1000`
	query := fmt.Sprintf(queryFmt, last_Proc_Req_GameId)
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func (db *DB) GetLastProcReqGameId() int {
	query := `SELECT last_Proc_Req_GameId FROM wager_bot_state WHERE id = 1000`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error: GetLastProcReqGameId Failed %v\n", err)
		return 0
	}

	var lastProcReqGameIds []int
	for rows.Next() {
		var lastProcReqGameId int
		if err := rows.Scan(&lastProcReqGameId); err != nil {
			log.Printf("Error: GetLastProcReqGameId Failed %v\n", err)
			return 0
		}
		lastProcReqGameIds = append(lastProcReqGameIds, lastProcReqGameId)
	}
	return lastProcReqGameIds[0]
}

func (db *DB) GetLastUserMentionUpdateTime() int64 {
	query := `SELECT lastusermentionsupdatetime FROM versus_state WHERE idx = 0`
	var rows *sql.Rows
	var err error

	for {
		rows, err = db.Query(query)
		if rows != nil && err == nil {
			break
		}

		log.Printf("Error: GetLastUserMentionUpdateTime Failed %v\n", err)
	}

	var lastusermentionsupdatetimes []int64
	for rows.Next() {
		var lastusermentionsupdatetime int64
		if err := rows.Scan(&lastusermentionsupdatetime); err != nil {
			log.Printf("Error: GetLastUserMentionUpdateTime Failed %v\n", err)
		}
		lastusermentionsupdatetimes = append(lastusermentionsupdatetimes, lastusermentionsupdatetime)
	}
	return lastusermentionsupdatetimes[0]
}
