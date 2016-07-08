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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

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
	file, err := os.Open(p)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	
	return bytes, nil
}

func RetrieveUploadFile(r *http.Request) error{

	//16 mb max (x << n == x*2^n ) (x >> n == x*2^(-n)
	r.ParseMultipartForm(32 << 19)

	file, header , err := r.FormFile("uploadfile")
        filename := header.Filename
        fmt.Println(filename)

	out, err := os.Create("./tmp/"+filename)

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return err
        }
	
        defer out.Close()

	if _, err = io.Copy(out, file);err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return err
        }

	return nil

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
