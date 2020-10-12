/*
 *  Copyright 2020 ChronoWave Authors
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  Package parser declares an expression parser with support for macro
 *  expansion.
 */

package embed

import (
	"context"
	"database/sql"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	seq int64
)

func openSqlite(dir string) error {
	var err error
	db, err = sql.Open("sqlite3", filepath.Join(dir, "db"))
	if err != nil {
		return err
	}

	create := []string{
		`CREATE TABLE IF NOT EXISTS wave
         (
           wid INTEGER PRIMARY KEY AUTOINCREMENT,
           beg INTEGER NOT NULL,
           end INTEGER NOT NULL,
           created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
         )`,
		`CREATE TABLE IF NOT EXISTS waveloc
         (
           path TEXT,
           key TEXT,
           wid INTEGER,
           created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
           PRIMARY KEY (path, key, wid)
         ) WITHOUT ROWID`,
	}

	for _, qry := range create {
		_, err = db.Exec(qry)
		if err != nil {
			return err
		}
	}

	indices := []string{
		`CREATE INDEX IF NOT EXISTS wave_time ON wave (beg, end)`,
	}

	for _, qry := range indices {
		_, err = db.Exec(qry)
		if err != nil {
			return err
		}
	}

	return resetSeq()
}

func closeSqlite() error {
	return db.Close()
}

func insertWaveLoc(path string, keys []string, wid int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	qry := `INSERT INTO waveloc (path, key, wid) VALUES (?, ?, ?) ON CONFLICT DO NOTHING`

	for _, key := range keys {
		if _, err = tx.Exec(qry, path, key, wid); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func insertWave(wid, beg, end int64) error {
	qry := `INSERT INTO wave (wid, beg, end) VALUES (?, ?, ?)`
	_, err := db.Exec(qry, wid, beg, end)
	return err
}

func resetSeq() error {
	rs, err := db.Query(`SELECT MAX(wid) FROM wave`)
	if err != nil {
		return err
	}
	defer rs.Close()

	if rs.Next() {
		var id sql.NullInt64
		if err = rs.Scan(&id); err != nil {
			return err
		}

		if id.Valid {
			seq = id.Int64
		}
	}

	return nil
}

func selectWave(beg, end int64) ([]int64, error) {
	qry := `SELECT wid FROM wave WHERE beg BETWEEN ? AND ? OR end BETWEEN ? AND ?`
	rows, err := db.Query(qry, beg, end, beg, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		tmp sql.NullInt64
		wid []int64
	)
	for rows.Next() {
		if err = rows.Scan(&tmp); err == nil && tmp.Valid {
			wid = append(wid, tmp.Int64)
		}
	}

	return wid, nil
}

func selectKey(path, key string) ([]int64, error) {
	qry := `SELECT DISTINCT wid FROM waveloc WHERE path = ? AND key = ?`
	rows, err := db.Query(qry, path, key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		tmp sql.NullInt64
		wid []int64
	)
	for rows.Next() {
		if err = rows.Scan(&tmp); err == nil && tmp.Valid {
			wid = append(wid, tmp.Int64)
		}
	}

	return wid, nil
}

func selectWidBeforeTime(time time.Time) ([]int64, error) {
	qry := `SELECT DISTINCT wid FROM wave WHERE created < ? UNION SELECT DISTINCT wid FROM waveloc WHERE created < ?`
	rows, err := db.Query(qry, time, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		tmp sql.NullInt64
		wid []int64
	)
	for rows.Next() {
		if err = rows.Scan(&tmp); err == nil && tmp.Valid {
			wid = append(wid, tmp.Int64)
		}
	}

	return wid, nil
}

func purgeBeforeTime(ctx context.Context, time time.Time) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM wave WHERE created < ?`, time)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM waveloc WHERE created < ?`, time)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
