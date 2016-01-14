// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/goto2/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type GooseDbVersion struct {
	Id        int64     `db:"id" json:"id"`
	VersionId int64     `db:"version_id" json:"versionId"`
	IsApplied int64     `db:"is_applied" json:"isApplied"`
	Tstamp    time.Time `db:"tstamp" json:"tstamp"`
}

// @Title getGooseDbVersionById
// @Description retrieves the goose_db_version information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    GooseDbVersion
// @Resource /api/2.0
// @Router /api/2.0/goose_db_version/{id} [get]
func getGooseDbVersionById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []GooseDbVersion{}
	arg := GooseDbVersion{Id: int64(id)}
	nstmt, err := db.PrepareNamed(`select * from goose_db_version where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getGooseDbVersions
// @Description retrieves the goose_db_version information for a certain id
// @Accept  application/json
// @Success 200 {array}    GooseDbVersion
// @Resource /api/2.0
// @Router /api/2.0/goose_db_version [get]
func getGooseDbVersions(db *sqlx.DB) (interface{}, error) {
	ret := []GooseDbVersion{}
	queryStr := "select * from goose_db_version"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postGooseDbVersion
// @Description enter a new goose_db_version
// @Accept  application/json
// @Param                 Body body     GooseDbVersion   true "GooseDbVersion object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/goose_db_version [post]
func postGooseDbVersion(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v GooseDbVersion
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO goose_db_version("
	sqlString += "version_id"
	sqlString += ",is_applied"
	sqlString += ",tstamp"
	sqlString += ") VALUES ("
	sqlString += ":version_id"
	sqlString += ",:is_applied"
	sqlString += ",:tstamp"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putGooseDbVersion
// @Description modify an existing goose_db_versionentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     GooseDbVersion   true "GooseDbVersion object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/goose_db_version/{id}  [put]
func putGooseDbVersion(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v GooseDbVersion
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "UPDATE goose_db_version SET "
	sqlString += "version_id = :version_id"
	sqlString += ",is_applied = :is_applied"
	sqlString += ",tstamp = :tstamp"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delGooseDbVersionById
// @Description deletes goose_db_version information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    GooseDbVersion
// @Resource /api/2.0
// @Router /api/2.0/goose_db_version/{id} [delete]
func delGooseDbVersion(id int, db *sqlx.DB) (interface{}, error) {
	arg := GooseDbVersion{Id: int64(id)}
	result, err := db.NamedExec("DELETE FROM goose_db_version WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
