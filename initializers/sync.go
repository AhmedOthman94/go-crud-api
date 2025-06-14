package initializers

import "go-crud-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
}
