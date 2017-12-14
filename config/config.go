package config

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Report struct {
		InputDir	string	`json:"inputDir"`
		OutputDir	string	`json:"outputDir"`
		ExeSecond	int		`json:"exeSecond"`
		ExeMinute	int		`json:"exeMinute"`
	} `json:"report"`

	Kafka struct {
		BrokerList	string	`json:"brokerList"`
		Group		string	`json:"group"`
		Topic		string	`json:"topic"`
	} `json:"kafka"`

	Redis struct {
		Hosts		string	`json:"hosts"`
	} `json:"redis"`

	Database struct {
		Host		string	`json:"host"`
		Port		string	`json:"port"`
		Db			string	`json:"db"`
		Username	string	`json:"username"`
		Password	string	`json:"password"`
	} `json:"database"`
}

var AppConfig *Config

func init() {
	f, err := os.Open("config.json")
	if nil != err {
		panic(err)
	}
	defer f.Close()

	jsonBuf, err := ioutil.ReadAll(f)
	cfg := new(Config)

	json.Unmarshal(jsonBuf, cfg)

	AppConfig = cfg

	buf, _ := json.MarshalIndent(cfg, "", "\t")
	log.Printf("config = %s\n", string(buf))
}
