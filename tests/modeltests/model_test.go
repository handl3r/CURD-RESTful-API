package modeltests

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/thaibuixuanDEV/forum/api/controllers"
	"github.com/thaibuixuanDEV/forum/api/models"
	"log"
	"os"
	"testing"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Getting an error when fetch ENV: %v", err)
	}
	Database()
	os.Exit(m.Run())
}

func Database() {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"),
		os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
	server.DB, err = gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Println("Can not connect to database!")
		log.Fatalf("Can not connect to DB: %v", err)
	} else {
		fmt.Printf("Connected to DB")
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("refesh user table successfully")
	return nil
}

func seedOneUser() (models.User, error) {
	refreshUserTable()

	user := models.User{
		Nickname: "nickname",
		Email:    "email@gmail.com",
		Password: "thaibuixuan",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatal("Can not seed a user")
	}
	return user, nil
}

func seedUsers() error {
	users := []models.User{
		models.User{
			Nickname: "nickname1",
			Email:    "email1@gmail.com",
			Password: "thaibuixuan",
		},
		models.User{
			Nickname: "nickname2",
			Email:    "email2@gmail.com",
			Password: "thaibuixuan",
		},
	}
	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndPostTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Println("Successfully refreshed tables")
	return nil
}

func seedOneUserAndPost() (models.Post, error) {
	user := models.User{
		Nickname: "nickname3",
		Email:    "email3@gmail.com",
		Password: "thaibuixuan",
	}
	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}

	post := models.Post{
		Title:    "title1",
		Content:  "content1",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func seedUserAndPosts() ([]models.User, []models.Post, error) {
	var err error
	users := []models.User{
		models.User{
			Nickname: "nickname1",
			Email:    "email1@gmail.com",
			Password: "thaibuixuan",
		},
		models.User{
			Nickname: "nickname2",
			Email:    "email2@gmail.com",
			Password: "thaibuixuan",
		},
	}
	posts := []models.Post{
		models.Post{
			Title:   "title1",
			Content: "content1",
		},
		models.Post{
			Title:   "title2",
			Content: "content2",
		},
	}

	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal("Can not create user")
		}

		posts[i].AuthorID = users[i].ID
		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatal("Can not create post")
		}
	}
	return users, posts, nil
}
