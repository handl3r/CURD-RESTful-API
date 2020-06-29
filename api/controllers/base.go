package controllers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/thaibuixuanDEV/forum/api/models"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DBName string) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser,
		DbPassword, DbHost, DbPort, DBName)
	server.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		fmt.Printf("Can not connect to %s database", DbDriver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("Connected to database %s", DbDriver)
	}

	server.DB.AutoMigrate(&models.User{}, &models.Post{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}
func (server *Server) Run(address string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(address, server.Router))
}
