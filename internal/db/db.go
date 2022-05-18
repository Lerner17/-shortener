package db

import (
	"github.com/Lerner17/shortener/internal/config"
	fileStorage "github.com/Lerner17/shortener/internal/db/file_storage"
	"github.com/Lerner17/shortener/internal/db/memdb"
)

type URLStorage interface {
	GetURL(string) (string, bool)
	CreateURL(string) (string, string)
}

func GetDB() URLStorage {
	cfg := config.GetConfig()

	if cfg.FileStoragePath == "" {
		return memdb.NewMemDB()
	} else {
		return fileStorage.NewFileStorage(cfg.FileStoragePath)
	}
}
