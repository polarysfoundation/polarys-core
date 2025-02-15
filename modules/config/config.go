package config

import (
	"log"
	"os"
	"path/filepath"
)

const (
	BasePath     = ".polarys"
	KeystorePath = "keystore"
)

func GetKeystorePath() string {
	return filepath.Join(getDir(BasePath, KeystorePath))
}

func getDir(dir string, subdir string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get home directory: %v", err)
	}
	subDirPath := filepath.Join(homeDir, dir, subdir)

	err = os.MkdirAll(subDirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create %s directory: %v", subdir, err)
	}

	return subDirPath
}
