package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/thaibuixuanDEV/forum/api/models"
	"log"
)

var users = []models.User{
	{
		Nickname: "thai0bui",
		Email:    "thaibui0@gmail.com",
		Password: "thai0bui",
	},
	{
		Nickname: "thai1bui",
		Email:    "thaibui1@gmail.com",
		Password: "thai1bui",
	},
}

var posts = []models.Post{
	{
		Title:   "Title 0",
		Content: "content0",
	},
	{
		Title:   "Title 1",
		Content: "content1",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("can not drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("can not migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attach foregin key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("can not seed users table: %v", err)
		}

		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("can not seed posts table: %v", err)
		}
	}
}
