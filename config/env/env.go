package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DataBase struct {
	Host     string
	Port     int64
	User     string
	Name     string
	Password string
	TimeZone string
}

var (
	GinMode       string
	DbConnection  DataBase
	AwesomeApiUrl string
)

func Load() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Running application without env file")
	}

	GinMode = os.Getenv("GIN_MODE")

	DbConnection = DataBase{}
	DbConnection.Host = os.Getenv("DB_HOST")
	DbConnection.User = os.Getenv("DB_USER")
	DbConnection.Name = os.Getenv("DB_NAME")
	DbConnection.Password = os.Getenv("DB_PASSWORD")
	DbConnection.TimeZone = os.Getenv("DB_TIME_ZONE")
	DbConnection.Port, _ = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
}
