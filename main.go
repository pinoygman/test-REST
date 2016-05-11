package main

import (
	"fmt"
	//"os"
	//"errors"
	//"flag"
	//"log"
	//"path/filepath"
	conf "github.build.ge.com/PredixSolutions/catalog-onboarding-backend/config"
	utils "github.build.ge.com/PredixSolutions/catalog-onboarding-backend/utils"
	//"github.com/gorilla/mux"	
)

func init(){

}

func main() {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(dir)
	defaultConfigPath := utils.GetPath([]string{"assets", "settings.json"})
	fmt.Printf("%+v\n", defaultConfigPath)
	_, err := conf.LoadConfig(defaultConfigPath)
	if err != nil {
		panic(fmt.Sprintf("Error creating server [%s]...", err.Error))
	}
	
	fmt.Printf("This's catalog onboarding publishing service.\n")
}
