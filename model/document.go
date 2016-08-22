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
	//"strings"

	//"github.com/aws/aws-sdk-go/service/s3/s3manager"
	//"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	//"github.build.ge.com/predixsolutions/catalog-onboarding-backend/utils"
	_ "github.com/lib/pq"
	"time"
	//"log"
	
	"fmt"
	//"io"
	//"net/http"
)

//question type
const (
//	DueDiligent = 1001
//	Architecture = 1002
//	Security = 1003
)

type Document struct {
	Guid              string                    `json:"_id"`
	Label             string                    `json:"label"`
	UploadId          string                    `json:"uploadId"`
	ContentType       string                    `json:"content_type"`
	FileName          string                    `json:"fileName"`
	CreatedDate       time.Time                 `json:"created_date"`
	CreatedBy         string                    `json:"created_by"`
}

var (
	//questionnaires map[string]Questionnaire
)

func (d *Document) Create() (*Document, error) {

	if d.Guid=="" {
		d.Guid=uuid.New()
	}
	
	tx := db.MustBegin()

	tx.MustExec(`INSERT INTO "pcs-document-tbl" (_id, label, uploadid, filename, createddate, createdby, contenttype) VALUES ($1, $2, $3, $4, $5, $6, $7)`, d.Guid,d.Label,d.UploadId,d.FileName, d.CreatedDate, d.CreatedBy, d.ContentType)
	
	tx.Commit()

	return d,nil
}

func (d *Document) Load() (*Document, error){
	//doc := Document{}
	fmt.Println(d.Guid)
	err := db.Get(d, `SELECT _id as guid, label, contenttype, uploadid, filename, createddate, createdby FROM "pcs-document-tbl" where _id=$1`,d.Guid)

	if err!=nil{
		return nil,err
	}
	
	return d,nil	
}

func GetDocumentsByProfileId(pId string) ([]Document, error){
	
	_qs := []Document{}
	db.Select(&_qs, `SELECT _id as guid, label, contenttype, uploadid, filename, createddate, createdby FROM "pcs-document-tbl" where createdby=$1`,pId)

	return _qs,nil

}

func DeleteDocById(guid string) (error){

	tx := db.MustBegin()

	tx.MustExec(`DELETE FROM "pcs-document-tbl" WHERE _id=$1`,guid)
	
	tx.Commit()
	
	return nil

}






