package common

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	PORT = os.Getenv("PORT")
}
