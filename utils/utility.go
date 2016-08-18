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


package utils

import (
	"net/mail"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/pborman/uuid"
	"log"
)

const(
	LOGFILE = "./logs/"
)

func init(){
	
	/*
	var buf bytes.Buffer

	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")

	fmt.Print(&buf)


	f, err := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
*/
}
/*
func writeLog(p interface{}) error {

	if err:=os.MkdirAll(LOGFILE,0777);err!=nil{
		return nil, err
	}

	t := time.Now()
	fmt.Println(t.Format("20060102150405"))
	
	f, err := os.OpenFile(LOGFILE+"/", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	

	
}*/

func EncodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}


func GetPath(paths []string) string {
	workDirectory, _ := os.Getwd()

	if len(paths) == 0 {
		return workDirectory
	}

	resultPath := workDirectory

	for _, path := range paths {
		resultPath += string(os.PathSeparator)
		resultPath += path
	}

	return resultPath
}

func ReadFile(path string) (content []byte, err error) {
	p:=GetPath(strings.Split(path,"/"))
	fmt.Println(p)
	file, err := os.Open(p);

	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	
	return bytes, nil
}

func RetrieveUploadFile(r *http.Request,s string) (string, error) {

	//16 mb max (x << n == x*2^n ) (x >> n == x*2^(-n)
	r.ParseMultipartForm(32 << 19)

	file, header , err := r.FormFile("pcs_file")
        filename := header.Filename
        //fmt.Println(filename)

	uid:=uuid.New()
	
	out, err := os.Create(s+uid+"."+filename)

	if err != nil {
		return "", err 
        }
	
        defer out.Close()

	if _, err = io.Copy(out, file);err != nil {
		log.Fatal(err)
		return "", err
        }

	return uid, nil

	/*
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
		fmt.Println(err)
		return err
        }
	
        defer file.Close()
	
        fmt.Fprintf(w, "%v", handler.Header)
	
        f, err := os.OpenFile("./tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		return err
        }
	
        defer f.Close()
	
        io.Copy(f, file)
	*/
}

func WriteResponse(w http.ResponseWriter, code int, object interface{}) {
	data, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	fmt.Fprintf(w, string(data))
}
