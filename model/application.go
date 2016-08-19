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
	"github.com/cloudfoundry-community/go-cfenv"
	"fmt"
	"encoding/json"
	"time"
	"github.com/pborman/uuid"
	"gopkg.in/redis.v4"
	"github.com/jmoiron/sqlx"
	sqltypes "github.com/jmoiron/sqlx/types"
	"errors"
	"log"
	"os"
	"strings"
)

//question type
const (
	
)

type Application struct {
	Guid              string                  `json:"_id"`
	ProfileId         string                  `json:"profileId"`
	Name              string                  `json:"applicationName"`
	Answers           sqltypes.JSONText       `json:"answers"`
	Notification      sqltypes.JSONText       `json:"notification"`
	CreatedDate       time.Time               `json:"created_date"`
	ModifiedDate      time.Time               `json:"modified_date"`
	ModifiedBy        string                  `json:"modified_by"`
	Status            string                  `json:"applicationStatus"`
}

var (
	client    *redis.Client
	db        *sqlx.DB
)

func init(){
	

	//init redis client
	
}

func (a *Application) Save() (*Application, error) {

	if a.Guid=="" {
		a.Guid=uuid.New()
	}

	_ea:=make(map[string]*Application)
	
	_b, err := client.Get(a.ProfileId).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist in redis")
		//do nothing
		//_ea=make(map[string]*Application)
		
	} else if err != nil {
		return nil, err.(error)
	} else {

		//fmt.Println(_b)
		json.Unmarshal([]byte(_b),&_ea)

		//fmt.Println(_ea)
		if _ea==nil{
			_ea=make(map[string]*Application)
		}
	}

	if a.CreatedDate.Equal(time.Time{}) {
		a.CreatedDate=time.Now()
	}
	
	a.ModifiedDate=time.Now()    
	a.ModifiedBy=CurrentProfile.ProfileId

	_ea[a.Guid]=a

	b,_:=json.Marshal(_ea)
	err2 := client.Set(CurrentProfile.ProfileId, string(b), 0).Err()

	if err2!=nil{
		return nil, err2
	}

	return a, nil

}

func (a *Application) Submit() (*Application,error){

	fmt.Println(a)
	
	_ea:=make(map[string]*Application)
	_b, err := client.Get(CurrentProfile.ProfileId).Result()
	if err == redis.Nil {
		return nil, errors.New("draft not found.")
		
	} else if err != nil {
		return nil, err
	} 

	json.Unmarshal([]byte(_b),&_ea)

	if _ea==nil{
		return nil, errors.New("draft not found.")
	}

	if _ea[a.Guid]==nil {
		return nil, errors.New("draft not found.")
	}

	a.Save()
	
	tx := db.MustBegin()

	tx.MustExec(`INSERT INTO "pcs-application-tbl" (_id, profileid, name, answers, status, notification, modifieddate, modifiedby) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,a.Guid,a.ProfileId,a.Name,a.Answers,a.Status,a.Notification,a.ModifiedDate,a.ModifiedBy)

	tx.Commit()

	DeleteDraftById(a.Guid)
	
	return a, nil
}

func InitRedis() error {

	if cfEnv, err := cfenv.Current(); err != nil {
		if err:=InitRedisClient("localhost","7991","8f5a2bd2-09db-4b6b-b6e7-2d191b07b11a");err!=nil{
			return err
		}
		
	} else {

		for k, _:= range cfEnv.Services {
			if strings.Contains(k,"redis") {
				o:=cfEnv.Services[k][0].Credentials
				err:=InitRedisClient(o["host"].(string),fmt.Sprintf("%.0f",o["port"]),o["password"].(string))
				if err!=nil {
					return err
				}
			}
		}
		
	}

	return nil
	
}

func InitRedisClient(_host string,_port string, _pwd string) error {

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s",_host,_port),
		Password: _pwd,
		DB:       0,
	})

	if _, err := client.Ping().Result();err!=nil{
		return err
	}

	InitQuestionType()
	
	return nil
}

func InitPostgresSql() error {
	_ref:=""
	if _ref=os.Getenv("SQLDSN");_ref=="" {
		_ref="host=localhost|port=7990|user=uc49c9583047d4173a217667509e17ddf|password=fb46202694704a7d994dd8e906666e6c|dbname=d13291d5f50c645f5b90d26b8a58e2f6b|connect_timeout=5|sslmode=disable"
	}
	
	_conn:=strings.Replace(_ref,"|"," ",-1)
	op, err := sqlx.Connect("postgres",_conn)
	
	if err != nil {
		log.Fatalln(err)
		return err
	}
	
	db=op
	return nil

}

func GetApplicationsByProfileId(pId string) ([]Application,error){
	_ap := []Application{}
	db.Select(&_ap, `SELECT _id as "guid", profileid, name, status, answers, notification,createddate, modifieddate, modifiedby FROM "pcs-application-tbl"`)// where "profileId"=$1`,pId)
	//created_date, last_modifed,
	return _ap,nil


}

func GetDraftsByProfileId(pId string) ([]Application ,error){

	_ea:=make(map[string]*Application)
	_b, err := client.Get(pId).Result()
	if err == redis.Nil {
		return nil, errors.New("draft not found 1.")
		
	} else if err != nil {
		return nil, err
	} 

	json.Unmarshal([]byte(_b),&_ea)

	if _ea==nil{
		return nil, errors.New("draft not found 2")
	}

	var _a  []Application
	for _, v := range _ea {
		_a=append(_a,*v)
	}
	return _a,nil
	
}

func DeleteDraftById(guid string) (error){

	_ea:=make(map[string]*Application)
	_b, err := client.Get(CurrentProfile.ProfileId).Result()
	if err == redis.Nil {
		return errors.New("draft not found 1.")
		
	} else if err != nil {
		return err
	} 

	json.Unmarshal([]byte(_b),&_ea)

	delete(_ea,guid)
	
	_v,_:=json.Marshal(_ea)

	err2 := client.Set(CurrentProfile.ProfileId, string(_v), 0).Err()

	if err2!=nil{
		return err2
	}

	return nil
}


func DeleteApplicationById(guid string) (error){

	if err:=DeleteDraftById(guid);err!=nil{
		return err
	}

	tx := db.MustBegin()

	tx.MustExec(`DELETE FROM "pcs-application-tbl" WHERE _id=$1`,guid)
	
	tx.Commit()

	return nil
}

