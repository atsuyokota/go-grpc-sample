package config

import (
	"fmt"
	"log"

	ini "gopkg.in/ini.v1"
)

type ConfigList struct {
	MongoURI string
	MongoDB  string
	LogFile  string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("server/config.ini")
	if err != nil {
		log.Fatalf("Failed to read file %v", err)
	}
	mongoURI := fmt.Sprintf(
		"mongodb://%v:%v",
		cfg.Section("mongo").Key("host").String(),
		cfg.Section("mongo").Key("port").String(),
	)

	Config = ConfigList{
		MongoURI: mongoURI,
		MongoDB:  cfg.Section("mongo").Key("db").String(),
		LogFile:  cfg.Section("log").Key("file").String(),
	}
}
