package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type config struct {
	Addr      string `json:"address"`
	Port      string `json:"port"`
	StaticDir string `json:"static_directory"`
	DBType    string `json:"database_type"`
	DBName    string `json:"database_name"`
	DBAddr    string `json:"database_address"`
	DBUser    string `json:"database_username"`
	DBPass    string `json:"database_password"`
}

func newConfig() *config {
	return &config{
		Addr:      "",
		Port:      "",
		StaticDir: "",
		DBType:    "",
		DBName:    "",
		DBAddr:    "",
		DBUser:    "",
		DBPass:    "",
	}
}

func (cfg *config) getAddressAndPort() string {
	return fmt.Sprintf("%v:%v", cfg.Addr, cfg.Port)
}

func (cfg *config) loadConfigFromPath(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to load configuration file from path: ", path)
	}
	defer f.Close()

	d := json.NewDecoder(f)
	err = d.Decode(&cfg)
	if err != nil {
		log.Fatal("Failed to decode json configuration file. Check your syntax!")
	}
}

func (cfg *config) getDatabaseConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=True", cfg.DBUser, cfg.DBPass, cfg.DBName)
}
