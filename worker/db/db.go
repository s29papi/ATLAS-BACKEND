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

func (db *DB) LastUserMentionsRequestTime() int64 {
	queryStatement := `SELECT time FROM usermentions_requesttime ORDER BY requestid DESC LIMIT 1`
	var rows *sql.Rows
	var err error

	for {
		rows, err = db.Query(queryStatement)
		if rows != nil && err == nil {
			break
		}

		log.Printf("Error: LastUserMentionsRequestTime Failed %v\n", err)
	}
	var lastUserMentionsRequestTimes []int64
	for rows.Next() {
		var lastUserMentionsRequestTime int64
		if err := rows.Scan(&lastUserMentionsRequestTime); err != nil {
			log.Printf("Error: LastUserMentionsRequestTime Failed %v\n", err)
		}
		lastUserMentionsRequestTimes = append(lastUserMentionsRequestTimes, lastUserMentionsRequestTime)
	}
	if len(lastUserMentionsRequestTimes) == 0 {
		return int64(0)
	}
	return lastUserMentionsRequestTimes[0]
}

// we create the game id per time, time in this case is the time specified by the user for the game to start
// we have a goroutine the manages this table i.e. finalized, framecast_hash e.t.c
// so we pass the current time into this function and we check on the time, 10 secs before the time
// we pass the current time into this function
func (db *DB) GameIdsByTime(t int64) []int64 {
	queryStatement := fmt.Sprintf("SELECT gameid FROM gameid_time WHERE time = %d AND finalized = FALSE", t)
	var rows *sql.Rows
	var err error

	for {
		rows, err = db.Query(queryStatement)
		if rows != nil && err == nil {
			break
		}
		log.Printf("Error: GameIdsByTime Failed %v\n", err)
	}
	var gameIds []int64
	for rows.Next() {
		var gameId int64
		if err := rows.Scan(&gameId); err != nil {
			log.Printf("Error: GameIdsByTime Failed %v\n", err)
		}
		gameIds = append(gameIds, gameId)
	}
	return gameIds
}

func (db *DB) LatestGameID() int64 {
	queryStatement := " SELECT gameid FROM gameid_time ORDER BY gameid DESC LIMIT 1"
	var rows *sql.Rows
	var err error

	for {
		rows, err = db.Query(queryStatement)
		if rows != nil && err == nil {
			break
		}
		log.Printf("Error: LatestGameID Failed %v\n", err)
	}
	var gameIds []int64
	for rows.Next() {
		var gameId int64
		if err := rows.Scan(&gameId); err != nil {
			log.Printf("Error: LatestGameID Failed %v\n", err)
		}
		gameIds = append(gameIds, gameId)
	}
	return gameIds[0]
}

func (db *DB) UnfinalizedCastHashByTime(t int64) []string {
	queryStatement := fmt.Sprintf("SELECT framecast_hash FROM gameid_time WHERE finalized = false AND time < %d", t)
	var rows *sql.Rows
	var err error

	for {
		rows, err = db.Query(queryStatement)
		if rows != nil && err == nil {
			break
		}
		log.Printf("Error: UnfinalizedCastHashByTime Failed %v\n", err)
	}
	var castHashes []string
	for rows.Next() {
		var castHash string
		if err := rows.Scan(&castHash); err != nil {
			log.Printf("Error: UnfinalizedCastHashByTime Failed %v\n", err)
		}
		castHashes = append(castHashes, castHash)
	}
	return castHashes
}

// func (db *DB) LastGameId() int64 {
// 	queryStatement := `SELECT gameid FROM gameid_time ORDER BY requestid DESC LIMIT 1`
// 	var rows *sql.Rows
// 	var err error
// }

// SELECT time FROM usermentions_requesttime ORDER BY requestid DESC LIMIT 1
