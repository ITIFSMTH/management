package db

import (
	"log"
	"management-backend/models"
	"management-backend/shared"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func InitDB(DBPath string) {
	db, err = gorm.Open(postgres.Open(DBPath), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.Setting{})

	db.AutoMigrate(&models.Theme{})

	db.AutoMigrate(&models.WorkerRole{})
	db.AutoMigrate(&models.Worker{})

	db.AutoMigrate(&models.Timeout{})
	db.AutoMigrate(&models.Shift{})

	db.AutoMigrate(&models.Operator{})
	db.AutoMigrate(&models.JuniorOperator{})
	db.AutoMigrate(&models.SeniorOperator{})

	db.AutoMigrate(&models.PollType{})
	db.AutoMigrate(&models.Poll{})
	db.AutoMigrate(&models.BudgetPoll{})

	db.AutoMigrate(&models.BudgetVote{})
	db.AutoMigrate(&models.RatingVote{})

	if db.Model(&models.Setting{}).First(&models.Setting{}).RowsAffected == 0 {
		db.Create(&shared.Setting)

		db.Create(&shared.Themes)

		db.Create(&shared.WorkerRoles)
		db.Create(&shared.Workers)
	}
}

func GetDB() *gorm.DB {
	return db
}
