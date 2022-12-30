package model

import (
	"github.com/jinzhu/gorm"
)

func SetAutoMigrate(db *gorm.DB) {
	// db.AutoMigrate(&File{})
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Userfile{})
	// db.AutoMigrate(&Share{})
	// db.AutoMigrate(&UserImage{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&File{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&UserImage{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Userfile{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Share{})
}
