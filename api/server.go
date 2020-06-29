package api

import (
	"fmt"
	"github.com/thaibuixuanDEV/forum/api/controllers"
	"github.com/joho/godotenv"
	"github.com/thaibuixuanDEV/forum/api/seed"
	"log"
	"os"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can not get ENV, not comming through %v", err)
	} else {
		fmt.Println("Getting ENV")
	}
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	seed.Load(server.DB)
	server.Run(":8080")
}
