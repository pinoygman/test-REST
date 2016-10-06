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

package config

import (
	"encoding/json"
	utils "github.build.ge.com/predixsolutions/catalog-onboarding-api/utils"
)

type Config struct {
	Port                     string `json:"port"`
//	DataPath                 string `json:"data_path"`
//	CatalogPath              string `json:"catalog_path"`
}

func Init(path string) (*Config, error) {
	c:=&Config{}

	bytes, err:= utils.ReadFile(path)

	if err != nil {
		return c, err
	}
	
	if err=json.Unmarshal(bytes, c);err != nil {
		return c, err
	}
	
	return c, nil
}
