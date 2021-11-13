package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	DbPassword string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini") // configファイル読み込み
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1) // プログラム終了
	}

	Config = ConfigList{
		DbPassword: cfg.Section("db").Key("password").String(),
	}
}
