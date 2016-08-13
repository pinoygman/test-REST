/*
 * Copyright (c) 2016 General Electric Company. All rights reserved.
 *
 * The copyright to the computer software herein is the property of
 * General Electric Company. The software may be used and/or copied only
 * with the written permission of General Electric Company or in accordance
 * with the terms and conditions stipulated in the agreement/contract
 * under which the software has been supplied.
 *
 * author: chia.chang@ge.com
 */

/*
CREATE TABLE "pcs-question-tbl"
(
  _id text,
  title text,
  name text,
  description text,
  type integer,
  "answerOptions" json
)
*/

package sql

import(
	"encoding/json"
	_ "github.com/lib/pq"
	//"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	db *sqlx.DB
)

func init(){
	op, err := sqlx.Connect("postgres","host=10.131.54.5 port=5432 user=uc49c9583047d4173a217667509e17ddf password=fb46202694704a7d994dd8e906666e6c dbname=d13291d5f50c645f5b90d26b8a58e2f6b connect_timeout=5 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	
	db=op
}

func AddAQuestion(_id string, _title string, _name string, _desc string, _type uint64,_ans map[string]interface{}) error {

	tx := db.MustBegin()
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person

	__ans,err:=json.Marshal(_ans)

	if err!=nil{
		return err
	}
	
	tx.MustExec("INSERT INTO pcs-question-tbl (_id, title, name, description, type, answerOptions) VALUES ($1, $2, $3, $4, $5, $6)", _id,_title,_name,_desc,_type,__ans)
	
	tx.Commit()

	return nil
}
