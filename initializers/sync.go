package initializers

import "github.com/AhmedOthman94/go-crud-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
}
