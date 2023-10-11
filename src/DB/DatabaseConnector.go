package DB

import (
	"cells-auth-server/src/Config"
	"cells-auth-server/src/Models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		Config.Cfg.DB.Host,
		Config.Cfg.DB.Port,
		Config.Cfg.DB.Database,
		Config.Cfg.DB.User,
		Config.Cfg.DB.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	return DB.AutoMigrate(Models.User{})
}
