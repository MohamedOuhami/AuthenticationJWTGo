package initializers

import (
	"github.com/MohamedOuhami/AuthenticationJWTGo/models"
)

func SyncDatabase() {
  // Used to sync the models with the tables present in the database
	DB.AutoMigrate(&models.User{})
}
