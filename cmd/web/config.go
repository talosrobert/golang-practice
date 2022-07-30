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
}

func newConfig() *config {
	return &config{
		Addr:      "",
		Port:      "",
		StaticDir: "",
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
