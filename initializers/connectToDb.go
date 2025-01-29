package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Making the Database global in the project
var (
  DB *gorm.DB
)

func ConnectToDB() {

  var err error
	dsn := os.Getenv("DB_URL") 

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to the database")
  }
}
