package main

import (
	"backend/db"
	"fmt"
)

func main() {
	//fmt.Println(algorithm.BuildLast("hahahahsgsdgsg"))
	_db, err := db.GetDatabase()
	is := &db.Testa{
		Name: "ttes",
		Is:   5,
	}

	_db.Create(&is)
	fmt.Println(_db, err)
}
