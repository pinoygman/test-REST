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

package model

import (
	"fmt"
	"os"
	"strconv"
	"github.com/pborman/uuid"
	_ "encoding/json"
	_ "github.com/lib/pq"
	//"database/sql"
	"github.com/jmoiron/sqlx"
	sqltypes "github.com/jmoiron/sqlx/types"
	"log"
	"strings"
)

//question type
const (
	
	Security = 1004
	Pricing = 1003
	Architecture = 1002
	Service = 1001
)


type Question struct {
	Guid            string                      `json:"_id"`
	Title           string                      `json:"title"`
	Name            string                      `json:"name"`
	Desc            string                      `json:"description"`
	Type            uint64                      `json:"type"`  //question type
        //AnswerOptions   map[string]interface{}      `json:"answerOptions"`
	AnswerOptions   sqltypes.JSONText           `json:"answerOptions"`
}

var (
	questionnaire map[string]*Question
	questionTypes map[string]string
	db *sqlx.DB
)

func init(){

	_ref:=""
	if _ref=os.Getenv("SQLDSN");_ref=="" {
		_ref="host=localhost|port=7990|user=uc49c9583047d4173a217667509e17ddf|password=fb46202694704a7d994dd8e906666e6c|dbname=d13291d5f50c645f5b90d26b8a58e2f6b|connect_timeout=5|sslmode=disable"
	}
	
	_conn:=strings.Replace(_ref,"|"," ",-1)
	op, err := sqlx.Connect("postgres",_conn)
	
	if err != nil {
		log.Fatalln(err)
	}
	
	db=op

	questionnaire=make(map[string]*Question)
	questionTypes=map[string]string{
		strconv.Itoa(Security):"Security",
		strconv.Itoa(Pricing):"Pricing",
		strconv.Itoa(Architecture):"Architecture",
		strconv.Itoa(Service):"Service",
	}
}

func (q *Question) load(guid string) (*Question, error){
	
	return questionnaire[guid],nil
}

func (q *Question) Save() (*Question, error) {

	if q.Guid=="" {
		q.Guid=uuid.New()
	}
	
	//_ans,err:=json.Marshal(q.AnswerOptions)

	//if err!=nil{
	//	return nil, err
	//}
	
	tx := db.MustBegin()

	tx.MustExec(`INSERT INTO "pcs-question-tbl" (_id, title, name, description, type, "answerOptions") VALUES ($1, $2, $3, $4, $5, $6)`, q.Guid,q.Title,q.Name,q.Desc,q.Type,q.AnswerOptions)
	
	tx.Commit()

	questionnaire[q.Guid]=q

	return q,nil
}

func (q *Question) Del() (string,error){

	delete(questionnaire,q.Guid)
	return "delete",nil

}

func InitQuestion(guid string) (*Question, error) {
	op:=&Question{}

	if guid!="" {
		return op.load(guid)
	}

	guid=uuid.New()
	questionnaire[guid]=op

	return questionnaire[guid], nil
}

func GetQuestionsByType(typeId uint64) (map[string]*Question, error){

	return questionnaire,nil

}

func GetQuestionTypes() (map[string]string, error) {

        return questionTypes,nil
}

func GetQuestions() ([]Question, error){
/*
	_q:=Question{}
	rows, _ := db.Queryx(`SELECT _id as "guid", title, name, description as "desc", type FROM "pcs-question-tbl"`)
	
	for rows.Next() {
		err := rows.StructScan(&_q)
		if err != nil {
			log.Fatalln(err)
		} 
		fmt.Printf("%#v\n", _q)
	}
*/
	
	_qs := []Question{}
	db.Select(&_qs, `SELECT _id as "guid", title, name, description as "desc", type,"answerOptions" as "answeroptions" FROM "pcs-question-tbl"`)

	//_ref1:=_qs[0]
	
	fmt.Println(_qs)

	return _qs,nil

}
