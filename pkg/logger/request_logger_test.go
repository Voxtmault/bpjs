package logger

import (
	"testing"

	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/storage"
)

var envPath = "../../.env"

func TestInitRequestLogger(t *testing.T) {
	cfg := config.New(envPath)
	if err := storage.InitMariaDB(&cfg.DBConfig); err != nil {
		t.Errorf("Error initializing MariaDB: %v", err)
		return
	}
	if err := storage.InitRedis(&cfg.RedisConfig); err != nil {
		t.Errorf("Error initializing Redis: %v", err)
		return
	}
	InitRequestLogger()
}
