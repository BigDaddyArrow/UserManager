package main

import (
	"encoding/json"
	"os"
)

type config struct {
	ServerHost string `json:"server_host"`
	ServerPort string `json:"server_port"`
	PgHost     string `json:"pg_host"`
	PgPort     string `json:"pg_port"`
	PgUser     string `json:"pg_user"`
	PgPass     string `json:"pg_pass"`
	PgDB       string `json:"pg_database"`
}

var cfg config

const conffile = "config.json"

func OpenCfg() error {
	f, err := os.Open(conffile)
	if err != nil {
		return err
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	readByte := make([]byte, stat.Size())

	_, err = f.Read(readByte)
	if err != nil {
		return err
	}

	err = json.Unmarshal(readByte, &cfg)
	if err != nil {
		return err
	}

	return err
}
