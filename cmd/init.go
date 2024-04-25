package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"team_4/internal/database"
	"team_4/internal/types"
)

var Log = log.New()

func init() {
	err := godotenv.Load()

	if err != nil {
		Log.Error("Error loading .env file")
	}

	log.Info("Loaded .env")

	database.Connect(os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_DATA"), os.Getenv("PG_PASS"))

	log.Info("Start migrating data")
	database.DB.AutoMigrate(&types.Account{}, &types.VerifyLink{})
}
