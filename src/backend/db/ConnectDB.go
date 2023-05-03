package db

import (
	"backend/models"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	sqlDb, err := sql.Open("mysql", "root:Fchrgrib2310*@tcp(localhost:3306)/chatgpt")

	//sqlDb, err := sql.Open("postgres", "user=postgres password=Fchrgrib2310* host=db.yeiajanscvcinnhtrzfw.supabase.co port=5432 dbname=postgres")

	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.Chat{},
		&models.QuestAns{},
		&models.ChatHistory{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
