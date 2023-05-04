package db

import (
	"backend/models"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	//sqlDb, err := sql.Open("mysql", "root:Fchrgrib2310*@tcp(localhost:3306)/chatgpt")

	sqlDb, err := sql.Open("mysql", "root:jE9nkshB0SbXDRqbj3KK@tcp(containers-us-west-178.railway.app:7240)/railway")

	//sqlDb, err := sql.Open("mysql", "cau16diyel4vzj2b:insrrqyxdjkfw6hv@tcp(lfmerukkeiac5y5w.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/ibo4dp9c0tuurvvg")

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
