package migration

import (
	"fmt"
	"log"
	articleModels "xzl_blog/internal/modules/article/models"
	userModels "xzl_blog/internal/modules/user/models"
	"xzl_blog/pkg/database"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done ..")
}
