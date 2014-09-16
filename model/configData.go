package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"scalarm_monitoring_daemon/utils"
	//"errors"
)

type ConfigData struct {
	InformationServiceAddress string
	Login                     string
	Password                  string
	Infrastructures           []string
	ScalarmCertificatePath    string
	ScalarmScheme             string
}

func ReadConfiguration() (*ConfigData, error) {
	log.Printf("readConfiguration")

	data, err := ioutil.ReadFile("config.json")
	utils.Check(err)

	var configData ConfigData
	err = json.Unmarshal(data, &configData)
	utils.Check(err)

	if configData.ScalarmScheme == "" {
		configData.ScalarmScheme = "https"
	}

	log.Printf("\tinformation service address: " + configData.InformationServiceAddress)
	log.Printf("\tlogin:                       " + configData.Login)
	log.Printf("\tpassword:                    " + configData.Password)
	log.Printf("\tinfrastructures:             %v", configData.Infrastructures)

	log.Printf("readConfiguration: OK")
	return &configData, nil
}
